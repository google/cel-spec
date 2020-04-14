// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test/v1/proto2/test_all_types.proto

package google_api_expr_test_v1_proto2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GlobalEnum int32

const (
	GlobalEnum_GOO GlobalEnum = 0
	GlobalEnum_GAR GlobalEnum = 1
	GlobalEnum_GAZ GlobalEnum = 2
)

var GlobalEnum_name = map[int32]string{
	0: "GOO",
	1: "GAR",
	2: "GAZ",
}

var GlobalEnum_value = map[string]int32{
	"GOO": 0,
	"GAR": 1,
	"GAZ": 2,
}

func (x GlobalEnum) Enum() *GlobalEnum {
	p := new(GlobalEnum)
	*p = x
	return p
}

func (x GlobalEnum) String() string {
	return proto.EnumName(GlobalEnum_name, int32(x))
}

func (x *GlobalEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GlobalEnum_value, data, "GlobalEnum")
	if err != nil {
		return err
	}
	*x = GlobalEnum(value)
	return nil
}

func (GlobalEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98cc99a98b5d98fe, []int{0}
}

type TestAllTypes_NestedEnum int32

const (
	TestAllTypes_FOO TestAllTypes_NestedEnum = 0
	TestAllTypes_BAR TestAllTypes_NestedEnum = 1
	TestAllTypes_BAZ TestAllTypes_NestedEnum = 2
)

var TestAllTypes_NestedEnum_name = map[int32]string{
	0: "FOO",
	1: "BAR",
	2: "BAZ",
}

var TestAllTypes_NestedEnum_value = map[string]int32{
	"FOO": 0,
	"BAR": 1,
	"BAZ": 2,
}

func (x TestAllTypes_NestedEnum) Enum() *TestAllTypes_NestedEnum {
	p := new(TestAllTypes_NestedEnum)
	*p = x
	return p
}

func (x TestAllTypes_NestedEnum) String() string {
	return proto.EnumName(TestAllTypes_NestedEnum_name, int32(x))
}

func (x *TestAllTypes_NestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestAllTypes_NestedEnum_value, data, "TestAllTypes_NestedEnum")
	if err != nil {
		return err
	}
	*x = TestAllTypes_NestedEnum(value)
	return nil
}

func (TestAllTypes_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98cc99a98b5d98fe, []int{0, 0}
}

type TestAllTypes struct {
	SingleInt32         *int32                `protobuf:"varint,1,opt,name=single_int32,json=singleInt32,def=-32" json:"single_int32,omitempty"`
	SingleInt64         *int64                `protobuf:"varint,2,opt,name=single_int64,json=singleInt64,def=-64" json:"single_int64,omitempty"`
	SingleUint32        *uint32               `protobuf:"varint,3,opt,name=single_uint32,json=singleUint32,def=32" json:"single_uint32,omitempty"`
	SingleUint64        *uint64               `protobuf:"varint,4,opt,name=single_uint64,json=singleUint64,def=64" json:"single_uint64,omitempty"`
	SingleSint32        *int32                `protobuf:"zigzag32,5,opt,name=single_sint32,json=singleSint32" json:"single_sint32,omitempty"`
	SingleSint64        *int64                `protobuf:"zigzag64,6,opt,name=single_sint64,json=singleSint64" json:"single_sint64,omitempty"`
	SingleFixed32       *uint32               `protobuf:"fixed32,7,opt,name=single_fixed32,json=singleFixed32" json:"single_fixed32,omitempty"`
	SingleFixed64       *uint64               `protobuf:"fixed64,8,opt,name=single_fixed64,json=singleFixed64" json:"single_fixed64,omitempty"`
	SingleSfixed32      *int32                `protobuf:"fixed32,9,opt,name=single_sfixed32,json=singleSfixed32" json:"single_sfixed32,omitempty"`
	SingleSfixed64      *int64                `protobuf:"fixed64,10,opt,name=single_sfixed64,json=singleSfixed64" json:"single_sfixed64,omitempty"`
	SingleFloat         *float32              `protobuf:"fixed32,11,opt,name=single_float,json=singleFloat,def=3" json:"single_float,omitempty"`
	SingleDouble        *float64              `protobuf:"fixed64,12,opt,name=single_double,json=singleDouble,def=6.4" json:"single_double,omitempty"`
	SingleBool          *bool                 `protobuf:"varint,13,opt,name=single_bool,json=singleBool,def=1" json:"single_bool,omitempty"`
	SingleString        *string               `protobuf:"bytes,14,opt,name=single_string,json=singleString,def=empty" json:"single_string,omitempty"`
	SingleBytes         []byte                `protobuf:"bytes,15,opt,name=single_bytes,json=singleBytes,def=none" json:"single_bytes,omitempty"`
	SingleAny           *any.Any              `protobuf:"bytes,100,opt,name=single_any,json=singleAny" json:"single_any,omitempty"`
	SingleDuration      *duration.Duration    `protobuf:"bytes,101,opt,name=single_duration,json=singleDuration" json:"single_duration,omitempty"`
	SingleTimestamp     *timestamp.Timestamp  `protobuf:"bytes,102,opt,name=single_timestamp,json=singleTimestamp" json:"single_timestamp,omitempty"`
	SingleStruct        *_struct.Struct       `protobuf:"bytes,103,opt,name=single_struct,json=singleStruct" json:"single_struct,omitempty"`
	SingleValue         *_struct.Value        `protobuf:"bytes,104,opt,name=single_value,json=singleValue" json:"single_value,omitempty"`
	SingleInt64Wrapper  *wrappers.Int64Value  `protobuf:"bytes,105,opt,name=single_int64_wrapper,json=singleInt64Wrapper" json:"single_int64_wrapper,omitempty"`
	SingleInt32Wrapper  *wrappers.Int32Value  `protobuf:"bytes,106,opt,name=single_int32_wrapper,json=singleInt32Wrapper" json:"single_int32_wrapper,omitempty"`
	SingleDoubleWrapper *wrappers.DoubleValue `protobuf:"bytes,107,opt,name=single_double_wrapper,json=singleDoubleWrapper" json:"single_double_wrapper,omitempty"`
	SingleFloatWrapper  *wrappers.FloatValue  `protobuf:"bytes,108,opt,name=single_float_wrapper,json=singleFloatWrapper" json:"single_float_wrapper,omitempty"`
	SingleUint64Wrapper *wrappers.UInt64Value `protobuf:"bytes,109,opt,name=single_uint64_wrapper,json=singleUint64Wrapper" json:"single_uint64_wrapper,omitempty"`
	SingleUint32Wrapper *wrappers.UInt32Value `protobuf:"bytes,110,opt,name=single_uint32_wrapper,json=singleUint32Wrapper" json:"single_uint32_wrapper,omitempty"`
	SingleStringWrapper *wrappers.StringValue `protobuf:"bytes,111,opt,name=single_string_wrapper,json=singleStringWrapper" json:"single_string_wrapper,omitempty"`
	SingleBoolWrapper   *wrappers.BoolValue   `protobuf:"bytes,112,opt,name=single_bool_wrapper,json=singleBoolWrapper" json:"single_bool_wrapper,omitempty"`
	SingleBytesWrapper  *wrappers.BytesValue  `protobuf:"bytes,113,opt,name=single_bytes_wrapper,json=singleBytesWrapper" json:"single_bytes_wrapper,omitempty"`
	// Types that are valid to be assigned to NestedType:
	//	*TestAllTypes_SingleNestedMessage
	//	*TestAllTypes_SingleNestedEnum
	NestedType            isTestAllTypes_NestedType     `protobuf_oneof:"nested_type"`
	RepeatedInt32         []int32                       `protobuf:"varint,31,rep,name=repeated_int32,json=repeatedInt32" json:"repeated_int32,omitempty"`
	RepeatedInt64         []int64                       `protobuf:"varint,32,rep,name=repeated_int64,json=repeatedInt64" json:"repeated_int64,omitempty"`
	RepeatedUint32        []uint32                      `protobuf:"varint,33,rep,name=repeated_uint32,json=repeatedUint32" json:"repeated_uint32,omitempty"`
	RepeatedUint64        []uint64                      `protobuf:"varint,34,rep,name=repeated_uint64,json=repeatedUint64" json:"repeated_uint64,omitempty"`
	RepeatedSint32        []int32                       `protobuf:"zigzag32,35,rep,name=repeated_sint32,json=repeatedSint32" json:"repeated_sint32,omitempty"`
	RepeatedSint64        []int64                       `protobuf:"zigzag64,36,rep,name=repeated_sint64,json=repeatedSint64" json:"repeated_sint64,omitempty"`
	RepeatedFixed32       []uint32                      `protobuf:"fixed32,37,rep,name=repeated_fixed32,json=repeatedFixed32" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64       []uint64                      `protobuf:"fixed64,38,rep,name=repeated_fixed64,json=repeatedFixed64" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32      []int32                       `protobuf:"fixed32,39,rep,name=repeated_sfixed32,json=repeatedSfixed32" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64      []int64                       `protobuf:"fixed64,40,rep,name=repeated_sfixed64,json=repeatedSfixed64" json:"repeated_sfixed64,omitempty"`
	RepeatedFloat         []float32                     `protobuf:"fixed32,41,rep,name=repeated_float,json=repeatedFloat" json:"repeated_float,omitempty"`
	RepeatedDouble        []float64                     `protobuf:"fixed64,42,rep,name=repeated_double,json=repeatedDouble" json:"repeated_double,omitempty"`
	RepeatedBool          []bool                        `protobuf:"varint,43,rep,name=repeated_bool,json=repeatedBool" json:"repeated_bool,omitempty"`
	RepeatedString        []string                      `protobuf:"bytes,44,rep,name=repeated_string,json=repeatedString" json:"repeated_string,omitempty"`
	RepeatedBytes         [][]byte                      `protobuf:"bytes,45,rep,name=repeated_bytes,json=repeatedBytes" json:"repeated_bytes,omitempty"`
	RepeatedNestedMessage []*TestAllTypes_NestedMessage `protobuf:"bytes,51,rep,name=repeated_nested_message,json=repeatedNestedMessage" json:"repeated_nested_message,omitempty"`
	RepeatedNestedEnum    []TestAllTypes_NestedEnum     `protobuf:"varint,52,rep,name=repeated_nested_enum,json=repeatedNestedEnum,enum=google.api.expr.test.v1.proto2.TestAllTypes_NestedEnum" json:"repeated_nested_enum,omitempty"`
	RepeatedStringPiece   []string                      `protobuf:"bytes,53,rep,name=repeated_string_piece,json=repeatedStringPiece" json:"repeated_string_piece,omitempty"`
	RepeatedCord          []string                      `protobuf:"bytes,54,rep,name=repeated_cord,json=repeatedCord" json:"repeated_cord,omitempty"`
	RepeatedLazyMessage   []*TestAllTypes_NestedMessage `protobuf:"bytes,55,rep,name=repeated_lazy_message,json=repeatedLazyMessage" json:"repeated_lazy_message,omitempty"`
	MapStringString       map[string]string             `protobuf:"bytes,61,rep,name=map_string_string,json=mapStringString" json:"map_string_string,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapInt64NestedType    map[int64]*NestedTestAllTypes `protobuf:"bytes,62,rep,name=map_int64_nested_type,json=mapInt64NestedType" json:"map_int64_nested_type,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral  struct{}                      `json:"-"`
	XXX_unrecognized      []byte                        `json:"-"`
	XXX_sizecache         int32                         `json:"-"`
}

func (m *TestAllTypes) Reset()         { *m = TestAllTypes{} }
func (m *TestAllTypes) String() string { return proto.CompactTextString(m) }
func (*TestAllTypes) ProtoMessage()    {}
func (*TestAllTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_98cc99a98b5d98fe, []int{0}
}

func (m *TestAllTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes.Unmarshal(m, b)
}
func (m *TestAllTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes.Marshal(b, m, deterministic)
}
func (m *TestAllTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes.Merge(m, src)
}
func (m *TestAllTypes) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes.Size(m)
}
func (m *TestAllTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes proto.InternalMessageInfo

const Default_TestAllTypes_SingleInt32 int32 = -32
const Default_TestAllTypes_SingleInt64 int64 = -64
const Default_TestAllTypes_SingleUint32 uint32 = 32
const Default_TestAllTypes_SingleUint64 uint64 = 64
const Default_TestAllTypes_SingleFloat float32 = 3
const Default_TestAllTypes_SingleDouble float64 = 6.4
const Default_TestAllTypes_SingleBool bool = true
const Default_TestAllTypes_SingleString string = "empty"

var Default_TestAllTypes_SingleBytes []byte = []byte("none")

const Default_TestAllTypes_SingleNestedEnum TestAllTypes_NestedEnum = TestAllTypes_BAR

func (m *TestAllTypes) GetSingleInt32() int32 {
	if m != nil && m.SingleInt32 != nil {
		return *m.SingleInt32
	}
	return Default_TestAllTypes_SingleInt32
}

func (m *TestAllTypes) GetSingleInt64() int64 {
	if m != nil && m.SingleInt64 != nil {
		return *m.SingleInt64
	}
	return Default_TestAllTypes_SingleInt64
}

func (m *TestAllTypes) GetSingleUint32() uint32 {
	if m != nil && m.SingleUint32 != nil {
		return *m.SingleUint32
	}
	return Default_TestAllTypes_SingleUint32
}

func (m *TestAllTypes) GetSingleUint64() uint64 {
	if m != nil && m.SingleUint64 != nil {
		return *m.SingleUint64
	}
	return Default_TestAllTypes_SingleUint64
}

func (m *TestAllTypes) GetSingleSint32() int32 {
	if m != nil && m.SingleSint32 != nil {
		return *m.SingleSint32
	}
	return 0
}

func (m *TestAllTypes) GetSingleSint64() int64 {
	if m != nil && m.SingleSint64 != nil {
		return *m.SingleSint64
	}
	return 0
}

func (m *TestAllTypes) GetSingleFixed32() uint32 {
	if m != nil && m.SingleFixed32 != nil {
		return *m.SingleFixed32
	}
	return 0
}

func (m *TestAllTypes) GetSingleFixed64() uint64 {
	if m != nil && m.SingleFixed64 != nil {
		return *m.SingleFixed64
	}
	return 0
}

func (m *TestAllTypes) GetSingleSfixed32() int32 {
	if m != nil && m.SingleSfixed32 != nil {
		return *m.SingleSfixed32
	}
	return 0
}

func (m *TestAllTypes) GetSingleSfixed64() int64 {
	if m != nil && m.SingleSfixed64 != nil {
		return *m.SingleSfixed64
	}
	return 0
}

func (m *TestAllTypes) GetSingleFloat() float32 {
	if m != nil && m.SingleFloat != nil {
		return *m.SingleFloat
	}
	return Default_TestAllTypes_SingleFloat
}

func (m *TestAllTypes) GetSingleDouble() float64 {
	if m != nil && m.SingleDouble != nil {
		return *m.SingleDouble
	}
	return Default_TestAllTypes_SingleDouble
}

func (m *TestAllTypes) GetSingleBool() bool {
	if m != nil && m.SingleBool != nil {
		return *m.SingleBool
	}
	return Default_TestAllTypes_SingleBool
}

func (m *TestAllTypes) GetSingleString() string {
	if m != nil && m.SingleString != nil {
		return *m.SingleString
	}
	return Default_TestAllTypes_SingleString
}

func (m *TestAllTypes) GetSingleBytes() []byte {
	if m != nil && m.SingleBytes != nil {
		return m.SingleBytes
	}
	return append([]byte(nil), Default_TestAllTypes_SingleBytes...)
}

func (m *TestAllTypes) GetSingleAny() *any.Any {
	if m != nil {
		return m.SingleAny
	}
	return nil
}

func (m *TestAllTypes) GetSingleDuration() *duration.Duration {
	if m != nil {
		return m.SingleDuration
	}
	return nil
}

func (m *TestAllTypes) GetSingleTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.SingleTimestamp
	}
	return nil
}

func (m *TestAllTypes) GetSingleStruct() *_struct.Struct {
	if m != nil {
		return m.SingleStruct
	}
	return nil
}

func (m *TestAllTypes) GetSingleValue() *_struct.Value {
	if m != nil {
		return m.SingleValue
	}
	return nil
}

func (m *TestAllTypes) GetSingleInt64Wrapper() *wrappers.Int64Value {
	if m != nil {
		return m.SingleInt64Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleInt32Wrapper() *wrappers.Int32Value {
	if m != nil {
		return m.SingleInt32Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleDoubleWrapper() *wrappers.DoubleValue {
	if m != nil {
		return m.SingleDoubleWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleFloatWrapper() *wrappers.FloatValue {
	if m != nil {
		return m.SingleFloatWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleUint64Wrapper() *wrappers.UInt64Value {
	if m != nil {
		return m.SingleUint64Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleUint32Wrapper() *wrappers.UInt32Value {
	if m != nil {
		return m.SingleUint32Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleStringWrapper() *wrappers.StringValue {
	if m != nil {
		return m.SingleStringWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleBoolWrapper() *wrappers.BoolValue {
	if m != nil {
		return m.SingleBoolWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleBytesWrapper() *wrappers.BytesValue {
	if m != nil {
		return m.SingleBytesWrapper
	}
	return nil
}

type isTestAllTypes_NestedType interface {
	isTestAllTypes_NestedType()
}

type TestAllTypes_SingleNestedMessage struct {
	SingleNestedMessage *TestAllTypes_NestedMessage `protobuf:"bytes,21,opt,name=single_nested_message,json=singleNestedMessage,oneof"`
}

type TestAllTypes_SingleNestedEnum struct {
	SingleNestedEnum TestAllTypes_NestedEnum `protobuf:"varint,22,opt,name=single_nested_enum,json=singleNestedEnum,enum=google.api.expr.test.v1.proto2.TestAllTypes_NestedEnum,oneof,def=1"`
}

func (*TestAllTypes_SingleNestedMessage) isTestAllTypes_NestedType() {}

func (*TestAllTypes_SingleNestedEnum) isTestAllTypes_NestedType() {}

func (m *TestAllTypes) GetNestedType() isTestAllTypes_NestedType {
	if m != nil {
		return m.NestedType
	}
	return nil
}

func (m *TestAllTypes) GetSingleNestedMessage() *TestAllTypes_NestedMessage {
	if x, ok := m.GetNestedType().(*TestAllTypes_SingleNestedMessage); ok {
		return x.SingleNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetSingleNestedEnum() TestAllTypes_NestedEnum {
	if x, ok := m.GetNestedType().(*TestAllTypes_SingleNestedEnum); ok {
		return x.SingleNestedEnum
	}
	return Default_TestAllTypes_SingleNestedEnum
}

func (m *TestAllTypes) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint32() []int32 {
	if m != nil {
		return m.RepeatedSint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint64() []int64 {
	if m != nil {
		return m.RepeatedSint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedMessage() []*TestAllTypes_NestedMessage {
	if m != nil {
		return m.RepeatedNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedEnum() []TestAllTypes_NestedEnum {
	if m != nil {
		return m.RepeatedNestedEnum
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedStringPiece() []string {
	if m != nil {
		return m.RepeatedStringPiece
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedCord() []string {
	if m != nil {
		return m.RepeatedCord
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedLazyMessage() []*TestAllTypes_NestedMessage {
	if m != nil {
		return m.RepeatedLazyMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapStringString() map[string]string {
	if m != nil {
		return m.MapStringString
	}
	return nil
}

func (m *TestAllTypes) GetMapInt64NestedType() map[int64]*NestedTestAllTypes {
	if m != nil {
		return m.MapInt64NestedType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestAllTypes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestAllTypes_SingleNestedMessage)(nil),
		(*TestAllTypes_SingleNestedEnum)(nil),
	}
}

type TestAllTypes_NestedMessage struct {
	Bb                   *int32   `protobuf:"varint,1,opt,name=bb" json:"bb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestAllTypes_NestedMessage) Reset()         { *m = TestAllTypes_NestedMessage{} }
func (m *TestAllTypes_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*TestAllTypes_NestedMessage) ProtoMessage()    {}
func (*TestAllTypes_NestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_98cc99a98b5d98fe, []int{0, 0}
}

func (m *TestAllTypes_NestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Unmarshal(m, b)
}
func (m *TestAllTypes_NestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Marshal(b, m, deterministic)
}
func (m *TestAllTypes_NestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes_NestedMessage.Merge(m, src)
}
func (m *TestAllTypes_NestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Size(m)
}
func (m *TestAllTypes_NestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes_NestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes_NestedMessage proto.InternalMessageInfo

func (m *TestAllTypes_NestedMessage) GetBb() int32 {
	if m != nil && m.Bb != nil {
		return *m.Bb
	}
	return 0
}

type NestedTestAllTypes struct {
	Child                *NestedTestAllTypes `protobuf:"bytes,1,opt,name=child" json:"child,omitempty"`
	Payload              *TestAllTypes       `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NestedTestAllTypes) Reset()         { *m = NestedTestAllTypes{} }
func (m *NestedTestAllTypes) String() string { return proto.CompactTextString(m) }
func (*NestedTestAllTypes) ProtoMessage()    {}
func (*NestedTestAllTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_98cc99a98b5d98fe, []int{1}
}

func (m *NestedTestAllTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedTestAllTypes.Unmarshal(m, b)
}
func (m *NestedTestAllTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NestedTestAllTypes.Marshal(b, m, deterministic)
}
func (m *NestedTestAllTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedTestAllTypes.Merge(m, src)
}
func (m *NestedTestAllTypes) XXX_Size() int {
	return xxx_messageInfo_NestedTestAllTypes.Size(m)
}
func (m *NestedTestAllTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedTestAllTypes.DiscardUnknown(m)
}

var xxx_messageInfo_NestedTestAllTypes proto.InternalMessageInfo

func (m *NestedTestAllTypes) GetChild() *NestedTestAllTypes {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *NestedTestAllTypes) GetPayload() *TestAllTypes {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.api.expr.test.v1.proto2.GlobalEnum", GlobalEnum_name, GlobalEnum_value)
	proto.RegisterEnum("google.api.expr.test.v1.proto2.TestAllTypes_NestedEnum", TestAllTypes_NestedEnum_name, TestAllTypes_NestedEnum_value)
	proto.RegisterType((*TestAllTypes)(nil), "google.api.expr.test.v1.proto2.TestAllTypes")
	proto.RegisterMapType((map[int64]*NestedTestAllTypes)(nil), "google.api.expr.test.v1.proto2.TestAllTypes.MapInt64NestedTypeEntry")
	proto.RegisterMapType((map[string]string)(nil), "google.api.expr.test.v1.proto2.TestAllTypes.MapStringStringEntry")
	proto.RegisterType((*TestAllTypes_NestedMessage)(nil), "google.api.expr.test.v1.proto2.TestAllTypes.NestedMessage")
	proto.RegisterType((*NestedTestAllTypes)(nil), "google.api.expr.test.v1.proto2.NestedTestAllTypes")
}

func init() {
	proto.RegisterFile("proto/test/v1/proto2/test_all_types.proto", fileDescriptor_98cc99a98b5d98fe)
}

var fileDescriptor_98cc99a98b5d98fe = []byte{
	// 1310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0x6d, 0x73, 0xd3, 0x46,
	0x10, 0xe6, 0x74, 0x09, 0x49, 0x2e, 0x71, 0xe2, 0x5c, 0x12, 0x38, 0x5c, 0x06, 0xae, 0xe1, 0x25,
	0xc7, 0x9b, 0x33, 0xc8, 0xaa, 0x68, 0x3d, 0x6d, 0x67, 0xe2, 0x42, 0x42, 0x3b, 0xa5, 0x30, 0x0a,
	0xb4, 0x33, 0xfd, 0xe2, 0x91, 0xe3, 0x4b, 0x70, 0x91, 0x25, 0x55, 0x96, 0x01, 0xf1, 0x83, 0xfa,
	0xe3, 0xfa, 0x2b, 0x3a, 0xba, 0x37, 0xeb, 0x05, 0xd3, 0x81, 0x7c, 0xc9, 0x48, 0xbb, 0xcf, 0x3e,
	0xbb, 0xfb, 0xdc, 0xad, 0x36, 0x46, 0x77, 0xe2, 0x24, 0x4a, 0xa3, 0xfd, 0x94, 0x4f, 0xd2, 0xfd,
	0xb7, 0x0f, 0xf7, 0xc5, 0x9b, 0x2d, 0x5e, 0xfb, 0x7e, 0x10, 0xf4, 0xd3, 0x2c, 0xe6, 0x93, 0xb6,
	0xb0, 0xe2, 0x6b, 0x67, 0x51, 0x74, 0x16, 0xf0, 0xb6, 0x1f, 0x8f, 0xda, 0xfc, 0x7d, 0x9c, 0xb4,
	0x73, 0x54, 0xfb, 0xed, 0x43, 0xe9, 0xb6, 0x5b, 0x57, 0xa4, 0x5f, 0x72, 0x0c, 0xa6, 0xa7, 0xfb,
	0x7e, 0x98, 0x49, 0x5f, 0xeb, 0x5a, 0xd5, 0x35, 0x9c, 0x26, 0x7e, 0x3a, 0x8a, 0x42, 0xe5, 0xbf,
	0x5a, 0xf5, 0x4f, 0xd2, 0x64, 0x7a, 0x92, 0x2a, 0xef, 0xf5, 0xaa, 0x37, 0x1d, 0x8d, 0xf9, 0x24,
	0xf5, 0xc7, 0xf1, 0x3c, 0xfa, 0x77, 0x89, 0x1f, 0xc7, 0x3c, 0x51, 0x95, 0xef, 0xfe, 0xdb, 0x42,
	0x6b, 0x2f, 0xf9, 0x24, 0x3d, 0x08, 0x82, 0x97, 0x79, 0x43, 0xf8, 0x36, 0x5a, 0x9b, 0x8c, 0xc2,
	0xb3, 0x80, 0xf7, 0x47, 0x61, 0xda, 0xb1, 0x09, 0xa0, 0x80, 0x2d, 0x76, 0xe1, 0x83, 0x8e, 0xed,
	0xad, 0x4a, 0xc7, 0xcf, 0xb9, 0xbd, 0x8c, 0x73, 0x1d, 0x62, 0x51, 0xc0, 0x60, 0x17, 0x3e, 0x70,
	0x9d, 0x02, 0xce, 0x75, 0xf0, 0x1e, 0x6a, 0x28, 0xdc, 0x54, 0x12, 0x42, 0x0a, 0x58, 0xa3, 0x6b,
	0x75, 0x6c, 0x4f, 0x11, 0xbc, 0x12, 0xf6, 0x0a, 0xd0, 0x75, 0xc8, 0x02, 0x05, 0x6c, 0xa1, 0x6b,
	0xb9, 0x4e, 0x11, 0xe8, 0x3a, 0xf8, 0x86, 0x01, 0x4e, 0x24, 0xe3, 0x22, 0x05, 0x6c, 0x53, 0x83,
	0x8e, 0x25, 0x5b, 0x19, 0xe4, 0x3a, 0xe4, 0x22, 0x05, 0x0c, 0x17, 0x41, 0xae, 0x83, 0x6f, 0xa1,
	0x75, 0x05, 0x3a, 0x1d, 0xbd, 0xe7, 0xc3, 0x8e, 0x4d, 0x96, 0x28, 0x60, 0x4b, 0x9e, 0x0a, 0x3d,
	0x94, 0xc6, 0x2a, 0xcc, 0x75, 0xc8, 0x32, 0x05, 0xec, 0x62, 0x09, 0x26, 0x3a, 0xdd, 0xd0, 0x29,
	0x35, 0xdd, 0x0a, 0x05, 0x6c, 0xc3, 0x53, 0xd1, 0xc7, 0xca, 0x5a, 0x03, 0xba, 0x0e, 0x41, 0x14,
	0xb0, 0x66, 0x19, 0xe8, 0x3a, 0xf8, 0xa6, 0xd1, 0xf8, 0x34, 0x88, 0xfc, 0x94, 0xac, 0x52, 0xc0,
	0xac, 0x2e, 0xe8, 0x68, 0x85, 0x0f, 0x73, 0x2b, 0x66, 0xa6, 0xd5, 0x61, 0x34, 0x1d, 0x04, 0x9c,
	0xac, 0x51, 0xc0, 0x40, 0x17, 0xba, 0x6d, 0xa3, 0xdc, 0x63, 0xe1, 0xc0, 0xb7, 0x90, 0x0a, 0xec,
	0x0f, 0xa2, 0x28, 0x20, 0x0d, 0x0a, 0xd8, 0x72, 0x77, 0x21, 0x4d, 0xa6, 0xdc, 0x43, 0xd2, 0xd1,
	0x8b, 0xa2, 0x00, 0xdf, 0x9d, 0x69, 0x97, 0x26, 0xa3, 0xf0, 0x8c, 0xac, 0x53, 0xc0, 0x56, 0xba,
	0x8b, 0x7c, 0x1c, 0xa7, 0x99, 0x91, 0x50, 0xb8, 0xf0, 0x9e, 0x29, 0x71, 0x90, 0xa5, 0x7c, 0x42,
	0x36, 0x28, 0x60, 0x6b, 0xdd, 0x85, 0x30, 0x0a, 0xb9, 0xae, 0xb2, 0x97, 0x3b, 0x70, 0x07, 0xa9,
	0x14, 0x7d, 0x3f, 0xcc, 0xc8, 0x90, 0x02, 0xb6, 0x6a, 0x6f, 0xb7, 0xd5, 0xdc, 0xe8, 0xdb, 0xd9,
	0x3e, 0x08, 0x33, 0x6f, 0x45, 0xe2, 0x0e, 0xc2, 0x0c, 0xf7, 0x8c, 0x52, 0x7a, 0x2a, 0x08, 0x17,
	0x91, 0x57, 0x6a, 0x91, 0x8f, 0x15, 0x40, 0x8b, 0xa8, 0xdf, 0xf1, 0x13, 0xd4, 0x54, 0x1c, 0x66,
	0x36, 0xc8, 0xa9, 0x20, 0x69, 0xd5, 0x48, 0x5e, 0x6a, 0x84, 0xa7, 0xf2, 0x1a, 0x03, 0xfe, 0xbe,
	0x28, 0xca, 0xf4, 0x24, 0x25, 0x67, 0x82, 0xe3, 0x72, 0x8d, 0xe3, 0x58, 0xb8, 0x0b, 0x32, 0x4d,
	0x4f, 0x52, 0xfc, 0x9d, 0x91, 0xe9, 0xad, 0x1f, 0x4c, 0x39, 0x79, 0x2d, 0x82, 0x2f, 0xd5, 0x82,
	0x7f, 0xcf, 0xbd, 0x5a, 0x38, 0xf1, 0x82, 0x9f, 0xa1, 0xed, 0xe2, 0xa0, 0xf5, 0xd5, 0x00, 0x93,
	0x91, 0xa0, 0xf8, 0xaa, 0x46, 0x21, 0xc6, 0x4e, 0xf2, 0xe0, 0xc2, 0x20, 0xfe, 0x21, 0xc3, 0xca,
	0x74, 0x1d, 0xdb, 0xd0, 0xfd, 0x35, 0x9f, 0xae, 0x63, 0x57, 0xe9, 0x3a, 0xb6, 0xa6, 0x7b, 0x81,
	0x76, 0x4a, 0x97, 0xcf, 0xf0, 0xbd, 0x11, 0x7c, 0x57, 0xeb, 0xe7, 0x24, 0x60, 0x92, 0x70, 0xab,
	0x78, 0x3b, 0xeb, 0x05, 0x8a, 0x4b, 0x6f, 0x08, 0x83, 0x39, 0x05, 0x8a, 0x21, 0x28, 0x15, 0x28,
	0x2c, 0xf5, 0x02, 0xa7, 0x65, 0xfd, 0xc6, 0x73, 0x0a, 0x7c, 0x55, 0x10, 0x70, 0xab, 0xf8, 0xe1,
	0xf9, 0x38, 0x63, 0x41, 0xc2, 0xf0, 0x13, 0x8c, 0x5a, 0xc3, 0xad, 0xe2, 0x37, 0xaf, 0xce, 0x28,
	0x07, 0xce, 0x30, 0x46, 0x73, 0x18, 0xe5, 0xf0, 0x95, 0x18, 0xa5, 0x49, 0x33, 0xfe, 0x82, 0xb6,
	0x0a, 0x93, 0x6e, 0xf8, 0xe2, 0x39, 0xf7, 0x3e, 0x1f, 0x7b, 0xc9, 0xb6, 0x39, 0xfb, 0x0e, 0xd4,
	0x0f, 0x44, 0x8c, 0xb8, 0x21, 0xfb, 0x7b, 0xce, 0x81, 0x88, 0x79, 0x2f, 0x1d, 0x88, 0xb0, 0x68,
	0xba, 0xd8, 0x34, 0x1b, 0xf2, 0x49, 0xca, 0x87, 0xfd, 0x31, 0x9f, 0x4c, 0xfc, 0x33, 0x4e, 0x76,
	0x04, 0x5f, 0xb7, 0xfd, 0xe9, 0x5d, 0xda, 0x2e, 0x6e, 0xab, 0xf6, 0x6f, 0x82, 0xe2, 0x99, 0x64,
	0x78, 0x7a, 0x41, 0x8b, 0x51, 0x32, 0xe3, 0x00, 0xe1, 0x72, 0x46, 0x1e, 0x4e, 0xc7, 0xe4, 0x12,
	0x05, 0x6c, 0xdd, 0x7e, 0xf4, 0x05, 0xe9, 0x9e, 0x84, 0xd3, 0x71, 0x17, 0xf6, 0x0e, 0xbc, 0xa7,
	0x17, 0xbc, 0x66, 0x31, 0x61, 0xee, 0xc8, 0xb7, 0x45, 0xc2, 0x63, 0xee, 0xe7, 0x89, 0xe4, 0x7e,
	0xba, 0x4e, 0x21, 0x5b, 0xf4, 0x1a, 0xda, 0x2a, 0xf7, 0x67, 0x05, 0xe6, 0x3a, 0x84, 0x52, 0xc8,
	0x60, 0x09, 0x26, 0x97, 0x8a, 0x81, 0xa9, 0x05, 0xfa, 0x35, 0x85, 0xac, 0xe1, 0x99, 0x68, 0xb3,
	0x3e, 0xcb, 0x40, 0xd7, 0x21, 0xbb, 0x14, 0xb2, 0x85, 0x32, 0xb0, 0xc2, 0xa8, 0x16, 0xe8, 0x0d,
	0x0a, 0xd9, 0xe6, 0x0c, 0x78, 0x5c, 0x67, 0x54, 0x4b, 0xf4, 0x26, 0x85, 0x0c, 0x97, 0x81, 0xae,
	0x83, 0xef, 0xa0, 0xa6, 0x01, 0xea, 0xcd, 0x77, 0x8b, 0x42, 0xb6, 0xe4, 0x19, 0x02, 0xbd, 0x4a,
	0x6b, 0x50, 0xd7, 0x21, 0xb7, 0x29, 0x64, 0x17, 0x2b, 0x50, 0xd7, 0xc1, 0xf7, 0xd0, 0xe6, 0x2c,
	0xbd, 0xa6, 0xdd, 0xa3, 0x90, 0x6d, 0x78, 0x86, 0xc3, 0xac, 0xd4, 0x3a, 0xd8, 0x75, 0x08, 0xa3,
	0x90, 0x35, 0xab, 0x60, 0xb9, 0xf6, 0x67, 0x45, 0x88, 0xc5, 0x7a, 0x87, 0x42, 0x66, 0xcd, 0xa4,
	0x97, 0x7b, 0xb5, 0xd8, 0xbf, 0xda, 0xac, 0x77, 0x29, 0x64, 0x60, 0xd6, 0xbf, 0x5a, 0xab, 0x37,
	0x90, 0x89, 0x94, 0x8b, 0xf5, 0x1e, 0x85, 0x6c, 0xd9, 0x5b, 0xd3, 0x46, 0xb1, 0x54, 0x4b, 0x6a,
	0xca, 0xb5, 0x7a, 0x9f, 0x42, 0xb6, 0x52, 0x50, 0x53, 0x6e, 0xd4, 0x62, 0x75, 0x72, 0xa7, 0x3e,
	0xa0, 0x90, 0xad, 0xcd, 0xaa, 0x93, 0xfb, 0x34, 0x41, 0x97, 0x0d, 0xac, 0x32, 0x48, 0x1d, 0x0a,
	0xcf, 0x37, 0x48, 0xde, 0x8e, 0xa6, 0x2e, 0x0f, 0xd2, 0x08, 0x6d, 0x57, 0x73, 0x8a, 0x51, 0x72,
	0x28, 0x3c, 0xc7, 0x28, 0x79, 0xb8, 0x9c, 0x4d, 0x4c, 0x91, 0x8b, 0x76, 0x2a, 0x72, 0xf5, 0xe3,
	0x11, 0x3f, 0xe1, 0xe4, 0x9b, 0x5c, 0xb4, 0x9e, 0xb5, 0x6c, 0x79, 0x5b, 0x65, 0xe1, 0x5e, 0xe4,
	0xee, 0xfc, 0xbf, 0x48, 0x13, 0x77, 0x12, 0x25, 0x43, 0xe2, 0x2a, 0x3c, 0x98, 0x9d, 0xc7, 0x4f,
	0x51, 0x32, 0xc4, 0x49, 0x21, 0x41, 0xe0, 0x7f, 0xc8, 0x8c, 0x7a, 0x8f, 0xce, 0xab, 0x5e, 0xcf,
	0x62, 0x60, 0x56, 0xdc, 0xaf, 0xfe, 0x87, 0x4c, 0xeb, 0x37, 0x46, 0x9b, 0x63, 0x3f, 0xd6, 0xfd,
	0xa8, 0x5b, 0xf0, 0x83, 0xc8, 0x77, 0xf0, 0x59, 0xf9, 0x9e, 0xf9, 0xb1, 0x6c, 0x5a, 0xfe, 0x7d,
	0x12, 0xa6, 0x49, 0xe6, 0x6d, 0x8c, 0xcb, 0x56, 0xfc, 0x0e, 0xed, 0xe4, 0xe9, 0xe4, 0xda, 0x53,
	0xe7, 0x95, 0xff, 0x6a, 0x21, 0x3f, 0x8a, 0x94, 0x8f, 0x3f, 0x37, 0xa5, 0xf8, 0x22, 0xc9, 0x56,
	0x73, 0x9b, 0xcc, 0x8a, 0xc7, 0x35, 0x47, 0xeb, 0x3a, 0x6a, 0x94, 0x2f, 0xce, 0x3a, 0xb2, 0x06,
	0x03, 0xf9, 0x53, 0xc2, 0xb3, 0x06, 0x83, 0x56, 0x0f, 0x6d, 0x7f, 0xac, 0x05, 0xdc, 0x44, 0xf0,
	0x0d, 0xcf, 0x04, 0x70, 0xc5, 0xcb, 0x1f, 0xf1, 0x36, 0x5a, 0x94, 0xff, 0x31, 0x59, 0xc2, 0x26,
	0x5f, 0xba, 0xd6, 0xb7, 0xa0, 0x95, 0xa1, 0xcb, 0x73, 0x6a, 0x2a, 0xd2, 0x40, 0x49, 0xf3, 0xb4,
	0x48, 0xb3, 0x6a, 0xdb, 0xff, 0xd7, 0xba, 0x62, 0x2c, 0x08, 0x50, 0x48, 0xbd, 0xbb, 0x87, 0x50,
	0xe1, 0xaa, 0x2e, 0x21, 0x78, 0xf8, 0xfc, 0x79, 0xf3, 0x42, 0xfe, 0xd0, 0x3b, 0xf0, 0x9a, 0x40,
	0x3e, 0xfc, 0xd9, 0xb4, 0x7a, 0x0d, 0xb4, 0x5a, 0xd0, 0x7d, 0xf7, 0x1f, 0x80, 0x70, 0x9d, 0x39,
	0x2f, 0xee, 0xe4, 0xf5, 0x28, 0x18, 0x8a, 0x82, 0xbf, 0xb0, 0x38, 0x41, 0x80, 0x0f, 0xd1, 0x52,
	0xec, 0x67, 0x41, 0xe4, 0x0f, 0x55, 0xa3, 0xf7, 0x3f, 0xe7, 0x8c, 0x3d, 0x1d, 0x7c, 0x77, 0x0f,
	0xa1, 0xa3, 0x20, 0x1a, 0xf8, 0x81, 0x6e, 0xf0, 0x48, 0x37, 0x78, 0xa4, 0x1b, 0x3c, 0xca, 0x1b,
	0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x63, 0x77, 0x31, 0x76, 0x24, 0x0f, 0x00, 0x00,
}
