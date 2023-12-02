// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: encumbrance/encumbrance.proto

package encumbrance

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

type EncumbranceWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Messsage string `protobuf:"bytes,2,opt,name=messsage,proto3" json:"messsage,omitempty"`
}

func (x *EncumbranceWriteResponse) Reset() {
	*x = EncumbranceWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceWriteResponse) ProtoMessage() {}

func (x *EncumbranceWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceWriteResponse.ProtoReflect.Descriptor instead.
func (*EncumbranceWriteResponse) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{0}
}

func (x *EncumbranceWriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EncumbranceWriteResponse) GetMesssage() string {
	if x != nil {
		return x.Messsage
	}
	return ""
}

type EncumbranceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *EncumbranceInfo) Reset() {
	*x = EncumbranceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceInfo) ProtoMessage() {}

func (x *EncumbranceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceInfo.ProtoReflect.Descriptor instead.
func (*EncumbranceInfo) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{1}
}

func (x *EncumbranceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EncumbranceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EncumbranceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EncumbranceInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *EncumbranceInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type EncumbranceFindIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EncumbranceFindIDRequest) Reset() {
	*x = EncumbranceFindIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceFindIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceFindIDRequest) ProtoMessage() {}

func (x *EncumbranceFindIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceFindIDRequest.ProtoReflect.Descriptor instead.
func (*EncumbranceFindIDRequest) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{2}
}

func (x *EncumbranceFindIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EncumbranceFindIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encumbrance *EncumbranceInfo `protobuf:"bytes,1,opt,name=encumbrance,proto3" json:"encumbrance,omitempty"`
}

func (x *EncumbranceFindIDResponse) Reset() {
	*x = EncumbranceFindIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceFindIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceFindIDResponse) ProtoMessage() {}

func (x *EncumbranceFindIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceFindIDResponse.ProtoReflect.Descriptor instead.
func (*EncumbranceFindIDResponse) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{3}
}

func (x *EncumbranceFindIDResponse) GetEncumbrance() *EncumbranceInfo {
	if x != nil {
		return x.Encumbrance
	}
	return nil
}

type EncumbranceFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *EncumbranceFindAllRequest) Reset() {
	*x = EncumbranceFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceFindAllRequest) ProtoMessage() {}

func (x *EncumbranceFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceFindAllRequest.ProtoReflect.Descriptor instead.
func (*EncumbranceFindAllRequest) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{4}
}

func (x *EncumbranceFindAllRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *EncumbranceFindAllRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type EncumbranceFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encumbrances []*EncumbranceInfo           `protobuf:"bytes,1,rep,name=encumbrances,proto3" json:"encumbrances,omitempty"`
	Paginations  *pagination.PaginationMaster `protobuf:"bytes,2,opt,name=paginations,proto3" json:"paginations,omitempty"`
}

func (x *EncumbranceFindAllResponse) Reset() {
	*x = EncumbranceFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceFindAllResponse) ProtoMessage() {}

func (x *EncumbranceFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceFindAllResponse.ProtoReflect.Descriptor instead.
func (*EncumbranceFindAllResponse) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{5}
}

func (x *EncumbranceFindAllResponse) GetEncumbrances() []*EncumbranceInfo {
	if x != nil {
		return x.Encumbrances
	}
	return nil
}

func (x *EncumbranceFindAllResponse) GetPaginations() *pagination.PaginationMaster {
	if x != nil {
		return x.Paginations
	}
	return nil
}

type EncumbranceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EncumbranceCreateRequest) Reset() {
	*x = EncumbranceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceCreateRequest) ProtoMessage() {}

func (x *EncumbranceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceCreateRequest.ProtoReflect.Descriptor instead.
func (*EncumbranceCreateRequest) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{6}
}

func (x *EncumbranceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EncumbranceCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EncumbranceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateResponse *EncumbranceWriteResponse `protobuf:"bytes,1,opt,name=create_response,json=createResponse,proto3" json:"create_response,omitempty"`
}

func (x *EncumbranceCreateResponse) Reset() {
	*x = EncumbranceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceCreateResponse) ProtoMessage() {}

func (x *EncumbranceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceCreateResponse.ProtoReflect.Descriptor instead.
func (*EncumbranceCreateResponse) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{7}
}

func (x *EncumbranceCreateResponse) GetCreateResponse() *EncumbranceWriteResponse {
	if x != nil {
		return x.CreateResponse
	}
	return nil
}

type EncumbranceDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EncumbranceDeleteRequest) Reset() {
	*x = EncumbranceDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceDeleteRequest) ProtoMessage() {}

func (x *EncumbranceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceDeleteRequest.ProtoReflect.Descriptor instead.
func (*EncumbranceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{8}
}

func (x *EncumbranceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EncumbranceDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteResponse *EncumbranceWriteResponse `protobuf:"bytes,1,opt,name=delete_response,json=deleteResponse,proto3" json:"delete_response,omitempty"`
}

func (x *EncumbranceDeleteResponse) Reset() {
	*x = EncumbranceDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceDeleteResponse) ProtoMessage() {}

func (x *EncumbranceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceDeleteResponse.ProtoReflect.Descriptor instead.
func (*EncumbranceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{9}
}

func (x *EncumbranceDeleteResponse) GetDeleteResponse() *EncumbranceWriteResponse {
	if x != nil {
		return x.DeleteResponse
	}
	return nil
}

type EncumbranceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EncumbranceUpdateRequest) Reset() {
	*x = EncumbranceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceUpdateRequest) ProtoMessage() {}

func (x *EncumbranceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceUpdateRequest.ProtoReflect.Descriptor instead.
func (*EncumbranceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{10}
}

func (x *EncumbranceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EncumbranceUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EncumbranceUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EncumbranceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateResponse *EncumbranceWriteResponse `protobuf:"bytes,1,opt,name=update_response,json=updateResponse,proto3" json:"update_response,omitempty"`
}

func (x *EncumbranceUpdateResponse) Reset() {
	*x = EncumbranceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encumbrance_encumbrance_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncumbranceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncumbranceUpdateResponse) ProtoMessage() {}

func (x *EncumbranceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encumbrance_encumbrance_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncumbranceUpdateResponse.ProtoReflect.Descriptor instead.
func (*EncumbranceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_encumbrance_encumbrance_proto_rawDescGZIP(), []int{11}
}

func (x *EncumbranceUpdateResponse) GetUpdateResponse() *EncumbranceWriteResponse {
	if x != nil {
		return x.UpdateResponse
	}
	return nil
}

var File_encumbrance_encumbrance_proto protoreflect.FileDescriptor

var file_encumbrance_encumbrance_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x65, 0x6e,
	0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x18,
	0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb1,
	0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x2a, 0x0a, 0x18, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x65, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d,
	0x0a, 0x19, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x69, 0x6e,
	0x64, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x65,
	0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0b, 0x65, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x45, 0x0a,
	0x19, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x1a, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x70, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b,
	0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x50, 0x0a, 0x18, 0x45, 0x6e,
	0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x19,
	0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x45,
	0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x19, 0x45, 0x6e, 0x63, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e,
	0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x18, 0x45, 0x6e, 0x63, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x19, 0x45, 0x6e, 0x63, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6e, 0x63, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61, 0x2f, 0x6b, 0x70,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encumbrance_encumbrance_proto_rawDescOnce sync.Once
	file_encumbrance_encumbrance_proto_rawDescData = file_encumbrance_encumbrance_proto_rawDesc
)

func file_encumbrance_encumbrance_proto_rawDescGZIP() []byte {
	file_encumbrance_encumbrance_proto_rawDescOnce.Do(func() {
		file_encumbrance_encumbrance_proto_rawDescData = protoimpl.X.CompressGZIP(file_encumbrance_encumbrance_proto_rawDescData)
	})
	return file_encumbrance_encumbrance_proto_rawDescData
}

var file_encumbrance_encumbrance_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_encumbrance_encumbrance_proto_goTypes = []interface{}{
	(*EncumbranceWriteResponse)(nil),    // 0: kpmasterproto.EncumbranceWriteResponse
	(*EncumbranceInfo)(nil),             // 1: kpmasterproto.EncumbranceInfo
	(*EncumbranceFindIDRequest)(nil),    // 2: kpmasterproto.EncumbranceFindIDRequest
	(*EncumbranceFindIDResponse)(nil),   // 3: kpmasterproto.EncumbranceFindIDResponse
	(*EncumbranceFindAllRequest)(nil),   // 4: kpmasterproto.EncumbranceFindAllRequest
	(*EncumbranceFindAllResponse)(nil),  // 5: kpmasterproto.EncumbranceFindAllResponse
	(*EncumbranceCreateRequest)(nil),    // 6: kpmasterproto.EncumbranceCreateRequest
	(*EncumbranceCreateResponse)(nil),   // 7: kpmasterproto.EncumbranceCreateResponse
	(*EncumbranceDeleteRequest)(nil),    // 8: kpmasterproto.EncumbranceDeleteRequest
	(*EncumbranceDeleteResponse)(nil),   // 9: kpmasterproto.EncumbranceDeleteResponse
	(*EncumbranceUpdateRequest)(nil),    // 10: kpmasterproto.EncumbranceUpdateRequest
	(*EncumbranceUpdateResponse)(nil),   // 11: kpmasterproto.EncumbranceUpdateResponse
	(*timestamppb.Timestamp)(nil),       // 12: google.protobuf.Timestamp
	(*pagination.PaginationMaster)(nil), // 13: kpmasterproto.PaginationMaster
}
var file_encumbrance_encumbrance_proto_depIdxs = []int32{
	12, // 0: kpmasterproto.EncumbranceInfo.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 1: kpmasterproto.EncumbranceFindIDResponse.encumbrance:type_name -> kpmasterproto.EncumbranceInfo
	1,  // 2: kpmasterproto.EncumbranceFindAllResponse.encumbrances:type_name -> kpmasterproto.EncumbranceInfo
	13, // 3: kpmasterproto.EncumbranceFindAllResponse.paginations:type_name -> kpmasterproto.PaginationMaster
	0,  // 4: kpmasterproto.EncumbranceCreateResponse.create_response:type_name -> kpmasterproto.EncumbranceWriteResponse
	0,  // 5: kpmasterproto.EncumbranceDeleteResponse.delete_response:type_name -> kpmasterproto.EncumbranceWriteResponse
	0,  // 6: kpmasterproto.EncumbranceUpdateResponse.update_response:type_name -> kpmasterproto.EncumbranceWriteResponse
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_encumbrance_encumbrance_proto_init() }
func file_encumbrance_encumbrance_proto_init() {
	if File_encumbrance_encumbrance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encumbrance_encumbrance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceWriteResponse); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceInfo); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceFindIDRequest); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceFindIDResponse); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceFindAllRequest); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceFindAllResponse); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceCreateRequest); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceCreateResponse); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceDeleteRequest); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceDeleteResponse); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceUpdateRequest); i {
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
		file_encumbrance_encumbrance_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncumbranceUpdateResponse); i {
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
			RawDescriptor: file_encumbrance_encumbrance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encumbrance_encumbrance_proto_goTypes,
		DependencyIndexes: file_encumbrance_encumbrance_proto_depIdxs,
		MessageInfos:      file_encumbrance_encumbrance_proto_msgTypes,
	}.Build()
	File_encumbrance_encumbrance_proto = out.File
	file_encumbrance_encumbrance_proto_rawDesc = nil
	file_encumbrance_encumbrance_proto_goTypes = nil
	file_encumbrance_encumbrance_proto_depIdxs = nil
}
