// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: server/proto/tuples.proto

package tuple_spaces

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

type Command int32

const (
	Command_GET      Command = 0
	Command_READ     Command = 1
	Command_WRITE    Command = 2
	Command_READ_ALL Command = 3
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "GET",
		1: "READ",
		2: "WRITE",
		3: "READ_ALL",
	}
	Command_value = map[string]int32{
		"GET":      0,
		"READ":     1,
		"WRITE":    2,
		"READ_ALL": 3,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_tuples_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_server_proto_tuples_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_tuples_proto_rawDescGZIP(), []int{0}
}

type Tuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Tuple) Reset() {
	*x = Tuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_tuples_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuple) ProtoMessage() {}

func (x *Tuple) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_tuples_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuple.ProtoReflect.Descriptor instead.
func (*Tuple) Descriptor() ([]byte, []int) {
	return file_server_proto_tuples_proto_rawDescGZIP(), []int{0}
}

func (x *Tuple) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type TupleSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tuples []*Tuple `protobuf:"bytes,2,rep,name=tuples,proto3" json:"tuples,omitempty"`
}

func (x *TupleSpace) Reset() {
	*x = TupleSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_tuples_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TupleSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TupleSpace) ProtoMessage() {}

func (x *TupleSpace) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_tuples_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TupleSpace.ProtoReflect.Descriptor instead.
func (*TupleSpace) Descriptor() ([]byte, []int) {
	return file_server_proto_tuples_proto_rawDescGZIP(), []int{1}
}

func (x *TupleSpace) GetTuples() []*Tuple {
	if x != nil {
		return x.Tuples
	}
	return nil
}

type RequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd   Command `protobuf:"varint,1,opt,name=cmd,proto3,enum=Command" json:"cmd,omitempty"`
	Tuple *Tuple  `protobuf:"bytes,2,opt,name=tuple,proto3" json:"tuple,omitempty"`
}

func (x *RequestData) Reset() {
	*x = RequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_tuples_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestData) ProtoMessage() {}

func (x *RequestData) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_tuples_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestData.ProtoReflect.Descriptor instead.
func (*RequestData) Descriptor() ([]byte, []int) {
	return file_server_proto_tuples_proto_rawDescGZIP(), []int{2}
}

func (x *RequestData) GetCmd() Command {
	if x != nil {
		return x.Cmd
	}
	return Command_GET
}

func (x *RequestData) GetTuple() *Tuple {
	if x != nil {
		return x.Tuple
	}
	return nil
}

var File_server_proto_tuples_proto protoreflect.FileDescriptor

var file_server_proto_tuples_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x05, 0x54,
	0x75, 0x70, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x0a,
	0x54, 0x75, 0x70, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x52, 0x06, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x2a, 0x35, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x41, 0x44, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f,
	0x74, 0x75, 0x70, 0x6c, 0x65, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_tuples_proto_rawDescOnce sync.Once
	file_server_proto_tuples_proto_rawDescData = file_server_proto_tuples_proto_rawDesc
)

func file_server_proto_tuples_proto_rawDescGZIP() []byte {
	file_server_proto_tuples_proto_rawDescOnce.Do(func() {
		file_server_proto_tuples_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_tuples_proto_rawDescData)
	})
	return file_server_proto_tuples_proto_rawDescData
}

var file_server_proto_tuples_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_proto_tuples_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_proto_tuples_proto_goTypes = []interface{}{
	(Command)(0),        // 0: Command
	(*Tuple)(nil),       // 1: Tuple
	(*TupleSpace)(nil),  // 2: TupleSpace
	(*RequestData)(nil), // 3: RequestData
}
var file_server_proto_tuples_proto_depIdxs = []int32{
	1, // 0: TupleSpace.tuples:type_name -> Tuple
	0, // 1: RequestData.cmd:type_name -> Command
	1, // 2: RequestData.tuple:type_name -> Tuple
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_proto_tuples_proto_init() }
func file_server_proto_tuples_proto_init() {
	if File_server_proto_tuples_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_tuples_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tuple); i {
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
		file_server_proto_tuples_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TupleSpace); i {
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
		file_server_proto_tuples_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestData); i {
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
			RawDescriptor: file_server_proto_tuples_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_proto_tuples_proto_goTypes,
		DependencyIndexes: file_server_proto_tuples_proto_depIdxs,
		EnumInfos:         file_server_proto_tuples_proto_enumTypes,
		MessageInfos:      file_server_proto_tuples_proto_msgTypes,
	}.Build()
	File_server_proto_tuples_proto = out.File
	file_server_proto_tuples_proto_rawDesc = nil
	file_server_proto_tuples_proto_goTypes = nil
	file_server_proto_tuples_proto_depIdxs = nil
}
