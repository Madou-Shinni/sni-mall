// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: proto/rbacManager.proto

package rbacManager

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

//管理员对应的结构体
type ManagerModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Mobile   string     `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email    string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Status   int64      `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	RoleId   int64      `protobuf:"varint,7,opt,name=roleId,proto3" json:"roleId,omitempty"`
	AddTime  int64      `protobuf:"varint,8,opt,name=addTime,proto3" json:"addTime,omitempty"`
	IsSuper  int64      `protobuf:"varint,9,opt,name=isSuper,proto3" json:"isSuper,omitempty"`
	Role     *RoleModel `protobuf:"bytes,10,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *ManagerModel) Reset() {
	*x = ManagerModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerModel) ProtoMessage() {}

func (x *ManagerModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerModel.ProtoReflect.Descriptor instead.
func (*ManagerModel) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{0}
}

func (x *ManagerModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManagerModel) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ManagerModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManagerModel) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ManagerModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerModel) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ManagerModel) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *ManagerModel) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *ManagerModel) GetIsSuper() int64 {
	if x != nil {
		return x.IsSuper
	}
	return 0
}

func (x *ManagerModel) GetRole() *RoleModel {
	if x != nil {
		return x.Role
	}
	return nil
}

//角色对应的结构体
type RoleModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      int64  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	AddTime     int64  `protobuf:"varint,5,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *RoleModel) Reset() {
	*x = RoleModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleModel) ProtoMessage() {}

func (x *RoleModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleModel.ProtoReflect.Descriptor instead.
func (*RoleModel) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{1}
}

func (x *RoleModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoleModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RoleModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoleModel) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RoleModel) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

type ManagerGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"` // 用来查询管理员是否存在
}

func (x *ManagerGetRequest) Reset() {
	*x = ManagerGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerGetRequest) ProtoMessage() {}

func (x *ManagerGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerGetRequest.ProtoReflect.Descriptor instead.
func (*ManagerGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{2}
}

func (x *ManagerGetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManagerGetRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ManagerGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerList []*ManagerModel `protobuf:"bytes,1,rep,name=ManagerList,proto3" json:"ManagerList,omitempty"`
}

func (x *ManagerGetResponse) Reset() {
	*x = ManagerGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerGetResponse) ProtoMessage() {}

func (x *ManagerGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerGetResponse.ProtoReflect.Descriptor instead.
func (*ManagerGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{3}
}

func (x *ManagerGetResponse) GetManagerList() []*ManagerModel {
	if x != nil {
		return x.ManagerList
	}
	return nil
}

type ManagerAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Status   int64  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	RoleId   int64  `protobuf:"varint,6,opt,name=roleId,proto3" json:"roleId,omitempty"`
	AddTime  int64  `protobuf:"varint,7,opt,name=addTime,proto3" json:"addTime,omitempty"`
	IsSuper  int64  `protobuf:"varint,8,opt,name=isSuper,proto3" json:"isSuper,omitempty"`
}

func (x *ManagerAddRequest) Reset() {
	*x = ManagerAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerAddRequest) ProtoMessage() {}

func (x *ManagerAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerAddRequest.ProtoReflect.Descriptor instead.
func (*ManagerAddRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{4}
}

func (x *ManagerAddRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ManagerAddRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManagerAddRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ManagerAddRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerAddRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ManagerAddRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *ManagerAddRequest) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *ManagerAddRequest) GetIsSuper() int64 {
	if x != nil {
		return x.IsSuper
	}
	return 0
}

type ManagerAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ManagerAddResponse) Reset() {
	*x = ManagerAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerAddResponse) ProtoMessage() {}

func (x *ManagerAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerAddResponse.ProtoReflect.Descriptor instead.
func (*ManagerAddResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{5}
}

func (x *ManagerAddResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ManagerAddResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ManagerUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Mobile   string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Status   int64  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	RoleId   int64  `protobuf:"varint,7,opt,name=roleId,proto3" json:"roleId,omitempty"`
	AddTime  int64  `protobuf:"varint,8,opt,name=addTime,proto3" json:"addTime,omitempty"`
	IsSuper  int64  `protobuf:"varint,9,opt,name=isSuper,proto3" json:"isSuper,omitempty"`
}

func (x *ManagerUpdateRequest) Reset() {
	*x = ManagerUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerUpdateRequest) ProtoMessage() {}

func (x *ManagerUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerUpdateRequest.ProtoReflect.Descriptor instead.
func (*ManagerUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{6}
}

func (x *ManagerUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManagerUpdateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ManagerUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManagerUpdateRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ManagerUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerUpdateRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ManagerUpdateRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *ManagerUpdateRequest) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *ManagerUpdateRequest) GetIsSuper() int64 {
	if x != nil {
		return x.IsSuper
	}
	return 0
}

type ManagerUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ManagerUpdateResponse) Reset() {
	*x = ManagerUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerUpdateResponse) ProtoMessage() {}

func (x *ManagerUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerUpdateResponse.ProtoReflect.Descriptor instead.
func (*ManagerUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{7}
}

func (x *ManagerUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ManagerUpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ManagerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ManagerDeleteRequest) Reset() {
	*x = ManagerDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerDeleteRequest) ProtoMessage() {}

func (x *ManagerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerDeleteRequest.ProtoReflect.Descriptor instead.
func (*ManagerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{8}
}

func (x *ManagerDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ManagerDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ManagerDeleteResponse) Reset() {
	*x = ManagerDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacManager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerDeleteResponse) ProtoMessage() {}

func (x *ManagerDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacManager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerDeleteResponse.ProtoReflect.Descriptor instead.
func (*ManagerDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacManager_proto_rawDescGZIP(), []int{9}
}

func (x *ManagerDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ManagerDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_rbacManager_proto protoreflect.FileDescriptor

var file_proto_rbacManager_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x62, 0x61, 0x63, 0x22,
	0x8d, 0x02, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x85, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf0,
	0x01, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x22, 0x4b, 0x0a, 0x15, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26,
	0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xab, 0x02, 0x0a, 0x0b, 0x52, 0x62, 0x61, 0x63, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x12, 0x17, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x12, 0x17, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x62, 0x61,
	0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rbacManager_proto_rawDescOnce sync.Once
	file_proto_rbacManager_proto_rawDescData = file_proto_rbacManager_proto_rawDesc
)

func file_proto_rbacManager_proto_rawDescGZIP() []byte {
	file_proto_rbacManager_proto_rawDescOnce.Do(func() {
		file_proto_rbacManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rbacManager_proto_rawDescData)
	})
	return file_proto_rbacManager_proto_rawDescData
}

var file_proto_rbacManager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_rbacManager_proto_goTypes = []interface{}{
	(*ManagerModel)(nil),          // 0: rbac.ManagerModel
	(*RoleModel)(nil),             // 1: rbac.RoleModel
	(*ManagerGetRequest)(nil),     // 2: rbac.ManagerGetRequest
	(*ManagerGetResponse)(nil),    // 3: rbac.ManagerGetResponse
	(*ManagerAddRequest)(nil),     // 4: rbac.ManagerAddRequest
	(*ManagerAddResponse)(nil),    // 5: rbac.ManagerAddResponse
	(*ManagerUpdateRequest)(nil),  // 6: rbac.ManagerUpdateRequest
	(*ManagerUpdateResponse)(nil), // 7: rbac.ManagerUpdateResponse
	(*ManagerDeleteRequest)(nil),  // 8: rbac.ManagerDeleteRequest
	(*ManagerDeleteResponse)(nil), // 9: rbac.ManagerDeleteResponse
}
var file_proto_rbacManager_proto_depIdxs = []int32{
	1, // 0: rbac.ManagerModel.role:type_name -> rbac.RoleModel
	0, // 1: rbac.ManagerGetResponse.ManagerList:type_name -> rbac.ManagerModel
	2, // 2: rbac.RbacManager.ManagerGet:input_type -> rbac.ManagerGetRequest
	4, // 3: rbac.RbacManager.ManagerAdd:input_type -> rbac.ManagerAddRequest
	6, // 4: rbac.RbacManager.ManagerUpdate:input_type -> rbac.ManagerUpdateRequest
	8, // 5: rbac.RbacManager.ManagerDelete:input_type -> rbac.ManagerDeleteRequest
	3, // 6: rbac.RbacManager.ManagerGet:output_type -> rbac.ManagerGetResponse
	5, // 7: rbac.RbacManager.ManagerAdd:output_type -> rbac.ManagerAddResponse
	7, // 8: rbac.RbacManager.ManagerUpdate:output_type -> rbac.ManagerUpdateResponse
	9, // 9: rbac.RbacManager.ManagerDelete:output_type -> rbac.ManagerDeleteResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_rbacManager_proto_init() }
func file_proto_rbacManager_proto_init() {
	if File_proto_rbacManager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rbacManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerModel); i {
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
		file_proto_rbacManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleModel); i {
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
		file_proto_rbacManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerGetRequest); i {
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
		file_proto_rbacManager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerGetResponse); i {
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
		file_proto_rbacManager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerAddRequest); i {
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
		file_proto_rbacManager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerAddResponse); i {
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
		file_proto_rbacManager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerUpdateRequest); i {
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
		file_proto_rbacManager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerUpdateResponse); i {
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
		file_proto_rbacManager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerDeleteRequest); i {
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
		file_proto_rbacManager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerDeleteResponse); i {
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
			RawDescriptor: file_proto_rbacManager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rbacManager_proto_goTypes,
		DependencyIndexes: file_proto_rbacManager_proto_depIdxs,
		MessageInfos:      file_proto_rbacManager_proto_msgTypes,
	}.Build()
	File_proto_rbacManager_proto = out.File
	file_proto_rbacManager_proto_rawDesc = nil
	file_proto_rbacManager_proto_goTypes = nil
	file_proto_rbacManager_proto_depIdxs = nil
}
