// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: realworld/v1/article.proto

package v1

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

type ListArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文章标签
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// 文章作者
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// 用户偏好
	Favorited string `protobuf:"bytes,3,opt,name=favorited,proto3" json:"favorited,omitempty"`
	// 取出条数
	Limit uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// 取数起止位置
	Offset uint64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListArticlesRequest) Reset() {
	*x = ListArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticlesRequest) ProtoMessage() {}

func (x *ListArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticlesRequest.ProtoReflect.Descriptor instead.
func (*ListArticlesRequest) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *ListArticlesRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ListArticlesRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ListArticlesRequest) GetFavorited() string {
	if x != nil {
		return x.Favorited
	}
	return ""
}

func (x *ListArticlesRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListArticlesRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListArticlesResponse) Reset() {
	*x = ListArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticlesResponse) ProtoMessage() {}

func (x *ListArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticlesResponse.ProtoReflect.Descriptor instead.
func (*ListArticlesResponse) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{1}
}

type FeedArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 取出条数
	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// 取数起止位置
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FeedArticlesRequest) Reset() {
	*x = FeedArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedArticlesRequest) ProtoMessage() {}

func (x *FeedArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedArticlesRequest.ProtoReflect.Descriptor instead.
func (*FeedArticlesRequest) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{2}
}

func (x *FeedArticlesRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FeedArticlesRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FeedArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedArticlesResponse) Reset() {
	*x = FeedArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedArticlesResponse) ProtoMessage() {}

func (x *FeedArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedArticlesResponse.ProtoReflect.Descriptor instead.
func (*FeedArticlesResponse) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{3}
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{4}
}

func (x *GetArticleRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetArticleResponse) Reset() {
	*x = GetArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleResponse) ProtoMessage() {}

func (x *GetArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleResponse.ProtoReflect.Descriptor instead.
func (*GetArticleResponse) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{5}
}

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *CreateArticleRequest_Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{6}
}

func (x *CreateArticleRequest) GetArticle() *CreateArticleRequest_Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{7}
}

type UpdateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文章标识
	Slug    string                        `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Article *UpdateArticleRequest_Article `protobuf:"bytes,2,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *UpdateArticleRequest) Reset() {
	*x = UpdateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleRequest) ProtoMessage() {}

func (x *UpdateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleRequest) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateArticleRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *UpdateArticleRequest) GetArticle() *UpdateArticleRequest_Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type UpdateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateArticleResponse) Reset() {
	*x = UpdateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleResponse) ProtoMessage() {}

func (x *UpdateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleResponse.ProtoReflect.Descriptor instead.
func (*UpdateArticleResponse) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{9}
}

type DeleteArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *DeleteArticleRequest) Reset() {
	*x = DeleteArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleRequest) ProtoMessage() {}

func (x *DeleteArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteArticleRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type DeleteArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteArticleResponse) Reset() {
	*x = DeleteArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleResponse) ProtoMessage() {}

func (x *DeleteArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleResponse.ProtoReflect.Descriptor instead.
func (*DeleteArticleResponse) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{11}
}

type CreateArticleRequest_Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标题
	// @required
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// 描述
	// @required
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// 内容
	// @required
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// 选中标签
	TagList []string `protobuf:"bytes,4,rep,name=tagList,proto3" json:"tagList,omitempty"`
}

func (x *CreateArticleRequest_Article) Reset() {
	*x = CreateArticleRequest_Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest_Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest_Article) ProtoMessage() {}

func (x *CreateArticleRequest_Article) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest_Article.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest_Article) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{6, 0}
}

func (x *CreateArticleRequest_Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleRequest_Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateArticleRequest_Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateArticleRequest_Article) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

type UpdateArticleRequest_Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标题
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// 描述
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// 内容
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpdateArticleRequest_Article) Reset() {
	*x = UpdateArticleRequest_Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realworld_v1_article_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleRequest_Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleRequest_Article) ProtoMessage() {}

func (x *UpdateArticleRequest_Article) ProtoReflect() protoreflect.Message {
	mi := &file_realworld_v1_article_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleRequest_Article.ProtoReflect.Descriptor instead.
func (*UpdateArticleRequest_Article) Descriptor() ([]byte, []int) {
	return file_realworld_v1_article_proto_rawDescGZIP(), []int{8, 0}
}

func (x *UpdateArticleRequest_Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateArticleRequest_Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateArticleRequest_Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_realworld_v1_article_proto protoreflect.FileDescriptor

var file_realworld_v1_article_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65,
	0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43,
	0x0a, 0x13, 0x46, 0x65, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x46, 0x65, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x6f, 0x0a, 0x07, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x12, 0x44, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x55, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x17, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc0, 0x05, 0x0a, 0x07,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x6c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65, 0x61,
	0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x71, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x12, 0x6d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x12, 0x72, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72,
	0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x72,
	0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x14, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x6c,
	0x75, 0x67, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65,
	0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x42, 0x29,
	0x5a, 0x27, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x67, 0x6f, 0x2d, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_realworld_v1_article_proto_rawDescOnce sync.Once
	file_realworld_v1_article_proto_rawDescData = file_realworld_v1_article_proto_rawDesc
)

func file_realworld_v1_article_proto_rawDescGZIP() []byte {
	file_realworld_v1_article_proto_rawDescOnce.Do(func() {
		file_realworld_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_realworld_v1_article_proto_rawDescData)
	})
	return file_realworld_v1_article_proto_rawDescData
}

var file_realworld_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_realworld_v1_article_proto_goTypes = []interface{}{
	(*ListArticlesRequest)(nil),          // 0: realworld.v1.ListArticlesRequest
	(*ListArticlesResponse)(nil),         // 1: realworld.v1.ListArticlesResponse
	(*FeedArticlesRequest)(nil),          // 2: realworld.v1.FeedArticlesRequest
	(*FeedArticlesResponse)(nil),         // 3: realworld.v1.FeedArticlesResponse
	(*GetArticleRequest)(nil),            // 4: realworld.v1.GetArticleRequest
	(*GetArticleResponse)(nil),           // 5: realworld.v1.GetArticleResponse
	(*CreateArticleRequest)(nil),         // 6: realworld.v1.CreateArticleRequest
	(*CreateArticleResponse)(nil),        // 7: realworld.v1.CreateArticleResponse
	(*UpdateArticleRequest)(nil),         // 8: realworld.v1.UpdateArticleRequest
	(*UpdateArticleResponse)(nil),        // 9: realworld.v1.UpdateArticleResponse
	(*DeleteArticleRequest)(nil),         // 10: realworld.v1.DeleteArticleRequest
	(*DeleteArticleResponse)(nil),        // 11: realworld.v1.DeleteArticleResponse
	(*CreateArticleRequest_Article)(nil), // 12: realworld.v1.CreateArticleRequest.Article
	(*UpdateArticleRequest_Article)(nil), // 13: realworld.v1.UpdateArticleRequest.Article
}
var file_realworld_v1_article_proto_depIdxs = []int32{
	12, // 0: realworld.v1.CreateArticleRequest.article:type_name -> realworld.v1.CreateArticleRequest.Article
	13, // 1: realworld.v1.UpdateArticleRequest.article:type_name -> realworld.v1.UpdateArticleRequest.Article
	0,  // 2: realworld.v1.Article.ListArticles:input_type -> realworld.v1.ListArticlesRequest
	2,  // 3: realworld.v1.Article.FeedArticles:input_type -> realworld.v1.FeedArticlesRequest
	4,  // 4: realworld.v1.Article.GetArticle:input_type -> realworld.v1.GetArticleRequest
	6,  // 5: realworld.v1.Article.CreateArticle:input_type -> realworld.v1.CreateArticleRequest
	8,  // 6: realworld.v1.Article.UpdateArticle:input_type -> realworld.v1.UpdateArticleRequest
	10, // 7: realworld.v1.Article.DeleteArticle:input_type -> realworld.v1.DeleteArticleRequest
	1,  // 8: realworld.v1.Article.ListArticles:output_type -> realworld.v1.ListArticlesResponse
	3,  // 9: realworld.v1.Article.FeedArticles:output_type -> realworld.v1.FeedArticlesResponse
	5,  // 10: realworld.v1.Article.GetArticle:output_type -> realworld.v1.GetArticleResponse
	7,  // 11: realworld.v1.Article.CreateArticle:output_type -> realworld.v1.CreateArticleResponse
	9,  // 12: realworld.v1.Article.UpdateArticle:output_type -> realworld.v1.UpdateArticleResponse
	11, // 13: realworld.v1.Article.DeleteArticle:output_type -> realworld.v1.DeleteArticleResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_realworld_v1_article_proto_init() }
func file_realworld_v1_article_proto_init() {
	if File_realworld_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_realworld_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticlesRequest); i {
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
		file_realworld_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticlesResponse); i {
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
		file_realworld_v1_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedArticlesRequest); i {
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
		file_realworld_v1_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedArticlesResponse); i {
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
		file_realworld_v1_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
		file_realworld_v1_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleResponse); i {
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
		file_realworld_v1_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_realworld_v1_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_realworld_v1_article_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleRequest); i {
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
		file_realworld_v1_article_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleResponse); i {
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
		file_realworld_v1_article_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleRequest); i {
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
		file_realworld_v1_article_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleResponse); i {
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
		file_realworld_v1_article_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest_Article); i {
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
		file_realworld_v1_article_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleRequest_Article); i {
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
			RawDescriptor: file_realworld_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_realworld_v1_article_proto_goTypes,
		DependencyIndexes: file_realworld_v1_article_proto_depIdxs,
		MessageInfos:      file_realworld_v1_article_proto_msgTypes,
	}.Build()
	File_realworld_v1_article_proto = out.File
	file_realworld_v1_article_proto_rawDesc = nil
	file_realworld_v1_article_proto_goTypes = nil
	file_realworld_v1_article_proto_depIdxs = nil
}