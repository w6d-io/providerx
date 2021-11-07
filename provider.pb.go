// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: provider.proto

package providerx

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OwnerType int32

const (
	OwnerType_USER  OwnerType = 0
	OwnerType_GROUP OwnerType = 1
)

// Enum value maps for OwnerType.
var (
	OwnerType_name = map[int32]string{
		0: "USER",
		1: "GROUP",
	}
	OwnerType_value = map[string]int32{
		"USER":  0,
		"GROUP": 1,
	}
)

func (x OwnerType) Enum() *OwnerType {
	p := new(OwnerType)
	*p = x
	return p
}

func (x OwnerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OwnerType) Descriptor() protoreflect.EnumDescriptor {
	return file_provider_proto_enumTypes[0].Descriptor()
}

func (OwnerType) Type() protoreflect.EnumType {
	return &file_provider_proto_enumTypes[0]
}

func (x OwnerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OwnerType.Descriptor instead.
func (OwnerType) EnumDescriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{0}
}

type ListProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiUrl      string `protobuf:"bytes,1,opt,name=apiUrl,proto3" json:"apiUrl,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Page        int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage     int32  `protobuf:"varint,4,opt,name=perPage,proto3" json:"perPage,omitempty"`
	IdUser      int32  `protobuf:"varint,5,opt,name=idUser,proto3" json:"idUser,omitempty"`
	AccessLevel int32  `protobuf:"varint,6,opt,name=accessLevel,proto3" json:"accessLevel,omitempty"`
}

func (x *ListProjectRequest) Reset() {
	*x = ListProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectRequest) ProtoMessage() {}

func (x *ListProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectRequest.ProtoReflect.Descriptor instead.
func (*ListProjectRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ListProjectRequest) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *ListProjectRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ListProjectRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProjectRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ListProjectRequest) GetIdUser() int32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *ListProjectRequest) GetAccessLevel() int32 {
	if x != nil {
		return x.AccessLevel
	}
	return 0
}

type ListProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Eof      bool       `protobuf:"varint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	Projects []*Project `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	Response bool       `protobuf:"varint,4,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ListProjectResponse) Reset() {
	*x = ListProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectResponse) ProtoMessage() {}

func (x *ListProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectResponse.ProtoReflect.Descriptor instead.
func (*ListProjectResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProjectResponse) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

func (x *ListProjectResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *ListProjectResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type SetWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlRepo string `protobuf:"bytes,1,opt,name=urlRepo,proto3" json:"urlRepo,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	IdRepo  int32  `protobuf:"varint,3,opt,name=idRepo,proto3" json:"idRepo,omitempty"`
}

func (x *SetWebhookRequest) Reset() {
	*x = SetWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetWebhookRequest) ProtoMessage() {}

func (x *SetWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetWebhookRequest.ProtoReflect.Descriptor instead.
func (*SetWebhookRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{2}
}

func (x *SetWebhookRequest) GetUrlRepo() string {
	if x != nil {
		return x.UrlRepo
	}
	return ""
}

func (x *SetWebhookRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetWebhookRequest) GetIdRepo() int32 {
	if x != nil {
		return x.IdRepo
	}
	return 0
}

type SetWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  bool  `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	WebhookId int32 `protobuf:"varint,2,opt,name=webhookId,proto3" json:"webhookId,omitempty"`
}

func (x *SetWebhookResponse) Reset() {
	*x = SetWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetWebhookResponse) ProtoMessage() {}

func (x *SetWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetWebhookResponse.ProtoReflect.Descriptor instead.
func (*SetWebhookResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{3}
}

func (x *SetWebhookResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

func (x *SetWebhookResponse) GetWebhookId() int32 {
	if x != nil {
		return x.WebhookId
	}
	return 0
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url            string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	HttpUrl        string `protobuf:"bytes,4,opt,name=httpUrl,proto3" json:"httpUrl,omitempty"`
	DefaultBranch  string `protobuf:"bytes,5,opt,name=defaultBranch,proto3" json:"defaultBranch,omitempty"`
	LastActivityAt string `protobuf:"bytes,6,opt,name=lastActivityAt,proto3" json:"lastActivityAt,omitempty"`
	FullPath       string `protobuf:"bytes,7,opt,name=fullPath,proto3" json:"fullPath,omitempty"`
	Private        bool   `protobuf:"varint,8,opt,name=private,proto3" json:"private,omitempty"`
	Admin          bool   `protobuf:"varint,9,opt,name=admin,proto3" json:"admin,omitempty"`
	Owner          *Owner `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{4}
}

func (x *Project) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Project) GetHttpUrl() string {
	if x != nil {
		return x.HttpUrl
	}
	return ""
}

func (x *Project) GetDefaultBranch() string {
	if x != nil {
		return x.DefaultBranch
	}
	return ""
}

func (x *Project) GetLastActivityAt() string {
	if x != nil {
		return x.LastActivityAt
	}
	return ""
}

func (x *Project) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

func (x *Project) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *Project) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *Project) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login     string    `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	AvatarUrl string    `protobuf:"bytes,3,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Type      OwnerType `protobuf:"varint,4,opt,name=type,proto3,enum=OwnerType" json:"type,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{5}
}

func (x *Owner) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Owner) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Owner) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Owner) GetType() OwnerType {
	if x != nil {
		return x.Type
	}
	return OwnerType_USER
}

type DeleteWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlRepo   string `protobuf:"bytes,1,opt,name=urlRepo,proto3" json:"urlRepo,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	IdRepo    int32  `protobuf:"varint,3,opt,name=idRepo,proto3" json:"idRepo,omitempty"`
	WebhookId int32  `protobuf:"varint,4,opt,name=webhookId,proto3" json:"webhookId,omitempty"`
}

func (x *DeleteWebhookRequest) Reset() {
	*x = DeleteWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebhookRequest) ProtoMessage() {}

func (x *DeleteWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebhookRequest.ProtoReflect.Descriptor instead.
func (*DeleteWebhookRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWebhookRequest) GetUrlRepo() string {
	if x != nil {
		return x.UrlRepo
	}
	return ""
}

func (x *DeleteWebhookRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteWebhookRequest) GetIdRepo() int32 {
	if x != nil {
		return x.IdRepo
	}
	return 0
}

func (x *DeleteWebhookRequest) GetWebhookId() int32 {
	if x != nil {
		return x.WebhookId
	}
	return 0
}

type DeleteWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Deleted  bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteWebhookResponse) Reset() {
	*x = DeleteWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebhookResponse) ProtoMessage() {}

func (x *DeleteWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebhookResponse.ProtoReflect.Descriptor instead.
func (*DeleteWebhookResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteWebhookResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

func (x *DeleteWebhookResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CheckAccessUserInProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlRepo string `protobuf:"bytes,1,opt,name=urlRepo,proto3" json:"urlRepo,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckAccessUserInProjectRequest) Reset() {
	*x = CheckAccessUserInProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessUserInProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessUserInProjectRequest) ProtoMessage() {}

func (x *CheckAccessUserInProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessUserInProjectRequest.ProtoReflect.Descriptor instead.
func (*CheckAccessUserInProjectRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{8}
}

func (x *CheckAccessUserInProjectRequest) GetUrlRepo() string {
	if x != nil {
		return x.UrlRepo
	}
	return ""
}

func (x *CheckAccessUserInProjectRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckAccessUserInProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Admin    bool `protobuf:"varint,2,opt,name=admin,proto3" json:"admin,omitempty"`
	Push     bool `protobuf:"varint,3,opt,name=push,proto3" json:"push,omitempty"`
	Pull     bool `protobuf:"varint,4,opt,name=pull,proto3" json:"pull,omitempty"`
}

func (x *CheckAccessUserInProjectResponse) Reset() {
	*x = CheckAccessUserInProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessUserInProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessUserInProjectResponse) ProtoMessage() {}

func (x *CheckAccessUserInProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessUserInProjectResponse.ProtoReflect.Descriptor instead.
func (*CheckAccessUserInProjectResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{9}
}

func (x *CheckAccessUserInProjectResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

func (x *CheckAccessUserInProjectResponse) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *CheckAccessUserInProjectResponse) GetPush() bool {
	if x != nil {
		return x.Push
	}
	return false
}

func (x *CheckAccessUserInProjectResponse) GetPull() bool {
	if x != nil {
		return x.Pull
	}
	return false
}

var File_provider_proto protoreflect.FileDescriptor

var file_provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaa, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x55, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x7d, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x66, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x22, 0x4e, 0x0a, 0x12, 0x53, 0x65, 0x74,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x74, 0x74, 0x70, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x74,
	0x74, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x1c, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x6b, 0x0a,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7c, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x1f, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x72,
	0x6c, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7c, 0x0a, 0x20, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x70, 0x75, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x70, 0x75, 0x6c, 0x6c, 0x2a, 0x20, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x01, 0x32, 0xa3, 0x02, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x12, 0x12, 0x2e, 0x53, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x15, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a,
	0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x36, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x78, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provider_proto_rawDescOnce sync.Once
	file_provider_proto_rawDescData = file_provider_proto_rawDesc
)

func file_provider_proto_rawDescGZIP() []byte {
	file_provider_proto_rawDescOnce.Do(func() {
		file_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_provider_proto_rawDescData)
	})
	return file_provider_proto_rawDescData
}

var file_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_provider_proto_goTypes = []interface{}{
	(OwnerType)(0),                           // 0: OwnerType
	(*ListProjectRequest)(nil),               // 1: ListProjectRequest
	(*ListProjectResponse)(nil),              // 2: ListProjectResponse
	(*SetWebhookRequest)(nil),                // 3: SetWebhookRequest
	(*SetWebhookResponse)(nil),               // 4: SetWebhookResponse
	(*Project)(nil),                          // 5: Project
	(*Owner)(nil),                            // 6: Owner
	(*DeleteWebhookRequest)(nil),             // 7: DeleteWebhookRequest
	(*DeleteWebhookResponse)(nil),            // 8: DeleteWebhookResponse
	(*CheckAccessUserInProjectRequest)(nil),  // 9: CheckAccessUserInProjectRequest
	(*CheckAccessUserInProjectResponse)(nil), // 10: CheckAccessUserInProjectResponse
}
var file_provider_proto_depIdxs = []int32{
	5,  // 0: ListProjectResponse.projects:type_name -> Project
	6,  // 1: Project.owner:type_name -> Owner
	0,  // 2: Owner.type:type_name -> OwnerType
	1,  // 3: Provider.LisProject:input_type -> ListProjectRequest
	3,  // 4: Provider.SetWebhook:input_type -> SetWebhookRequest
	7,  // 5: Provider.DeleteWebhook:input_type -> DeleteWebhookRequest
	9,  // 6: Provider.CheckAccessUserInProject:input_type -> CheckAccessUserInProjectRequest
	2,  // 7: Provider.LisProject:output_type -> ListProjectResponse
	4,  // 8: Provider.SetWebhook:output_type -> SetWebhookResponse
	8,  // 9: Provider.DeleteWebhook:output_type -> DeleteWebhookResponse
	10, // 10: Provider.CheckAccessUserInProject:output_type -> CheckAccessUserInProjectResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_provider_proto_init() }
func file_provider_proto_init() {
	if File_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectRequest); i {
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
		file_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectResponse); i {
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
		file_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetWebhookRequest); i {
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
		file_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetWebhookResponse); i {
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
		file_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_provider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
		file_provider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWebhookRequest); i {
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
		file_provider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWebhookResponse); i {
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
		file_provider_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccessUserInProjectRequest); i {
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
		file_provider_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccessUserInProjectResponse); i {
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
			RawDescriptor: file_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_proto_goTypes,
		DependencyIndexes: file_provider_proto_depIdxs,
		EnumInfos:         file_provider_proto_enumTypes,
		MessageInfos:      file_provider_proto_msgTypes,
	}.Build()
	File_provider_proto = out.File
	file_provider_proto_rawDesc = nil
	file_provider_proto_goTypes = nil
	file_provider_proto_depIdxs = nil
}
