// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type E int32

const (
	E_ZERO E = 0
)

var E_name = map[int32]string{
	0: "ZERO",
}

var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}

func (x *E) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(E_value, data, "E")
	if err != nil {
		return err
	}
	*x = E(value)
	return nil
}

func (E) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

type M_Subenum int32

const (
	M_M_ZERO M_Subenum = 0
)

var M_Subenum_name = map[int32]string{
	0: "M_ZERO",
}

var M_Subenum_value = map[string]int32{
	"M_ZERO": 0,
}

func (x M_Subenum) Enum() *M_Subenum {
	p := new(M_Subenum)
	*p = x
	return p
}

func (x M_Subenum) String() string {
	return proto.EnumName(M_Subenum_name, int32(x))
}

func (x *M_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Subenum_value, data, "M_Subenum")
	if err != nil {
		return err
	}
	*x = M_Subenum(value)
	return nil
}

func (M_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

type M_Submessage_Submessage_Subenum int32

const (
	M_Submessage_M_SUBMESSAGE_ZERO M_Submessage_Submessage_Subenum = 0
)

var M_Submessage_Submessage_Subenum_name = map[int32]string{
	0: "M_SUBMESSAGE_ZERO",
}

var M_Submessage_Submessage_Subenum_value = map[string]int32{
	"M_SUBMESSAGE_ZERO": 0,
}

func (x M_Submessage_Submessage_Subenum) Enum() *M_Submessage_Submessage_Subenum {
	p := new(M_Submessage_Submessage_Subenum)
	*p = x
	return p
}

func (x M_Submessage_Submessage_Subenum) String() string {
	return proto.EnumName(M_Submessage_Submessage_Subenum_name, int32(x))
}

func (x *M_Submessage_Submessage_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Submessage_Submessage_Subenum_value, data, "M_Submessage_Submessage_Subenum")
	if err != nil {
		return err
	}
	*x = M_Submessage_Submessage_Subenum(value)
	return nil
}

func (M_Submessage_Submessage_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 1, 0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2 *M2 `protobuf:"bytes,1,opt,name=m2" json:"m2,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*M_OneofInt32
	//	*M_OneofInt64
	OneofField           isM_OneofField `protobuf_oneof:"oneof_field"`
	Grouping             *M_Grouping    `protobuf:"group,4,opt,name=Grouping,json=grouping" json:"grouping,omitempty"`
	DefaultField         *string        `protobuf:"bytes,6,opt,name=default_field,json=defaultField,def=def" json:"default_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}
func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

const Default_M_DefaultField string = "def"

type isM_OneofField interface {
	isM_OneofField()
}

type M_OneofInt32 struct {
	OneofInt32 int32 `protobuf:"varint,2,opt,name=oneof_int32,json=oneofInt32,oneof" json:"oneof_int32,omitempty"`
}
type M_OneofInt64 struct {
	OneofInt64 int64 `protobuf:"varint,3,opt,name=oneof_int64,json=oneofInt64,oneof" json:"oneof_int64,omitempty"`
}

func (*M_OneofInt32) isM_OneofField() {}
func (*M_OneofInt64) isM_OneofField() {}

func (m *M) GetOneofField() isM_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func (m *M) GetOneofInt32() int32 {
	if x, ok := m.GetOneofField().(*M_OneofInt32); ok {
		return x.OneofInt32
	}
	return 0
}

func (m *M) GetOneofInt64() int64 {
	if x, ok := m.GetOneofField().(*M_OneofInt64); ok {
		return x.OneofInt64
	}
	return 0
}

func (m *M) GetGrouping() *M_Grouping {
	if m != nil {
		return m.Grouping
	}
	return nil
}

func (m *M) GetDefaultField() string {
	if m != nil && m.DefaultField != nil {
		return *m.DefaultField
	}
	return Default_M_DefaultField
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_OneofInt32)(nil),
		(*M_OneofInt64)(nil),
	}
}

type M_Grouping struct {
	GroupField           *string  `protobuf:"bytes,5,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M_Grouping) Reset()         { *m = M_Grouping{} }
func (m *M_Grouping) String() string { return proto.CompactTextString(m) }
func (*M_Grouping) ProtoMessage()    {}
func (*M_Grouping) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}
func (m *M_Grouping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Grouping.Unmarshal(m, b)
}
func (m *M_Grouping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Grouping.Marshal(b, m, deterministic)
}
func (m *M_Grouping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Grouping.Merge(m, src)
}
func (m *M_Grouping) XXX_Size() int {
	return xxx_messageInfo_M_Grouping.Size(m)
}
func (m *M_Grouping) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Grouping.DiscardUnknown(m)
}

var xxx_messageInfo_M_Grouping proto.InternalMessageInfo

func (m *M_Grouping) GetGroupField() string {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return ""
}

type M_Submessage struct {
	// Types that are valid to be assigned to SubmessageOneofField:
	//	*M_Submessage_SubmessageOneofInt32
	//	*M_Submessage_SubmessageOneofInt64
	SubmessageOneofField isM_Submessage_SubmessageOneofField `protobuf_oneof:"submessage_oneof_field"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *M_Submessage) Reset()         { *m = M_Submessage{} }
func (m *M_Submessage) String() string { return proto.CompactTextString(m) }
func (*M_Submessage) ProtoMessage()    {}
func (*M_Submessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 1}
}
func (m *M_Submessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Submessage.Unmarshal(m, b)
}
func (m *M_Submessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Submessage.Marshal(b, m, deterministic)
}
func (m *M_Submessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Submessage.Merge(m, src)
}
func (m *M_Submessage) XXX_Size() int {
	return xxx_messageInfo_M_Submessage.Size(m)
}
func (m *M_Submessage) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Submessage.DiscardUnknown(m)
}

var xxx_messageInfo_M_Submessage proto.InternalMessageInfo

type isM_Submessage_SubmessageOneofField interface {
	isM_Submessage_SubmessageOneofField()
}

type M_Submessage_SubmessageOneofInt32 struct {
	SubmessageOneofInt32 int32 `protobuf:"varint,1,opt,name=submessage_oneof_int32,json=submessageOneofInt32,oneof" json:"submessage_oneof_int32,omitempty"`
}
type M_Submessage_SubmessageOneofInt64 struct {
	SubmessageOneofInt64 int64 `protobuf:"varint,2,opt,name=submessage_oneof_int64,json=submessageOneofInt64,oneof" json:"submessage_oneof_int64,omitempty"`
}

func (*M_Submessage_SubmessageOneofInt32) isM_Submessage_SubmessageOneofField() {}
func (*M_Submessage_SubmessageOneofInt64) isM_Submessage_SubmessageOneofField() {}

func (m *M_Submessage) GetSubmessageOneofField() isM_Submessage_SubmessageOneofField {
	if m != nil {
		return m.SubmessageOneofField
	}
	return nil
}

func (m *M_Submessage) GetSubmessageOneofInt32() int32 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt32); ok {
		return x.SubmessageOneofInt32
	}
	return 0
}

func (m *M_Submessage) GetSubmessageOneofInt64() int64 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt64); ok {
		return x.SubmessageOneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M_Submessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_Submessage_SubmessageOneofInt32)(nil),
		(*M_Submessage_SubmessageOneofInt64)(nil),
	}
}

var E_ExtensionField = &proto.ExtensionDesc{
	ExtendedType:  (*M2)(nil),
	ExtensionType: (*string)(nil),
	Field:         1,
	Name:          "goproto.test.import_public.sub.extension_field",
	Tag:           "bytes,1,opt,name=extension_field",
	Filename:      "import_public/sub/a.proto",
}

func init() {
	proto.RegisterEnum("goproto.test.import_public.sub.E", E_name, E_value)
	proto.RegisterEnum("goproto.test.import_public.sub.M_Subenum", M_Subenum_name, M_Subenum_value)
	proto.RegisterEnum("goproto.test.import_public.sub.M_Submessage_Submessage_Subenum", M_Submessage_Submessage_Subenum_name, M_Submessage_Submessage_Subenum_value)
	proto.RegisterType((*M)(nil), "goproto.test.import_public.sub.M")
	proto.RegisterType((*M_Grouping)(nil), "goproto.test.import_public.sub.M.Grouping")
	proto.RegisterType((*M_Submessage)(nil), "goproto.test.import_public.sub.M.Submessage")
	proto.RegisterExtension(E_ExtensionField)
}

func init() { proto.RegisterFile("import_public/sub/a.proto", fileDescriptor_382f7805394b5c4e) }

var fileDescriptor_382f7805394b5c4e = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xab, 0x24, 0xed, 0xd2, 0x93, 0x65, 0x7f, 0xc4, 0x3a, 0xbc, 0x5e, 0x6c, 0x5e, 0xae,
	0x4c, 0x4b, 0x64, 0xf0, 0x82, 0x2f, 0x7a, 0xb7, 0x40, 0xd2, 0x8e, 0x61, 0x0a, 0x36, 0xbb, 0xe9,
	0x8d, 0xb1, 0x12, 0x59, 0x33, 0xc4, 0x52, 0xb0, 0x24, 0xd8, 0x23, 0xec, 0xbd, 0xf6, 0x62, 0xc3,
	0xb2, 0x9d, 0x34, 0x2c, 0xdb, 0x7a, 0x27, 0x9f, 0xf3, 0xfd, 0xbe, 0xc3, 0xf9, 0x7c, 0xe0, 0x5d,
	0x51, 0x6e, 0x65, 0xa5, 0xd3, 0xad, 0xa1, 0x9b, 0x62, 0xe5, 0x2b, 0x43, 0xfd, 0x8c, 0x6c, 0x2b,
	0xa9, 0x25, 0x7e, 0xcf, 0xa5, 0x7d, 0x10, 0xcd, 0x94, 0x26, 0x07, 0x3a, 0xa2, 0x0c, 0xbd, 0x3c,
	0x82, 0xd2, 0x06, 0x9d, 0xfc, 0x1c, 0x00, 0x8a, 0x70, 0x00, 0xbd, 0x32, 0x70, 0x90, 0x8b, 0xbc,
	0x51, 0x30, 0x21, 0xff, 0x76, 0x23, 0x51, 0x10, 0xf7, 0xca, 0x00, 0x7f, 0x84, 0x91, 0x14, 0x4c,
	0xe6, 0x69, 0x21, 0xf4, 0xa7, 0xc0, 0xe9, 0xb9, 0xc8, 0x3b, 0xbd, 0x3b, 0x89, 0xc1, 0x16, 0xbf,
	0xd4, 0xb5, 0x03, 0x49, 0x38, 0x73, 0xfa, 0x2e, 0xf2, 0xfa, 0x8f, 0x25, 0xe1, 0x0c, 0x2f, 0x61,
	0xc8, 0x2b, 0x69, 0xb6, 0x85, 0xe0, 0xce, 0xc0, 0x45, 0x1e, 0x04, 0x57, 0xff, 0x9d, 0x4f, 0x6e,
	0x5b, 0x22, 0xde, 0xb1, 0xd8, 0x83, 0xf1, 0x9a, 0xe5, 0x99, 0xd9, 0xe8, 0x34, 0x2f, 0xd8, 0x66,
	0xed, 0x9c, 0xb9, 0xc8, 0x3b, 0xbf, 0xe9, 0xaf, 0x59, 0x1e, 0x3f, 0x6f, 0x3b, 0xcb, 0xba, 0x71,
	0x79, 0x0d, 0xc3, 0x8e, 0xc7, 0x1f, 0x60, 0x64, 0x1d, 0x5a, 0xe6, 0xb4, 0x66, 0x62, 0xb0, 0xa5,
	0x46, 0xfc, 0x0b, 0x01, 0x24, 0x86, 0x96, 0x4c, 0xa9, 0x8c, 0x33, 0x1c, 0xc2, 0x5b, 0xb5, 0xfb,
	0x4a, 0x1f, 0xaf, 0x8f, 0xda, 0xf5, 0xdf, 0xec, 0xfb, 0xf7, 0xfb, 0x20, 0xfe, 0xc2, 0x85, 0x33,
	0x1b, 0x5b, 0xff, 0x38, 0x17, 0xce, 0x26, 0xd7, 0x80, 0xf7, 0xd3, 0xd3, 0xc4, 0x50, 0x26, 0x4c,
	0x89, 0x2f, 0xe0, 0x75, 0x94, 0x26, 0xdf, 0xe6, 0xd1, 0x22, 0x49, 0x3e, 0xdf, 0x2e, 0xd2, 0x87,
	0x45, 0x7c, 0xff, 0xea, 0x64, 0xee, 0x1c, 0x19, 0x62, 0xf7, 0x9a, 0x5c, 0xc0, 0xb3, 0x8e, 0x05,
	0x38, 0x8b, 0x3a, 0x60, 0xdc, 0xfd, 0x1e, 0xab, 0xba, 0x1a, 0x03, 0x5a, 0xe0, 0x21, 0x0c, 0x9a,
	0xee, 0xcd, 0x57, 0x78, 0xc9, 0x7e, 0x68, 0x26, 0x54, 0x21, 0x45, 0xa3, 0xc0, 0x4f, 0x38, 0x0d,
	0x1b, 0xc4, 0x79, 0xfc, 0x62, 0x87, 0xda, 0x1c, 0xe7, 0x77, 0x0f, 0x4b, 0x5e, 0xe8, 0xef, 0x86,
	0x92, 0x95, 0x2c, 0x7d, 0x5d, 0x49, 0x31, 0x35, 0xca, 0xb7, 0x5e, 0xd4, 0xe4, 0xcd, 0x63, 0x35,
	0xe5, 0x4c, 0x4c, 0xb9, 0xe4, 0xd2, 0xaf, 0xed, 0xd7, 0x99, 0xce, 0xfc, 0x3f, 0x0e, 0xf7, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x04, 0x14, 0xdb, 0x07, 0x03, 0x00, 0x00,
}
