// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: apps/slb/pb/slb.proto

package slb

import (
	resource "github.com/infraboard/cmdb/apps/resource"
	request "github.com/infraboard/mcube/http/request"
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

type SLB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        *resource.Base        `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information,omitempty"`
	Describe    *Describe             `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe,omitempty"`
}

func (x *SLB) Reset() {
	*x = SLB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_slb_pb_slb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SLB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SLB) ProtoMessage() {}

func (x *SLB) ProtoReflect() protoreflect.Message {
	mi := &file_apps_slb_pb_slb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SLB.ProtoReflect.Descriptor instead.
func (*SLB) Descriptor() ([]byte, []int) {
	return file_apps_slb_pb_slb_proto_rawDescGZIP(), []int{0}
}

func (x *SLB) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *SLB) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *SLB) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 负载均衡实例的域名，仅公网传统型负载均衡实例才提供该字段
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain"`
	// 状态变化的时间
	// @gotags: json:"status_at"
	StatusAt int64 `protobuf:"varint,2,opt,name=status_at,json=statusAt,proto3" json:"status_at"`
	// IP版本，可以设置为ipv4或者ipv6
	// @gotags: json:"ip_version"
	IpVersion string `protobuf:"bytes,3,opt,name=ip_version,json=ipVersion,proto3" json:"ip_version"`
	// 私网负载均衡实例的网络类型 vpc：专有网络实例 classic：经典网络实例
	// @gotags: json:"network_type"
	NetworkType string `protobuf:"bytes,4,opt,name=network_type,json=networkType,proto3" json:"network_type"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_slb_pb_slb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_slb_pb_slb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Describe.ProtoReflect.Descriptor instead.
func (*Describe) Descriptor() ([]byte, []int) {
	return file_apps_slb_pb_slb_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Describe) GetStatusAt() int64 {
	if x != nil {
		return x.StatusAt
	}
	return 0
}

func (x *Describe) GetIpVersion() string {
	if x != nil {
		return x.IpVersion
	}
	return ""
}

func (x *Describe) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 列表项
	Items []*SLB `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// 总数量
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_slb_pb_slb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_slb_pb_slb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_apps_slb_pb_slb_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetItems() []*SLB {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type QuerySLBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QuerySLBRequest) Reset() {
	*x = QuerySLBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_slb_pb_slb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySLBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySLBRequest) ProtoMessage() {}

func (x *QuerySLBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_slb_pb_slb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySLBRequest.ProtoReflect.Descriptor instead.
func (*QuerySLBRequest) Descriptor() ([]byte, []int) {
	return file_apps_slb_pb_slb_proto_rawDescGZIP(), []int{3}
}

func (x *QuerySLBRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_apps_slb_pb_slb_proto protoreflect.FileDescriptor

var file_apps_slb_pb_slb_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x6c, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6c,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62, 0x1a, 0x1f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01,
	0x0a, 0x03, 0x53, 0x4c, 0x42, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x81, 0x01,
	0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x4b, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62, 0x2e, 0x53, 0x4c,
	0x42, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x49,
	0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x4c, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x98, 0x01, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x4c, 0x42,
	0x12, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62, 0x2e, 0x53, 0x4c, 0x42, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62,
	0x2e, 0x53, 0x4c, 0x42, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x4c, 0x42, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x4c,
	0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x6c, 0x62, 0x2e, 0x53,
	0x65, 0x74, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d,
	0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x6c, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_slb_pb_slb_proto_rawDescOnce sync.Once
	file_apps_slb_pb_slb_proto_rawDescData = file_apps_slb_pb_slb_proto_rawDesc
)

func file_apps_slb_pb_slb_proto_rawDescGZIP() []byte {
	file_apps_slb_pb_slb_proto_rawDescOnce.Do(func() {
		file_apps_slb_pb_slb_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_slb_pb_slb_proto_rawDescData)
	})
	return file_apps_slb_pb_slb_proto_rawDescData
}

var file_apps_slb_pb_slb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_slb_pb_slb_proto_goTypes = []interface{}{
	(*SLB)(nil),                  // 0: infraboard.cmdb.slb.SLB
	(*Describe)(nil),             // 1: infraboard.cmdb.slb.Describe
	(*Set)(nil),                  // 2: infraboard.cmdb.slb.Set
	(*QuerySLBRequest)(nil),      // 3: infraboard.cmdb.slb.QuerySLBRequest
	(*resource.Base)(nil),        // 4: infraboard.cmdb.resource.Base
	(*resource.Information)(nil), // 5: infraboard.cmdb.resource.Information
	(*request.PageRequest)(nil),  // 6: infraboard.mcube.page.PageRequest
}
var file_apps_slb_pb_slb_proto_depIdxs = []int32{
	4, // 0: infraboard.cmdb.slb.SLB.base:type_name -> infraboard.cmdb.resource.Base
	5, // 1: infraboard.cmdb.slb.SLB.information:type_name -> infraboard.cmdb.resource.Information
	1, // 2: infraboard.cmdb.slb.SLB.describe:type_name -> infraboard.cmdb.slb.Describe
	0, // 3: infraboard.cmdb.slb.Set.items:type_name -> infraboard.cmdb.slb.SLB
	6, // 4: infraboard.cmdb.slb.QuerySLBRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 5: infraboard.cmdb.slb.Service.SyncSLB:input_type -> infraboard.cmdb.slb.SLB
	3, // 6: infraboard.cmdb.slb.Service.QuerySLB:input_type -> infraboard.cmdb.slb.QuerySLBRequest
	0, // 7: infraboard.cmdb.slb.Service.SyncSLB:output_type -> infraboard.cmdb.slb.SLB
	2, // 8: infraboard.cmdb.slb.Service.QuerySLB:output_type -> infraboard.cmdb.slb.Set
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_slb_pb_slb_proto_init() }
func file_apps_slb_pb_slb_proto_init() {
	if File_apps_slb_pb_slb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_slb_pb_slb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SLB); i {
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
		file_apps_slb_pb_slb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Describe); i {
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
		file_apps_slb_pb_slb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_apps_slb_pb_slb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySLBRequest); i {
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
			RawDescriptor: file_apps_slb_pb_slb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_slb_pb_slb_proto_goTypes,
		DependencyIndexes: file_apps_slb_pb_slb_proto_depIdxs,
		MessageInfos:      file_apps_slb_pb_slb_proto_msgTypes,
	}.Build()
	File_apps_slb_pb_slb_proto = out.File
	file_apps_slb_pb_slb_proto_rawDesc = nil
	file_apps_slb_pb_slb_proto_goTypes = nil
	file_apps_slb_pb_slb_proto_depIdxs = nil
}
