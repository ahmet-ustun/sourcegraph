// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: internalapi.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_proto_rawDescGZIP(), []int{0}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawUnified *RawUnified `protobuf:"bytes,1,opt,name=raw_unified,json=rawUnified,proto3" json:"raw_unified,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigResponse) GetRawUnified() *RawUnified {
	if x != nil {
		return x.RawUnified
	}
	return nil
}

type RawUnified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Site               string              `protobuf:"bytes,2,opt,name=site,proto3" json:"site,omitempty"`
	ServiceConnections *ServiceConnections `protobuf:"bytes,3,opt,name=service_connections,json=serviceConnections,proto3" json:"service_connections,omitempty"`
}

func (x *RawUnified) Reset() {
	*x = RawUnified{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawUnified) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawUnified) ProtoMessage() {}

func (x *RawUnified) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawUnified.ProtoReflect.Descriptor instead.
func (*RawUnified) Descriptor() ([]byte, []int) {
	return file_internalapi_proto_rawDescGZIP(), []int{2}
}

func (x *RawUnified) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RawUnified) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *RawUnified) GetServiceConnections() *ServiceConnections {
	if x != nil {
		return x.ServiceConnections
	}
	return nil
}

// ServiceConnections represents configuration about how the deployment
// internally connects to services. These are settings that need to be
// propagated from the frontend to other services, so that the frontend
// can be the source of truth for all configuration.
type ServiceConnections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// GitServers is the addresses of gitserver instances that should be
	// talked to.
	GitServers []string `protobuf:"bytes,1,rep,name=git_servers,json=gitServers,proto3" json:"git_servers,omitempty"`
	// PostgresDSN is the PostgreSQL DB data source name.
	// eg: "postgres://sg@pgsql/sourcegraph?sslmode=false"
	PostgresDsn string `protobuf:"bytes,2,opt,name=postgres_dsn,json=postgresDsn,proto3" json:"postgres_dsn,omitempty"`
	// CodeIntelPostgresDSN is the PostgreSQL DB data source name for the
	// code intel database.
	// eg: "postgres://sg@pgsql/sourcegraph_codeintel?sslmode=false"
	CodeIntelPostgresDsn string `protobuf:"bytes,3,opt,name=code_intel_postgres_dsn,json=codeIntelPostgresDsn,proto3" json:"code_intel_postgres_dsn,omitempty"`
	// CodeInsightsDSN is the PostgreSQL DB data source name for the
	// code insights database.
	// eg: "postgres://sg@pgsql/sourcegraph_codeintel?sslmode=false"
	CodeInsightsDsn string `protobuf:"bytes,4,opt,name=code_insights_dsn,json=codeInsightsDsn,proto3" json:"code_insights_dsn,omitempty"`
	// Searchers is the addresses of searcher instances that should be talked to.
	Searchers []string `protobuf:"bytes,5,rep,name=searchers,proto3" json:"searchers,omitempty"`
	// Symbols is the addresses of symbol instances that should be talked to.
	Symbols []string `protobuf:"bytes,6,rep,name=symbols,proto3" json:"symbols,omitempty"`
	// Embeddings is the addresses of embeddings instances that should be talked
	// to.
	Embeddings []string `protobuf:"bytes,7,rep,name=embeddings,proto3" json:"embeddings,omitempty"`
	// Qdrant is the address of the Qdrant instance (or empty if disabled)
	Qdrant string `protobuf:"bytes,8,opt,name=qdrant,proto3" json:"qdrant,omitempty"`
	// Zoekts is the addresses of Zoekt instances to talk to.
	Zoekts []string `protobuf:"bytes,9,rep,name=zoekts,proto3" json:"zoekts,omitempty"`
	// ZoektListTTL is the TTL of the internal cache that Zoekt clients use to
	// cache the list of indexed repository. After TTL is over, new list will
	// get requested from Zoekt shards.
	ZoektListTtl *durationpb.Duration `protobuf:"bytes,10,opt,name=zoekt_list_ttl,json=zoektListTtl,proto3" json:"zoekt_list_ttl,omitempty"`
}

func (x *ServiceConnections) Reset() {
	*x = ServiceConnections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConnections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConnections) ProtoMessage() {}

func (x *ServiceConnections) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConnections.ProtoReflect.Descriptor instead.
func (*ServiceConnections) Descriptor() ([]byte, []int) {
	return file_internalapi_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceConnections) GetGitServers() []string {
	if x != nil {
		return x.GitServers
	}
	return nil
}

func (x *ServiceConnections) GetPostgresDsn() string {
	if x != nil {
		return x.PostgresDsn
	}
	return ""
}

func (x *ServiceConnections) GetCodeIntelPostgresDsn() string {
	if x != nil {
		return x.CodeIntelPostgresDsn
	}
	return ""
}

func (x *ServiceConnections) GetCodeInsightsDsn() string {
	if x != nil {
		return x.CodeInsightsDsn
	}
	return ""
}

func (x *ServiceConnections) GetSearchers() []string {
	if x != nil {
		return x.Searchers
	}
	return nil
}

func (x *ServiceConnections) GetSymbols() []string {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *ServiceConnections) GetEmbeddings() []string {
	if x != nil {
		return x.Embeddings
	}
	return nil
}

func (x *ServiceConnections) GetQdrant() string {
	if x != nil {
		return x.Qdrant
	}
	return ""
}

func (x *ServiceConnections) GetZoekts() []string {
	if x != nil {
		return x.Zoekts
	}
	return nil
}

func (x *ServiceConnections) GetZoektListTtl() *durationpb.Duration {
	if x != nil {
		return x.ZoektListTtl
	}
	return nil
}

var File_internalapi_proto protoreflect.FileDescriptor

var file_internalapi_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x55, 0x6e,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x22, 0x89, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x77, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x74, 0x65, 0x12, 0x57, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x84, 0x03,
	0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x5f, 0x64, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x44, 0x73, 0x6e, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f,
	0x64, 0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x44, 0x73, 0x6e, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x5f, 0x64, 0x73, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x44, 0x73, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x7a,
	0x6f, 0x65, 0x6b, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x65,
	0x6b, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x7a, 0x6f, 0x65, 0x6b, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x7a, 0x6f, 0x65, 0x6b, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x74, 0x6c, 0x32, 0x6b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_proto_rawDescOnce sync.Once
	file_internalapi_proto_rawDescData = file_internalapi_proto_rawDesc
)

func file_internalapi_proto_rawDescGZIP() []byte {
	file_internalapi_proto_rawDescOnce.Do(func() {
		file_internalapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_proto_rawDescData)
	})
	return file_internalapi_proto_rawDescData
}

var file_internalapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_proto_goTypes = []interface{}{
	(*GetConfigRequest)(nil),    // 0: api.internalapi.v1.GetConfigRequest
	(*GetConfigResponse)(nil),   // 1: api.internalapi.v1.GetConfigResponse
	(*RawUnified)(nil),          // 2: api.internalapi.v1.RawUnified
	(*ServiceConnections)(nil),  // 3: api.internalapi.v1.ServiceConnections
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_internalapi_proto_depIdxs = []int32{
	2, // 0: api.internalapi.v1.GetConfigResponse.raw_unified:type_name -> api.internalapi.v1.RawUnified
	3, // 1: api.internalapi.v1.RawUnified.service_connections:type_name -> api.internalapi.v1.ServiceConnections
	4, // 2: api.internalapi.v1.ServiceConnections.zoekt_list_ttl:type_name -> google.protobuf.Duration
	0, // 3: api.internalapi.v1.ConfigService.GetConfig:input_type -> api.internalapi.v1.GetConfigRequest
	1, // 4: api.internalapi.v1.ConfigService.GetConfig:output_type -> api.internalapi.v1.GetConfigResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internalapi_proto_init() }
func file_internalapi_proto_init() {
	if File_internalapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawUnified); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConnections); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_proto_goTypes,
		DependencyIndexes: file_internalapi_proto_depIdxs,
		MessageInfos:      file_internalapi_proto_msgTypes,
	}.Build()
	File_internalapi_proto = out.File
	file_internalapi_proto_rawDesc = nil
	file_internalapi_proto_goTypes = nil
	file_internalapi_proto_depIdxs = nil
}
