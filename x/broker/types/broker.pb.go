// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: broker/broker.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Source struct {
	PoolId       uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	InterqueryId string `protobuf:"bytes,2,opt,name=interquery_id,json=interqueryId,proto3" json:"interquery_id,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{0}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Source.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return m.Size()
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *Source) GetInterqueryId() string {
	if m != nil {
		return m.InterqueryId
	}
	return ""
}

func (m *Source) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Broker struct {
	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId string    `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	Pools        []*Source `protobuf:"bytes,3,rep,name=pools,proto3" json:"pools,omitempty"`
	Status       string    `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Broker) Reset()         { *m = Broker{} }
func (m *Broker) String() string { return proto.CompactTextString(m) }
func (*Broker) ProtoMessage()    {}
func (*Broker) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{1}
}
func (m *Broker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Broker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Broker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Broker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Broker.Merge(m, src)
}
func (m *Broker) XXX_Size() int {
	return m.Size()
}
func (m *Broker) XXX_DiscardUnknown() {
	xxx_messageInfo_Broker.DiscardUnknown(m)
}

var xxx_messageInfo_Broker proto.InternalMessageInfo

func (m *Broker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Broker) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Broker) GetPools() []*Source {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *Broker) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Transfer struct {
	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Channel  string      `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Sequence uint64      `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status   string      `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Token    *types.Coin `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Sender   string      `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string      `protobuf:"bytes,7,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{2}
}
func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return m.Size()
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transfer) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Transfer) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Transfer) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Transfer) GetToken() *types.Coin {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Transfer) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Transfer) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "defundlabs.defund.broker.Source")
	proto.RegisterType((*Broker)(nil), "defundlabs.defund.broker.Broker")
	proto.RegisterType((*Transfer)(nil), "defundlabs.defund.broker.Transfer")
}

func init() { proto.RegisterFile("broker/broker.proto", fileDescriptor_09a300fef54c4624) }

var fileDescriptor_09a300fef54c4624 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x8d, 0xd3, 0x64, 0xd3, 0xba, 0x2d, 0x07, 0x53, 0xc1, 0x92, 0xc3, 0x12, 0x85, 0x4b, 0x0e,
	0x60, 0xab, 0x45, 0xe2, 0x80, 0xc4, 0x25, 0x48, 0x48, 0xb9, 0x2e, 0x9c, 0x38, 0x80, 0x76, 0xbd,
	0xd3, 0xd4, 0xea, 0xc6, 0x93, 0xda, 0xde, 0x8a, 0xfc, 0x05, 0x7f, 0xc1, 0xaf, 0x70, 0x41, 0xea,
	0x91, 0x13, 0x42, 0xc9, 0x1f, 0xf0, 0x05, 0xc8, 0xf6, 0x92, 0x52, 0xa9, 0x39, 0xed, 0xbc, 0x7d,
	0x4f, 0x6f, 0xde, 0xcc, 0x98, 0x3e, 0x2c, 0x0d, 0x5e, 0x82, 0x11, 0xf1, 0xc3, 0x97, 0x06, 0x1d,
	0xb2, 0xb4, 0x82, 0xf3, 0x46, 0x57, 0x75, 0x51, 0x5a, 0x1e, 0x4b, 0x1e, 0xf9, 0xe1, 0xc9, 0x1c,
	0xe7, 0x18, 0x44, 0xc2, 0x57, 0x51, 0x3f, 0xcc, 0x24, 0xda, 0x05, 0x5a, 0x51, 0x16, 0x16, 0xc4,
	0xf5, 0x69, 0x09, 0xae, 0x38, 0x15, 0x12, 0x95, 0x8e, 0xfc, 0xf8, 0x13, 0x4d, 0xde, 0x63, 0x63,
	0x24, 0xb0, 0xc7, 0x74, 0xb0, 0x44, 0xac, 0x3f, 0xab, 0x2a, 0x25, 0x23, 0x32, 0xe9, 0xe5, 0x89,
	0x87, 0xb3, 0x8a, 0x3d, 0xa3, 0xc7, 0x4a, 0x3b, 0x30, 0x57, 0x0d, 0x98, 0x95, 0xa7, 0xbb, 0x23,
	0x32, 0x39, 0xc8, 0x8f, 0x6e, 0x7f, 0xce, 0x2a, 0xf6, 0x88, 0x26, 0xd6, 0x15, 0xae, 0xb1, 0xe9,
	0x5e, 0x60, 0x5b, 0x34, 0xfe, 0x46, 0x68, 0x32, 0x0d, 0x01, 0xd9, 0x03, 0xda, 0x6d, 0xbd, 0x0f,
	0xf2, 0xae, 0xaa, 0xd8, 0x1b, 0x7a, 0x2c, 0x51, 0x6b, 0x90, 0x4e, 0xa1, 0xde, 0xfa, 0x4e, 0xd3,
	0x3f, 0xbf, 0x9e, 0x9e, 0xac, 0x8a, 0x45, 0xfd, 0x7a, 0x7c, 0x87, 0x1e, 0xe7, 0x47, 0xb7, 0x78,
	0x56, 0xb1, 0x57, 0xb4, 0xef, 0x03, 0xfa, 0x86, 0x7b, 0x93, 0xc3, 0xb3, 0x11, 0xdf, 0xb5, 0x19,
	0x1e, 0x07, 0xcc, 0xa3, 0xfc, 0xbf, 0xa4, 0xbd, 0x3b, 0x49, 0x7f, 0x10, 0xba, 0xff, 0xc1, 0x14,
	0xda, 0x9e, 0xdf, 0x93, 0x35, 0xa5, 0x03, 0x79, 0x51, 0x68, 0x0d, 0x75, 0x3b, 0xfd, 0x3f, 0xc8,
	0x86, 0x74, 0xdf, 0xc2, 0x55, 0x03, 0x5a, 0x42, 0x18, 0xbd, 0x97, 0x6f, 0xf1, 0xae, 0x56, 0x4c,
	0xd0, 0xbe, 0xc3, 0x4b, 0xd0, 0x69, 0x7f, 0x44, 0x26, 0x87, 0x67, 0x4f, 0x78, 0x3c, 0x12, 0xf7,
	0x47, 0xe2, 0xed, 0x91, 0xf8, 0x5b, 0x54, 0x3a, 0x8f, 0xba, 0x60, 0x04, 0xba, 0x02, 0x93, 0x26,
	0xad, 0x51, 0x40, 0xbe, 0xb9, 0x01, 0x09, 0xea, 0x1a, 0x4c, 0x3a, 0x08, 0xcc, 0x16, 0x4f, 0xdf,
	0x7d, 0x5f, 0x67, 0xe4, 0x66, 0x9d, 0x91, 0xdf, 0xeb, 0x8c, 0x7c, 0xdd, 0x64, 0x9d, 0x9b, 0x4d,
	0xd6, 0xf9, 0xb9, 0xc9, 0x3a, 0x1f, 0x9f, 0xcf, 0x95, 0xbb, 0x68, 0x4a, 0x2e, 0x71, 0x21, 0xe2,
	0xa6, 0x5e, 0xf8, 0xad, 0xb5, 0xb5, 0xf8, 0xd2, 0xbe, 0x38, 0xe1, 0x56, 0x4b, 0xb0, 0x65, 0x12,
	0x1e, 0xca, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x38, 0xcd, 0xf4, 0xbf, 0x8f, 0x02, 0x00,
	0x00,
}

func (m *Source) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Source) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Source) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InterqueryId) > 0 {
		i -= len(m.InterqueryId)
		copy(dAtA[i:], m.InterqueryId)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.InterqueryId)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintBroker(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Broker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Broker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Broker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBroker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Transfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x32
	}
	if m.Token != nil {
		{
			size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBroker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if m.Sequence != 0 {
		i = encodeVarintBroker(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBroker(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBroker(dAtA []byte, offset int, v uint64) int {
	offset -= sovBroker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Source) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovBroker(uint64(m.PoolId))
	}
	l = len(m.InterqueryId)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func (m *Broker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovBroker(uint64(l))
		}
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func (m *Transfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovBroker(uint64(m.Sequence))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	if m.Token != nil {
		l = m.Token.Size()
		n += 1 + l + sovBroker(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovBroker(uint64(l))
	}
	return n
}

func sovBroker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBroker(x uint64) (n int) {
	return sovBroker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Source) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Source: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Source: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterqueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterqueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBroker
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Broker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Broker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Broker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, &Source{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBroker
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Transfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Transfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Token == nil {
				m.Token = &types.Coin{}
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBroker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBroker
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBroker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBroker
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBroker
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBroker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBroker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBroker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBroker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBroker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBroker = fmt.Errorf("proto: unexpected end of group")
)
