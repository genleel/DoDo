// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/videoService.proto

package videoService

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

type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime int64 `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
	UserId     int64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{0}
}

func (x *FeedReq) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

func (x *FeedReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author        *User  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PlayUrl       string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl      string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	FavoriteCount int64  `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
	CommentCount  int64  `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	IsFavorite    bool   `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{2}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FeedRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextTime  int64    `protobuf:"varint,1,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"`
	VideoList []*Video `protobuf:"bytes,2,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *FeedRsp) Reset() {
	*x = FeedRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRsp) ProtoMessage() {}

func (x *FeedRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRsp.ProtoReflect.Descriptor instead.
func (*FeedRsp) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{3}
}

func (x *FeedRsp) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

func (x *FeedRsp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type GetVideoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetVideoReq) Reset() {
	*x = GetVideoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoReq) ProtoMessage() {}

func (x *GetVideoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoReq.ProtoReflect.Descriptor instead.
func (*GetVideoReq) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{4}
}

func (x *GetVideoReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *GetVideoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetVideoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video *Video `protobuf:"bytes,3,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *GetVideoRsp) Reset() {
	*x = GetVideoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRsp) ProtoMessage() {}

func (x *GetVideoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRsp.ProtoReflect.Descriptor instead.
func (*GetVideoRsp) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{5}
}

func (x *GetVideoRsp) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	FileSize int64  `protobuf:"varint,4,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	FileExt  string `protobuf:"bytes,5,opt,name=fileExt,proto3" json:"fileExt,omitempty"`
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{6}
}

func (x *PublishReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PublishReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishReq) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *PublishReq) GetFileExt() string {
	if x != nil {
		return x.FileExt
	}
	return ""
}

type PublishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishRsp) Reset() {
	*x = PublishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRsp) ProtoMessage() {}

func (x *PublishRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRsp.ProtoReflect.Descriptor instead.
func (*PublishRsp) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{7}
}

type PublishListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CurId  int64 `protobuf:"varint,2,opt,name=curId,proto3" json:"curId,omitempty"`
}

func (x *PublishListReq) Reset() {
	*x = PublishListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListReq) ProtoMessage() {}

func (x *PublishListReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListReq.ProtoReflect.Descriptor instead.
func (*PublishListReq) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{8}
}

func (x *PublishListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishListReq) GetCurId() int64 {
	if x != nil {
		return x.CurId
	}
	return 0
}

type PublishListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video []*Video `protobuf:"bytes,1,rep,name=video,proto3" json:"video,omitempty"`
}

func (x *PublishListRsp) Reset() {
	*x = PublishListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListRsp) ProtoMessage() {}

func (x *PublishListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListRsp.ProtoReflect.Descriptor instead.
func (*PublishListRsp) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{9}
}

func (x *PublishListRsp) GetVideo() []*Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type VideoIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *VideoIdReq) Reset() {
	*x = VideoIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoIdReq) ProtoMessage() {}

func (x *VideoIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoIdReq.ProtoReflect.Descriptor instead.
func (*VideoIdReq) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{10}
}

func (x *VideoIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type VideoIdRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId []int64 `protobuf:"varint,1,rep,packed,name=videoId,proto3" json:"videoId,omitempty"`
}

func (x *VideoIdRsp) Reset() {
	*x = VideoIdRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_videoService_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoIdRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoIdRsp) ProtoMessage() {}

func (x *VideoIdRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_videoService_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoIdRsp.ProtoReflect.Descriptor instead.
func (*VideoIdRsp) Descriptor() ([]byte, []int) {
	return file_proto_videoService_proto_rawDescGZIP(), []int{11}
}

func (x *VideoIdRsp) GetVideoId() []int64 {
	if x != nil {
		return x.VideoId
	}
	return nil
}

var File_proto_videoService_proto protoreflect.FileDescriptor

var file_proto_videoService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x22, 0xfe, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x5a, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3f, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x29, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x22,
	0x0c, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x22, 0x3e, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x75, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12,
	0x29, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x26, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x32, 0xe3, 0x02, 0x0a, 0x0c, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x46, 0x65, 0x65,
	0x64, 0x12, 0x15, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x19, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x12, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_videoService_proto_rawDescOnce sync.Once
	file_proto_videoService_proto_rawDescData = file_proto_videoService_proto_rawDesc
)

func file_proto_videoService_proto_rawDescGZIP() []byte {
	file_proto_videoService_proto_rawDescOnce.Do(func() {
		file_proto_videoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_videoService_proto_rawDescData)
	})
	return file_proto_videoService_proto_rawDescData
}

var file_proto_videoService_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_videoService_proto_goTypes = []interface{}{
	(*FeedReq)(nil),        // 0: videoService.FeedReq
	(*User)(nil),           // 1: videoService.User
	(*Video)(nil),          // 2: videoService.Video
	(*FeedRsp)(nil),        // 3: videoService.FeedRsp
	(*GetVideoReq)(nil),    // 4: videoService.GetVideoReq
	(*GetVideoRsp)(nil),    // 5: videoService.GetVideoRsp
	(*PublishReq)(nil),     // 6: videoService.PublishReq
	(*PublishRsp)(nil),     // 7: videoService.PublishRsp
	(*PublishListReq)(nil), // 8: videoService.PublishListReq
	(*PublishListRsp)(nil), // 9: videoService.PublishListRsp
	(*VideoIdReq)(nil),     // 10: videoService.VideoIdReq
	(*VideoIdRsp)(nil),     // 11: videoService.VideoIdRsp
}
var file_proto_videoService_proto_depIdxs = []int32{
	1,  // 0: videoService.Video.author:type_name -> videoService.User
	2,  // 1: videoService.FeedRsp.video_list:type_name -> videoService.Video
	2,  // 2: videoService.GetVideoRsp.video:type_name -> videoService.Video
	2,  // 3: videoService.PublishListRsp.video:type_name -> videoService.Video
	0,  // 4: videoService.VideoService.Feed:input_type -> videoService.FeedReq
	4,  // 5: videoService.VideoService.GetVideo:input_type -> videoService.GetVideoReq
	6,  // 6: videoService.VideoService.Publish:input_type -> videoService.PublishReq
	8,  // 7: videoService.VideoService.GetPublishList:input_type -> videoService.PublishListReq
	10, // 8: videoService.VideoService.GetVideoIdList:input_type -> videoService.VideoIdReq
	3,  // 9: videoService.VideoService.Feed:output_type -> videoService.FeedRsp
	5,  // 10: videoService.VideoService.GetVideo:output_type -> videoService.GetVideoRsp
	7,  // 11: videoService.VideoService.Publish:output_type -> videoService.PublishRsp
	9,  // 12: videoService.VideoService.GetPublishList:output_type -> videoService.PublishListRsp
	11, // 13: videoService.VideoService.GetVideoIdList:output_type -> videoService.VideoIdRsp
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_videoService_proto_init() }
func file_proto_videoService_proto_init() {
	if File_proto_videoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_videoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_proto_videoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_videoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_videoService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedRsp); i {
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
		file_proto_videoService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoReq); i {
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
		file_proto_videoService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRsp); i {
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
		file_proto_videoService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReq); i {
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
		file_proto_videoService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRsp); i {
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
		file_proto_videoService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListReq); i {
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
		file_proto_videoService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListRsp); i {
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
		file_proto_videoService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoIdReq); i {
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
		file_proto_videoService_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoIdRsp); i {
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
			RawDescriptor: file_proto_videoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_videoService_proto_goTypes,
		DependencyIndexes: file_proto_videoService_proto_depIdxs,
		MessageInfos:      file_proto_videoService_proto_msgTypes,
	}.Build()
	File_proto_videoService_proto = out.File
	file_proto_videoService_proto_rawDesc = nil
	file_proto_videoService_proto_goTypes = nil
	file_proto_videoService_proto_depIdxs = nil
}
