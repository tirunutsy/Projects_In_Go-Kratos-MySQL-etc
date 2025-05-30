// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/helloworld/v1/customer.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CustomerData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age           int32                  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Gender        string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Emails        []*Emaildata           `protobuf:"bytes,6,rep,name=emails,proto3" json:"emails,omitempty"`
	Phones        []*Phones              `protobuf:"bytes,7,rep,name=phones,proto3" json:"phones,omitempty"`
	Addresses     []*Address             `protobuf:"bytes,8,rep,name=addresses,proto3" json:"addresses,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerData) Reset() {
	*x = CustomerData{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerData) ProtoMessage() {}

func (x *CustomerData) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerData.ProtoReflect.Descriptor instead.
func (*CustomerData) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CustomerData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CustomerData) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CustomerData) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CustomerData) GetEmails() []*Emaildata {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *CustomerData) GetPhones() []*Phones {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *CustomerData) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *CustomerData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CustomerData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Emaildata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	IsPrimary     bool                   `protobuf:"varint,4,opt,name=is_primary,json=isPrimary,proto3" json:"is_primary,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Emaildata) Reset() {
	*x = Emaildata{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Emaildata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emaildata) ProtoMessage() {}

func (x *Emaildata) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emaildata.ProtoReflect.Descriptor instead.
func (*Emaildata) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{1}
}

func (x *Emaildata) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Emaildata) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Emaildata) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Emaildata) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

type Phones struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Number        string                 `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	IsPrimary     bool                   `protobuf:"varint,4,opt,name=is_primary,json=isPrimary,proto3" json:"is_primary,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Phones) Reset() {
	*x = Phones{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Phones) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phones) ProtoMessage() {}

func (x *Phones) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phones.ProtoReflect.Descriptor instead.
func (*Phones) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{2}
}

func (x *Phones) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Phones) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Phones) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Phones) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Pincode       string                 `protobuf:"bytes,3,opt,name=pincode,proto3" json:"pincode,omitempty"`
	Type          string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetPincode() string {
	if x != nil {
		return x.Pincode
	}
	return ""
}

func (x *Address) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstName     string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age           int32                  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Emails        []*Emaildata           `protobuf:"bytes,5,rep,name=emails,proto3" json:"emails,omitempty"`
	Phones        []*Phones              `protobuf:"bytes,6,rep,name=phones,proto3" json:"phones,omitempty"`
	Addresses     []*Address             `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCustomerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateCustomerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateCustomerRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateCustomerRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateCustomerRequest) GetEmails() []*Emaildata {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *CreateCustomerRequest) GetPhones() []*Phones {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *CreateCustomerRequest) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type CreateCustomerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *CustomerData          `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCustomerReply) Reset() {
	*x = CreateCustomerReply{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerReply) ProtoMessage() {}

func (x *CreateCustomerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerReply.ProtoReflect.Descriptor instead.
func (*CreateCustomerReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCustomerReply) GetCustomer() *CustomerData {
	if x != nil {
		return x.Customer
	}
	return nil
}

type UpdateCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age           int32                  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Gender        string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Emails        []*Emaildata           `protobuf:"bytes,6,rep,name=emails,proto3" json:"emails,omitempty"`
	Phones        []*Phones              `protobuf:"bytes,7,rep,name=phones,proto3" json:"phones,omitempty"`
	Addresses     []*Address             `protobuf:"bytes,8,rep,name=addresses,proto3" json:"addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCustomerRequest) Reset() {
	*x = UpdateCustomerRequest{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerRequest) ProtoMessage() {}

func (x *UpdateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCustomerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCustomerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateCustomerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateCustomerRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateCustomerRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateCustomerRequest) GetEmails() []*Emaildata {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *UpdateCustomerRequest) GetPhones() []*Phones {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *UpdateCustomerRequest) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type UpdateCustomerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *CustomerData          `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCustomerReply) Reset() {
	*x = UpdateCustomerReply{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerReply) ProtoMessage() {}

func (x *UpdateCustomerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerReply.ProtoReflect.Descriptor instead.
func (*UpdateCustomerReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCustomerReply) GetCustomer() *CustomerData {
	if x != nil {
		return x.Customer
	}
	return nil
}

type DeleteCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCustomerRequest) Reset() {
	*x = DeleteCustomerRequest{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerRequest) ProtoMessage() {}

func (x *DeleteCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerRequest.ProtoReflect.Descriptor instead.
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCustomerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCustomerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Example field to indicate success or failure
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCustomerReply) Reset() {
	*x = DeleteCustomerReply{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCustomerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerReply) ProtoMessage() {}

func (x *DeleteCustomerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerReply.ProtoReflect.Descriptor instead.
func (*DeleteCustomerReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCustomerReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{10}
}

func (x *GetCustomerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCustomerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *CustomerData          `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerReply) Reset() {
	*x = GetCustomerReply{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerReply) ProtoMessage() {}

func (x *GetCustomerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerReply.ProtoReflect.Descriptor instead.
func (*GetCustomerReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{11}
}

func (x *GetCustomerReply) GetCustomer() *CustomerData {
	if x != nil {
		return x.Customer
	}
	return nil
}

type ListCustomersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCustomersRequest) Reset() {
	*x = ListCustomersRequest{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersRequest) ProtoMessage() {}

func (x *ListCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersRequest.ProtoReflect.Descriptor instead.
func (*ListCustomersRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{12}
}

type ListCustomersReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customers     []*CustomerData        `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCustomersReply) Reset() {
	*x = ListCustomersReply{}
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCustomersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersReply) ProtoMessage() {}

func (x *ListCustomersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_customer_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersReply.ProtoReflect.Descriptor instead.
func (*ListCustomersReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_customer_proto_rawDescGZIP(), []int{13}
}

func (x *ListCustomersReply) GetCustomers() []*CustomerData {
	if x != nil {
		return x.Customers
	}
	return nil
}

var File_api_helloworld_v1_customer_proto protoreflect.FileDescriptor

const file_api_helloworld_v1_customer_proto_rawDesc = "" +
	"\n" +
	" api/helloworld/v1/customer.proto\x12\x11api.helloworld.v1\x1a\x1cgoogle/api/annotations.proto\"\xe5\x02\n" +
	"\fCustomerData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1d\n" +
	"\n" +
	"first_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x10\n" +
	"\x03age\x18\x04 \x01(\x05R\x03age\x12\x16\n" +
	"\x06gender\x18\x05 \x01(\tR\x06gender\x124\n" +
	"\x06emails\x18\x06 \x03(\v2\x1c.api.helloworld.v1.EmaildataR\x06emails\x121\n" +
	"\x06phones\x18\a \x03(\v2\x19.api.helloworld.v1.PhonesR\x06phones\x128\n" +
	"\taddresses\x18\b \x03(\v2\x1a.api.helloworld.v1.AddressR\taddresses\x12\x1d\n" +
	"\n" +
	"created_at\x18\t \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\n" +
	" \x01(\tR\tupdatedAt\"d\n" +
	"\tEmaildata\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\x12\x1d\n" +
	"\n" +
	"is_primary\x18\x04 \x01(\bR\tisPrimary\"c\n" +
	"\x06Phones\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x16\n" +
	"\x06number\x18\x02 \x01(\tR\x06number\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\x12\x1d\n" +
	"\n" +
	"is_primary\x18\x04 \x01(\bR\tisPrimary\"a\n" +
	"\aAddress\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x12\x18\n" +
	"\apincode\x18\x03 \x01(\tR\apincode\x12\x12\n" +
	"\x04type\x18\x04 \x01(\tR\x04type\"\xa0\x02\n" +
	"\x15CreateCustomerRequest\x12\x1d\n" +
	"\n" +
	"first_name\x18\x01 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x02 \x01(\tR\blastName\x12\x10\n" +
	"\x03age\x18\x03 \x01(\x05R\x03age\x12\x16\n" +
	"\x06gender\x18\x04 \x01(\tR\x06gender\x124\n" +
	"\x06emails\x18\x05 \x03(\v2\x1c.api.helloworld.v1.EmaildataR\x06emails\x121\n" +
	"\x06phones\x18\x06 \x03(\v2\x19.api.helloworld.v1.PhonesR\x06phones\x128\n" +
	"\taddresses\x18\a \x03(\v2\x1a.api.helloworld.v1.AddressR\taddresses\"R\n" +
	"\x13CreateCustomerReply\x12;\n" +
	"\bcustomer\x18\x01 \x01(\v2\x1f.api.helloworld.v1.CustomerDataR\bcustomer\"\xb0\x02\n" +
	"\x15UpdateCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1d\n" +
	"\n" +
	"first_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x10\n" +
	"\x03age\x18\x04 \x01(\x05R\x03age\x12\x16\n" +
	"\x06gender\x18\x05 \x01(\tR\x06gender\x124\n" +
	"\x06emails\x18\x06 \x03(\v2\x1c.api.helloworld.v1.EmaildataR\x06emails\x121\n" +
	"\x06phones\x18\a \x03(\v2\x19.api.helloworld.v1.PhonesR\x06phones\x128\n" +
	"\taddresses\x18\b \x03(\v2\x1a.api.helloworld.v1.AddressR\taddresses\"R\n" +
	"\x13UpdateCustomerReply\x12;\n" +
	"\bcustomer\x18\x01 \x01(\v2\x1f.api.helloworld.v1.CustomerDataR\bcustomer\"'\n" +
	"\x15DeleteCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"/\n" +
	"\x13DeleteCustomerReply\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"$\n" +
	"\x12GetCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"O\n" +
	"\x10GetCustomerReply\x12;\n" +
	"\bcustomer\x18\x01 \x01(\v2\x1f.api.helloworld.v1.CustomerDataR\bcustomer\"\x16\n" +
	"\x14ListCustomersRequest\"S\n" +
	"\x12ListCustomersReply\x12=\n" +
	"\tcustomers\x18\x01 \x03(\v2\x1f.api.helloworld.v1.CustomerDataR\tcustomers2\xf6\x04\n" +
	"\bCustomer\x12{\n" +
	"\x0eCreateCustomer\x12(.api.helloworld.v1.CreateCustomerRequest\x1a&.api.helloworld.v1.CreateCustomerReply\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/v1/customer\x12\x80\x01\n" +
	"\x0eUpdateCustomer\x12(.api.helloworld.v1.UpdateCustomerRequest\x1a&.api.helloworld.v1.UpdateCustomerReply\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\x1a\x11/v1/customer/{id}\x12}\n" +
	"\x0eDeleteCustomer\x12(.api.helloworld.v1.DeleteCustomerRequest\x1a&.api.helloworld.v1.DeleteCustomerReply\"\x19\x82\xd3\xe4\x93\x02\x13*\x11/v1/customer/{id}\x12t\n" +
	"\vGetCustomer\x12%.api.helloworld.v1.GetCustomerRequest\x1a#.api.helloworld.v1.GetCustomerReply\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/customer/{id}\x12u\n" +
	"\fListCustomer\x12'.api.helloworld.v1.ListCustomersRequest\x1a%.api.helloworld.v1.ListCustomersReply\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/v1/customersB<\n" +
	"\x11api.helloworld.v1P\x01Z%customer-service/api/helloworld/v1;v1b\x06proto3"

var (
	file_api_helloworld_v1_customer_proto_rawDescOnce sync.Once
	file_api_helloworld_v1_customer_proto_rawDescData []byte
)

func file_api_helloworld_v1_customer_proto_rawDescGZIP() []byte {
	file_api_helloworld_v1_customer_proto_rawDescOnce.Do(func() {
		file_api_helloworld_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_helloworld_v1_customer_proto_rawDesc), len(file_api_helloworld_v1_customer_proto_rawDesc)))
	})
	return file_api_helloworld_v1_customer_proto_rawDescData
}

var file_api_helloworld_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_helloworld_v1_customer_proto_goTypes = []any{
	(*CustomerData)(nil),          // 0: api.helloworld.v1.CustomerData
	(*Emaildata)(nil),             // 1: api.helloworld.v1.Emaildata
	(*Phones)(nil),                // 2: api.helloworld.v1.Phones
	(*Address)(nil),               // 3: api.helloworld.v1.Address
	(*CreateCustomerRequest)(nil), // 4: api.helloworld.v1.CreateCustomerRequest
	(*CreateCustomerReply)(nil),   // 5: api.helloworld.v1.CreateCustomerReply
	(*UpdateCustomerRequest)(nil), // 6: api.helloworld.v1.UpdateCustomerRequest
	(*UpdateCustomerReply)(nil),   // 7: api.helloworld.v1.UpdateCustomerReply
	(*DeleteCustomerRequest)(nil), // 8: api.helloworld.v1.DeleteCustomerRequest
	(*DeleteCustomerReply)(nil),   // 9: api.helloworld.v1.DeleteCustomerReply
	(*GetCustomerRequest)(nil),    // 10: api.helloworld.v1.GetCustomerRequest
	(*GetCustomerReply)(nil),      // 11: api.helloworld.v1.GetCustomerReply
	(*ListCustomersRequest)(nil),  // 12: api.helloworld.v1.ListCustomersRequest
	(*ListCustomersReply)(nil),    // 13: api.helloworld.v1.ListCustomersReply
}
var file_api_helloworld_v1_customer_proto_depIdxs = []int32{
	1,  // 0: api.helloworld.v1.CustomerData.emails:type_name -> api.helloworld.v1.Emaildata
	2,  // 1: api.helloworld.v1.CustomerData.phones:type_name -> api.helloworld.v1.Phones
	3,  // 2: api.helloworld.v1.CustomerData.addresses:type_name -> api.helloworld.v1.Address
	1,  // 3: api.helloworld.v1.CreateCustomerRequest.emails:type_name -> api.helloworld.v1.Emaildata
	2,  // 4: api.helloworld.v1.CreateCustomerRequest.phones:type_name -> api.helloworld.v1.Phones
	3,  // 5: api.helloworld.v1.CreateCustomerRequest.addresses:type_name -> api.helloworld.v1.Address
	0,  // 6: api.helloworld.v1.CreateCustomerReply.customer:type_name -> api.helloworld.v1.CustomerData
	1,  // 7: api.helloworld.v1.UpdateCustomerRequest.emails:type_name -> api.helloworld.v1.Emaildata
	2,  // 8: api.helloworld.v1.UpdateCustomerRequest.phones:type_name -> api.helloworld.v1.Phones
	3,  // 9: api.helloworld.v1.UpdateCustomerRequest.addresses:type_name -> api.helloworld.v1.Address
	0,  // 10: api.helloworld.v1.UpdateCustomerReply.customer:type_name -> api.helloworld.v1.CustomerData
	0,  // 11: api.helloworld.v1.GetCustomerReply.customer:type_name -> api.helloworld.v1.CustomerData
	0,  // 12: api.helloworld.v1.ListCustomersReply.customers:type_name -> api.helloworld.v1.CustomerData
	4,  // 13: api.helloworld.v1.Customer.CreateCustomer:input_type -> api.helloworld.v1.CreateCustomerRequest
	6,  // 14: api.helloworld.v1.Customer.UpdateCustomer:input_type -> api.helloworld.v1.UpdateCustomerRequest
	8,  // 15: api.helloworld.v1.Customer.DeleteCustomer:input_type -> api.helloworld.v1.DeleteCustomerRequest
	10, // 16: api.helloworld.v1.Customer.GetCustomer:input_type -> api.helloworld.v1.GetCustomerRequest
	12, // 17: api.helloworld.v1.Customer.ListCustomer:input_type -> api.helloworld.v1.ListCustomersRequest
	5,  // 18: api.helloworld.v1.Customer.CreateCustomer:output_type -> api.helloworld.v1.CreateCustomerReply
	7,  // 19: api.helloworld.v1.Customer.UpdateCustomer:output_type -> api.helloworld.v1.UpdateCustomerReply
	9,  // 20: api.helloworld.v1.Customer.DeleteCustomer:output_type -> api.helloworld.v1.DeleteCustomerReply
	11, // 21: api.helloworld.v1.Customer.GetCustomer:output_type -> api.helloworld.v1.GetCustomerReply
	13, // 22: api.helloworld.v1.Customer.ListCustomer:output_type -> api.helloworld.v1.ListCustomersReply
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_customer_proto_init() }
func file_api_helloworld_v1_customer_proto_init() {
	if File_api_helloworld_v1_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_helloworld_v1_customer_proto_rawDesc), len(file_api_helloworld_v1_customer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_helloworld_v1_customer_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_customer_proto_depIdxs,
		MessageInfos:      file_api_helloworld_v1_customer_proto_msgTypes,
	}.Build()
	File_api_helloworld_v1_customer_proto = out.File
	file_api_helloworld_v1_customer_proto_goTypes = nil
	file_api_helloworld_v1_customer_proto_depIdxs = nil
}
