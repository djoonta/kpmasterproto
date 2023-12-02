// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: occupation/occupation.proto

package occupation

import (
	pagination "github.com/djoonta/kpmasterproto/pagination"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusBlacklist int32

const (
	StatusBlacklist_APPROVED  StatusBlacklist = 0
	StatusBlacklist_BLACKLIST StatusBlacklist = 1
	StatusBlacklist_REVIEW    StatusBlacklist = 2
)

// Enum value maps for StatusBlacklist.
var (
	StatusBlacklist_name = map[int32]string{
		0: "APPROVED",
		1: "BLACKLIST",
		2: "REVIEW",
	}
	StatusBlacklist_value = map[string]int32{
		"APPROVED":  0,
		"BLACKLIST": 1,
		"REVIEW":    2,
	}
)

func (x StatusBlacklist) Enum() *StatusBlacklist {
	p := new(StatusBlacklist)
	*p = x
	return p
}

func (x StatusBlacklist) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusBlacklist) Descriptor() protoreflect.EnumDescriptor {
	return file_occupation_occupation_proto_enumTypes[0].Descriptor()
}

func (StatusBlacklist) Type() protoreflect.EnumType {
	return &file_occupation_occupation_proto_enumTypes[0]
}

func (x StatusBlacklist) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusBlacklist.Descriptor instead.
func (StatusBlacklist) EnumDescriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{0}
}

type OccupationWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Messsage string `protobuf:"bytes,2,opt,name=messsage,proto3" json:"messsage,omitempty"`
}

func (x *OccupationWriteResponse) Reset() {
	*x = OccupationWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationWriteResponse) ProtoMessage() {}

func (x *OccupationWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationWriteResponse.ProtoReflect.Descriptor instead.
func (*OccupationWriteResponse) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{0}
}

func (x *OccupationWriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *OccupationWriteResponse) GetMesssage() string {
	if x != nil {
		return x.Messsage
	}
	return ""
}

type OccupationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StatusBlacklist StatusBlacklist        `protobuf:"varint,4,opt,name=status_blacklist,json=statusBlacklist,proto3,enum=kpmasterproto.StatusBlacklist" json:"status_blacklist,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OccupationInfo) Reset() {
	*x = OccupationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationInfo) ProtoMessage() {}

func (x *OccupationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationInfo.ProtoReflect.Descriptor instead.
func (*OccupationInfo) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{1}
}

func (x *OccupationInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OccupationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OccupationInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OccupationInfo) GetStatusBlacklist() StatusBlacklist {
	if x != nil {
		return x.StatusBlacklist
	}
	return StatusBlacklist_APPROVED
}

func (x *OccupationInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type OccupationFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *OccupationFindAllRequest) Reset() {
	*x = OccupationFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationFindAllRequest) ProtoMessage() {}

func (x *OccupationFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationFindAllRequest.ProtoReflect.Descriptor instead.
func (*OccupationFindAllRequest) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{2}
}

func (x *OccupationFindAllRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *OccupationFindAllRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type OccupationFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Occupations []*OccupationInfo            `protobuf:"bytes,1,rep,name=occupations,proto3" json:"occupations,omitempty"`
	Paginations *pagination.PaginationMaster `protobuf:"bytes,2,opt,name=paginations,proto3" json:"paginations,omitempty"`
}

func (x *OccupationFindAllResponse) Reset() {
	*x = OccupationFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationFindAllResponse) ProtoMessage() {}

func (x *OccupationFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationFindAllResponse.ProtoReflect.Descriptor instead.
func (*OccupationFindAllResponse) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{3}
}

func (x *OccupationFindAllResponse) GetOccupations() []*OccupationInfo {
	if x != nil {
		return x.Occupations
	}
	return nil
}

func (x *OccupationFindAllResponse) GetPaginations() *pagination.PaginationMaster {
	if x != nil {
		return x.Paginations
	}
	return nil
}

type OccupationFindIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OccupationFindIDRequest) Reset() {
	*x = OccupationFindIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationFindIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationFindIDRequest) ProtoMessage() {}

func (x *OccupationFindIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationFindIDRequest.ProtoReflect.Descriptor instead.
func (*OccupationFindIDRequest) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{4}
}

func (x *OccupationFindIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OccupationFindIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Occupation *OccupationInfo `protobuf:"bytes,1,opt,name=occupation,proto3" json:"occupation,omitempty"`
}

func (x *OccupationFindIDResponse) Reset() {
	*x = OccupationFindIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationFindIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationFindIDResponse) ProtoMessage() {}

func (x *OccupationFindIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationFindIDResponse.ProtoReflect.Descriptor instead.
func (*OccupationFindIDResponse) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{5}
}

func (x *OccupationFindIDResponse) GetOccupation() *OccupationInfo {
	if x != nil {
		return x.Occupation
	}
	return nil
}

type OccupationCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description     string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	StatusBlacklist StatusBlacklist `protobuf:"varint,3,opt,name=status_blacklist,json=statusBlacklist,proto3,enum=kpmasterproto.StatusBlacklist" json:"status_blacklist,omitempty"`
}

func (x *OccupationCreateRequest) Reset() {
	*x = OccupationCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationCreateRequest) ProtoMessage() {}

func (x *OccupationCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationCreateRequest.ProtoReflect.Descriptor instead.
func (*OccupationCreateRequest) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{6}
}

func (x *OccupationCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OccupationCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OccupationCreateRequest) GetStatusBlacklist() StatusBlacklist {
	if x != nil {
		return x.StatusBlacklist
	}
	return StatusBlacklist_APPROVED
}

type OccupationCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateResponse *OccupationWriteResponse `protobuf:"bytes,1,opt,name=create_response,json=createResponse,proto3" json:"create_response,omitempty"`
}

func (x *OccupationCreateResponse) Reset() {
	*x = OccupationCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationCreateResponse) ProtoMessage() {}

func (x *OccupationCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationCreateResponse.ProtoReflect.Descriptor instead.
func (*OccupationCreateResponse) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{7}
}

func (x *OccupationCreateResponse) GetCreateResponse() *OccupationWriteResponse {
	if x != nil {
		return x.CreateResponse
	}
	return nil
}

type OccupationDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OccupationDeleteRequest) Reset() {
	*x = OccupationDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationDeleteRequest) ProtoMessage() {}

func (x *OccupationDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationDeleteRequest.ProtoReflect.Descriptor instead.
func (*OccupationDeleteRequest) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{8}
}

func (x *OccupationDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OccupationDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteResponse *OccupationWriteResponse `protobuf:"bytes,1,opt,name=delete_response,json=deleteResponse,proto3" json:"delete_response,omitempty"`
}

func (x *OccupationDeleteResponse) Reset() {
	*x = OccupationDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationDeleteResponse) ProtoMessage() {}

func (x *OccupationDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationDeleteResponse.ProtoReflect.Descriptor instead.
func (*OccupationDeleteResponse) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{9}
}

func (x *OccupationDeleteResponse) GetDeleteResponse() *OccupationWriteResponse {
	if x != nil {
		return x.DeleteResponse
	}
	return nil
}

type OccupationUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StatusBlacklist StatusBlacklist `protobuf:"varint,4,opt,name=status_blacklist,json=statusBlacklist,proto3,enum=kpmasterproto.StatusBlacklist" json:"status_blacklist,omitempty"`
}

func (x *OccupationUpdateRequest) Reset() {
	*x = OccupationUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationUpdateRequest) ProtoMessage() {}

func (x *OccupationUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationUpdateRequest.ProtoReflect.Descriptor instead.
func (*OccupationUpdateRequest) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{10}
}

func (x *OccupationUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OccupationUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OccupationUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OccupationUpdateRequest) GetStatusBlacklist() StatusBlacklist {
	if x != nil {
		return x.StatusBlacklist
	}
	return StatusBlacklist_APPROVED
}

type OccupationUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateResponse *OccupationWriteResponse `protobuf:"bytes,1,opt,name=update_response,json=updateResponse,proto3" json:"update_response,omitempty"`
}

func (x *OccupationUpdateResponse) Reset() {
	*x = OccupationUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occupation_occupation_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupationUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupationUpdateResponse) ProtoMessage() {}

func (x *OccupationUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_occupation_occupation_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupationUpdateResponse.ProtoReflect.Descriptor instead.
func (*OccupationUpdateResponse) Descriptor() ([]byte, []int) {
	return file_occupation_occupation_proto_rawDescGZIP(), []int{11}
}

func (x *OccupationUpdateResponse) GetUpdateResponse() *OccupationWriteResponse {
	if x != nil {
		return x.UpdateResponse
	}
	return nil
}

var File_occupation_occupation_proto protoreflect.FileDescriptor

var file_occupation_occupation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b,
	0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x17, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x0e,
	0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x18, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x9f, 0x01, 0x0a, 0x19, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x41, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x29, 0x0a, 0x17, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a,
	0x18, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x6f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x17, 0x4f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x18, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x70, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a,
	0x18, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x17, 0x4f,
	0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x10,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x18, 0x4f, 0x63, 0x63, 0x75, 0x70,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b,
	0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x3a, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x4c, 0x49,
	0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61, 0x2f, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_occupation_occupation_proto_rawDescOnce sync.Once
	file_occupation_occupation_proto_rawDescData = file_occupation_occupation_proto_rawDesc
)

func file_occupation_occupation_proto_rawDescGZIP() []byte {
	file_occupation_occupation_proto_rawDescOnce.Do(func() {
		file_occupation_occupation_proto_rawDescData = protoimpl.X.CompressGZIP(file_occupation_occupation_proto_rawDescData)
	})
	return file_occupation_occupation_proto_rawDescData
}

var file_occupation_occupation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_occupation_occupation_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_occupation_occupation_proto_goTypes = []interface{}{
	(StatusBlacklist)(0),                // 0: kpmasterproto.StatusBlacklist
	(*OccupationWriteResponse)(nil),     // 1: kpmasterproto.OccupationWriteResponse
	(*OccupationInfo)(nil),              // 2: kpmasterproto.OccupationInfo
	(*OccupationFindAllRequest)(nil),    // 3: kpmasterproto.OccupationFindAllRequest
	(*OccupationFindAllResponse)(nil),   // 4: kpmasterproto.OccupationFindAllResponse
	(*OccupationFindIDRequest)(nil),     // 5: kpmasterproto.OccupationFindIDRequest
	(*OccupationFindIDResponse)(nil),    // 6: kpmasterproto.OccupationFindIDResponse
	(*OccupationCreateRequest)(nil),     // 7: kpmasterproto.OccupationCreateRequest
	(*OccupationCreateResponse)(nil),    // 8: kpmasterproto.OccupationCreateResponse
	(*OccupationDeleteRequest)(nil),     // 9: kpmasterproto.OccupationDeleteRequest
	(*OccupationDeleteResponse)(nil),    // 10: kpmasterproto.OccupationDeleteResponse
	(*OccupationUpdateRequest)(nil),     // 11: kpmasterproto.OccupationUpdateRequest
	(*OccupationUpdateResponse)(nil),    // 12: kpmasterproto.OccupationUpdateResponse
	(*timestamppb.Timestamp)(nil),       // 13: google.protobuf.Timestamp
	(*pagination.PaginationMaster)(nil), // 14: kpmasterproto.PaginationMaster
}
var file_occupation_occupation_proto_depIdxs = []int32{
	0,  // 0: kpmasterproto.OccupationInfo.status_blacklist:type_name -> kpmasterproto.StatusBlacklist
	13, // 1: kpmasterproto.OccupationInfo.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 2: kpmasterproto.OccupationFindAllResponse.occupations:type_name -> kpmasterproto.OccupationInfo
	14, // 3: kpmasterproto.OccupationFindAllResponse.paginations:type_name -> kpmasterproto.PaginationMaster
	2,  // 4: kpmasterproto.OccupationFindIDResponse.occupation:type_name -> kpmasterproto.OccupationInfo
	0,  // 5: kpmasterproto.OccupationCreateRequest.status_blacklist:type_name -> kpmasterproto.StatusBlacklist
	1,  // 6: kpmasterproto.OccupationCreateResponse.create_response:type_name -> kpmasterproto.OccupationWriteResponse
	1,  // 7: kpmasterproto.OccupationDeleteResponse.delete_response:type_name -> kpmasterproto.OccupationWriteResponse
	0,  // 8: kpmasterproto.OccupationUpdateRequest.status_blacklist:type_name -> kpmasterproto.StatusBlacklist
	1,  // 9: kpmasterproto.OccupationUpdateResponse.update_response:type_name -> kpmasterproto.OccupationWriteResponse
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_occupation_occupation_proto_init() }
func file_occupation_occupation_proto_init() {
	if File_occupation_occupation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_occupation_occupation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationWriteResponse); i {
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
		file_occupation_occupation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationInfo); i {
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
		file_occupation_occupation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationFindAllRequest); i {
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
		file_occupation_occupation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationFindAllResponse); i {
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
		file_occupation_occupation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationFindIDRequest); i {
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
		file_occupation_occupation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationFindIDResponse); i {
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
		file_occupation_occupation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationCreateRequest); i {
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
		file_occupation_occupation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationCreateResponse); i {
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
		file_occupation_occupation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationDeleteRequest); i {
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
		file_occupation_occupation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationDeleteResponse); i {
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
		file_occupation_occupation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationUpdateRequest); i {
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
		file_occupation_occupation_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupationUpdateResponse); i {
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
			RawDescriptor: file_occupation_occupation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_occupation_occupation_proto_goTypes,
		DependencyIndexes: file_occupation_occupation_proto_depIdxs,
		EnumInfos:         file_occupation_occupation_proto_enumTypes,
		MessageInfos:      file_occupation_occupation_proto_msgTypes,
	}.Build()
	File_occupation_occupation_proto = out.File
	file_occupation_occupation_proto_rawDesc = nil
	file_occupation_occupation_proto_goTypes = nil
	file_occupation_occupation_proto_depIdxs = nil
}
