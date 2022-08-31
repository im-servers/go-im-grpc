// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: im_server.proto

package im_server

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

type SnedMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户的deviceToken,用于确定是哪条连接
	DeviceTokens []string `protobuf:"bytes,1,rep,name=deviceTokens,proto3" json:"deviceTokens,omitempty"`
	// 消息体
	Msg []byte `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SnedMsgReq) Reset() {
	*x = SnedMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnedMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnedMsgReq) ProtoMessage() {}

func (x *SnedMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnedMsgReq.ProtoReflect.Descriptor instead.
func (*SnedMsgReq) Descriptor() ([]byte, []int) {
	return file_im_server_proto_rawDescGZIP(), []int{0}
}

func (x *SnedMsgReq) GetDeviceTokens() []string {
	if x != nil {
		return x.DeviceTokens
	}
	return nil
}

func (x *SnedMsgReq) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_im_server_proto_rawDescGZIP(), []int{1}
}

var File_im_server_proto protoreflect.FileDescriptor

var file_im_server_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x69, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x0a,
	0x53, 0x6e, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4d, 0x73, 0x67,
	0x22, 0x0b, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x32, 0x42, 0x0a,
	0x08, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x6e, 0x65,
	0x64, 0x4d, 0x73, 0x67, 0x12, 0x15, 0x2e, 0x69, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x53, 0x6e, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x69, 0x6d,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x69, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_server_proto_rawDescOnce sync.Once
	file_im_server_proto_rawDescData = file_im_server_proto_rawDesc
)

func file_im_server_proto_rawDescGZIP() []byte {
	file_im_server_proto_rawDescOnce.Do(func() {
		file_im_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_server_proto_rawDescData)
	})
	return file_im_server_proto_rawDescData
}

var file_im_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_im_server_proto_goTypes = []interface{}{
	(*SnedMsgReq)(nil), // 0: im_server.SnedMsgReq
	(*EmptyResp)(nil),  // 1: im_server.EmptyResp
}
var file_im_server_proto_depIdxs = []int32{
	0, // 0: im_server.IMServer.snedMsg:input_type -> im_server.SnedMsgReq
	1, // 1: im_server.IMServer.snedMsg:output_type -> im_server.EmptyResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_server_proto_init() }
func file_im_server_proto_init() {
	if File_im_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnedMsgReq); i {
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
		file_im_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResp); i {
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
			RawDescriptor: file_im_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_server_proto_goTypes,
		DependencyIndexes: file_im_server_proto_depIdxs,
		MessageInfos:      file_im_server_proto_msgTypes,
	}.Build()
	File_im_server_proto = out.File
	file_im_server_proto_rawDesc = nil
	file_im_server_proto_goTypes = nil
	file_im_server_proto_depIdxs = nil
}
