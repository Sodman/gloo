// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: function_spec.proto

package aws

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Function Spec for Functions on AWS Lambda Upstreams
type FunctionSpec struct {
	// The Name of the Lambda Function as it appears in the AWS Lambda Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// The Qualifier for the Lambda Function. Qualifiers act as a kind of version
	// for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info.
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
}

func (m *FunctionSpec) Reset()                    { *m = FunctionSpec{} }
func (m *FunctionSpec) String() string            { return proto.CompactTextString(m) }
func (*FunctionSpec) ProtoMessage()               {}
func (*FunctionSpec) Descriptor() ([]byte, []int) { return fileDescriptorFunctionSpec, []int{0} }

func (m *FunctionSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *FunctionSpec) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionSpec)(nil), "gloo.api.v1.FunctionSpec")
}
func (this *FunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionSpec)
	if !ok {
		that2, ok := that.(FunctionSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.Qualifier != that1.Qualifier {
		return false
	}
	return true
}

func init() { proto.RegisterFile("function_spec.proto", fileDescriptorFunctionSpec) }

var fileDescriptorFunctionSpec = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2b, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x2e, 0x48, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4e, 0xcf, 0xc9, 0xcf, 0xd7, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58, 0x10, 0x25, 0x4a, 0x81, 0x5c, 0x3c, 0x6e, 0x50, 0x9d,
	0xc1, 0x05, 0xa9, 0xc9, 0x42, 0xca, 0x5c, 0xbc, 0x70, 0x93, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0x60, 0x82, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0x32, 0x5c, 0x9c,
	0x85, 0xa5, 0x89, 0x39, 0x99, 0x69, 0x99, 0xa9, 0x45, 0x12, 0x4c, 0x60, 0x05, 0x08, 0x01, 0x27,
	0xdd, 0x15, 0x8f, 0xe4, 0x18, 0xa3, 0xd4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x8b, 0xf3, 0x73, 0xf2, 0x75, 0x33, 0xf3, 0xf5, 0x41, 0xce, 0xd1, 0x2f, 0xc8, 0x4e,
	0xd7, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0x4f, 0x2c, 0x2f, 0x4e, 0x62, 0x03, 0x3b,
	0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x7c, 0xd2, 0x23, 0xc2, 0x00, 0x00, 0x00,
}