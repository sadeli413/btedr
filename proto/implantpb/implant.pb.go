// Protobuf service communicating with implants
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative implant.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: implantpb/implant.proto

package implantpb

import (
	taskpb "github.com/devilsec/btedr/proto/taskpb"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implantpb_implant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_implantpb_implant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_implantpb_implant_proto_rawDescGZIP(), []int{0}
}

// A task request
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Implantid string `protobuf:"bytes,1,opt,name=implantid,proto3" json:"implantid,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implantpb_implant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_implantpb_implant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_implantpb_implant_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetImplantid() string {
	if x != nil {
		return x.Implantid
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taskid    string      `protobuf:"bytes,1,opt,name=taskid,proto3" json:"taskid,omitempty"`
	Implantid string      `protobuf:"bytes,2,opt,name=implantid,proto3" json:"implantid,omitempty"`
	Type      taskpb.Type `protobuf:"varint,3,opt,name=type,proto3,enum=taskpb.Type" json:"type,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implantpb_implant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_implantpb_implant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_implantpb_implant_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetTaskid() string {
	if x != nil {
		return x.Taskid
	}
	return ""
}

func (x *Result) GetImplantid() string {
	if x != nil {
		return x.Implantid
	}
	return ""
}

func (x *Result) GetType() taskpb.Type {
	if x != nil {
		return x.Type
	}
	return taskpb.Type(0)
}

var File_implantpb_implant_proto protoreflect.FileDescriptor

var file_implantpb_implant_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69, 0x6d, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x74, 0x65, 0x64, 0x72,
	0x70, 0x62, 0x1a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x66, 0x0a, 0x0a, 0x49, 0x6d, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x52, 0x50, 0x43, 0x12, 0x29, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x10, 0x2e, 0x62, 0x74, 0x65, 0x64, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x2d, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0f, 0x2e, 0x62, 0x74, 0x65, 0x64, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x1a, 0x0e, 0x2e, 0x62, 0x74, 0x65, 0x64, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x6c, 0x73, 0x65, 0x63, 0x2f, 0x62, 0x74, 0x65, 0x64, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_implantpb_implant_proto_rawDescOnce sync.Once
	file_implantpb_implant_proto_rawDescData = file_implantpb_implant_proto_rawDesc
)

func file_implantpb_implant_proto_rawDescGZIP() []byte {
	file_implantpb_implant_proto_rawDescOnce.Do(func() {
		file_implantpb_implant_proto_rawDescData = protoimpl.X.CompressGZIP(file_implantpb_implant_proto_rawDescData)
	})
	return file_implantpb_implant_proto_rawDescData
}

var file_implantpb_implant_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_implantpb_implant_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: btedrpb.Empty
	(*Request)(nil),     // 1: btedrpb.Request
	(*Result)(nil),      // 2: btedrpb.Result
	(taskpb.Type)(0),    // 3: taskpb.Type
	(*taskpb.Task)(nil), // 4: taskpb.Task
}
var file_implantpb_implant_proto_depIdxs = []int32{
	3, // 0: btedrpb.Result.type:type_name -> taskpb.Type
	1, // 1: btedrpb.ImplantRPC.GetTask:input_type -> btedrpb.Request
	2, // 2: btedrpb.ImplantRPC.TaskResult:input_type -> btedrpb.Result
	4, // 3: btedrpb.ImplantRPC.GetTask:output_type -> taskpb.Task
	0, // 4: btedrpb.ImplantRPC.TaskResult:output_type -> btedrpb.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_implantpb_implant_proto_init() }
func file_implantpb_implant_proto_init() {
	if File_implantpb_implant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_implantpb_implant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_implantpb_implant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_implantpb_implant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_implantpb_implant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_implantpb_implant_proto_goTypes,
		DependencyIndexes: file_implantpb_implant_proto_depIdxs,
		MessageInfos:      file_implantpb_implant_proto_msgTypes,
	}.Build()
	File_implantpb_implant_proto = out.File
	file_implantpb_implant_proto_rawDesc = nil
	file_implantpb_implant_proto_goTypes = nil
	file_implantpb_implant_proto_depIdxs = nil
}
