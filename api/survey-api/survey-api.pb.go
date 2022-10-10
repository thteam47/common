// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.7.1
// source: survey-api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type StringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx   *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Value string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringRequest) Reset() {
	*x = StringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringRequest) ProtoMessage() {}

func (x *StringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringRequest.ProtoReflect.Descriptor instead.
func (*StringRequest) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{1}
}

func (x *StringRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *StringRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type StringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StringResponse) Reset() {
	*x = StringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringResponse) ProtoMessage() {}

func (x *StringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringResponse.ProtoReflect.Descriptor instead.
func (*StringResponse) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{2}
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{3}
}

func (x *MessageResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Answers  []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
	Choose   string   `protobuf:"bytes,3,opt,name=choose,proto3" json:"choose,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{4}
}

func (x *Content) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Content) GetAnswers() []string {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *Content) GetChoose() string {
	if x != nil {
		return x.Choose
	}
	return ""
}

type Survey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId     string     `protobuf:"bytes,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
	UserIdCreate string     `protobuf:"bytes,2,opt,name=user_id_create,json=userIdCreate,proto3" json:"user_id_create,omitempty"`
	UserIdVerify string     `protobuf:"bytes,3,opt,name=user_id_verify,json=userIdVerify,proto3" json:"user_id_verify,omitempty"`
	UserIdJoin   string     `protobuf:"bytes,4,opt,name=user_id_join,json=userIdJoin,proto3" json:"user_id_join,omitempty"`
	Contents     []*Content `protobuf:"bytes,5,rep,name=contents,proto3" json:"contents,omitempty"`
	Status       string     `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime   int32      `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime   int32      `protobuf:"varint,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Survey) Reset() {
	*x = Survey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Survey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Survey) ProtoMessage() {}

func (x *Survey) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Survey.ProtoReflect.Descriptor instead.
func (*Survey) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{5}
}

func (x *Survey) GetSurveyId() string {
	if x != nil {
		return x.SurveyId
	}
	return ""
}

func (x *Survey) GetUserIdCreate() string {
	if x != nil {
		return x.UserIdCreate
	}
	return ""
}

func (x *Survey) GetUserIdVerify() string {
	if x != nil {
		return x.UserIdVerify
	}
	return ""
}

func (x *Survey) GetUserIdJoin() string {
	if x != nil {
		return x.UserIdJoin
	}
	return ""
}

func (x *Survey) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *Survey) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Survey) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Survey) GetUpdateTime() int32 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type SurveyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx  *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Data *Survey  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SurveyRequest) Reset() {
	*x = SurveyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyRequest) ProtoMessage() {}

func (x *SurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyRequest.ProtoReflect.Descriptor instead.
func (*SurveyRequest) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{6}
}

func (x *SurveyRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *SurveyRequest) GetData() *Survey {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateSurveyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx   *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Data  *Survey  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Value string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateSurveyRequest) Reset() {
	*x = UpdateSurveyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSurveyRequest) ProtoMessage() {}

func (x *UpdateSurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSurveyRequest.ProtoReflect.Descriptor instead.
func (*UpdateSurveyRequest) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSurveyRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *UpdateSurveyRequest) GetData() *Survey {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateSurveyRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx    *Context    `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Data   *Pagination `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Search string      `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{8}
}

func (x *ListRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *ListRequest) GetData() *Pagination {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{9}
}

func (x *Pagination) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListSurveyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Survey `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListSurveyResponse) Reset() {
	*x = ListSurveyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSurveyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSurveyResponse) ProtoMessage() {}

func (x *ListSurveyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_survey_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSurveyResponse.ProtoReflect.Descriptor instead.
func (*ListSurveyResponse) Descriptor() ([]byte, []int) {
	return file_survey_api_proto_rawDescGZIP(), []int{10}
}

func (x *ListSurveyResponse) GetData() []*Survey {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListSurveyResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_survey_api_proto protoreflect.FileDescriptor

var file_survey_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x63,
	0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63,
	0x74, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x57, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x06, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x78, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x3a, 0x0a,
	0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x52, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x8b, 0x07,
	0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x67, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12,
	0x19, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2e, 0x1a, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x3a, 0x01, 0x2a,
	0x12, 0x6e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73,
	0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d,
	0x12, 0x86, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69,
	0x6e, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x7b,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x42, 0x79, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x49, 0x64, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x12, 0x7c, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b,
	0x2a, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_survey_api_proto_rawDescOnce sync.Once
	file_survey_api_proto_rawDescData = file_survey_api_proto_rawDesc
)

func file_survey_api_proto_rawDescGZIP() []byte {
	file_survey_api_proto_rawDescOnce.Do(func() {
		file_survey_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_survey_api_proto_rawDescData)
	})
	return file_survey_api_proto_rawDescData
}

var file_survey_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_survey_api_proto_goTypes = []interface{}{
	(*Context)(nil),             // 0: survey_api.Context
	(*StringRequest)(nil),       // 1: survey_api.StringRequest
	(*StringResponse)(nil),      // 2: survey_api.StringResponse
	(*MessageResponse)(nil),     // 3: survey_api.MessageResponse
	(*Content)(nil),             // 4: survey_api.Content
	(*Survey)(nil),              // 5: survey_api.Survey
	(*SurveyRequest)(nil),       // 6: survey_api.SurveyRequest
	(*UpdateSurveyRequest)(nil), // 7: survey_api.UpdateSurveyRequest
	(*ListRequest)(nil),         // 8: survey_api.ListRequest
	(*Pagination)(nil),          // 9: survey_api.Pagination
	(*ListSurveyResponse)(nil),  // 10: survey_api.ListSurveyResponse
}
var file_survey_api_proto_depIdxs = []int32{
	0,  // 0: survey_api.StringRequest.ctx:type_name -> survey_api.Context
	4,  // 1: survey_api.Survey.contents:type_name -> survey_api.Content
	0,  // 2: survey_api.SurveyRequest.ctx:type_name -> survey_api.Context
	5,  // 3: survey_api.SurveyRequest.data:type_name -> survey_api.Survey
	0,  // 4: survey_api.UpdateSurveyRequest.ctx:type_name -> survey_api.Context
	5,  // 5: survey_api.UpdateSurveyRequest.data:type_name -> survey_api.Survey
	0,  // 6: survey_api.ListRequest.ctx:type_name -> survey_api.Context
	9,  // 7: survey_api.ListRequest.data:type_name -> survey_api.Pagination
	5,  // 8: survey_api.ListSurveyResponse.data:type_name -> survey_api.Survey
	6,  // 9: survey_api.SurveyService.CreateSurvey:input_type -> survey_api.SurveyRequest
	7,  // 10: survey_api.SurveyService.UpdateSurveyById:input_type -> survey_api.UpdateSurveyRequest
	1,  // 11: survey_api.SurveyService.GetSurveyById:input_type -> survey_api.StringRequest
	1,  // 12: survey_api.SurveyService.GetSurveyByUserJoin:input_type -> survey_api.StringRequest
	1,  // 13: survey_api.SurveyService.GetSurveyByUserCreate:input_type -> survey_api.StringRequest
	1,  // 14: survey_api.SurveyService.ApproveBySurveyId:input_type -> survey_api.StringRequest
	1,  // 15: survey_api.SurveyService.DeleteSurveyById:input_type -> survey_api.StringRequest
	5,  // 16: survey_api.SurveyService.CreateSurvey:output_type -> survey_api.Survey
	2,  // 17: survey_api.SurveyService.UpdateSurveyById:output_type -> survey_api.StringResponse
	5,  // 18: survey_api.SurveyService.GetSurveyById:output_type -> survey_api.Survey
	10, // 19: survey_api.SurveyService.GetSurveyByUserJoin:output_type -> survey_api.ListSurveyResponse
	10, // 20: survey_api.SurveyService.GetSurveyByUserCreate:output_type -> survey_api.ListSurveyResponse
	2,  // 21: survey_api.SurveyService.ApproveBySurveyId:output_type -> survey_api.StringResponse
	2,  // 22: survey_api.SurveyService.DeleteSurveyById:output_type -> survey_api.StringResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_survey_api_proto_init() }
func file_survey_api_proto_init() {
	if File_survey_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_survey_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_survey_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringRequest); i {
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
		file_survey_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringResponse); i {
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
		file_survey_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
		file_survey_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_survey_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Survey); i {
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
		file_survey_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyRequest); i {
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
		file_survey_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSurveyRequest); i {
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
		file_survey_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_survey_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_survey_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSurveyResponse); i {
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
			RawDescriptor: file_survey_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_survey_api_proto_goTypes,
		DependencyIndexes: file_survey_api_proto_depIdxs,
		MessageInfos:      file_survey_api_proto_msgTypes,
	}.Build()
	File_survey_api_proto = out.File
	file_survey_api_proto_rawDesc = nil
	file_survey_api_proto_goTypes = nil
	file_survey_api_proto_depIdxs = nil
}
