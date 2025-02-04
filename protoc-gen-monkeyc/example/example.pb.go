// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: example.proto

package example

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

type GlobalEnum int32

const (
	GlobalEnum_A GlobalEnum = 0
	GlobalEnum_B GlobalEnum = 1
)

// Enum value maps for GlobalEnum.
var (
	GlobalEnum_name = map[int32]string{
		0: "A",
		1: "B",
	}
	GlobalEnum_value = map[string]int32{
		"A": 0,
		"B": 1,
	}
)

func (x GlobalEnum) Enum() *GlobalEnum {
	p := new(GlobalEnum)
	*p = x
	return p
}

func (x GlobalEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GlobalEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_enumTypes[0].Descriptor()
}

func (GlobalEnum) Type() protoreflect.EnumType {
	return &file_example_proto_enumTypes[0]
}

func (x GlobalEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GlobalEnum.Descriptor instead.
func (GlobalEnum) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

type ExampleMessage_LocalEnum int32

const (
	ExampleMessage_LA ExampleMessage_LocalEnum = 0
	ExampleMessage_LB ExampleMessage_LocalEnum = 1
)

// Enum value maps for ExampleMessage_LocalEnum.
var (
	ExampleMessage_LocalEnum_name = map[int32]string{
		0: "LA",
		1: "LB",
	}
	ExampleMessage_LocalEnum_value = map[string]int32{
		"LA": 0,
		"LB": 1,
	}
)

func (x ExampleMessage_LocalEnum) Enum() *ExampleMessage_LocalEnum {
	p := new(ExampleMessage_LocalEnum)
	*p = x
	return p
}

func (x ExampleMessage_LocalEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleMessage_LocalEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_enumTypes[1].Descriptor()
}

func (ExampleMessage_LocalEnum) Type() protoreflect.EnumType {
	return &file_example_proto_enumTypes[1]
}

func (x ExampleMessage_LocalEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleMessage_LocalEnum.Descriptor instead.
func (ExampleMessage_LocalEnum) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I32   int32                        `protobuf:"varint,1,opt,name=i32,proto3" json:"i32,omitempty"`
	I64   int64                        `protobuf:"varint,2,opt,name=i64,proto3" json:"i64,omitempty"`
	U32   uint32                       `protobuf:"varint,3,opt,name=u32,proto3" json:"u32,omitempty"`
	U64   uint64                       `protobuf:"varint,4,opt,name=u64,proto3" json:"u64,omitempty"`
	S32   int32                        `protobuf:"zigzag32,5,opt,name=s32,proto3" json:"s32,omitempty"`
	S64   int64                        `protobuf:"zigzag64,6,opt,name=s64,proto3" json:"s64,omitempty"`
	F32   uint32                       `protobuf:"fixed32,7,opt,name=f32,proto3" json:"f32,omitempty"`
	F64   uint64                       `protobuf:"fixed64,8,opt,name=f64,proto3" json:"f64,omitempty"`
	Sf32  int32                        `protobuf:"fixed32,9,opt,name=sf32,proto3" json:"sf32,omitempty"`
	Sf64  int64                        `protobuf:"fixed64,10,opt,name=sf64,proto3" json:"sf64,omitempty"`
	Fl    float32                      `protobuf:"fixed32,11,opt,name=fl,proto3" json:"fl,omitempty"` // double dl = 12; // double is not supported
	Str   string                       `protobuf:"bytes,13,opt,name=str,proto3" json:"str,omitempty"`
	Byt   []byte                       `protobuf:"bytes,14,opt,name=byt,proto3" json:"byt,omitempty"`
	B     bool                         `protobuf:"varint,15,opt,name=b,proto3" json:"b,omitempty"`
	Ge    GlobalEnum                   `protobuf:"varint,16,opt,name=ge,proto3,enum=example.GlobalEnum" json:"ge,omitempty"`
	Le    ExampleMessage_LocalEnum     `protobuf:"varint,17,opt,name=le,proto3,enum=example.ExampleMessage_LocalEnum" json:"le,omitempty"`
	Gm    *GlobalMessage               `protobuf:"bytes,18,opt,name=gm,proto3" json:"gm,omitempty"`
	Lm    *ExampleMessage_LocalMessage `protobuf:"bytes,19,opt,name=lm,proto3" json:"lm,omitempty"`
	Ri64  []int64                      `protobuf:"varint,20,rep,name=ri64,proto3" json:"ri64,omitempty"`
	Rf32  []int32                      `protobuf:"fixed32,21,rep,name=rf32,proto3" json:"rf32,omitempty"`
	Rf64  []int64                      `protobuf:"fixed64,22,rep,name=rf64,proto3" json:"rf64,omitempty"`
	Rstr  []string                     `protobuf:"bytes,23,rep,name=rstr,proto3" json:"rstr,omitempty"`
	Rgm   []*GlobalMessage             `protobuf:"bytes,24,rep,name=rgm,proto3" json:"rgm,omitempty"`
	Rpi64 []int64                      `protobuf:"varint,25,rep,packed,name=rpi64,proto3" json:"rpi64,omitempty"`
	Rpf32 []int32                      `protobuf:"fixed32,26,rep,packed,name=rpf32,proto3" json:"rpf32,omitempty"`
	Rpf64 []int64                      `protobuf:"fixed64,27,rep,packed,name=rpf64,proto3" json:"rpf64,omitempty"`
}

func (x *ExampleMessage) Reset() {
	*x = ExampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage) ProtoMessage() {}

func (x *ExampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage.ProtoReflect.Descriptor instead.
func (*ExampleMessage) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleMessage) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

func (x *ExampleMessage) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

func (x *ExampleMessage) GetU32() uint32 {
	if x != nil {
		return x.U32
	}
	return 0
}

func (x *ExampleMessage) GetU64() uint64 {
	if x != nil {
		return x.U64
	}
	return 0
}

func (x *ExampleMessage) GetS32() int32 {
	if x != nil {
		return x.S32
	}
	return 0
}

func (x *ExampleMessage) GetS64() int64 {
	if x != nil {
		return x.S64
	}
	return 0
}

func (x *ExampleMessage) GetF32() uint32 {
	if x != nil {
		return x.F32
	}
	return 0
}

func (x *ExampleMessage) GetF64() uint64 {
	if x != nil {
		return x.F64
	}
	return 0
}

func (x *ExampleMessage) GetSf32() int32 {
	if x != nil {
		return x.Sf32
	}
	return 0
}

func (x *ExampleMessage) GetSf64() int64 {
	if x != nil {
		return x.Sf64
	}
	return 0
}

func (x *ExampleMessage) GetFl() float32 {
	if x != nil {
		return x.Fl
	}
	return 0
}

func (x *ExampleMessage) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *ExampleMessage) GetByt() []byte {
	if x != nil {
		return x.Byt
	}
	return nil
}

func (x *ExampleMessage) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

func (x *ExampleMessage) GetGe() GlobalEnum {
	if x != nil {
		return x.Ge
	}
	return GlobalEnum_A
}

func (x *ExampleMessage) GetLe() ExampleMessage_LocalEnum {
	if x != nil {
		return x.Le
	}
	return ExampleMessage_LA
}

func (x *ExampleMessage) GetGm() *GlobalMessage {
	if x != nil {
		return x.Gm
	}
	return nil
}

func (x *ExampleMessage) GetLm() *ExampleMessage_LocalMessage {
	if x != nil {
		return x.Lm
	}
	return nil
}

func (x *ExampleMessage) GetRi64() []int64 {
	if x != nil {
		return x.Ri64
	}
	return nil
}

func (x *ExampleMessage) GetRf32() []int32 {
	if x != nil {
		return x.Rf32
	}
	return nil
}

func (x *ExampleMessage) GetRf64() []int64 {
	if x != nil {
		return x.Rf64
	}
	return nil
}

func (x *ExampleMessage) GetRstr() []string {
	if x != nil {
		return x.Rstr
	}
	return nil
}

func (x *ExampleMessage) GetRgm() []*GlobalMessage {
	if x != nil {
		return x.Rgm
	}
	return nil
}

func (x *ExampleMessage) GetRpi64() []int64 {
	if x != nil {
		return x.Rpi64
	}
	return nil
}

func (x *ExampleMessage) GetRpf32() []int32 {
	if x != nil {
		return x.Rpf32
	}
	return nil
}

func (x *ExampleMessage) GetRpf64() []int64 {
	if x != nil {
		return x.Rpf64
	}
	return nil
}

type GlobalMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	G1 int32 `protobuf:"varint,1,opt,name=G1,proto3" json:"G1,omitempty"`
}

func (x *GlobalMessage) Reset() {
	*x = GlobalMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalMessage) ProtoMessage() {}

func (x *GlobalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalMessage.ProtoReflect.Descriptor instead.
func (*GlobalMessage) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalMessage) GetG1() int32 {
	if x != nil {
		return x.G1
	}
	return 0
}

type ExampleMessage_LocalMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L1 string `protobuf:"bytes,1,opt,name=L1,proto3" json:"L1,omitempty"`
}

func (x *ExampleMessage_LocalMessage) Reset() {
	*x = ExampleMessage_LocalMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage_LocalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage_LocalMessage) ProtoMessage() {}

func (x *ExampleMessage_LocalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage_LocalMessage.ProtoReflect.Descriptor instead.
func (*ExampleMessage_LocalMessage) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExampleMessage_LocalMessage) GetL1() string {
	if x != nil {
		return x.L1
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xc5, 0x05, 0x0a, 0x0e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x36, 0x34, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x33,
	0x32, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x03, 0x73, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x03, 0x73, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x33, 0x32, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x03, 0x66, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x36, 0x34,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x66, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x66, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x04, 0x73, 0x66, 0x33, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x66, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x04, 0x73,
	0x66, 0x36, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x02, 0x66, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x79, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x62, 0x79, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x01, 0x62, 0x12, 0x23, 0x0a, 0x02, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x02, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x02, 0x6c, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x02, 0x6c, 0x65, 0x12, 0x26, 0x0a,
	0x02, 0x67, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x02, 0x67, 0x6d, 0x12, 0x34, 0x0a, 0x02, 0x6c, 0x6d, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x6c, 0x6d, 0x12, 0x16, 0x0a, 0x04, 0x72,
	0x69, 0x36, 0x34, 0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x00, 0x52, 0x04, 0x72,
	0x69, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x04, 0x72, 0x66, 0x33, 0x32, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x0f, 0x42, 0x02, 0x10, 0x00, 0x52, 0x04, 0x72, 0x66, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x04, 0x72,
	0x66, 0x36, 0x34, 0x18, 0x16, 0x20, 0x03, 0x28, 0x10, 0x42, 0x02, 0x10, 0x00, 0x52, 0x04, 0x72,
	0x66, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x73, 0x74, 0x72, 0x18, 0x17, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x73, 0x74, 0x72, 0x12, 0x28, 0x0a, 0x03, 0x72, 0x67, 0x6d, 0x18, 0x18,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x72, 0x67,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x70, 0x69, 0x36, 0x34, 0x18, 0x19, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x05, 0x72, 0x70, 0x69, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x70, 0x66, 0x33, 0x32,
	0x18, 0x1a, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x05, 0x72, 0x70, 0x66, 0x33, 0x32, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x70, 0x66, 0x36, 0x34, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x10, 0x52, 0x05, 0x72, 0x70,
	0x66, 0x36, 0x34, 0x1a, 0x1e, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4c, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x4c, 0x31, 0x22, 0x1b, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x06, 0x0a, 0x02, 0x4c, 0x41, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x42, 0x10, 0x01,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x47, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x47,
	0x31, 0x2a, 0x1a, 0x0a, 0x0a, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_proto_goTypes = []any{
	(GlobalEnum)(0),                     // 0: example.GlobalEnum
	(ExampleMessage_LocalEnum)(0),       // 1: example.ExampleMessage.LocalEnum
	(*ExampleMessage)(nil),              // 2: example.ExampleMessage
	(*GlobalMessage)(nil),               // 3: example.GlobalMessage
	(*ExampleMessage_LocalMessage)(nil), // 4: example.ExampleMessage.LocalMessage
}
var file_example_proto_depIdxs = []int32{
	0, // 0: example.ExampleMessage.ge:type_name -> example.GlobalEnum
	1, // 1: example.ExampleMessage.le:type_name -> example.ExampleMessage.LocalEnum
	3, // 2: example.ExampleMessage.gm:type_name -> example.GlobalMessage
	4, // 3: example.ExampleMessage.lm:type_name -> example.ExampleMessage.LocalMessage
	3, // 4: example.ExampleMessage.rgm:type_name -> example.GlobalMessage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExampleMessage); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GlobalMessage); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExampleMessage_LocalMessage); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		EnumInfos:         file_example_proto_enumTypes,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
