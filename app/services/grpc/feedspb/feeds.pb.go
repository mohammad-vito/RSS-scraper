// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: app/services/grpc/feedspb/feeds.proto

package feedspb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetFeedListRequest_OrderBy int32

const (
	GetFeedListRequest_DEFAULT      GetFeedListRequest_OrderBy = 0
	GetFeedListRequest_DISPLAY_NAME GetFeedListRequest_OrderBy = 1
	GetFeedListRequest_CREATED_AT   GetFeedListRequest_OrderBy = 2
)

// Enum value maps for GetFeedListRequest_OrderBy.
var (
	GetFeedListRequest_OrderBy_name = map[int32]string{
		0: "DEFAULT",
		1: "DISPLAY_NAME",
		2: "CREATED_AT",
	}
	GetFeedListRequest_OrderBy_value = map[string]int32{
		"DEFAULT":      0,
		"DISPLAY_NAME": 1,
		"CREATED_AT":   2,
	}
)

func (x GetFeedListRequest_OrderBy) Enum() *GetFeedListRequest_OrderBy {
	p := new(GetFeedListRequest_OrderBy)
	*p = x
	return p
}

func (x GetFeedListRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetFeedListRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_app_services_grpc_feedspb_feeds_proto_enumTypes[0].Descriptor()
}

func (GetFeedListRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_app_services_grpc_feedspb_feeds_proto_enumTypes[0]
}

func (x GetFeedListRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetFeedListRequest_OrderBy.Descriptor instead.
func (GetFeedListRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP(), []int{3, 0}
}

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Link            string               `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Description     *string              `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	LastPublishedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_published_at,json=lastPublishedAt,proto3,oneof" json:"last_published_at,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP(), []int{0}
}

func (x *Feed) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Feed) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Feed) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Feed) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Feed) GetLastPublishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.LastPublishedAt
	}
	return nil
}

type GetFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFeedRequest) Reset() {
	*x = GetFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedRequest) ProtoMessage() {}

func (x *GetFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedRequest.ProtoReflect.Descriptor instead.
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeedRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *GetFeedResponse) Reset() {
	*x = GetFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedResponse) ProtoMessage() {}

func (x *GetFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedResponse.ProtoReflect.Descriptor instead.
func (*GetFeedResponse) Descriptor() ([]byte, []int) {
	return file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP(), []int{2}
}

func (x *GetFeedResponse) GetFeed() *Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type GetFeedListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit   int32                      `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32                      `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	OrderBy GetFeedListRequest_OrderBy `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=feeds.GetFeedListRequest_OrderBy" json:"order_by,omitempty"`
}

func (x *GetFeedListRequest) Reset() {
	*x = GetFeedListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedListRequest) ProtoMessage() {}

func (x *GetFeedListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedListRequest.ProtoReflect.Descriptor instead.
func (*GetFeedListRequest) Descriptor() ([]byte, []int) {
	return file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeedListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetFeedListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetFeedListRequest) GetOrderBy() GetFeedListRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return GetFeedListRequest_DEFAULT
}

type GetFeedListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feeds []*Feed `protobuf:"bytes,1,rep,name=feeds,proto3" json:"feeds,omitempty"`
}

func (x *GetFeedListResponse) Reset() {
	*x = GetFeedListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedListResponse) ProtoMessage() {}

func (x *GetFeedListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_grpc_feedspb_feeds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedListResponse.ProtoReflect.Descriptor instead.
func (*GetFeedListResponse) Descriptor() ([]byte, []int) {
	return file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP(), []int{4}
}

func (x *GetFeedListResponse) GetFeeds() []*Feed {
	if x != nil {
		return x.Feeds
	}
	return nil
}

var File_app_services_grpc_feedspb_feeds_proto protoreflect.FileDescriptor

var file_app_services_grpc_feedspb_feeds_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x70, 0x62, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xda, 0x01, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x11, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x66, 0x65,
	0x65, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x38, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x02, 0x22,
	0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x32, 0x8d, 0x01, 0x0a, 0x0b, 0x46, 0x65,
	0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x12, 0x15, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_services_grpc_feedspb_feeds_proto_rawDescOnce sync.Once
	file_app_services_grpc_feedspb_feeds_proto_rawDescData = file_app_services_grpc_feedspb_feeds_proto_rawDesc
)

func file_app_services_grpc_feedspb_feeds_proto_rawDescGZIP() []byte {
	file_app_services_grpc_feedspb_feeds_proto_rawDescOnce.Do(func() {
		file_app_services_grpc_feedspb_feeds_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_services_grpc_feedspb_feeds_proto_rawDescData)
	})
	return file_app_services_grpc_feedspb_feeds_proto_rawDescData
}

var file_app_services_grpc_feedspb_feeds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_services_grpc_feedspb_feeds_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_services_grpc_feedspb_feeds_proto_goTypes = []interface{}{
	(GetFeedListRequest_OrderBy)(0), // 0: feeds.GetFeedListRequest.OrderBy
	(*Feed)(nil),                    // 1: feeds.Feed
	(*GetFeedRequest)(nil),          // 2: feeds.GetFeedRequest
	(*GetFeedResponse)(nil),         // 3: feeds.GetFeedResponse
	(*GetFeedListRequest)(nil),      // 4: feeds.GetFeedListRequest
	(*GetFeedListResponse)(nil),     // 5: feeds.GetFeedListResponse
	(*timestamp.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_app_services_grpc_feedspb_feeds_proto_depIdxs = []int32{
	6, // 0: feeds.Feed.last_published_at:type_name -> google.protobuf.Timestamp
	1, // 1: feeds.GetFeedResponse.feed:type_name -> feeds.Feed
	0, // 2: feeds.GetFeedListRequest.order_by:type_name -> feeds.GetFeedListRequest.OrderBy
	1, // 3: feeds.GetFeedListResponse.feeds:type_name -> feeds.Feed
	2, // 4: feeds.FeedService.GetFeed:input_type -> feeds.GetFeedRequest
	4, // 5: feeds.FeedService.GetFeedList:input_type -> feeds.GetFeedListRequest
	3, // 6: feeds.FeedService.GetFeed:output_type -> feeds.GetFeedResponse
	5, // 7: feeds.FeedService.GetFeedList:output_type -> feeds.GetFeedListResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_app_services_grpc_feedspb_feeds_proto_init() }
func file_app_services_grpc_feedspb_feeds_proto_init() {
	if File_app_services_grpc_feedspb_feeds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_services_grpc_feedspb_feeds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed); i {
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
		file_app_services_grpc_feedspb_feeds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedRequest); i {
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
		file_app_services_grpc_feedspb_feeds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedResponse); i {
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
		file_app_services_grpc_feedspb_feeds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedListRequest); i {
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
		file_app_services_grpc_feedspb_feeds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedListResponse); i {
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
	file_app_services_grpc_feedspb_feeds_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_services_grpc_feedspb_feeds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_services_grpc_feedspb_feeds_proto_goTypes,
		DependencyIndexes: file_app_services_grpc_feedspb_feeds_proto_depIdxs,
		EnumInfos:         file_app_services_grpc_feedspb_feeds_proto_enumTypes,
		MessageInfos:      file_app_services_grpc_feedspb_feeds_proto_msgTypes,
	}.Build()
	File_app_services_grpc_feedspb_feeds_proto = out.File
	file_app_services_grpc_feedspb_feeds_proto_rawDesc = nil
	file_app_services_grpc_feedspb_feeds_proto_goTypes = nil
	file_app_services_grpc_feedspb_feeds_proto_depIdxs = nil
}
