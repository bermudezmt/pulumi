// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource.proto

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _struct "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SupportsFeatureRequest allows a client to test if the resource monitor supports a certain feature, which it may use
// to control the format or types of messages it sends.
type SupportsFeatureRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportsFeatureRequest) Reset()         { *m = SupportsFeatureRequest{} }
func (m *SupportsFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*SupportsFeatureRequest) ProtoMessage()    {}
func (*SupportsFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{0}
}
func (m *SupportsFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportsFeatureRequest.Unmarshal(m, b)
}
func (m *SupportsFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportsFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *SupportsFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportsFeatureRequest.Merge(dst, src)
}
func (m *SupportsFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_SupportsFeatureRequest.Size(m)
}
func (m *SupportsFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportsFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupportsFeatureRequest proto.InternalMessageInfo

func (m *SupportsFeatureRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SupportsFeatureResponse struct {
	HasSupport           bool     `protobuf:"varint,1,opt,name=hasSupport" json:"hasSupport,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportsFeatureResponse) Reset()         { *m = SupportsFeatureResponse{} }
func (m *SupportsFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*SupportsFeatureResponse) ProtoMessage()    {}
func (*SupportsFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{1}
}
func (m *SupportsFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportsFeatureResponse.Unmarshal(m, b)
}
func (m *SupportsFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportsFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *SupportsFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportsFeatureResponse.Merge(dst, src)
}
func (m *SupportsFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_SupportsFeatureResponse.Size(m)
}
func (m *SupportsFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportsFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupportsFeatureResponse proto.InternalMessageInfo

func (m *SupportsFeatureResponse) GetHasSupport() bool {
	if m != nil {
		return m.HasSupport
	}
	return false
}

// ReadResourceRequest contains enough information to uniquely qualify and read a resource's state.
type ReadResourceRequest struct {
	Id                      string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                    string          `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Name                    string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Parent                  string          `protobuf:"bytes,4,opt,name=parent" json:"parent,omitempty"`
	Properties              *_struct.Struct `protobuf:"bytes,5,opt,name=properties" json:"properties,omitempty"`
	Dependencies            []string        `protobuf:"bytes,6,rep,name=dependencies" json:"dependencies,omitempty"`
	Provider                string          `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Version                 string          `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
	AcceptSecrets           bool            `protobuf:"varint,9,opt,name=acceptSecrets" json:"acceptSecrets,omitempty"`
	AdditionalSecretOutputs []string        `protobuf:"bytes,10,rep,name=additionalSecretOutputs" json:"additionalSecretOutputs,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}        `json:"-"`
	XXX_unrecognized        []byte          `json:"-"`
	XXX_sizecache           int32           `json:"-"`
}

func (m *ReadResourceRequest) Reset()         { *m = ReadResourceRequest{} }
func (m *ReadResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ReadResourceRequest) ProtoMessage()    {}
func (*ReadResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{2}
}
func (m *ReadResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResourceRequest.Unmarshal(m, b)
}
func (m *ReadResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResourceRequest.Marshal(b, m, deterministic)
}
func (dst *ReadResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResourceRequest.Merge(dst, src)
}
func (m *ReadResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ReadResourceRequest.Size(m)
}
func (m *ReadResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResourceRequest proto.InternalMessageInfo

func (m *ReadResourceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadResourceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReadResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadResourceRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReadResourceRequest) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ReadResourceRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *ReadResourceRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ReadResourceRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ReadResourceRequest) GetAcceptSecrets() bool {
	if m != nil {
		return m.AcceptSecrets
	}
	return false
}

func (m *ReadResourceRequest) GetAdditionalSecretOutputs() []string {
	if m != nil {
		return m.AdditionalSecretOutputs
	}
	return nil
}

// ReadResourceResponse contains the result of reading a resource's state.
type ReadResourceResponse struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Properties           *_struct.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadResourceResponse) Reset()         { *m = ReadResourceResponse{} }
func (m *ReadResourceResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResourceResponse) ProtoMessage()    {}
func (*ReadResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{3}
}
func (m *ReadResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResourceResponse.Unmarshal(m, b)
}
func (m *ReadResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResourceResponse.Marshal(b, m, deterministic)
}
func (dst *ReadResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResourceResponse.Merge(dst, src)
}
func (m *ReadResourceResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResourceResponse.Size(m)
}
func (m *ReadResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResourceResponse proto.InternalMessageInfo

func (m *ReadResourceResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ReadResourceResponse) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

// RegisterResourceRequest contains information about a resource object that was newly allocated.
type RegisterResourceRequest struct {
	Type                    string                                                   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name                    string                                                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Parent                  string                                                   `protobuf:"bytes,3,opt,name=parent" json:"parent,omitempty"`
	Custom                  bool                                                     `protobuf:"varint,4,opt,name=custom" json:"custom,omitempty"`
	Object                  *_struct.Struct                                          `protobuf:"bytes,5,opt,name=object" json:"object,omitempty"`
	Protect                 bool                                                     `protobuf:"varint,6,opt,name=protect" json:"protect,omitempty"`
	Dependencies            []string                                                 `protobuf:"bytes,7,rep,name=dependencies" json:"dependencies,omitempty"`
	Provider                string                                                   `protobuf:"bytes,8,opt,name=provider" json:"provider,omitempty"`
	PropertyDependencies    map[string]*RegisterResourceRequest_PropertyDependencies `protobuf:"bytes,9,rep,name=propertyDependencies" json:"propertyDependencies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DeleteBeforeReplace     bool                                                     `protobuf:"varint,10,opt,name=deleteBeforeReplace" json:"deleteBeforeReplace,omitempty"`
	Version                 string                                                   `protobuf:"bytes,11,opt,name=version" json:"version,omitempty"`
	IgnoreChanges           []string                                                 `protobuf:"bytes,12,rep,name=ignoreChanges" json:"ignoreChanges,omitempty"`
	AcceptSecrets           bool                                                     `protobuf:"varint,13,opt,name=acceptSecrets" json:"acceptSecrets,omitempty"`
	AdditionalSecretOutputs []string                                                 `protobuf:"bytes,14,rep,name=additionalSecretOutputs" json:"additionalSecretOutputs,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                                 `json:"-"`
	XXX_unrecognized        []byte                                                   `json:"-"`
	XXX_sizecache           int32                                                    `json:"-"`
}

func (m *RegisterResourceRequest) Reset()         { *m = RegisterResourceRequest{} }
func (m *RegisterResourceRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterResourceRequest) ProtoMessage()    {}
func (*RegisterResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{4}
}
func (m *RegisterResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceRequest.Unmarshal(m, b)
}
func (m *RegisterResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceRequest.Merge(dst, src)
}
func (m *RegisterResourceRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceRequest.Size(m)
}
func (m *RegisterResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceRequest proto.InternalMessageInfo

func (m *RegisterResourceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RegisterResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterResourceRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RegisterResourceRequest) GetCustom() bool {
	if m != nil {
		return m.Custom
	}
	return false
}

func (m *RegisterResourceRequest) GetObject() *_struct.Struct {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *RegisterResourceRequest) GetProtect() bool {
	if m != nil {
		return m.Protect
	}
	return false
}

func (m *RegisterResourceRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *RegisterResourceRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *RegisterResourceRequest) GetPropertyDependencies() map[string]*RegisterResourceRequest_PropertyDependencies {
	if m != nil {
		return m.PropertyDependencies
	}
	return nil
}

func (m *RegisterResourceRequest) GetDeleteBeforeReplace() bool {
	if m != nil {
		return m.DeleteBeforeReplace
	}
	return false
}

func (m *RegisterResourceRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegisterResourceRequest) GetIgnoreChanges() []string {
	if m != nil {
		return m.IgnoreChanges
	}
	return nil
}

func (m *RegisterResourceRequest) GetAcceptSecrets() bool {
	if m != nil {
		return m.AcceptSecrets
	}
	return false
}

func (m *RegisterResourceRequest) GetAdditionalSecretOutputs() []string {
	if m != nil {
		return m.AdditionalSecretOutputs
	}
	return nil
}

// PropertyDependencies describes the resources that a particular property depends on.
type RegisterResourceRequest_PropertyDependencies struct {
	Urns                 []string `protobuf:"bytes,1,rep,name=urns" json:"urns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResourceRequest_PropertyDependencies) Reset() {
	*m = RegisterResourceRequest_PropertyDependencies{}
}
func (m *RegisterResourceRequest_PropertyDependencies) String() string {
	return proto.CompactTextString(m)
}
func (*RegisterResourceRequest_PropertyDependencies) ProtoMessage() {}
func (*RegisterResourceRequest_PropertyDependencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{4, 0}
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Unmarshal(m, b)
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceRequest_PropertyDependencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Merge(dst, src)
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Size(m)
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceRequest_PropertyDependencies proto.InternalMessageInfo

func (m *RegisterResourceRequest_PropertyDependencies) GetUrns() []string {
	if m != nil {
		return m.Urns
	}
	return nil
}

// RegisterResourceResponse is returned by the engine after a resource has finished being initialized.  It includes the
// auto-assigned URN, the provider-assigned ID, and any other properties initialized by the engine.
type RegisterResourceResponse struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Object               *_struct.Struct `protobuf:"bytes,3,opt,name=object" json:"object,omitempty"`
	Stable               bool            `protobuf:"varint,4,opt,name=stable" json:"stable,omitempty"`
	Stables              []string        `protobuf:"bytes,5,rep,name=stables" json:"stables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterResourceResponse) Reset()         { *m = RegisterResourceResponse{} }
func (m *RegisterResourceResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResourceResponse) ProtoMessage()    {}
func (*RegisterResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{5}
}
func (m *RegisterResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceResponse.Unmarshal(m, b)
}
func (m *RegisterResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceResponse.Merge(dst, src)
}
func (m *RegisterResourceResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceResponse.Size(m)
}
func (m *RegisterResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceResponse proto.InternalMessageInfo

func (m *RegisterResourceResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *RegisterResourceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterResourceResponse) GetObject() *_struct.Struct {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *RegisterResourceResponse) GetStable() bool {
	if m != nil {
		return m.Stable
	}
	return false
}

func (m *RegisterResourceResponse) GetStables() []string {
	if m != nil {
		return m.Stables
	}
	return nil
}

// RegisterResourceOutputsRequest adds extra resource outputs created by the program after registration has occurred.
type RegisterResourceOutputsRequest struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Outputs              *_struct.Struct `protobuf:"bytes,2,opt,name=outputs" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterResourceOutputsRequest) Reset()         { *m = RegisterResourceOutputsRequest{} }
func (m *RegisterResourceOutputsRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterResourceOutputsRequest) ProtoMessage()    {}
func (*RegisterResourceOutputsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_12d170a7510bd808, []int{6}
}
func (m *RegisterResourceOutputsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceOutputsRequest.Unmarshal(m, b)
}
func (m *RegisterResourceOutputsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceOutputsRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceOutputsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceOutputsRequest.Merge(dst, src)
}
func (m *RegisterResourceOutputsRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceOutputsRequest.Size(m)
}
func (m *RegisterResourceOutputsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceOutputsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceOutputsRequest proto.InternalMessageInfo

func (m *RegisterResourceOutputsRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *RegisterResourceOutputsRequest) GetOutputs() *_struct.Struct {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*SupportsFeatureRequest)(nil), "pulumirpc.SupportsFeatureRequest")
	proto.RegisterType((*SupportsFeatureResponse)(nil), "pulumirpc.SupportsFeatureResponse")
	proto.RegisterType((*ReadResourceRequest)(nil), "pulumirpc.ReadResourceRequest")
	proto.RegisterType((*ReadResourceResponse)(nil), "pulumirpc.ReadResourceResponse")
	proto.RegisterType((*RegisterResourceRequest)(nil), "pulumirpc.RegisterResourceRequest")
	proto.RegisterMapType((map[string]*RegisterResourceRequest_PropertyDependencies)(nil), "pulumirpc.RegisterResourceRequest.PropertyDependenciesEntry")
	proto.RegisterType((*RegisterResourceRequest_PropertyDependencies)(nil), "pulumirpc.RegisterResourceRequest.PropertyDependencies")
	proto.RegisterType((*RegisterResourceResponse)(nil), "pulumirpc.RegisterResourceResponse")
	proto.RegisterType((*RegisterResourceOutputsRequest)(nil), "pulumirpc.RegisterResourceOutputsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceMonitor service

type ResourceMonitorClient interface {
	SupportsFeature(ctx context.Context, in *SupportsFeatureRequest, opts ...grpc.CallOption) (*SupportsFeatureResponse, error)
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error)
	RegisterResource(ctx context.Context, in *RegisterResourceRequest, opts ...grpc.CallOption) (*RegisterResourceResponse, error)
	RegisterResourceOutputs(ctx context.Context, in *RegisterResourceOutputsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type resourceMonitorClient struct {
	cc *grpc.ClientConn
}

func NewResourceMonitorClient(cc *grpc.ClientConn) ResourceMonitorClient {
	return &resourceMonitorClient{cc}
}

func (c *resourceMonitorClient) SupportsFeature(ctx context.Context, in *SupportsFeatureRequest, opts ...grpc.CallOption) (*SupportsFeatureResponse, error) {
	out := new(SupportsFeatureResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/SupportsFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error) {
	out := new(ReadResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/ReadResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterResource(ctx context.Context, in *RegisterResourceRequest, opts ...grpc.CallOption) (*RegisterResourceResponse, error) {
	out := new(RegisterResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterResourceOutputs(ctx context.Context, in *RegisterResourceOutputsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterResourceOutputs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceMonitor service

type ResourceMonitorServer interface {
	SupportsFeature(context.Context, *SupportsFeatureRequest) (*SupportsFeatureResponse, error)
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error)
	RegisterResource(context.Context, *RegisterResourceRequest) (*RegisterResourceResponse, error)
	RegisterResourceOutputs(context.Context, *RegisterResourceOutputsRequest) (*empty.Empty, error)
}

func RegisterResourceMonitorServer(s *grpc.Server, srv ResourceMonitorServer) {
	s.RegisterService(&_ResourceMonitor_serviceDesc, srv)
}

func _ResourceMonitor_SupportsFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportsFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).SupportsFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/SupportsFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).SupportsFeature(ctx, req.(*SupportsFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_ReadResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).ReadResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/ReadResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).ReadResource(ctx, req.(*ReadResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterResource(ctx, req.(*RegisterResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterResourceOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResourceOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterResourceOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterResourceOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterResourceOutputs(ctx, req.(*RegisterResourceOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceMonitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.ResourceMonitor",
	HandlerType: (*ResourceMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SupportsFeature",
			Handler:    _ResourceMonitor_SupportsFeature_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _ResourceMonitor_Invoke_Handler,
		},
		{
			MethodName: "ReadResource",
			Handler:    _ResourceMonitor_ReadResource_Handler,
		},
		{
			MethodName: "RegisterResource",
			Handler:    _ResourceMonitor_RegisterResource_Handler,
		},
		{
			MethodName: "RegisterResourceOutputs",
			Handler:    _ResourceMonitor_RegisterResourceOutputs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}

func init() { proto.RegisterFile("resource.proto", fileDescriptor_resource_12d170a7510bd808) }

var fileDescriptor_resource_12d170a7510bd808 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xae, 0x9d, 0xd6, 0x49, 0x4e, 0xdb, 0xb4, 0x9a, 0x56, 0xc9, 0xd4, 0xf7, 0xaa, 0x37, 0xd7,
	0xf7, 0x2e, 0x02, 0x8b, 0x14, 0xca, 0xa2, 0x05, 0x21, 0x21, 0x01, 0x45, 0x62, 0x51, 0x01, 0xee,
	0x0e, 0x09, 0x24, 0xc7, 0x3e, 0x4d, 0x4d, 0x13, 0xcf, 0x30, 0x33, 0x8e, 0x94, 0x1d, 0x6f, 0xc2,
	0xab, 0xf0, 0x18, 0xbc, 0x05, 0xaf, 0x80, 0x3c, 0xb6, 0x43, 0x62, 0x3b, 0x49, 0x61, 0x37, 0xe7,
	0xd7, 0x33, 0xdf, 0xf7, 0x9d, 0x19, 0x43, 0x4b, 0xa0, 0x64, 0xb1, 0xf0, 0xb1, 0xcf, 0x05, 0x53,
	0x8c, 0x34, 0x79, 0x3c, 0x8a, 0xc7, 0xa1, 0xe0, 0xbe, 0xfd, 0xd7, 0x90, 0xb1, 0xe1, 0x08, 0x4f,
	0x74, 0x60, 0x10, 0x5f, 0x9f, 0xe0, 0x98, 0xab, 0x69, 0x9a, 0x67, 0xff, 0x5d, 0x0c, 0x4a, 0x25,
	0x62, 0x5f, 0x65, 0xd1, 0x16, 0x17, 0x6c, 0x12, 0x06, 0x28, 0x52, 0xdb, 0xe9, 0x41, 0xfb, 0x2a,
	0xe6, 0x9c, 0x09, 0x25, 0x5f, 0xa1, 0xa7, 0x62, 0x81, 0x2e, 0x7e, 0x8e, 0x51, 0x2a, 0xd2, 0x02,
	0x33, 0x0c, 0xa8, 0xd1, 0x35, 0x7a, 0x4d, 0xd7, 0x0c, 0x03, 0xe7, 0x31, 0x74, 0x4a, 0x99, 0x92,
	0xb3, 0x48, 0x22, 0x39, 0x06, 0xb8, 0xf1, 0x64, 0x16, 0xd5, 0x25, 0x0d, 0x77, 0xce, 0xe3, 0x7c,
	0x37, 0xe1, 0xc0, 0x45, 0x2f, 0x70, 0xb3, 0x13, 0x2d, 0xf9, 0x04, 0x21, 0xb0, 0xa9, 0xa6, 0x1c,
	0xa9, 0xa9, 0x3d, 0x7a, 0x9d, 0xf8, 0x22, 0x6f, 0x8c, 0xb4, 0x96, 0xfa, 0x92, 0x35, 0x69, 0x83,
	0xc5, 0x3d, 0x81, 0x91, 0xa2, 0x9b, 0xda, 0x9b, 0x59, 0xe4, 0x0c, 0x80, 0x0b, 0xc6, 0x51, 0xa8,
	0x10, 0x25, 0xdd, 0xea, 0x1a, 0xbd, 0xed, 0xd3, 0x4e, 0x3f, 0xc5, 0xa3, 0x9f, 0xe3, 0xd1, 0xbf,
	0xd2, 0x78, 0xb8, 0x73, 0xa9, 0xc4, 0x81, 0x9d, 0x00, 0x39, 0x46, 0x01, 0x46, 0x7e, 0x52, 0x6a,
	0x75, 0x6b, 0xbd, 0xa6, 0xbb, 0xe0, 0x23, 0x36, 0x34, 0x72, 0xec, 0x68, 0x5d, 0x7f, 0x76, 0x66,
	0x13, 0x0a, 0xf5, 0x09, 0x0a, 0x19, 0xb2, 0x88, 0x36, 0x74, 0x28, 0x37, 0xc9, 0xff, 0xb0, 0xeb,
	0xf9, 0x3e, 0x72, 0x75, 0x85, 0xbe, 0x40, 0x25, 0x69, 0x53, 0xa3, 0xb3, 0xe8, 0x24, 0xe7, 0xd0,
	0xf1, 0x82, 0x20, 0x54, 0x21, 0x8b, 0xbc, 0x51, 0xea, 0x7c, 0x13, 0x2b, 0x1e, 0x2b, 0x49, 0x41,
	0x6f, 0x65, 0x59, 0xd8, 0xf1, 0xe0, 0x70, 0x11, 0xd9, 0x8c, 0x92, 0x7d, 0xa8, 0xc5, 0x22, 0xca,
	0xb0, 0x4d, 0x96, 0x05, 0x70, 0xcc, 0x3b, 0x83, 0xe3, 0xfc, 0xd8, 0x82, 0x8e, 0x8b, 0xc3, 0x50,
	0x2a, 0x14, 0x45, 0x06, 0x73, 0xc6, 0x8c, 0x0a, 0xc6, 0xcc, 0x4a, 0xc6, 0x6a, 0x0b, 0x8c, 0xb5,
	0xc1, 0xf2, 0x63, 0xa9, 0xd8, 0x58, 0x33, 0xd9, 0x70, 0x33, 0x8b, 0x9c, 0x80, 0xc5, 0x06, 0x9f,
	0xd0, 0x57, 0xeb, 0x58, 0xcc, 0xd2, 0x12, 0x06, 0x92, 0x50, 0x52, 0x61, 0xe9, 0x4e, 0xb9, 0x59,
	0xe2, 0xb6, 0xbe, 0x86, 0xdb, 0x46, 0x81, 0x5b, 0x0e, 0x87, 0x19, 0x18, 0xd3, 0x97, 0xf3, 0x7d,
	0x9a, 0xdd, 0x5a, 0x6f, 0xfb, 0xf4, 0x69, 0x7f, 0x36, 0x96, 0xfd, 0x25, 0x20, 0xf5, 0xdf, 0x56,
	0x94, 0x5f, 0x44, 0x4a, 0x4c, 0xdd, 0xca, 0xce, 0xe4, 0x01, 0x1c, 0x04, 0x38, 0x42, 0x85, 0xcf,
	0xf1, 0x9a, 0x25, 0x63, 0xc6, 0x47, 0x9e, 0x8f, 0x14, 0xf4, 0xb9, 0xaa, 0x42, 0xf3, 0xfa, 0xdb,
	0x2e, 0xe9, 0x2f, 0x1c, 0x46, 0x4c, 0xe0, 0x8b, 0x1b, 0x2f, 0x1a, 0xa2, 0xa4, 0x3b, 0xfa, 0xf8,
	0x8b, 0xce, 0xb2, 0x4a, 0x77, 0x7f, 0x53, 0xa5, 0xad, 0x95, 0x2a, 0xb5, 0xef, 0xc3, 0x61, 0x15,
	0x08, 0x89, 0x54, 0x62, 0x11, 0x49, 0x6a, 0xe8, 0x72, 0xbd, 0xb6, 0xbf, 0x18, 0x70, 0xb4, 0x14,
	0xb1, 0x44, 0xd7, 0xb7, 0x38, 0xcd, 0x75, 0x7d, 0x8b, 0x53, 0x72, 0x09, 0x5b, 0x13, 0x6f, 0x14,
	0x63, 0x26, 0xe9, 0xb3, 0x3f, 0x24, 0xc4, 0x4d, 0xbb, 0x3c, 0x31, 0xcf, 0x0d, 0xe7, 0xab, 0x01,
	0xb4, 0x5c, 0xbb, 0x74, 0xb2, 0xd2, 0x6b, 0xcc, 0x9c, 0x5d, 0x63, 0xbf, 0xc4, 0x5b, 0xbb, 0x9b,
	0x78, 0xdb, 0x60, 0x49, 0xe5, 0x0d, 0x46, 0x98, 0x4f, 0x41, 0x6a, 0x25, 0xb4, 0xa6, 0xab, 0xe4,
	0x32, 0x4b, 0x10, 0xca, 0x4d, 0x07, 0xe1, 0xb8, 0xb8, 0xc1, 0x0c, 0xeb, 0x7c, 0x32, 0xcb, 0xdb,
	0x7c, 0x08, 0x75, 0x96, 0xd1, 0xb5, 0x66, 0xfa, 0xf3, 0xbc, 0xd3, 0x6f, 0x35, 0xd8, 0xcb, 0xfb,
	0x5f, 0xb2, 0x28, 0x54, 0x4c, 0x90, 0xf7, 0xb0, 0x57, 0x78, 0x07, 0xc8, 0xbf, 0x73, 0x98, 0x57,
	0xbf, 0x26, 0xb6, 0xb3, 0x2a, 0x25, 0x45, 0xd6, 0xd9, 0x20, 0xcf, 0xc0, 0x7a, 0x1d, 0x4d, 0xd8,
	0x2d, 0x12, 0x3a, 0x97, 0x9f, 0xba, 0xf2, 0x4e, 0x47, 0x15, 0x91, 0x59, 0x83, 0x77, 0xb0, 0x33,
	0x7f, 0x1d, 0x92, 0xe3, 0x05, 0x35, 0x94, 0x5e, 0x20, 0xfb, 0x9f, 0xa5, 0xf1, 0x59, 0xcb, 0x0f,
	0xb0, 0x5f, 0x84, 0x9a, 0x38, 0xeb, 0x45, 0x66, 0xff, 0xb7, 0x32, 0x67, 0xd6, 0xfe, 0x63, 0xf9,
	0x72, 0xcd, 0x98, 0x24, 0xf7, 0x56, 0x74, 0x58, 0x64, 0xdb, 0x6e, 0x97, 0xa8, 0xbc, 0x48, 0x7e,
	0x09, 0x9c, 0x8d, 0x81, 0xa5, 0x3d, 0x8f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x86, 0x35,
	0xc8, 0x4f, 0x08, 0x00, 0x00,
}
