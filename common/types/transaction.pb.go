// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bottos-project/bottos/common/types/transaction.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	github.com/bottos-project/bottos/common/types/transaction.proto
	github.com/bottos-project/bottos/common/types/basic-transaction.proto
	github.com/bottos-project/bottos/common/types/block.proto

It has these top-level messages:
	Transaction
	DerivedTransaction
	HandledTransaction
	BasicTransaction
	Block
	Header
*/
package types

import proto "github.com/golang/protobuf/proto"
import log "github.com/cihub/seelog"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = log.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Transaction struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       []byte `protobuf:"bytes,8,opt,name=param,proto3" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   []byte `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Transaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Transaction) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *Transaction) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *Transaction) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *Transaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Transaction) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *Transaction) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Transaction) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *Transaction) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *Transaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DerivedTransaction struct {
	Transaction *Transaction          `protobuf:"bytes,1,opt,name=transaction" json:"transaction"`
	DerivedTrx  []*DerivedTransaction `protobuf:"bytes,2,rep,name=derived_trx,json=derivedTrx" json:"derived_trx"`
}

func (m *DerivedTransaction) Reset()                    { *m = DerivedTransaction{} }
func (m *DerivedTransaction) String() string            { return proto.CompactTextString(m) }
func (*DerivedTransaction) ProtoMessage()               {}
func (*DerivedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DerivedTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *DerivedTransaction) GetDerivedTrx() []*DerivedTransaction {
	if m != nil {
		return m.DerivedTrx
	}
	return nil
}

type HandledTransaction struct {
	Transaction *Transaction          `protobuf:"bytes,1,opt,name=transaction" json:"transaction"`
	DerivedTrx  []*DerivedTransaction `protobuf:"bytes,2,rep,name=derived_trx,json=derivedTrx" json:"derived_trx"`
}

func (m *HandledTransaction) Reset()                    { *m = HandledTransaction{} }
func (m *HandledTransaction) String() string            { return proto.CompactTextString(m) }
func (*HandledTransaction) ProtoMessage()               {}
func (*HandledTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HandledTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *HandledTransaction) GetDerivedTrx() []*DerivedTransaction {
	if m != nil {
		return m.DerivedTrx
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "types.Transaction")
	proto.RegisterType((*DerivedTransaction)(nil), "types.DerivedTransaction")
	proto.RegisterType((*HandledTransaction)(nil), "types.HandledTransaction")
}

func init() {
	proto.RegisterFile("github.com/bottos-project/bottos/common/types/transaction.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x3f, 0x4f, 0xeb, 0x30,
	0x14, 0xc5, 0x95, 0xfe, 0x49, 0x9b, 0x9b, 0xbe, 0xc5, 0x7a, 0x7a, 0xcf, 0x20, 0x90, 0x42, 0xa7,
	0x2c, 0x4d, 0xa4, 0xc2, 0xc4, 0x82, 0x40, 0x0c, 0x0c, 0x88, 0x21, 0x62, 0x62, 0xa9, 0x1c, 0xc7,
	0xa4, 0x46, 0xb1, 0x1d, 0xd9, 0x4e, 0x55, 0xbe, 0x00, 0x1f, 0x83, 0xcf, 0x8a, 0xea, 0xa4, 0x6d,
	0x24, 0x26, 0x26, 0xc6, 0xdf, 0x39, 0xf7, 0xd8, 0x57, 0x47, 0x17, 0x6e, 0x4a, 0x6e, 0xd7, 0x4d,
	0x9e, 0x50, 0x25, 0xd2, 0x5c, 0x59, 0xab, 0xcc, 0xa2, 0xd6, 0xea, 0x8d, 0x51, 0xdb, 0x61, 0x4a,
	0x95, 0x10, 0x4a, 0xa6, 0xf6, 0xbd, 0x66, 0x26, 0xb5, 0x9a, 0x48, 0x43, 0xa8, 0xe5, 0x4a, 0x26,
	0xb5, 0x56, 0x56, 0xa1, 0xb1, 0x33, 0xe6, 0x9f, 0x03, 0x08, 0x9f, 0x8f, 0x26, 0xc2, 0x30, 0xd9,
	0x30, 0x6d, 0xb8, 0x92, 0xd8, 0x8b, 0xbc, 0xf8, 0x4f, 0xb6, 0x47, 0x74, 0x0e, 0x40, 0x1b, 0x6d,
	0x94, 0x5e, 0xc9, 0x46, 0xe0, 0x81, 0x33, 0x83, 0x56, 0x79, 0x6a, 0x04, 0xba, 0x80, 0x59, 0x67,
	0x57, 0x24, 0x67, 0x15, 0x1e, 0xba, 0x81, 0xb0, 0xd5, 0x1e, 0x77, 0x12, 0x3a, 0x85, 0x69, 0xc5,
	0x5f, 0x99, 0xe5, 0x82, 0xe1, 0x51, 0xe4, 0xc5, 0xa3, 0xec, 0xc0, 0xe8, 0x1f, 0xf8, 0x86, 0xc9,
	0x82, 0x69, 0x3c, 0x8e, 0xbc, 0x38, 0xc8, 0x3a, 0xda, 0x65, 0xa8, 0x92, 0x56, 0x13, 0x6a, 0xb1,
	0xef, 0x9c, 0x03, 0xef, 0x32, 0x82, 0xd9, 0xb5, 0x2a, 0xf0, 0xa4, 0xcd, 0xb4, 0x84, 0xfe, 0xc2,
	0xb8, 0x26, 0x9a, 0x08, 0x3c, 0x8d, 0xbc, 0x78, 0x96, 0xb5, 0x80, 0xfe, 0xc3, 0xc4, 0xf0, 0x72,
	0x45, 0xaa, 0x12, 0x07, 0x6e, 0x37, 0xdf, 0xf0, 0xf2, 0xb6, 0x2a, 0xd1, 0x19, 0x04, 0x86, 0x97,
	0x92, 0xd8, 0x46, 0x33, 0x0c, 0x2e, 0x72, 0x14, 0xe6, 0x1f, 0x1e, 0xa0, 0x7b, 0xa6, 0xf9, 0x86,
	0x15, 0xfd, 0x9e, 0xae, 0x20, 0xec, 0x75, 0xea, 0xba, 0x0a, 0x97, 0x28, 0x71, 0xa5, 0x26, 0xbd,
	0xc1, 0xac, 0x3f, 0x86, 0xae, 0x21, 0x2c, 0xda, 0xb7, 0x56, 0x56, 0x6f, 0xf1, 0x20, 0x1a, 0xc6,
	0xe1, 0xf2, 0xa4, 0x4b, 0x7d, 0xff, 0x25, 0x83, 0x62, 0xaf, 0x6d, 0xdd, 0x22, 0x0f, 0x44, 0x16,
	0xd5, 0x2f, 0x2f, 0x72, 0x97, 0xbe, 0x2c, 0x7e, 0x74, 0x7c, 0xb9, 0xef, 0x2e, 0xee, 0xf2, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x19, 0x0e, 0x5f, 0xd8, 0xb4, 0x02, 0x00, 0x00,
}
