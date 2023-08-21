package tst

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"os"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"

	"github.com/sourcegraph/sourcegraph/lib/errors"
)

func (gh *GitHubClient) selectOrg(ctx context.Context) (*github.Organization, error) {
	org, resp, err := gh.c.Organizations.Get(ctx, "william-templates")
	if resp.StatusCode >= 299 {
		return nil, errors.Newf("failed to find org: %s", "william-templates")
	} else if err != nil {
		return nil, err
	}
	return org, err
}

func (gh *GitHubClient) createOrg(ctx context.Context, name string) (*github.Organization, error) {
	newOrg := github.Organization{
		Login: &name,
	}

	org, resp, err := gh.c.Admin.CreateOrg(ctx, &newOrg, gh.cfg.User)
	if resp.StatusCode >= 400 {
		io.Copy(os.Stdout, resp.Body)
		return nil, errors.Newf("failed to create org %q - GitHub status code %d: %v", name, resp.StatusCode, err)
	}
	return org, err
}

func (gh *GitHubClient) orgUsers(ctx context.Context, org *github.Organization) ([]*github.User, error) {
	users, _, err := gh.c.Organizations.ListMembers(ctx, org.GetLogin(), &github.ListMembersOptions{})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (gh *GitHubClient) createUser(ctx context.Context, name, email string) (*github.User, error) {
	user, resp, err := gh.c.Admin.CreateUser(ctx, name, email)
	if resp.StatusCode >= 400 {
		return nil, errors.Newf("failed to create user %q - GitHub status code %d: %v", name, resp.StatusCode, err)
	}
	return user, nil
}

func (gh *GitHubClient) getUser(ctx context.Context, name string) (*github.User, error) {
	user, resp, err := gh.c.Users.Get(ctx, name)
	if resp.StatusCode >= 400 {
		return nil, errors.Newf("failed to get user %q - GitHub status code %d: %v", name, resp.StatusCode, err)
	}
	return user, nil
}

func (gh *GitHubClient) deleteUser(ctx context.Context, username string) error {
	resp, err := gh.c.Admin.DeleteUser(ctx, username)
	if resp.StatusCode >= 400 {
		return errors.Newf("failed to delete user %q - GitHub status code %d: %v", username, resp.StatusCode, err)
	}

	return nil
}

func (gh *GitHubClient) getOrCreateTeam(ctx context.Context, org *github.Organization, name string) (*github.Team, error) {
	team, resp, err := gh.c.Teams.GetTeamBySlug(ctx, org.GetLogin(), name)

	switch resp.StatusCode {
	case 200:
		return team, nil
	case 404:
		newTeam := github.NewTeam{
			Name:        name,
			Description: strp("auto created team"),
			Privacy:     strp("closed"),
		}
		team, _, err = gh.c.Teams.CreateTeam(ctx, org.GetLogin(), newTeam)
	}
	return team, err
}

func (gh *GitHubClient) deleteTeam(ctx context.Context, org *github.Organization, name string) error {
	resp, err := gh.c.Teams.DeleteTeamBySlug(ctx, org.GetLogin(), name)
	if resp.StatusCode >= 400 {
		return errors.Newf("failed to delete team %q - GitHub status code %d: %v", name, resp.StatusCode, err)
	}
	return err
}

func (gh *GitHubClient) assignTeamMembership(ctx context.Context, org *github.Organization, team *github.Team, users ...*github.User) (*github.Team, error) {
	for _, u := range users {
		_, resp, err := gh.c.Teams.GetTeamMembershipByID(ctx, org.GetID(), team.GetID(), u.GetLogin())
		if resp.StatusCode == 200 {
			// user is already part of this team
			return team, nil
		} else if resp.StatusCode >= 500 {
			return nil, errors.Newf("server error[%d]: %v", resp.StatusCode, err)
		}

		_, _, err = gh.c.Teams.AddTeamMembershipByID(ctx, org.GetID(), team.GetID(), u.GetLogin(), &github.TeamAddTeamMembershipOptions{
			Role: "member",
		})

		if err != nil {
			return nil, err
		}

	}
	return team, nil
}

func (gh *GitHubClient) getRepo(ctx context.Context, owner, repoName string) (*github.Repository, error) {
	repo, resp, err := gh.c.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, errors.Newf("failed to get repo %q - GitHub response code %d: err", repoName, resp.StatusCode, err)
	}

	return repo, nil
}

func (gh *GitHubClient) forkRepo(ctx context.Context, org *github.Organization, owner, repoName string) error {
	_, resp, err := gh.c.Repositories.CreateFork(ctx, owner, repoName, &github.RepositoryCreateForkOptions{
		Organization:      org.GetLogin(),
		Name:              repoName,
		DefaultBranchOnly: true,
	})
	if resp.StatusCode == 202 {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func (gh *GitHubClient) updateRepo(ctx context.Context, org *github.Organization, repo *github.Repository) (*github.Repository, error) {
	result, resp, err := gh.c.Repositories.Edit(ctx, org.GetLogin(), repo.GetName(), repo)

	if resp.StatusCode >= 400 {
		return nil, errors.Newf("failed to edit repository %q - github response status code %d: %v", repo.GetName(), resp.StatusCode, err)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gh *GitHubClient) deleteRepo(ctx context.Context, org *github.Organization, repo *github.Repository) error {
	resp, err := gh.c.Repositories.Delete(ctx, org.GetLogin(), repo.GetName())

	if resp.StatusCode >= 400 {
		return errors.Newf("failed to edit repository %q - github response status code %d: %v", repo.GetName(), resp.StatusCode, err)
	}

	if err != nil {
		return err
	}

	return nil
}

func (gh *GitHubClient) updateTeamRepoPermissions(ctx context.Context, org *github.Organization, team *github.Team, repo *github.Repository) error {
	resp, err := gh.c.Teams.AddTeamRepoByID(ctx, org.GetID(), team.GetID(), org.GetLogin(), repo.GetName(), &github.TeamAddTeamRepoOptions{})
	if resp.StatusCode != 204 {
		return errors.Newf("failed to update repo %q permissions for team %q: %v", repo.GetName(), team.GetSlug(), err)
	}
	return nil
}

func NewGitHubClient(ctx context.Context, cfg Config) (*GitHubClient, error) {
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.CodeHost.Token},
	))

	tc.Transport.(*oauth2.Transport).Base = http.DefaultTransport
	tc.Transport.(*oauth2.Transport).Base.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	gh, err := github.NewEnterpriseClient(cfg.CodeHost.URL, cfg.CodeHost.URL, tc)
	if err != nil {
		return nil, err
	}

	c := GitHubClient{
		cfg: &cfg.CodeHost,
		c:   gh,
	}
	return &c, nil
}
