// Code generated by protoc-gen-go.
// source: addressbook.proto
// DO NOT EDIT!

/*
Package protobuf_intro is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	PhoneNumber
	AddressBook
*/
package protobuf_intro

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

var PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}
func (PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Person struct {
	Name   string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32          `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string         `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type PhoneNumber struct {
	Number    string    `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	PhoneType PhoneType `protobuf:"varint,2,opt,name=phoneType,enum=protobuf.intro.PhoneType" json:"phoneType,omitempty"`
}

func (m *PhoneNumber) Reset()                    { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()               {}
func (*PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *PhoneNumber) GetPhoneType() PhoneType {
	if m != nil {
		return m.PhoneType
	}
	return PhoneType_MOBILE
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "protobuf.intro.Person")
	proto.RegisterType((*PhoneNumber)(nil), "protobuf.intro.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "protobuf.intro.AddressBook")
	proto.RegisterEnum("protobuf.intro.PhoneType", PhoneType_name, PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x69, 0xba, 0x98, 0x09, 0x84, 0x38, 0x48, 0x89, 0x78, 0x29, 0x39, 0x15, 0x85, 0x3d,
	0xb4, 0x07, 0x4f, 0x1e, 0x0c, 0x14, 0x14, 0xad, 0x29, 0x8b, 0xe0, 0x4d, 0x48, 0xc8, 0x8a, 0xc1,
	0x66, 0x37, 0x6c, 0x5a, 0xc4, 0x7f, 0x6f, 0x76, 0xba, 0xf8, 0x7d, 0xca, 0x7b, 0x93, 0xb7, 0x6f,
	0xde, 0x1b, 0x38, 0x2a, 0xeb, 0xda, 0xc8, 0xbe, 0xaf, 0xb4, 0x7e, 0xe5, 0x9d, 0xd1, 0x5b, 0x8d,
	0x31, 0x7d, 0xaa, 0xdd, 0x33, 0x6f, 0xd4, 0xd6, 0xe8, 0xec, 0x0d, 0xd8, 0x5a, 0x9a, 0x5e, 0x2b,
	0x44, 0x08, 0x54, 0xd9, 0xca, 0xd4, 0x9b, 0x7a, 0xb3, 0x50, 0x10, 0xc6, 0x18, 0xfc, 0xa6, 0x4e,
	0xfd, 0x61, 0x32, 0x16, 0x03, 0xc2, 0x63, 0x18, 0xcb, 0xb6, 0x6c, 0x36, 0xe9, 0x88, 0x44, 0x7b,
	0x82, 0x0b, 0x60, 0xdd, 0x8b, 0x56, 0xb2, 0x4f, 0x83, 0xe9, 0x68, 0x16, 0xcd, 0x4f, 0xf9, 0xcf,
	0x25, 0x7c, 0x6d, 0xff, 0xde, 0xef, 0xda, 0x4a, 0x1a, 0xe1, 0xa4, 0xd9, 0x13, 0x44, 0xdf, 0xc6,
	0x38, 0x01, 0xa6, 0x08, 0xb9, 0xfd, 0x8e, 0xe1, 0x05, 0x84, 0xf4, 0xe0, 0xe1, 0xbd, 0x93, 0x14,
	0x24, 0x9e, 0x9f, 0xfc, 0x6b, 0x6f, 0x05, 0xe2, 0x4b, 0x9b, 0x5d, 0x42, 0x74, 0xb5, 0x6f, 0x9f,
	0x0f, 0xed, 0x91, 0x0f, 0x19, 0xa5, 0xee, 0x36, 0xb6, 0x9f, 0xcd, 0x38, 0xf9, 0x63, 0x42, 0x57,
	0x10, 0x4e, 0x75, 0x76, 0x0e, 0xe1, 0xa7, 0x2d, 0x02, 0xb0, 0x55, 0x91, 0xdf, 0xdc, 0x2d, 0x93,
	0x03, 0x3c, 0x84, 0xe0, 0xba, 0x58, 0x2d, 0x13, 0xcf, 0xa2, 0xc7, 0x42, 0xdc, 0x26, 0x7e, 0x9e,
	0xc0, 0xaf, 0xb3, 0x56, 0x8c, 0xf8, 0xe2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x94, 0xfc, 0x5d, 0x7e,
	0x82, 0x01, 0x00, 0x00,
}