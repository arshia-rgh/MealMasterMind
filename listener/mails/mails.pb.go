// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: mails.proto

package mails

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

type MailRequestResetLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Link    string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *MailRequestResetLink) Reset() {
	*x = MailRequestResetLink{}
	mi := &file_mails_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailRequestResetLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailRequestResetLink) ProtoMessage() {}

func (x *MailRequestResetLink) ProtoReflect() protoreflect.Message {
	mi := &file_mails_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailRequestResetLink.ProtoReflect.Descriptor instead.
func (*MailRequestResetLink) Descriptor() ([]byte, []int) {
	return file_mails_proto_rawDescGZIP(), []int{0}
}

func (x *MailRequestResetLink) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MailRequestResetLink) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *MailRequestResetLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type MailResponseResetLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MailResponseResetLink) Reset() {
	*x = MailResponseResetLink{}
	mi := &file_mails_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailResponseResetLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailResponseResetLink) ProtoMessage() {}

func (x *MailResponseResetLink) ProtoReflect() protoreflect.Message {
	mi := &file_mails_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailResponseResetLink.ProtoReflect.Descriptor instead.
func (*MailResponseResetLink) Descriptor() ([]byte, []int) {
	return file_mails_proto_rawDescGZIP(), []int{1}
}

func (x *MailResponseResetLink) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_mails_proto protoreflect.FileDescriptor

var file_mails_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x5a, 0x0a, 0x14, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x22, 0x2f, 0x0a, 0x15, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x51, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x45, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mails_proto_rawDescOnce sync.Once
	file_mails_proto_rawDescData = file_mails_proto_rawDesc
)

func file_mails_proto_rawDescGZIP() []byte {
	file_mails_proto_rawDescOnce.Do(func() {
		file_mails_proto_rawDescData = protoimpl.X.CompressGZIP(file_mails_proto_rawDescData)
	})
	return file_mails_proto_rawDescData
}

var file_mails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mails_proto_goTypes = []any{
	(*MailRequestResetLink)(nil),  // 0: mails.MailRequestResetLink
	(*MailResponseResetLink)(nil), // 1: mails.MailResponseResetLink
}
var file_mails_proto_depIdxs = []int32{
	0, // 0: mails.SendMail.SendMail:input_type -> mails.MailRequestResetLink
	1, // 1: mails.SendMail.SendMail:output_type -> mails.MailResponseResetLink
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mails_proto_init() }
func file_mails_proto_init() {
	if File_mails_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mails_proto_goTypes,
		DependencyIndexes: file_mails_proto_depIdxs,
		MessageInfos:      file_mails_proto_msgTypes,
	}.Build()
	File_mails_proto = out.File
	file_mails_proto_rawDesc = nil
	file_mails_proto_goTypes = nil
	file_mails_proto_depIdxs = nil
}
