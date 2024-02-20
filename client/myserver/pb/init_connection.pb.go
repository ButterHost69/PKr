// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: init_connection.proto

package pb

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

type OTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	IpAddress string `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Otp       int32  `protobuf:"varint,3,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *OTP) Reset() {
	*x = OTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTP) ProtoMessage() {}

func (x *OTP) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTP.ProtoReflect.Descriptor instead.
func (*OTP) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{0}
}

func (x *OTP) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *OTP) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *OTP) GetOtp() int32 {
	if x != nil {
		return x.Otp
	}
	return 0
}

type OTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IfOtpCorrect   bool   `protobuf:"varint,1,opt,name=if_otp_correct,json=ifOtpCorrect,proto3" json:"if_otp_correct,omitempty"`
	ConnectionSlug string `protobuf:"bytes,2,opt,name=connection_slug,json=connectionSlug,proto3" json:"connection_slug,omitempty"`
	PublicKey      string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *OTPResponse) Reset() {
	*x = OTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPResponse) ProtoMessage() {}

func (x *OTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPResponse.ProtoReflect.Descriptor instead.
func (*OTPResponse) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{1}
}

func (x *OTPResponse) GetIfOtpCorrect() bool {
	if x != nil {
		return x.IfOtpCorrect
	}
	return false
}

func (x *OTPResponse) GetConnectionSlug() string {
	if x != nil {
		return x.ConnectionSlug
	}
	return ""
}

func (x *OTPResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_init_connection_proto protoreflect.FileDescriptor

var file_init_connection_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x7b, 0x0a, 0x0b,
	0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x66, 0x5f, 0x6f, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x66, 0x4f, 0x74, 0x70, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x32, 0x51, 0x0a, 0x0e, 0x49, 0x6e, 0x69,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x09, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x54, 0x50, 0x1a, 0x1c,
	0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_init_connection_proto_rawDescOnce sync.Once
	file_init_connection_proto_rawDescData = file_init_connection_proto_rawDesc
)

func file_init_connection_proto_rawDescGZIP() []byte {
	file_init_connection_proto_rawDescOnce.Do(func() {
		file_init_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_init_connection_proto_rawDescData)
	})
	return file_init_connection_proto_rawDescData
}

var file_init_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_init_connection_proto_goTypes = []interface{}{
	(*OTP)(nil),         // 0: init_connection.OTP
	(*OTPResponse)(nil), // 1: init_connection.OTPResponse
}
var file_init_connection_proto_depIdxs = []int32{
	0, // 0: init_connection.InitConnection.VerifyOTP:input_type -> init_connection.OTP
	1, // 1: init_connection.InitConnection.VerifyOTP:output_type -> init_connection.OTPResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_init_connection_proto_init() }
func file_init_connection_proto_init() {
	if File_init_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_init_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTP); i {
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
		file_init_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPResponse); i {
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
			RawDescriptor: file_init_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_init_connection_proto_goTypes,
		DependencyIndexes: file_init_connection_proto_depIdxs,
		MessageInfos:      file_init_connection_proto_msgTypes,
	}.Build()
	File_init_connection_proto = out.File
	file_init_connection_proto_rawDesc = nil
	file_init_connection_proto_goTypes = nil
	file_init_connection_proto_depIdxs = nil
}
