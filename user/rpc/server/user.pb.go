// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: user.proto

package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Register 相关消息体
type RegisterRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Username         string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Mobile           string                 `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Avatar           *string                `protobuf:"bytes,3,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
	Password         string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	VerificationCode string                 `protobuf:"bytes,5,opt,name=verificationCode,proto3" json:"verificationCode,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterRequest) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Token         string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// FindById 相关消息体
type FindByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByIdRequest) Reset() {
	*x = FindByIdRequest{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdRequest) ProtoMessage() {}

func (x *FindByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdRequest.ProtoReflect.Descriptor instead.
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *FindByIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserInfo      *UserInfoResponse      `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByIdResponse) Reset() {
	*x = FindByIdResponse{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdResponse) ProtoMessage() {}

func (x *FindByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdResponse.ProtoReflect.Descriptor instead.
func (*FindByIdResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *FindByIdResponse) GetUserInfo() *UserInfoResponse {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

// FindByMobile 相关消息体
type FindByMobileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mobile        string                 `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByMobileRequest) Reset() {
	*x = FindByMobileRequest{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByMobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByMobileRequest) ProtoMessage() {}

func (x *FindByMobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByMobileRequest.ProtoReflect.Descriptor instead.
func (*FindByMobileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *FindByMobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type FindByMobileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserInfo      *UserInfoResponse      `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByMobileResponse) Reset() {
	*x = FindByMobileResponse{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByMobileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByMobileResponse) ProtoMessage() {}

func (x *FindByMobileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByMobileResponse.ProtoReflect.Descriptor instead.
func (*FindByMobileResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *FindByMobileResponse) GetUserInfo() *UserInfoResponse {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

// SendSms 相关消息体
type SendSmsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mobile        string                 `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`   // 手机号
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 短信内容
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendSmsRequest) Reset() {
	*x = SendSmsRequest{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsRequest) ProtoMessage() {}

func (x *SendSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsRequest.ProtoReflect.Descriptor instead.
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *SendSmsRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SendSmsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendSmsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendSmsResponse) Reset() {
	*x = SendSmsResponse{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendSmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsResponse) ProtoMessage() {}

func (x *SendSmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsResponse.ProtoReflect.Descriptor instead.
func (*SendSmsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

type MobileLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mobile        string                 `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Verification  string                 `protobuf:"bytes,3,opt,name=verification,proto3" json:"verification,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MobileLoginRequest) Reset() {
	*x = MobileLoginRequest{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MobileLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileLoginRequest) ProtoMessage() {}

func (x *MobileLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileLoginRequest.ProtoReflect.Descriptor instead.
func (*MobileLoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *MobileLoginRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *MobileLoginRequest) GetVerification() string {
	if x != nil {
		return x.Verification
	}
	return ""
}

type UsernameLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsernameLoginRequest) Reset() {
	*x = UsernameLoginRequest{}
	mi := &file_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsernameLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameLoginRequest) ProtoMessage() {}

func (x *UsernameLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameLoginRequest.ProtoReflect.Descriptor instead.
func (*UsernameLoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *UsernameLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsernameLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mobile        string                 `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	mi := &file_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *UserInfoResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfoResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UserInfoUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Mobile        string                 `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoUpdateRequest) Reset() {
	*x = UserInfoUpdateRequest{}
	mi := &file_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoUpdateRequest) ProtoMessage() {}

func (x *UserInfoUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserInfoUpdateRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *UserInfoUpdateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfoUpdateRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfoUpdateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type IsUsernameExistRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsUsernameExistRequest) Reset() {
	*x = IsUsernameExistRequest{}
	mi := &file_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsUsernameExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUsernameExistRequest) ProtoMessage() {}

func (x *IsUsernameExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUsernameExistRequest.ProtoReflect.Descriptor instead.
func (*IsUsernameExistRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *IsUsernameExistRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type IsUsernameExistResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exist         bool                   `protobuf:"varint,1,opt,name=Exist,proto3" json:"Exist,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsUsernameExistResponse) Reset() {
	*x = IsUsernameExistResponse{}
	mi := &file_user_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsUsernameExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUsernameExistResponse) ProtoMessage() {}

func (x *IsUsernameExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUsernameExistResponse.ProtoReflect.Descriptor instead.
func (*IsUsernameExistResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{14}
}

func (x *IsUsernameExistResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\x06server\"\xb5\x01\n" +
	"\x0fRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x16\n" +
	"\x06mobile\x18\x02 \x01(\tR\x06mobile\x12\x1b\n" +
	"\x06avatar\x18\x03 \x01(\tH\x00R\x06avatar\x88\x01\x01\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\x12*\n" +
	"\x10verificationCode\x18\x05 \x01(\tR\x10verificationCodeB\t\n" +
	"\a_avatar\"@\n" +
	"\x10RegisterResponse\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\x03R\x06userId\x12\x14\n" +
	"\x05token\x18\x03 \x01(\tR\x05token\")\n" +
	"\x0fFindByIdRequest\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\x03R\x06userId\"H\n" +
	"\x10FindByIdResponse\x124\n" +
	"\buserInfo\x18\x01 \x01(\v2\x18.server.UserInfoResponseR\buserInfo\"-\n" +
	"\x13FindByMobileRequest\x12\x16\n" +
	"\x06mobile\x18\x01 \x01(\tR\x06mobile\"L\n" +
	"\x14FindByMobileResponse\x124\n" +
	"\buserInfo\x18\x01 \x01(\v2\x18.server.UserInfoResponseR\buserInfo\"B\n" +
	"\x0eSendSmsRequest\x12\x16\n" +
	"\x06mobile\x18\x01 \x01(\tR\x06mobile\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\"\x11\n" +
	"\x0fSendSmsResponse\"P\n" +
	"\x12MobileLoginRequest\x12\x16\n" +
	"\x06mobile\x18\x02 \x01(\tR\x06mobile\x12\"\n" +
	"\fverification\x18\x03 \x01(\tR\fverification\"N\n" +
	"\x14UsernameLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"=\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\x03R\x06userId\"v\n" +
	"\x10UserInfoResponse\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\x03R\x06userId\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x16\n" +
	"\x06mobile\x18\x03 \x01(\tR\x06mobile\x12\x16\n" +
	"\x06avatar\x18\x04 \x01(\tR\x06avatar\"\x7f\n" +
	"\x15UserInfoUpdateRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x16\n" +
	"\x06mobile\x18\x02 \x01(\tR\x06mobile\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x16\n" +
	"\x06avatar\x18\x04 \x01(\tR\x06avatar\"4\n" +
	"\x16IsUsernameExistRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\"/\n" +
	"\x17IsUsernameExistResponse\x12\x14\n" +
	"\x05Exist\x18\x01 \x01(\bR\x05Exist2\xb2\x04\n" +
	"\x04User\x12=\n" +
	"\bRegister\x12\x17.server.RegisterRequest\x1a\x18.server.RegisterResponse\x12=\n" +
	"\bFindById\x12\x17.server.FindByIdRequest\x1a\x18.server.FindByIdResponse\x12I\n" +
	"\fFindByMobile\x12\x1b.server.FindByMobileRequest\x1a\x1c.server.FindByMobileResponse\x12:\n" +
	"\aSendSms\x12\x16.server.SendSmsRequest\x1a\x17.server.SendSmsResponse\x12@\n" +
	"\vMobileLogin\x12\x1a.server.MobileLoginRequest\x1a\x15.server.LoginResponse\x12D\n" +
	"\rUsernameLogin\x12\x1c.server.UsernameLoginRequest\x1a\x15.server.LoginResponse\x12I\n" +
	"\x0eUserInfoUpdate\x12\x1d.server.UserInfoUpdateRequest\x1a\x18.server.UserInfoResponse\x12R\n" +
	"\x0fIsUsernameExist\x12\x1e.server.IsUsernameExistRequest\x1a\x1f.server.IsUsernameExistResponseB\n" +
	"Z\b./serverb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_user_proto_goTypes = []any{
	(*RegisterRequest)(nil),         // 0: server.RegisterRequest
	(*RegisterResponse)(nil),        // 1: server.RegisterResponse
	(*FindByIdRequest)(nil),         // 2: server.FindByIdRequest
	(*FindByIdResponse)(nil),        // 3: server.FindByIdResponse
	(*FindByMobileRequest)(nil),     // 4: server.FindByMobileRequest
	(*FindByMobileResponse)(nil),    // 5: server.FindByMobileResponse
	(*SendSmsRequest)(nil),          // 6: server.SendSmsRequest
	(*SendSmsResponse)(nil),         // 7: server.SendSmsResponse
	(*MobileLoginRequest)(nil),      // 8: server.MobileLoginRequest
	(*UsernameLoginRequest)(nil),    // 9: server.UsernameLoginRequest
	(*LoginResponse)(nil),           // 10: server.LoginResponse
	(*UserInfoResponse)(nil),        // 11: server.UserInfoResponse
	(*UserInfoUpdateRequest)(nil),   // 12: server.UserInfoUpdateRequest
	(*IsUsernameExistRequest)(nil),  // 13: server.IsUsernameExistRequest
	(*IsUsernameExistResponse)(nil), // 14: server.IsUsernameExistResponse
}
var file_user_proto_depIdxs = []int32{
	11, // 0: server.FindByIdResponse.userInfo:type_name -> server.UserInfoResponse
	11, // 1: server.FindByMobileResponse.userInfo:type_name -> server.UserInfoResponse
	0,  // 2: server.User.Register:input_type -> server.RegisterRequest
	2,  // 3: server.User.FindById:input_type -> server.FindByIdRequest
	4,  // 4: server.User.FindByMobile:input_type -> server.FindByMobileRequest
	6,  // 5: server.User.SendSms:input_type -> server.SendSmsRequest
	8,  // 6: server.User.MobileLogin:input_type -> server.MobileLoginRequest
	9,  // 7: server.User.UsernameLogin:input_type -> server.UsernameLoginRequest
	12, // 8: server.User.UserInfoUpdate:input_type -> server.UserInfoUpdateRequest
	13, // 9: server.User.IsUsernameExist:input_type -> server.IsUsernameExistRequest
	1,  // 10: server.User.Register:output_type -> server.RegisterResponse
	3,  // 11: server.User.FindById:output_type -> server.FindByIdResponse
	5,  // 12: server.User.FindByMobile:output_type -> server.FindByMobileResponse
	7,  // 13: server.User.SendSms:output_type -> server.SendSmsResponse
	10, // 14: server.User.MobileLogin:output_type -> server.LoginResponse
	10, // 15: server.User.UsernameLogin:output_type -> server.LoginResponse
	11, // 16: server.User.UserInfoUpdate:output_type -> server.UserInfoResponse
	14, // 17: server.User.IsUsernameExist:output_type -> server.IsUsernameExistResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
