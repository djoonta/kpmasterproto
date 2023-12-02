// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bank/bank.proto

package bank

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

type BankWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Messsage string `protobuf:"bytes,2,opt,name=messsage,proto3" json:"messsage,omitempty"`
}

func (x *BankWriteResponse) Reset() {
	*x = BankWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankWriteResponse) ProtoMessage() {}

func (x *BankWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankWriteResponse.ProtoReflect.Descriptor instead.
func (*BankWriteResponse) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{0}
}

func (x *BankWriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BankWriteResponse) GetMesssage() string {
	if x != nil {
		return x.Messsage
	}
	return ""
}

type BankInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Code        string                 `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	SwiftCode   string                 `protobuf:"bytes,5,opt,name=swift_code,json=swiftCode,proto3" json:"swift_code,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *BankInfo) Reset() {
	*x = BankInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankInfo) ProtoMessage() {}

func (x *BankInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankInfo.ProtoReflect.Descriptor instead.
func (*BankInfo) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{1}
}

func (x *BankInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BankInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BankInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BankInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BankInfo) GetSwiftCode() string {
	if x != nil {
		return x.SwiftCode
	}
	return ""
}

func (x *BankInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type BankFindIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BankFindIDRequest) Reset() {
	*x = BankFindIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankFindIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankFindIDRequest) ProtoMessage() {}

func (x *BankFindIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankFindIDRequest.ProtoReflect.Descriptor instead.
func (*BankFindIDRequest) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{2}
}

func (x *BankFindIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BankFindIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bank *BankInfo `protobuf:"bytes,1,opt,name=bank,proto3" json:"bank,omitempty"`
}

func (x *BankFindIDResponse) Reset() {
	*x = BankFindIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankFindIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankFindIDResponse) ProtoMessage() {}

func (x *BankFindIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankFindIDResponse.ProtoReflect.Descriptor instead.
func (*BankFindIDResponse) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{3}
}

func (x *BankFindIDResponse) GetBank() *BankInfo {
	if x != nil {
		return x.Bank
	}
	return nil
}

type BankFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *BankFindAllRequest) Reset() {
	*x = BankFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankFindAllRequest) ProtoMessage() {}

func (x *BankFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankFindAllRequest.ProtoReflect.Descriptor instead.
func (*BankFindAllRequest) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{4}
}

func (x *BankFindAllRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *BankFindAllRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type BankFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banks       []*BankInfo                  `protobuf:"bytes,1,rep,name=banks,proto3" json:"banks,omitempty"`
	Paginations *pagination.PaginationMaster `protobuf:"bytes,2,opt,name=paginations,proto3" json:"paginations,omitempty"`
}

func (x *BankFindAllResponse) Reset() {
	*x = BankFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankFindAllResponse) ProtoMessage() {}

func (x *BankFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankFindAllResponse.ProtoReflect.Descriptor instead.
func (*BankFindAllResponse) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{5}
}

func (x *BankFindAllResponse) GetBanks() []*BankInfo {
	if x != nil {
		return x.Banks
	}
	return nil
}

func (x *BankFindAllResponse) GetPaginations() *pagination.PaginationMaster {
	if x != nil {
		return x.Paginations
	}
	return nil
}

type BankCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	SwiftCode   string `protobuf:"bytes,5,opt,name=swift_code,json=swiftCode,proto3" json:"swift_code,omitempty"`
}

func (x *BankCreateRequest) Reset() {
	*x = BankCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCreateRequest) ProtoMessage() {}

func (x *BankCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCreateRequest.ProtoReflect.Descriptor instead.
func (*BankCreateRequest) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{6}
}

func (x *BankCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BankCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BankCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BankCreateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BankCreateRequest) GetSwiftCode() string {
	if x != nil {
		return x.SwiftCode
	}
	return ""
}

type BankCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateResponse *BankWriteResponse `protobuf:"bytes,1,opt,name=create_response,json=createResponse,proto3" json:"create_response,omitempty"`
}

func (x *BankCreateResponse) Reset() {
	*x = BankCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCreateResponse) ProtoMessage() {}

func (x *BankCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCreateResponse.ProtoReflect.Descriptor instead.
func (*BankCreateResponse) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{7}
}

func (x *BankCreateResponse) GetCreateResponse() *BankWriteResponse {
	if x != nil {
		return x.CreateResponse
	}
	return nil
}

type BankDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BankDeleteRequest) Reset() {
	*x = BankDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankDeleteRequest) ProtoMessage() {}

func (x *BankDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankDeleteRequest.ProtoReflect.Descriptor instead.
func (*BankDeleteRequest) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{8}
}

func (x *BankDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BankDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteResponse *BankWriteResponse `protobuf:"bytes,1,opt,name=delete_response,json=deleteResponse,proto3" json:"delete_response,omitempty"`
}

func (x *BankDeleteResponse) Reset() {
	*x = BankDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankDeleteResponse) ProtoMessage() {}

func (x *BankDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankDeleteResponse.ProtoReflect.Descriptor instead.
func (*BankDeleteResponse) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{9}
}

func (x *BankDeleteResponse) GetDeleteResponse() *BankWriteResponse {
	if x != nil {
		return x.DeleteResponse
	}
	return nil
}

type BankUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	SwiftCode   string `protobuf:"bytes,5,opt,name=swift_code,json=swiftCode,proto3" json:"swift_code,omitempty"`
}

func (x *BankUpdateRequest) Reset() {
	*x = BankUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankUpdateRequest) ProtoMessage() {}

func (x *BankUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankUpdateRequest.ProtoReflect.Descriptor instead.
func (*BankUpdateRequest) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{10}
}

func (x *BankUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BankUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BankUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BankUpdateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BankUpdateRequest) GetSwiftCode() string {
	if x != nil {
		return x.SwiftCode
	}
	return ""
}

type BankUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateResponse *BankWriteResponse `protobuf:"bytes,1,opt,name=update_response,json=updateResponse,proto3" json:"update_response,omitempty"`
}

func (x *BankUpdateResponse) Reset() {
	*x = BankUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankUpdateResponse) ProtoMessage() {}

func (x *BankUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankUpdateResponse.ProtoReflect.Descriptor instead.
func (*BankUpdateResponse) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{11}
}

func (x *BankUpdateResponse) GetUpdateResponse() *BankWriteResponse {
	if x != nil {
		return x.UpdateResponse
	}
	return nil
}

var File_bank_bank_proto protoreflect.FileDescriptor

var file_bank_bank_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x08, 0x42, 0x61,
	0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x77, 0x69, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x77, 0x69, 0x66, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x42, 0x61,
	0x6e, 0x6b, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x12, 0x42, 0x61, 0x6e, 0x6b, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x61,
	0x6e, 0x6b, 0x22, 0x3e, 0x0a, 0x12, 0x42, 0x61, 0x6e, 0x6b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x6e, 0x6b, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x61,
	0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x70, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8c, 0x01, 0x0a,
	0x11, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x77, 0x69, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x77, 0x69, 0x66, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x42,
	0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x70, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11,
	0x42, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5f, 0x0a, 0x12, 0x42, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x77, 0x69, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x77, 0x69, 0x66, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x5f, 0x0a, 0x12, 0x42, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61, 0x2f, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bank_bank_proto_rawDescOnce sync.Once
	file_bank_bank_proto_rawDescData = file_bank_bank_proto_rawDesc
)

func file_bank_bank_proto_rawDescGZIP() []byte {
	file_bank_bank_proto_rawDescOnce.Do(func() {
		file_bank_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_bank_proto_rawDescData)
	})
	return file_bank_bank_proto_rawDescData
}

var file_bank_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_bank_bank_proto_goTypes = []interface{}{
	(*BankWriteResponse)(nil),           // 0: kpmasterproto.BankWriteResponse
	(*BankInfo)(nil),                    // 1: kpmasterproto.BankInfo
	(*BankFindIDRequest)(nil),           // 2: kpmasterproto.BankFindIDRequest
	(*BankFindIDResponse)(nil),          // 3: kpmasterproto.BankFindIDResponse
	(*BankFindAllRequest)(nil),          // 4: kpmasterproto.BankFindAllRequest
	(*BankFindAllResponse)(nil),         // 5: kpmasterproto.BankFindAllResponse
	(*BankCreateRequest)(nil),           // 6: kpmasterproto.BankCreateRequest
	(*BankCreateResponse)(nil),          // 7: kpmasterproto.BankCreateResponse
	(*BankDeleteRequest)(nil),           // 8: kpmasterproto.BankDeleteRequest
	(*BankDeleteResponse)(nil),          // 9: kpmasterproto.BankDeleteResponse
	(*BankUpdateRequest)(nil),           // 10: kpmasterproto.BankUpdateRequest
	(*BankUpdateResponse)(nil),          // 11: kpmasterproto.BankUpdateResponse
	(*timestamppb.Timestamp)(nil),       // 12: google.protobuf.Timestamp
	(*pagination.PaginationMaster)(nil), // 13: kpmasterproto.PaginationMaster
}
var file_bank_bank_proto_depIdxs = []int32{
	12, // 0: kpmasterproto.BankInfo.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 1: kpmasterproto.BankFindIDResponse.bank:type_name -> kpmasterproto.BankInfo
	1,  // 2: kpmasterproto.BankFindAllResponse.banks:type_name -> kpmasterproto.BankInfo
	13, // 3: kpmasterproto.BankFindAllResponse.paginations:type_name -> kpmasterproto.PaginationMaster
	0,  // 4: kpmasterproto.BankCreateResponse.create_response:type_name -> kpmasterproto.BankWriteResponse
	0,  // 5: kpmasterproto.BankDeleteResponse.delete_response:type_name -> kpmasterproto.BankWriteResponse
	0,  // 6: kpmasterproto.BankUpdateResponse.update_response:type_name -> kpmasterproto.BankWriteResponse
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_bank_bank_proto_init() }
func file_bank_bank_proto_init() {
	if File_bank_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankWriteResponse); i {
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
		file_bank_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankInfo); i {
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
		file_bank_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankFindIDRequest); i {
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
		file_bank_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankFindIDResponse); i {
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
		file_bank_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankFindAllRequest); i {
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
		file_bank_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankFindAllResponse); i {
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
		file_bank_bank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCreateRequest); i {
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
		file_bank_bank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCreateResponse); i {
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
		file_bank_bank_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankDeleteRequest); i {
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
		file_bank_bank_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankDeleteResponse); i {
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
		file_bank_bank_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankUpdateRequest); i {
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
		file_bank_bank_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankUpdateResponse); i {
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
			RawDescriptor: file_bank_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bank_bank_proto_goTypes,
		DependencyIndexes: file_bank_bank_proto_depIdxs,
		MessageInfos:      file_bank_bank_proto_msgTypes,
	}.Build()
	File_bank_bank_proto = out.File
	file_bank_bank_proto_rawDesc = nil
	file_bank_bank_proto_goTypes = nil
	file_bank_bank_proto_depIdxs = nil
}
