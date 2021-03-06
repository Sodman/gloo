// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/wasm/v3/wasm.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
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

// Configuration for a Wasm VM.
// [#next-free-field: 7]
type VmConfig struct {
	// An ID which will be used along with a hash of the wasm code (or the name of the registered Null
	// VM plugin) to determine which VM will be used for the plugin. All plugins which use the same
	// *vm_id* and code will use the same VM. May be left blank. Sharing a VM between plugins can
	// reduce memory utilization and make sharing of data easier which may have security implications.
	// See ref: "TODO: add ref" for details.
	VmId string `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// The Wasm runtime type (either "v8" or "null" for code compiled into Envoy).
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The Wasm code that Envoy will execute.
	Code *v3.AsyncDataSource `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The Wasm configuration used in initialization of a new VM
	// (proxy_on_start). `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *types.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Allow the wasm file to include pre-compiled code on VMs which support it.
	// Warning: this should only be enable for trusted sources as the precompiled code is not
	// verified.
	AllowPrecompiled bool `protobuf:"varint,5,opt,name=allow_precompiled,json=allowPrecompiled,proto3" json:"allow_precompiled,omitempty"`
	// If true and the code needs to be remotely fetched and it is not in the cache then NACK the configuration
	// update and do a background fetch to fill the cache, otherwise fetch the code asynchronously and enter
	// warming state.
	NackOnCodeCacheMiss  bool     `protobuf:"varint,6,opt,name=nack_on_code_cache_miss,json=nackOnCodeCacheMiss,proto3" json:"nack_on_code_cache_miss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmConfig) Reset()         { *m = VmConfig{} }
func (m *VmConfig) String() string { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()    {}
func (*VmConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_300f8be9338032c2, []int{0}
}
func (m *VmConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmConfig.Unmarshal(m, b)
}
func (m *VmConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmConfig.Marshal(b, m, deterministic)
}
func (m *VmConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmConfig.Merge(m, src)
}
func (m *VmConfig) XXX_Size() int {
	return xxx_messageInfo_VmConfig.Size(m)
}
func (m *VmConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VmConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VmConfig proto.InternalMessageInfo

func (m *VmConfig) GetVmId() string {
	if m != nil {
		return m.VmId
	}
	return ""
}

func (m *VmConfig) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *VmConfig) GetCode() *v3.AsyncDataSource {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *VmConfig) GetConfiguration() *types.Any {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *VmConfig) GetAllowPrecompiled() bool {
	if m != nil {
		return m.AllowPrecompiled
	}
	return false
}

func (m *VmConfig) GetNackOnCodeCacheMiss() bool {
	if m != nil {
		return m.NackOnCodeCacheMiss
	}
	return false
}

// Base Configuration for Wasm Plugins e.g. filters and services.
// [#next-free-field: 6]
type PluginConfig struct {
	// A unique name for a filters/services in a VM for use in identifying the filter/service if
	// multiple filters/services are handled by the same *vm_id* and *root_id* and for
	// logging/debugging.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a set of filters/services in a VM which will share a RootContext and Contexts
	// if applicable (e.g. an Wasm HttpFilter and an Wasm AccessLog). If left blank, all
	// filters/services with a blank root_id with the same *vm_id* will share Context(s).
	RootId string `protobuf:"bytes,2,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// Configuration for finding or starting VM.
	//
	// Types that are valid to be assigned to Vm:
	//	*PluginConfig_VmConfig
	Vm isPluginConfig_Vm `protobuf_oneof:"vm"`
	// Filter/service configuration used to configure or reconfigure a plugin
	// (proxy_on_configuration).
	// `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *types.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// If there is a fatal error on the VM (e.g. exception, abort(), on_start or on_configure return false),
	// then all plugins associated with the VM will either fail closed (by default), e.g. by returning an HTTP 503 error,
	// or fail open (if 'fail_open' is set to true) by bypassing the filter. Note: when on_start or on_configure return false
	// during xDS updates the xDS configuration will be rejected and when on_start or on_configuration return false on initial
	// startup the proxy will not start.
	FailOpen             bool     `protobuf:"varint,5,opt,name=fail_open,json=failOpen,proto3" json:"fail_open,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginConfig) Reset()         { *m = PluginConfig{} }
func (m *PluginConfig) String() string { return proto.CompactTextString(m) }
func (*PluginConfig) ProtoMessage()    {}
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_300f8be9338032c2, []int{1}
}
func (m *PluginConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginConfig.Unmarshal(m, b)
}
func (m *PluginConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginConfig.Marshal(b, m, deterministic)
}
func (m *PluginConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginConfig.Merge(m, src)
}
func (m *PluginConfig) XXX_Size() int {
	return xxx_messageInfo_PluginConfig.Size(m)
}
func (m *PluginConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PluginConfig proto.InternalMessageInfo

type isPluginConfig_Vm interface {
	isPluginConfig_Vm()
	Equal(interface{}) bool
}

type PluginConfig_VmConfig struct {
	VmConfig *VmConfig `protobuf:"bytes,3,opt,name=vm_config,json=vmConfig,proto3,oneof" json:"vm_config,omitempty"`
}

func (*PluginConfig_VmConfig) isPluginConfig_Vm() {}

func (m *PluginConfig) GetVm() isPluginConfig_Vm {
	if m != nil {
		return m.Vm
	}
	return nil
}

func (m *PluginConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginConfig) GetRootId() string {
	if m != nil {
		return m.RootId
	}
	return ""
}

func (m *PluginConfig) GetVmConfig() *VmConfig {
	if x, ok := m.GetVm().(*PluginConfig_VmConfig); ok {
		return x.VmConfig
	}
	return nil
}

func (m *PluginConfig) GetConfiguration() *types.Any {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *PluginConfig) GetFailOpen() bool {
	if m != nil {
		return m.FailOpen
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PluginConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PluginConfig_VmConfig)(nil),
	}
}

// WasmService is configured as a built-in *envoy.wasm_service* :ref:`WasmService
// <config_wasm_service>` This opaque configuration will be used to create a Wasm Service.
type WasmService struct {
	// General plugin configuration.
	Config *PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// If true, create a single VM rather than creating one VM per worker. Such a singleton can
	// not be used with filters.
	Singleton            bool     `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WasmService) Reset()         { *m = WasmService{} }
func (m *WasmService) String() string { return proto.CompactTextString(m) }
func (*WasmService) ProtoMessage()    {}
func (*WasmService) Descriptor() ([]byte, []int) {
	return fileDescriptor_300f8be9338032c2, []int{2}
}
func (m *WasmService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmService.Unmarshal(m, b)
}
func (m *WasmService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmService.Marshal(b, m, deterministic)
}
func (m *WasmService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmService.Merge(m, src)
}
func (m *WasmService) XXX_Size() int {
	return xxx_messageInfo_WasmService.Size(m)
}
func (m *WasmService) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmService.DiscardUnknown(m)
}

var xxx_messageInfo_WasmService proto.InternalMessageInfo

func (m *WasmService) GetConfig() *PluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *WasmService) GetSingleton() bool {
	if m != nil {
		return m.Singleton
	}
	return false
}

func init() {
	proto.RegisterType((*VmConfig)(nil), "envoy.extensions.wasm.v3.VmConfig")
	proto.RegisterType((*PluginConfig)(nil), "envoy.extensions.wasm.v3.PluginConfig")
	proto.RegisterType((*WasmService)(nil), "envoy.extensions.wasm.v3.WasmService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/wasm/v3/wasm.proto", fileDescriptor_300f8be9338032c2)
}

var fileDescriptor_300f8be9338032c2 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xc1, 0x4d, 0x93, 0x2d, 0x48, 0xb0, 0xad, 0x54, 0x53, 0x10, 0x84, 0x48, 0x54, 0x95,
	0x10, 0xbb, 0x12, 0xe1, 0x02, 0x07, 0xa4, 0xa6, 0x1c, 0xe8, 0x01, 0x25, 0x72, 0xa5, 0x22, 0x71,
	0xb1, 0x36, 0xeb, 0x8d, 0xbb, 0xc4, 0xbb, 0x63, 0xed, 0xda, 0x6e, 0xf3, 0x0b, 0x7c, 0x03, 0x1f,
	0xc0, 0x4f, 0x71, 0xe1, 0x33, 0x38, 0xa1, 0x5d, 0xdb, 0x0a, 0x1c, 0x22, 0x55, 0xe2, 0xe4, 0x99,
	0x79, 0xf3, 0x3c, 0x6f, 0xde, 0x68, 0xd1, 0x65, 0x26, 0xcb, 0xab, 0x6a, 0x41, 0x38, 0x28, 0x6a,
	0x21, 0x87, 0x57, 0x12, 0x68, 0x96, 0x03, 0xd0, 0xc2, 0xc0, 0x57, 0xc1, 0x4b, 0xdb, 0x64, 0xac,
	0x90, 0x54, 0xdc, 0x94, 0xc2, 0x68, 0x96, 0x53, 0xa1, 0x6b, 0x58, 0xfb, 0x54, 0x5b, 0x09, 0xda,
	0xd2, 0x6b, 0x66, 0x15, 0xad, 0x27, 0xfe, 0x4b, 0x0a, 0x03, 0x25, 0xe0, 0xc8, 0x37, 0x91, 0x4d,
	0x13, 0xf1, 0x60, 0x3d, 0x39, 0x7a, 0xd6, 0xd0, 0x39, 0xe8, 0xa5, 0xcc, 0x28, 0x07, 0x23, 0x1c,
	0x75, 0xc1, 0xac, 0x68, 0xa8, 0x47, 0x8f, 0x32, 0x80, 0x2c, 0x17, 0xd4, 0x67, 0x8b, 0x6a, 0x49,
	0x99, 0x5e, 0xb7, 0xd0, 0x61, 0xcd, 0x72, 0x99, 0xb2, 0x52, 0xd0, 0x2e, 0x68, 0x81, 0x83, 0x0c,
	0x32, 0xf0, 0x21, 0x75, 0x51, 0x53, 0x1d, 0x7f, 0xef, 0xa1, 0xc1, 0xa5, 0x3a, 0xf3, 0x93, 0xf0,
	0x3e, 0xda, 0xa9, 0x55, 0x22, 0xd3, 0x28, 0x18, 0x05, 0x27, 0xc3, 0x38, 0xac, 0xd5, 0x79, 0x8a,
	0x9f, 0xa3, 0x5d, 0x53, 0xe9, 0x52, 0x2a, 0x11, 0xf5, 0x5c, 0x79, 0xba, 0xfb, 0x7b, 0x1a, 0x9a,
	0xde, 0x28, 0x88, 0xbb, 0x3a, 0x7e, 0x8b, 0x42, 0x0e, 0xa9, 0x88, 0xee, 0x8e, 0x82, 0x93, 0xbd,
	0xd7, 0x2f, 0x48, 0xb3, 0x58, 0x23, 0x9f, 0x38, 0xf9, 0xa4, 0x9e, 0x90, 0x53, 0xbb, 0xd6, 0xfc,
	0x03, 0x2b, 0xd9, 0x05, 0x54, 0x86, 0x8b, 0xd8, 0x53, 0xf0, 0x3b, 0x74, 0xbf, 0xe9, 0xab, 0x0c,
	0x2b, 0x25, 0xe8, 0x28, 0xf4, 0xff, 0x38, 0x20, 0xcd, 0x86, 0xa4, 0xdb, 0x90, 0x9c, 0xea, 0x75,
	0xfc, 0x6f, 0x2b, 0x7e, 0x89, 0x1e, 0xb2, 0x3c, 0x87, 0xeb, 0xa4, 0x30, 0x82, 0x83, 0x2a, 0x64,
	0x2e, 0xd2, 0x68, 0x67, 0x14, 0x9c, 0x0c, 0xe2, 0x07, 0x1e, 0x98, 0x6f, 0xea, 0xf8, 0x0d, 0x3a,
	0xd4, 0x8c, 0xaf, 0x12, 0xd0, 0x89, 0x1b, 0x9c, 0x70, 0xc6, 0xaf, 0x44, 0xa2, 0xa4, 0xb5, 0x51,
	0xdf, 0x53, 0xf6, 0x1d, 0x3c, 0xd3, 0x67, 0x90, 0x8a, 0x33, 0x87, 0x7d, 0x92, 0xd6, 0x8e, 0x7f,
	0x06, 0xe8, 0xde, 0x3c, 0xaf, 0x32, 0xa9, 0x5b, 0x8b, 0x30, 0x0a, 0x35, 0x53, 0xa2, 0x73, 0xc8,
	0xc5, 0xf8, 0x10, 0xed, 0x1a, 0x80, 0xd2, 0x19, 0xe7, 0x1d, 0x8a, 0xfb, 0x2e, 0x3d, 0x4f, 0xf1,
	0x29, 0x1a, 0xd6, 0x2a, 0x69, 0x44, 0xb7, 0xe6, 0x8c, 0xc9, 0xb6, 0xab, 0x93, 0xee, 0x0c, 0x1f,
	0xef, 0xc4, 0x83, 0xba, 0x3b, 0xc9, 0xff, 0xf8, 0xf3, 0x18, 0x0d, 0x97, 0x4c, 0xe6, 0x09, 0x14,
	0x42, 0xb7, 0xbe, 0x0c, 0x5c, 0x61, 0x56, 0x08, 0x3d, 0x0d, 0x51, 0xaf, 0x56, 0xe3, 0x15, 0xda,
	0xfb, 0xcc, 0xac, 0xba, 0x10, 0xa6, 0x96, 0x5c, 0xe0, 0xf7, 0xa8, 0xdf, 0xaa, 0x0d, 0xfc, 0x98,
	0xe3, 0xed, 0x6a, 0xff, 0x76, 0x25, 0x6e, 0x59, 0xf8, 0x09, 0x1a, 0x5a, 0xa9, 0xb3, 0x5c, 0x94,
	0xa0, 0xbd, 0x17, 0x83, 0x78, 0x53, 0x98, 0x7e, 0x0b, 0x7e, 0xfc, 0x7a, 0x1a, 0xa0, 0x63, 0x09,
	0xcd, 0x6f, 0x0b, 0x03, 0x37, 0xeb, 0xad, 0x13, 0xa6, 0x43, 0xa7, 0x6c, 0xee, 0xf6, 0x9b, 0x07,
	0x5f, 0x66, 0xb7, 0x7b, 0x84, 0xc5, 0x2a, 0xbb, 0xdd, 0x43, 0x5c, 0xf4, 0xbd, 0x73, 0x93, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x57, 0xc1, 0xa8, 0xde, 0x03, 0x00, 0x00,
}

func (this *VmConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VmConfig)
	if !ok {
		that2, ok := that.(VmConfig)
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
	if this.VmId != that1.VmId {
		return false
	}
	if this.Runtime != that1.Runtime {
		return false
	}
	if !this.Code.Equal(that1.Code) {
		return false
	}
	if !this.Configuration.Equal(that1.Configuration) {
		return false
	}
	if this.AllowPrecompiled != that1.AllowPrecompiled {
		return false
	}
	if this.NackOnCodeCacheMiss != that1.NackOnCodeCacheMiss {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PluginConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PluginConfig)
	if !ok {
		that2, ok := that.(PluginConfig)
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
	if this.Name != that1.Name {
		return false
	}
	if this.RootId != that1.RootId {
		return false
	}
	if that1.Vm == nil {
		if this.Vm != nil {
			return false
		}
	} else if this.Vm == nil {
		return false
	} else if !this.Vm.Equal(that1.Vm) {
		return false
	}
	if !this.Configuration.Equal(that1.Configuration) {
		return false
	}
	if this.FailOpen != that1.FailOpen {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PluginConfig_VmConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PluginConfig_VmConfig)
	if !ok {
		that2, ok := that.(PluginConfig_VmConfig)
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
	if !this.VmConfig.Equal(that1.VmConfig) {
		return false
	}
	return true
}
func (this *WasmService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WasmService)
	if !ok {
		that2, ok := that.(WasmService)
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
	if !this.Config.Equal(that1.Config) {
		return false
	}
	if this.Singleton != that1.Singleton {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
