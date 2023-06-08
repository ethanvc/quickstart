// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: mytest.proto

package main

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

type DataV0Proto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Class string `protobuf:"bytes,3,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *DataV0Proto) Reset() {
	*x = DataV0Proto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mytest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataV0Proto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataV0Proto) ProtoMessage() {}

func (x *DataV0Proto) ProtoReflect() protoreflect.Message {
	mi := &file_mytest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataV0Proto.ProtoReflect.Descriptor instead.
func (*DataV0Proto) Descriptor() ([]byte, []int) {
	return file_mytest_proto_rawDescGZIP(), []int{0}
}

func (x *DataV0Proto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataV0Proto) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *DataV0Proto) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

var File_mytest_proto protoreflect.FileDescriptor

var file_mytest_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x79, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x49, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x56, 0x30, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x42,
	0x07, 0x5a, 0x05, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mytest_proto_rawDescOnce sync.Once
	file_mytest_proto_rawDescData = file_mytest_proto_rawDesc
)

func file_mytest_proto_rawDescGZIP() []byte {
	file_mytest_proto_rawDescOnce.Do(func() {
		file_mytest_proto_rawDescData = protoimpl.X.CompressGZIP(file_mytest_proto_rawDescData)
	})
	return file_mytest_proto_rawDescData
}

var file_mytest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mytest_proto_goTypes = []interface{}{
	(*DataV0Proto)(nil), // 0: main.DataV0Proto
}
var file_mytest_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mytest_proto_init() }
func file_mytest_proto_init() {
	if File_mytest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mytest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataV0Proto); i {
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
			RawDescriptor: file_mytest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mytest_proto_goTypes,
		DependencyIndexes: file_mytest_proto_depIdxs,
		MessageInfos:      file_mytest_proto_msgTypes,
	}.Build()
	File_mytest_proto = out.File
	file_mytest_proto_rawDesc = nil
	file_mytest_proto_goTypes = nil
	file_mytest_proto_depIdxs = nil
}