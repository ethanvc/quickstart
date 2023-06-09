// docker run --rm -it -v $(pwd):/workdir protoc_image --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=. --go-grpc_out=. coordinator_migration_svr.proto
// tips: use --go_out=paths=source_relative:. to generate file without directory structure.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: coordinator_migration_svr.proto

package coordinator_migration_svr

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

// The request message containing the user's name.
type GetUserDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetUserDetailReq) Reset() {
	*x = GetUserDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_migration_svr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailReq) ProtoMessage() {}

func (x *GetUserDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_migration_svr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailReq.ProtoReflect.Descriptor instead.
func (*GetUserDetailReq) Descriptor() ([]byte, []int) {
	return file_coordinator_migration_svr_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserDetailReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// The response message containing the greetings
type GetUserDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserDetailResp) Reset() {
	*x = GetUserDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_migration_svr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailResp) ProtoMessage() {}

func (x *GetUserDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_migration_svr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailResp.ProtoReflect.Descriptor instead.
func (*GetUserDetailResp) Descriptor() ([]byte, []int) {
	return file_coordinator_migration_svr_proto_rawDescGZIP(), []int{1}
}

var File_coordinator_migration_svr_proto protoreflect.FileDescriptor

var file_coordinator_migration_svr_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x72, 0x22, 0x24, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x32, 0x87, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x76, 0x72, 0x12, 0x6c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x74, 0x68, 0x61, 0x6e, 0x76, 0x63, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x76, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coordinator_migration_svr_proto_rawDescOnce sync.Once
	file_coordinator_migration_svr_proto_rawDescData = file_coordinator_migration_svr_proto_rawDesc
)

func file_coordinator_migration_svr_proto_rawDescGZIP() []byte {
	file_coordinator_migration_svr_proto_rawDescOnce.Do(func() {
		file_coordinator_migration_svr_proto_rawDescData = protoimpl.X.CompressGZIP(file_coordinator_migration_svr_proto_rawDescData)
	})
	return file_coordinator_migration_svr_proto_rawDescData
}

var file_coordinator_migration_svr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coordinator_migration_svr_proto_goTypes = []interface{}{
	(*GetUserDetailReq)(nil),  // 0: coordinator_migration_svr.GetUserDetailReq
	(*GetUserDetailResp)(nil), // 1: coordinator_migration_svr.GetUserDetailResp
}
var file_coordinator_migration_svr_proto_depIdxs = []int32{
	0, // 0: coordinator_migration_svr.CoordinatorMigrationSvr.GetUserDetail:input_type -> coordinator_migration_svr.GetUserDetailReq
	1, // 1: coordinator_migration_svr.CoordinatorMigrationSvr.GetUserDetail:output_type -> coordinator_migration_svr.GetUserDetailResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coordinator_migration_svr_proto_init() }
func file_coordinator_migration_svr_proto_init() {
	if File_coordinator_migration_svr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coordinator_migration_svr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailReq); i {
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
		file_coordinator_migration_svr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailResp); i {
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
			RawDescriptor: file_coordinator_migration_svr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coordinator_migration_svr_proto_goTypes,
		DependencyIndexes: file_coordinator_migration_svr_proto_depIdxs,
		MessageInfos:      file_coordinator_migration_svr_proto_msgTypes,
	}.Build()
	File_coordinator_migration_svr_proto = out.File
	file_coordinator_migration_svr_proto_rawDesc = nil
	file_coordinator_migration_svr_proto_goTypes = nil
	file_coordinator_migration_svr_proto_depIdxs = nil
}
