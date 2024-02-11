// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.8
// source: backend/api/proto/admin.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Role      string `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdateAt  string `protobuf:"bytes,6,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	DeletedAt string `protobuf:"bytes,7,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	Verified  bool   `protobuf:"varint,8,opt,name=Verified,proto3" json:"Verified,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_proto_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_proto_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_backend_api_proto_admin_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *User) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *User) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string `protobuf:"bytes,1,opt,name=Error_message,json=ErrorMessage,proto3" json:"Error_message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_proto_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_proto_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_backend_api_proto_admin_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_backend_api_proto_admin_proto protoreflect.FileDescriptor

var file_backend_api_proto_admin_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22,
	0x34, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_api_proto_admin_proto_rawDescOnce sync.Once
	file_backend_api_proto_admin_proto_rawDescData = file_backend_api_proto_admin_proto_rawDesc
)

func file_backend_api_proto_admin_proto_rawDescGZIP() []byte {
	file_backend_api_proto_admin_proto_rawDescOnce.Do(func() {
		file_backend_api_proto_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_api_proto_admin_proto_rawDescData)
	})
	return file_backend_api_proto_admin_proto_rawDescData
}

var file_backend_api_proto_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backend_api_proto_admin_proto_goTypes = []interface{}{
	(*User)(nil),          // 0: admin.User
	(*ErrorResponse)(nil), // 1: admin.ErrorResponse
}
var file_backend_api_proto_admin_proto_depIdxs = []int32{
	0, // 0: admin.CreateTeacher.CreateUser:input_type -> admin.User
	1, // 1: admin.CreateTeacher.CreateUser:output_type -> admin.ErrorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_api_proto_admin_proto_init() }
func file_backend_api_proto_admin_proto_init() {
	if File_backend_api_proto_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_api_proto_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_backend_api_proto_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_backend_api_proto_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_api_proto_admin_proto_goTypes,
		DependencyIndexes: file_backend_api_proto_admin_proto_depIdxs,
		MessageInfos:      file_backend_api_proto_admin_proto_msgTypes,
	}.Build()
	File_backend_api_proto_admin_proto = out.File
	file_backend_api_proto_admin_proto_rawDesc = nil
	file_backend_api_proto_admin_proto_goTypes = nil
	file_backend_api_proto_admin_proto_depIdxs = nil
}