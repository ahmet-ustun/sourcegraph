package com.sourcegraph.cody.config

import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.ui.ValidationInfo
import com.intellij.openapi.ui.setEmptyState
import com.intellij.ui.components.JBTextField
import com.intellij.ui.components.fields.ExtendableTextField
import com.intellij.ui.dsl.builder.MAX_LINE_LENGTH_NO_WRAP
import com.intellij.ui.layout.LayoutBuilder
import com.intellij.ui.layout.applyToComponent
import com.sourcegraph.cody.api.SourcegraphApiRequestExecutor
import com.sourcegraph.cody.api.SourcegraphAuthenticationException
import com.sourcegraph.cody.api.SourcegraphSecurityUtil
import com.sourcegraph.cody.config.DialogValidationUtils.custom
import com.sourcegraph.cody.config.DialogValidationUtils.notBlank
import com.sourcegraph.common.AuthorizationUtil
import java.net.UnknownHostException
import javax.swing.JComponent

internal class CodyTokenCredentialsUi(
    private val serverTextField: ExtendableTextField,
    private val customRequestHeadersField: ExtendableTextField,
    val factory: SourcegraphApiRequestExecutor.Factory,
    val isAccountUnique: UniqueLoginPredicate
) : CodyCredentialsUi() {

  private val tokenTextField = JBTextField()
  private var fixedLogin: String? = null

  fun setToken(token: String) {
    tokenTextField.text = token
  }

  override fun LayoutBuilder.centerPanel() {
    row("Server: ") { serverTextField(pushX, growX) }
    row("Token: ") { tokenTextField(constraints = arrayOf(pushX, growX)) }
    row("Custom request headers: ") {
      customRequestHeadersField(pushX, growX)
          .comment(
              """Any custom headers to send with every request to Sourcegraph.<br>
                  |Use any number of pairs: "header1, value1, header2, value2, ...".<br>
                  |Whitespace around commas doesn't matter.
              """
                  .trimMargin(),
              MAX_LINE_LENGTH_NO_WRAP)
          .applyToComponent { this.setEmptyState("Client-ID, client-one, X-Extra, some metadata") }
    }
  }

  override fun getPreferredFocusableComponent(): JComponent = tokenTextField

  override fun getValidator(): () -> ValidationInfo? = {
    notBlank(tokenTextField, "Token cannot be empty")
        ?: custom(tokenTextField, "Invalid access token") {
          AuthorizationUtil.isValidAccessToken(tokenTextField.text)
        }
  }

  override fun createExecutor() = factory.create(tokenTextField.text)

  override fun acquireLoginAndToken(
      server: SourcegraphServerPath,
      executor: SourcegraphApiRequestExecutor,
      indicator: ProgressIndicator
  ): Pair<String, String> {
    val login = acquireLogin(server, executor, indicator, isAccountUnique, fixedLogin)
    return login to tokenTextField.text
  }

  override fun handleAcquireError(error: Throwable): ValidationInfo =
      when (error) {
        is SourcegraphParseException ->
            ValidationInfo(error.message ?: "Invalid server url", serverTextField)
        else -> handleError(error)
      }

  override fun setBusy(busy: Boolean) {
    tokenTextField.isEnabled = !busy
  }

  fun setFixedLogin(fixedLogin: String?) {
    this.fixedLogin = fixedLogin
  }

  companion object {

    fun acquireLogin(
        server: SourcegraphServerPath,
        executor: SourcegraphApiRequestExecutor,
        indicator: ProgressIndicator,
        isAccountUnique: UniqueLoginPredicate,
        fixedLogin: String?
    ): String {
      val accountDetails =
          SourcegraphSecurityUtil.loadCurrentUserDetails(executor, indicator, server)

      val login = accountDetails.username
      if (fixedLogin != null && fixedLogin != login)
          throw SourcegraphAuthenticationException("Token should match username \"$fixedLogin\"")
      if (!isAccountUnique(login, server)) throw LoginNotUniqueException(login)

      return login
    }

    fun handleError(error: Throwable): ValidationInfo =
        when (error) {
          is LoginNotUniqueException ->
              ValidationInfo("Account '${error.login}' already added").withOKEnabled()
          is UnknownHostException -> ValidationInfo("Server is unreachable").withOKEnabled()
          is SourcegraphAuthenticationException ->
              ValidationInfo("Incorrect credentials.\n" + error.message.orEmpty()).withOKEnabled()
          else ->
              ValidationInfo("Invalid authentication data.\n" + error.message.orEmpty())
                  .withOKEnabled()
        }
  }
}
