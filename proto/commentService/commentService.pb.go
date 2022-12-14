// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: proto/commentService.proto

package commentService

import (
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

type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{0}
}

func (x *IdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CountRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountRsp) Reset() {
	*x = CountRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRsp) ProtoMessage() {}

func (x *CountRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRsp.ProtoReflect.Descriptor instead.
func (*CountRsp) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{1}
}

func (x *CountRsp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	VideoId     int64  `protobuf:"varint,3,opt,name=videoId,proto3" json:"videoId,omitempty"`
	CommentText string `protobuf:"bytes,4,opt,name=commentText,proto3" json:"commentText,omitempty"`
	CreateDate  string `protobuf:"bytes,5,opt,name=createDate,proto3" json:"createDate,omitempty"`
	Cancel      int64  `protobuf:"varint,6,opt,name=cancel,proto3" json:"cancel,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *Comment) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *Comment) GetCancel() int64 {
	if x != nil {
		return x.Cancel
	}
	return 0
}

type FeedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FollowCount    int64  `protobuf:"varint,3,opt,name=followCount,proto3" json:"followCount,omitempty"`
	FollowerCount  int64  `protobuf:"varint,4,opt,name=followerCount,proto3" json:"followerCount,omitempty"`
	IsFollow       bool   `protobuf:"varint,5,opt,name=isFollow,proto3" json:"isFollow,omitempty"`
	TotalFavorited int64  `protobuf:"varint,6,opt,name=totalFavorited,proto3" json:"totalFavorited,omitempty"`
	FavoriteCount  int64  `protobuf:"varint,7,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"`
}

func (x *FeedUser) Reset() {
	*x = FeedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedUser) ProtoMessage() {}

func (x *FeedUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedUser.ProtoReflect.Descriptor instead.
func (*FeedUser) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{3}
}

func (x *FeedUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeedUser) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *FeedUser) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *FeedUser) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *FeedUser) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *FeedUser) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

type CommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserInfo   *FeedUser `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	Content    string    `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreateDate string    `protobuf:"bytes,4,opt,name=createDate,proto3" json:"createDate,omitempty"`
}

func (x *CommentInfo) Reset() {
	*x = CommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentInfo) ProtoMessage() {}

func (x *CommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentInfo.ProtoReflect.Descriptor instead.
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{4}
}

func (x *CommentInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentInfo) GetUserInfo() *FeedUser {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *CommentInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentInfo) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type CommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentReq) Reset() {
	*x = CommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReq) ProtoMessage() {}

func (x *CommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReq.ProtoReflect.Descriptor instead.
func (*CommentReq) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{5}
}

func (x *CommentReq) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentInfo *CommentInfo `protobuf:"bytes,3,opt,name=commentInfo,proto3" json:"commentInfo,omitempty"`
}

func (x *CommentRsp) Reset() {
	*x = CommentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRsp) ProtoMessage() {}

func (x *CommentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRsp.ProtoReflect.Descriptor instead.
func (*CommentRsp) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{6}
}

func (x *CommentRsp) GetCommentInfo() *CommentInfo {
	if x != nil {
		return x.CommentInfo
	}
	return nil
}

type DelRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelRsp) Reset() {
	*x = DelRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRsp) ProtoMessage() {}

func (x *DelRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRsp.ProtoReflect.Descriptor instead.
func (*DelRsp) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{7}
}

type VideoUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *VideoUserReq) Reset() {
	*x = VideoUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoUserReq) ProtoMessage() {}

func (x *VideoUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoUserReq.ProtoReflect.Descriptor instead.
func (*VideoUserReq) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{8}
}

func (x *VideoUserReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *VideoUserReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CommentListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentInfo []*CommentInfo `protobuf:"bytes,1,rep,name=commentInfo,proto3" json:"commentInfo,omitempty"`
}

func (x *CommentListRsp) Reset() {
	*x = CommentListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commentService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRsp) ProtoMessage() {}

func (x *CommentListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commentService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRsp.ProtoReflect.Descriptor instead.
func (*CommentListRsp) Descriptor() ([]byte, []int) {
	return file_proto_commentService_proto_rawDescGZIP(), []int{9}
}

func (x *CommentListRsp) GetCommentInfo() []*CommentInfo {
	if x != nil {
		return x.CommentInfo
	}
	return nil
}

var File_proto_commentService_proto protoreflect.FileDescriptor

var file_proto_commentService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x17, 0x0a, 0x05,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x22,
	0xe0, 0x01, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x3f, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x40, 0x0a, 0x0c, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x3d,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x9f, 0x02,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x45, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_commentService_proto_rawDescOnce sync.Once
	file_proto_commentService_proto_rawDescData = file_proto_commentService_proto_rawDesc
)

func file_proto_commentService_proto_rawDescGZIP() []byte {
	file_proto_commentService_proto_rawDescOnce.Do(func() {
		file_proto_commentService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_commentService_proto_rawDescData)
	})
	return file_proto_commentService_proto_rawDescData
}

var file_proto_commentService_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_commentService_proto_goTypes = []interface{}{
	(*IdReq)(nil),          // 0: commentService.IdReq
	(*CountRsp)(nil),       // 1: commentService.CountRsp
	(*Comment)(nil),        // 2: commentService.Comment
	(*FeedUser)(nil),       // 3: commentService.FeedUser
	(*CommentInfo)(nil),    // 4: commentService.CommentInfo
	(*CommentReq)(nil),     // 5: commentService.CommentReq
	(*CommentRsp)(nil),     // 6: commentService.CommentRsp
	(*DelRsp)(nil),         // 7: commentService.DelRsp
	(*VideoUserReq)(nil),   // 8: commentService.VideoUserReq
	(*CommentListRsp)(nil), // 9: commentService.CommentListRsp
}
var file_proto_commentService_proto_depIdxs = []int32{
	3, // 0: commentService.CommentInfo.userInfo:type_name -> commentService.FeedUser
	2, // 1: commentService.CommentReq.comment:type_name -> commentService.Comment
	4, // 2: commentService.CommentRsp.commentInfo:type_name -> commentService.CommentInfo
	4, // 3: commentService.CommentListRsp.commentInfo:type_name -> commentService.CommentInfo
	0, // 4: commentService.CommentService.CountFromVideoId:input_type -> commentService.IdReq
	5, // 5: commentService.CommentService.Send:input_type -> commentService.CommentReq
	0, // 6: commentService.CommentService.Delete:input_type -> commentService.IdReq
	8, // 7: commentService.CommentService.GetList:input_type -> commentService.VideoUserReq
	1, // 8: commentService.CommentService.CountFromVideoId:output_type -> commentService.CountRsp
	6, // 9: commentService.CommentService.Send:output_type -> commentService.CommentRsp
	7, // 10: commentService.CommentService.Delete:output_type -> commentService.DelRsp
	9, // 11: commentService.CommentService.GetList:output_type -> commentService.CommentListRsp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_commentService_proto_init() }
func file_proto_commentService_proto_init() {
	if File_proto_commentService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_commentService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_proto_commentService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRsp); i {
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
		file_proto_commentService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_proto_commentService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedUser); i {
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
		file_proto_commentService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentInfo); i {
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
		file_proto_commentService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReq); i {
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
		file_proto_commentService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRsp); i {
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
		file_proto_commentService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRsp); i {
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
		file_proto_commentService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoUserReq); i {
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
		file_proto_commentService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRsp); i {
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
			RawDescriptor: file_proto_commentService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_commentService_proto_goTypes,
		DependencyIndexes: file_proto_commentService_proto_depIdxs,
		MessageInfos:      file_proto_commentService_proto_msgTypes,
	}.Build()
	File_proto_commentService_proto = out.File
	file_proto_commentService_proto_rawDesc = nil
	file_proto_commentService_proto_goTypes = nil
	file_proto_commentService_proto_depIdxs = nil
}
