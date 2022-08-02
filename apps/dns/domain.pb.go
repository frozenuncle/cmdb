// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: apps/dns/pb/domain.proto

package dns

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

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        *resource.Base        `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information,omitempty"`
	Describe    *Describe             `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe,omitempty"`
	Records     *RecordSet            `protobuf:"bytes,4,opt,name=records,proto3" json:"records,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_dns_pb_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_apps_dns_pb_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_apps_dns_pb_domain_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Domain) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *Domain) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

func (x *Domain) GetRecords() *RecordSet {
	if x != nil {
		return x.Records
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域名的等级
	// @gotags: json:"grade"
	Grade string `protobuf:"bytes,1,opt,name=grade,proto3" json:"grade"`
	// 域名的 ID
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 当前域名允许的最小的 TTL
	// @gotags: json:"min_ttl"
	MinTtl int64 `protobuf:"varint,3,opt,name=min_ttl,json=minTtl,proto3" json:"min_ttl"`
	// 域名
	// @gotags: json:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// 域名所有者的邮箱帐号
	// @gotags: json:"owner"
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner"`
	// punycode编码后的域名
	// @gotags: json:"punycode"
	Punycode string `protobuf:"bytes,6,opt,name=punycode,proto3" json:"punycode"`
	// 域名下的解析记录默认的 TTL 值
	// @gotags: json:"ttl"
	Ttl int64 `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_dns_pb_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_dns_pb_domain_proto_msgTypes[1]
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
	return file_apps_dns_pb_domain_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *Describe) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Describe) GetMinTtl() int64 {
	if x != nil {
		return x.MinTtl
	}
	return 0
}

func (x *Describe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Describe) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Describe) GetPunycode() string {
	if x != nil {
		return x.Punycode
	}
	return ""
}

func (x *Describe) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type DomainSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表项
	// @gotags: json:"items"
	Items []*Domain `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *DomainSet) Reset() {
	*x = DomainSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_dns_pb_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainSet) ProtoMessage() {}

func (x *DomainSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_dns_pb_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainSet.ProtoReflect.Descriptor instead.
func (*DomainSet) Descriptor() ([]byte, []int) {
	return file_apps_dns_pb_domain_proto_rawDescGZIP(), []int{2}
}

func (x *DomainSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DomainSet) GetItems() []*Domain {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueryDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryDomainRequest) Reset() {
	*x = QueryDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_dns_pb_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDomainRequest) ProtoMessage() {}

func (x *QueryDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_dns_pb_domain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDomainRequest.ProtoReflect.Descriptor instead.
func (*QueryDomainRequest) Descriptor() ([]byte, []int) {
	return file_apps_dns_pb_domain_proto_rawDescGZIP(), []int{3}
}

func (x *QueryDomainRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 记录的暂停、启用状态，1和0分别代表启用和暂停
	// @gotags: json:"status"
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
	// 解析记录的ID
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 解析记录的线路编号
	// @gotags: json:"line"
	Line string `protobuf:"bytes,3,opt,name=line,proto3" json:"line"`
	// 记录的优先级，非 MX 记录的话，该值为0
	// @gotags: json:"mx"
	Mx int64 `protobuf:"varint,4,opt,name=mx,proto3" json:"mx"`
	// 子域名
	// @gotags: json:"name"
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	// 解析记录的备注信息
	// @gotags: json:"remark"
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark"`
	// 记录的 TTL 值
	// @gotags: json:"ttl"
	Ttl int64 `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl"`
	// 解析记录的类型
	// @gotags: json:"type"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	// 解析记录的最后修改时间
	// @gotags: json:"updated_on"
	UpdatedOn int64 `protobuf:"varint,9,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on"`
	// 记录的值
	// @gotags: json:"value"
	Value string `protobuf:"bytes,10,opt,name=value,proto3" json:"value"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_dns_pb_domain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_apps_dns_pb_domain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_apps_dns_pb_domain_proto_rawDescGZIP(), []int{4}
}

func (x *Record) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

func (x *Record) GetMx() int64 {
	if x != nil {
		return x.Mx
	}
	return 0
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Record) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Record) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *Record) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Record) GetUpdatedOn() int64 {
	if x != nil {
		return x.UpdatedOn
	}
	return 0
}

func (x *Record) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RecordSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 列表项
	// @gotags: json:"items"
	Items []*Record `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *RecordSet) Reset() {
	*x = RecordSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_dns_pb_domain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSet) ProtoMessage() {}

func (x *RecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_dns_pb_domain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSet.ProtoReflect.Descriptor instead.
func (*RecordSet) Descriptor() ([]byte, []int) {
	return file_apps_dns_pb_domain_proto_rawDescGZIP(), []int{5}
}

func (x *RecordSet) GetItems() []*Record {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RecordSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_apps_dns_pb_domain_proto protoreflect.FileDescriptor

var file_apps_dns_pb_domain_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x32,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x08, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e,
	0x5f, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x54,
	0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x75, 0x6e, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x75, 0x6e, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x57, 0x0a, 0x09, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x34, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0xdb, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6d, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x57, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xb9, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53,
	0x65, 0x74, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d,
	0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_dns_pb_domain_proto_rawDescOnce sync.Once
	file_apps_dns_pb_domain_proto_rawDescData = file_apps_dns_pb_domain_proto_rawDesc
)

func file_apps_dns_pb_domain_proto_rawDescGZIP() []byte {
	file_apps_dns_pb_domain_proto_rawDescOnce.Do(func() {
		file_apps_dns_pb_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_dns_pb_domain_proto_rawDescData)
	})
	return file_apps_dns_pb_domain_proto_rawDescData
}

var file_apps_dns_pb_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_dns_pb_domain_proto_goTypes = []interface{}{
	(*Domain)(nil),               // 0: infraboard.cmdb.domain.Domain
	(*Describe)(nil),             // 1: infraboard.cmdb.domain.Describe
	(*DomainSet)(nil),            // 2: infraboard.cmdb.domain.DomainSet
	(*QueryDomainRequest)(nil),   // 3: infraboard.cmdb.domain.QueryDomainRequest
	(*Record)(nil),               // 4: infraboard.cmdb.domain.Record
	(*RecordSet)(nil),            // 5: infraboard.cmdb.domain.RecordSet
	(*resource.Base)(nil),        // 6: infraboard.cmdb.resource.Base
	(*resource.Information)(nil), // 7: infraboard.cmdb.resource.Information
	(*request.PageRequest)(nil),  // 8: infraboard.mcube.page.PageRequest
}
var file_apps_dns_pb_domain_proto_depIdxs = []int32{
	6, // 0: infraboard.cmdb.domain.Domain.base:type_name -> infraboard.cmdb.resource.Base
	7, // 1: infraboard.cmdb.domain.Domain.information:type_name -> infraboard.cmdb.resource.Information
	1, // 2: infraboard.cmdb.domain.Domain.describe:type_name -> infraboard.cmdb.domain.Describe
	5, // 3: infraboard.cmdb.domain.Domain.records:type_name -> infraboard.cmdb.domain.RecordSet
	0, // 4: infraboard.cmdb.domain.DomainSet.items:type_name -> infraboard.cmdb.domain.Domain
	8, // 5: infraboard.cmdb.domain.QueryDomainRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4, // 6: infraboard.cmdb.domain.RecordSet.items:type_name -> infraboard.cmdb.domain.Record
	0, // 7: infraboard.cmdb.domain.Service.SyncDomain:input_type -> infraboard.cmdb.domain.Domain
	3, // 8: infraboard.cmdb.domain.Service.QueryDomain:input_type -> infraboard.cmdb.domain.QueryDomainRequest
	0, // 9: infraboard.cmdb.domain.Service.SyncDomain:output_type -> infraboard.cmdb.domain.Domain
	2, // 10: infraboard.cmdb.domain.Service.QueryDomain:output_type -> infraboard.cmdb.domain.DomainSet
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_dns_pb_domain_proto_init() }
func file_apps_dns_pb_domain_proto_init() {
	if File_apps_dns_pb_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_dns_pb_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_apps_dns_pb_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_dns_pb_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainSet); i {
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
		file_apps_dns_pb_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDomainRequest); i {
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
		file_apps_dns_pb_domain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_apps_dns_pb_domain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSet); i {
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
			RawDescriptor: file_apps_dns_pb_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_dns_pb_domain_proto_goTypes,
		DependencyIndexes: file_apps_dns_pb_domain_proto_depIdxs,
		MessageInfos:      file_apps_dns_pb_domain_proto_msgTypes,
	}.Build()
	File_apps_dns_pb_domain_proto = out.File
	file_apps_dns_pb_domain_proto_rawDesc = nil
	file_apps_dns_pb_domain_proto_goTypes = nil
	file_apps_dns_pb_domain_proto_depIdxs = nil
}