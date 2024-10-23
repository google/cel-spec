// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.0
// source: cel/expr/conformance/test/simple.proto

package test

import (
	expr "cel.dev/expr"
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

type SimpleTestFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Section     []*SimpleTestSection `protobuf:"bytes,3,rep,name=section,proto3" json:"section,omitempty"`
}

func (x *SimpleTestFile) Reset() {
	*x = SimpleTestFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleTestFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleTestFile) ProtoMessage() {}

func (x *SimpleTestFile) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleTestFile.ProtoReflect.Descriptor instead.
func (*SimpleTestFile) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_simple_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleTestFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleTestFile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SimpleTestFile) GetSection() []*SimpleTestSection {
	if x != nil {
		return x.Section
	}
	return nil
}

type SimpleTestSection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Test        []*SimpleTest `protobuf:"bytes,3,rep,name=test,proto3" json:"test,omitempty"`
}

func (x *SimpleTestSection) Reset() {
	*x = SimpleTestSection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleTestSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleTestSection) ProtoMessage() {}

func (x *SimpleTestSection) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleTestSection.ProtoReflect.Descriptor instead.
func (*SimpleTestSection) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_simple_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleTestSection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleTestSection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SimpleTestSection) GetTest() []*SimpleTest {
	if x != nil {
		return x.Test
	}
	return nil
}

type SimpleTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Expr          string                     `protobuf:"bytes,3,opt,name=expr,proto3" json:"expr,omitempty"`
	DisableMacros bool                       `protobuf:"varint,4,opt,name=disable_macros,json=disableMacros,proto3" json:"disable_macros,omitempty"`
	DisableCheck  bool                       `protobuf:"varint,5,opt,name=disable_check,json=disableCheck,proto3" json:"disable_check,omitempty"`
	TypeEnv       []*expr.Decl               `protobuf:"bytes,6,rep,name=type_env,json=typeEnv,proto3" json:"type_env,omitempty"`
	Container     string                     `protobuf:"bytes,13,opt,name=container,proto3" json:"container,omitempty"`
	Locale        string                     `protobuf:"bytes,14,opt,name=locale,proto3" json:"locale,omitempty"`
	Bindings      map[string]*expr.ExprValue `protobuf:"bytes,7,rep,name=bindings,proto3" json:"bindings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to ResultMatcher:
	//
	//	*SimpleTest_Value
	//	*SimpleTest_EvalError
	//	*SimpleTest_AnyEvalErrors
	//	*SimpleTest_Unknown
	//	*SimpleTest_AnyUnknowns
	ResultMatcher isSimpleTest_ResultMatcher `protobuf_oneof:"result_matcher"`
}

func (x *SimpleTest) Reset() {
	*x = SimpleTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleTest) ProtoMessage() {}

func (x *SimpleTest) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleTest.ProtoReflect.Descriptor instead.
func (*SimpleTest) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_simple_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleTest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleTest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SimpleTest) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *SimpleTest) GetDisableMacros() bool {
	if x != nil {
		return x.DisableMacros
	}
	return false
}

func (x *SimpleTest) GetDisableCheck() bool {
	if x != nil {
		return x.DisableCheck
	}
	return false
}

func (x *SimpleTest) GetTypeEnv() []*expr.Decl {
	if x != nil {
		return x.TypeEnv
	}
	return nil
}

func (x *SimpleTest) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *SimpleTest) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *SimpleTest) GetBindings() map[string]*expr.ExprValue {
	if x != nil {
		return x.Bindings
	}
	return nil
}

func (m *SimpleTest) GetResultMatcher() isSimpleTest_ResultMatcher {
	if m != nil {
		return m.ResultMatcher
	}
	return nil
}

func (x *SimpleTest) GetValue() *expr.Value {
	if x, ok := x.GetResultMatcher().(*SimpleTest_Value); ok {
		return x.Value
	}
	return nil
}

func (x *SimpleTest) GetEvalError() *expr.ErrorSet {
	if x, ok := x.GetResultMatcher().(*SimpleTest_EvalError); ok {
		return x.EvalError
	}
	return nil
}

func (x *SimpleTest) GetAnyEvalErrors() *ErrorSetMatcher {
	if x, ok := x.GetResultMatcher().(*SimpleTest_AnyEvalErrors); ok {
		return x.AnyEvalErrors
	}
	return nil
}

func (x *SimpleTest) GetUnknown() *expr.UnknownSet {
	if x, ok := x.GetResultMatcher().(*SimpleTest_Unknown); ok {
		return x.Unknown
	}
	return nil
}

func (x *SimpleTest) GetAnyUnknowns() *UnknownSetMatcher {
	if x, ok := x.GetResultMatcher().(*SimpleTest_AnyUnknowns); ok {
		return x.AnyUnknowns
	}
	return nil
}

type isSimpleTest_ResultMatcher interface {
	isSimpleTest_ResultMatcher()
}

type SimpleTest_Value struct {
	Value *expr.Value `protobuf:"bytes,8,opt,name=value,proto3,oneof"`
}

type SimpleTest_EvalError struct {
	EvalError *expr.ErrorSet `protobuf:"bytes,9,opt,name=eval_error,json=evalError,proto3,oneof"`
}

type SimpleTest_AnyEvalErrors struct {
	AnyEvalErrors *ErrorSetMatcher `protobuf:"bytes,10,opt,name=any_eval_errors,json=anyEvalErrors,proto3,oneof"`
}

type SimpleTest_Unknown struct {
	Unknown *expr.UnknownSet `protobuf:"bytes,11,opt,name=unknown,proto3,oneof"`
}

type SimpleTest_AnyUnknowns struct {
	AnyUnknowns *UnknownSetMatcher `protobuf:"bytes,12,opt,name=any_unknowns,json=anyUnknowns,proto3,oneof"`
}

func (*SimpleTest_Value) isSimpleTest_ResultMatcher() {}

func (*SimpleTest_EvalError) isSimpleTest_ResultMatcher() {}

func (*SimpleTest_AnyEvalErrors) isSimpleTest_ResultMatcher() {}

func (*SimpleTest_Unknown) isSimpleTest_ResultMatcher() {}

func (*SimpleTest_AnyUnknowns) isSimpleTest_ResultMatcher() {}

type ErrorSetMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []*expr.ErrorSet `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ErrorSetMatcher) Reset() {
	*x = ErrorSetMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorSetMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorSetMatcher) ProtoMessage() {}

func (x *ErrorSetMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorSetMatcher.ProtoReflect.Descriptor instead.
func (*ErrorSetMatcher) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_simple_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorSetMatcher) GetErrors() []*expr.ErrorSet {
	if x != nil {
		return x.Errors
	}
	return nil
}

type UnknownSetMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unknowns []*expr.UnknownSet `protobuf:"bytes,1,rep,name=unknowns,proto3" json:"unknowns,omitempty"`
}

func (x *UnknownSetMatcher) Reset() {
	*x = UnknownSetMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnknownSetMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownSetMatcher) ProtoMessage() {}

func (x *UnknownSetMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_simple_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownSetMatcher.ProtoReflect.Descriptor instead.
func (*UnknownSetMatcher) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_simple_proto_rawDescGZIP(), []int{4}
}

func (x *UnknownSetMatcher) GetUnknowns() []*expr.UnknownSet {
	if x != nil {
		return x.Unknowns
	}
	return nil
}

var File_cel_expr_conformance_test_simple_proto protoreflect.FileDescriptor

var file_cel_expr_conformance_test_simple_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78,
	0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x65, 0x6c,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x46, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0xf1,
	0x05, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72,
	0x2e, 0x44, 0x65, 0x63, 0x6c, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x76, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33,
	0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x09, 0x65, 0x76, 0x61, 0x6c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x54, 0x0a, 0x0f, 0x61, 0x6e, 0x79, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63,
	0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x6e, 0x79, 0x45,
	0x76, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x65, 0x6c,
	0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x51, 0x0a, 0x0c, 0x61,
	0x6e, 0x79, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x0b, 0x61, 0x6e, 0x79, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x73, 0x1a, 0x50,
	0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x22, 0x45, 0x0a, 0x11, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x08,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x73, 0x42, 0x4b, 0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e,
	0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x63, 0x65, 0x6c, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x65, 0x78, 0x70,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cel_expr_conformance_test_simple_proto_rawDescOnce sync.Once
	file_cel_expr_conformance_test_simple_proto_rawDescData = file_cel_expr_conformance_test_simple_proto_rawDesc
)

func file_cel_expr_conformance_test_simple_proto_rawDescGZIP() []byte {
	file_cel_expr_conformance_test_simple_proto_rawDescOnce.Do(func() {
		file_cel_expr_conformance_test_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_cel_expr_conformance_test_simple_proto_rawDescData)
	})
	return file_cel_expr_conformance_test_simple_proto_rawDescData
}

var file_cel_expr_conformance_test_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cel_expr_conformance_test_simple_proto_goTypes = []any{
	(*SimpleTestFile)(nil),    // 0: cel.expr.conformance.test.SimpleTestFile
	(*SimpleTestSection)(nil), // 1: cel.expr.conformance.test.SimpleTestSection
	(*SimpleTest)(nil),        // 2: cel.expr.conformance.test.SimpleTest
	(*ErrorSetMatcher)(nil),   // 3: cel.expr.conformance.test.ErrorSetMatcher
	(*UnknownSetMatcher)(nil), // 4: cel.expr.conformance.test.UnknownSetMatcher
	nil,                       // 5: cel.expr.conformance.test.SimpleTest.BindingsEntry
	(*expr.Decl)(nil),         // 6: cel.expr.Decl
	(*expr.Value)(nil),        // 7: cel.expr.Value
	(*expr.ErrorSet)(nil),     // 8: cel.expr.ErrorSet
	(*expr.UnknownSet)(nil),   // 9: cel.expr.UnknownSet
	(*expr.ExprValue)(nil),    // 10: cel.expr.ExprValue
}
var file_cel_expr_conformance_test_simple_proto_depIdxs = []int32{
	1,  // 0: cel.expr.conformance.test.SimpleTestFile.section:type_name -> cel.expr.conformance.test.SimpleTestSection
	2,  // 1: cel.expr.conformance.test.SimpleTestSection.test:type_name -> cel.expr.conformance.test.SimpleTest
	6,  // 2: cel.expr.conformance.test.SimpleTest.type_env:type_name -> cel.expr.Decl
	5,  // 3: cel.expr.conformance.test.SimpleTest.bindings:type_name -> cel.expr.conformance.test.SimpleTest.BindingsEntry
	7,  // 4: cel.expr.conformance.test.SimpleTest.value:type_name -> cel.expr.Value
	8,  // 5: cel.expr.conformance.test.SimpleTest.eval_error:type_name -> cel.expr.ErrorSet
	3,  // 6: cel.expr.conformance.test.SimpleTest.any_eval_errors:type_name -> cel.expr.conformance.test.ErrorSetMatcher
	9,  // 7: cel.expr.conformance.test.SimpleTest.unknown:type_name -> cel.expr.UnknownSet
	4,  // 8: cel.expr.conformance.test.SimpleTest.any_unknowns:type_name -> cel.expr.conformance.test.UnknownSetMatcher
	8,  // 9: cel.expr.conformance.test.ErrorSetMatcher.errors:type_name -> cel.expr.ErrorSet
	9,  // 10: cel.expr.conformance.test.UnknownSetMatcher.unknowns:type_name -> cel.expr.UnknownSet
	10, // 11: cel.expr.conformance.test.SimpleTest.BindingsEntry.value:type_name -> cel.expr.ExprValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_cel_expr_conformance_test_simple_proto_init() }
func file_cel_expr_conformance_test_simple_proto_init() {
	if File_cel_expr_conformance_test_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cel_expr_conformance_test_simple_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SimpleTestFile); i {
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
		file_cel_expr_conformance_test_simple_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SimpleTestSection); i {
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
		file_cel_expr_conformance_test_simple_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SimpleTest); i {
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
		file_cel_expr_conformance_test_simple_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ErrorSetMatcher); i {
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
		file_cel_expr_conformance_test_simple_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UnknownSetMatcher); i {
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
	file_cel_expr_conformance_test_simple_proto_msgTypes[2].OneofWrappers = []any{
		(*SimpleTest_Value)(nil),
		(*SimpleTest_EvalError)(nil),
		(*SimpleTest_AnyEvalErrors)(nil),
		(*SimpleTest_Unknown)(nil),
		(*SimpleTest_AnyUnknowns)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cel_expr_conformance_test_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cel_expr_conformance_test_simple_proto_goTypes,
		DependencyIndexes: file_cel_expr_conformance_test_simple_proto_depIdxs,
		MessageInfos:      file_cel_expr_conformance_test_simple_proto_msgTypes,
	}.Build()
	File_cel_expr_conformance_test_simple_proto = out.File
	file_cel_expr_conformance_test_simple_proto_rawDesc = nil
	file_cel_expr_conformance_test_simple_proto_goTypes = nil
	file_cel_expr_conformance_test_simple_proto_depIdxs = nil
}
