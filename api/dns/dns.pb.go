// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: dns/dns.proto

package dns

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SetHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *SetHostnameRequest) Reset() {
	*x = SetHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHostnameRequest) ProtoMessage() {}

func (x *SetHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHostnameRequest.ProtoReflect.Descriptor instead.
func (*SetHostnameRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{0}
}

func (x *SetHostnameRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type SetHostnameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SetHostnameResponse) Reset() {
	*x = SetHostnameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHostnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHostnameResponse) ProtoMessage() {}

func (x *SetHostnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHostnameResponse.ProtoReflect.Descriptor instead.
func (*SetHostnameResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{1}
}

func (x *SetHostnameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SetHostnameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetDNSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDNSRequest) Reset() {
	*x = GetDNSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDNSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDNSRequest) ProtoMessage() {}

func (x *GetDNSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDNSRequest.ProtoReflect.Descriptor instead.
func (*GetDNSRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{2}
}

type GetDNSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServers []string `protobuf:"bytes,1,rep,name=dns_servers,json=dnsServers,proto3" json:"dns_servers,omitempty"`
}

func (x *GetDNSResponse) Reset() {
	*x = GetDNSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDNSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDNSResponse) ProtoMessage() {}

func (x *GetDNSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDNSResponse.ProtoReflect.Descriptor instead.
func (*GetDNSResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{3}
}

func (x *GetDNSResponse) GetDnsServers() []string {
	if x != nil {
		return x.DnsServers
	}
	return nil
}

type AddDNSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServer string `protobuf:"bytes,1,opt,name=dns_server,json=dnsServer,proto3" json:"dns_server,omitempty"`
}

func (x *AddDNSRequest) Reset() {
	*x = AddDNSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDNSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDNSRequest) ProtoMessage() {}

func (x *AddDNSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDNSRequest.ProtoReflect.Descriptor instead.
func (*AddDNSRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{4}
}

func (x *AddDNSRequest) GetDnsServer() string {
	if x != nil {
		return x.DnsServer
	}
	return ""
}

type AddDNSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddDNSResponse) Reset() {
	*x = AddDNSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDNSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDNSResponse) ProtoMessage() {}

func (x *AddDNSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDNSResponse.ProtoReflect.Descriptor instead.
func (*AddDNSResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{5}
}

func (x *AddDNSResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddDNSResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveDNSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServer string `protobuf:"bytes,1,opt,name=dns_server,json=dnsServer,proto3" json:"dns_server,omitempty"`
}

func (x *RemoveDNSRequest) Reset() {
	*x = RemoveDNSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDNSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDNSRequest) ProtoMessage() {}

func (x *RemoveDNSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDNSRequest.ProtoReflect.Descriptor instead.
func (*RemoveDNSRequest) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveDNSRequest) GetDnsServer() string {
	if x != nil {
		return x.DnsServer
	}
	return ""
}

type RemoveDNSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveDNSResponse) Reset() {
	*x = RemoveDNSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_dns_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDNSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDNSResponse) ProtoMessage() {}

func (x *RemoveDNSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_dns_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDNSResponse.ProtoReflect.Descriptor instead.
func (*RemoveDNSResponse) Descriptor() ([]byte, []int) {
	return file_dns_dns_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveDNSResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveDNSResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_dns_dns_proto protoreflect.FileDescriptor

var file_dns_dns_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x64, 0x6e, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x30, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xcb, 0x02, 0x0a, 0x0a, 0x44, 0x4e, 0x53, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64,
	0x6e, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x1a, 0x05,
	0x2f, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x12, 0x12, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x64, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x64, 0x6e, 0x73,
	0x12, 0x4b, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53, 0x12, 0x12, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x04, 0x2f, 0x64, 0x6e,
	0x73, 0x3a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x54, 0x0a,
	0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x4e, 0x53, 0x12, 0x15, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x4e,
	0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x1a, 0x04, 0x2f, 0x64, 0x6e, 0x73, 0x3a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x64, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dns_dns_proto_rawDescOnce sync.Once
	file_dns_dns_proto_rawDescData = file_dns_dns_proto_rawDesc
)

func file_dns_dns_proto_rawDescGZIP() []byte {
	file_dns_dns_proto_rawDescOnce.Do(func() {
		file_dns_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_dns_proto_rawDescData)
	})
	return file_dns_dns_proto_rawDescData
}

var file_dns_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dns_dns_proto_goTypes = []interface{}{
	(*SetHostnameRequest)(nil),  // 0: dns.SetHostnameRequest
	(*SetHostnameResponse)(nil), // 1: dns.SetHostnameResponse
	(*GetDNSRequest)(nil),       // 2: dns.GetDNSRequest
	(*GetDNSResponse)(nil),      // 3: dns.GetDNSResponse
	(*AddDNSRequest)(nil),       // 4: dns.AddDNSRequest
	(*AddDNSResponse)(nil),      // 5: dns.AddDNSResponse
	(*RemoveDNSRequest)(nil),    // 6: dns.RemoveDNSRequest
	(*RemoveDNSResponse)(nil),   // 7: dns.RemoveDNSResponse
}
var file_dns_dns_proto_depIdxs = []int32{
	0, // 0: dns.DNSService.SetHostname:input_type -> dns.SetHostnameRequest
	2, // 1: dns.DNSService.GetDNS:input_type -> dns.GetDNSRequest
	4, // 2: dns.DNSService.AddDNS:input_type -> dns.AddDNSRequest
	6, // 3: dns.DNSService.RemoveDNS:input_type -> dns.RemoveDNSRequest
	1, // 4: dns.DNSService.SetHostname:output_type -> dns.SetHostnameResponse
	3, // 5: dns.DNSService.GetDNS:output_type -> dns.GetDNSResponse
	5, // 6: dns.DNSService.AddDNS:output_type -> dns.AddDNSResponse
	7, // 7: dns.DNSService.RemoveDNS:output_type -> dns.RemoveDNSResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dns_dns_proto_init() }
func file_dns_dns_proto_init() {
	if File_dns_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dns_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHostnameRequest); i {
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
		file_dns_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHostnameResponse); i {
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
		file_dns_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDNSRequest); i {
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
		file_dns_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDNSResponse); i {
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
		file_dns_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDNSRequest); i {
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
		file_dns_dns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDNSResponse); i {
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
		file_dns_dns_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDNSRequest); i {
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
		file_dns_dns_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDNSResponse); i {
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
			RawDescriptor: file_dns_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dns_dns_proto_goTypes,
		DependencyIndexes: file_dns_dns_proto_depIdxs,
		MessageInfos:      file_dns_dns_proto_msgTypes,
	}.Build()
	File_dns_dns_proto = out.File
	file_dns_dns_proto_rawDesc = nil
	file_dns_dns_proto_goTypes = nil
	file_dns_dns_proto_depIdxs = nil
}