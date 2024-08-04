// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: user.proto

package __

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

type FetchAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchAll) Reset() {
	*x = FetchAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAll) ProtoMessage() {}

func (x *FetchAll) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAll.ProtoReflect.Descriptor instead.
func (*FetchAll) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Dob       string `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Address   string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Password  string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignupRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *SignupRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *SignupRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *SignupRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Dob       string `protobuf:"bytes,4,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender    string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Profile) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Profile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Profile) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Profile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userlist []*Profile `protobuf:"bytes,1,rep,name=userlist,proto3" json:"userlist,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *Users) GetUserlist() []*Profile {
	if x != nil {
		return x.Userlist
	}
	return nil
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *CommonResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CommonResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *Email) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserID) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UNoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UNoParam) Reset() {
	*x = UNoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UNoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UNoParam) ProtoMessage() {}

func (x *UNoParam) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UNoParam.ProtoReflect.Descriptor instead.
func (*UNoParam) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x22, 0x0a, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c,
	0x22, 0xd5, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f,
	0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x22, 0x34, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x1d, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0a, 0x0a, 0x08, 0x55, 0x4e, 0x6f,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x32, 0x93, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x4e, 0x6f, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x34, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a,
	0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(*FetchAll)(nil),       // 0: userpb.FetchAll
	(*SignupRequest)(nil),  // 1: userpb.SignupRequest
	(*Profile)(nil),        // 2: userpb.Profile
	(*Users)(nil),          // 3: userpb.Users
	(*CommonResponse)(nil), // 4: userpb.CommonResponse
	(*Email)(nil),          // 5: userpb.Email
	(*UserID)(nil),         // 6: userpb.UserID
	(*UNoParam)(nil),       // 7: userpb.UNoParam
}
var file_user_proto_depIdxs = []int32{
	2, // 0: userpb.Users.userlist:type_name -> userpb.Profile
	1, // 1: userpb.UserService.EditUser:input_type -> userpb.SignupRequest
	7, // 2: userpb.UserService.FindAllUsers:input_type -> userpb.UNoParam
	6, // 3: userpb.UserService.DeleteUser:input_type -> userpb.UserID
	6, // 4: userpb.UserService.FindUserByID:input_type -> userpb.UserID
	5, // 5: userpb.UserService.FindUserByEmail:input_type -> userpb.Email
	4, // 6: userpb.UserService.EditUser:output_type -> userpb.CommonResponse
	3, // 7: userpb.UserService.FindAllUsers:output_type -> userpb.Users
	4, // 8: userpb.UserService.DeleteUser:output_type -> userpb.CommonResponse
	2, // 9: userpb.UserService.FindUserByID:output_type -> userpb.Profile
	2, // 10: userpb.UserService.FindUserByEmail:output_type -> userpb.Profile
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAll); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UNoParam); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
