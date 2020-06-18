// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/meta.proto

package service

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// RequestMetaHeader contains information about request meta headers
// (should be embedded into message)
type RequestMetaHeader struct {
	// TTL must be larger than zero, it decreased in every NeoFS Node
	TTL uint32 `protobuf:"varint,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
	// Epoch for user can be empty, because node sets epoch to the actual value
	Epoch uint64 `protobuf:"varint,2,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	// Version defines protocol version
	// TODO: not used for now, should be implemented in future
	Version uint32 `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	// Raw determines whether the request is raw or not
	Raw bool `protobuf:"varint,4,opt,name=Raw,proto3" json:"Raw,omitempty"`
	// ExtendedHeader carries extended headers of the request
	RequestExtendedHeader `protobuf:"bytes,5,opt,name=ExtendedHeader,proto3,embedded=ExtendedHeader" json:"ExtendedHeader"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *RequestMetaHeader) Reset()         { *m = RequestMetaHeader{} }
func (m *RequestMetaHeader) String() string { return proto.CompactTextString(m) }
func (*RequestMetaHeader) ProtoMessage()    {}
func (*RequestMetaHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638867e7b43457c, []int{0}
}
func (m *RequestMetaHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestMetaHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RequestMetaHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetaHeader.Merge(m, src)
}
func (m *RequestMetaHeader) XXX_Size() int {
	return m.Size()
}
func (m *RequestMetaHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetaHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetaHeader proto.InternalMessageInfo

func (m *RequestMetaHeader) GetTTL() uint32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *RequestMetaHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *RequestMetaHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RequestMetaHeader) GetRaw() bool {
	if m != nil {
		return m.Raw
	}
	return false
}

// ResponseMetaHeader contains meta information based on request processing by server
// (should be embedded into message)
type ResponseMetaHeader struct {
	// Current NeoFS epoch on server
	Epoch uint64 `protobuf:"varint,1,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	// Version defines protocol version
	// TODO: not used for now, should be implemented in future
	Version              uint32   `protobuf:"varint,2,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMetaHeader) Reset()         { *m = ResponseMetaHeader{} }
func (m *ResponseMetaHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseMetaHeader) ProtoMessage()    {}
func (*ResponseMetaHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638867e7b43457c, []int{1}
}
func (m *ResponseMetaHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseMetaHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ResponseMetaHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMetaHeader.Merge(m, src)
}
func (m *ResponseMetaHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResponseMetaHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMetaHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMetaHeader proto.InternalMessageInfo

func (m *ResponseMetaHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ResponseMetaHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// RequestExtendedHeader contains extended headers of request
type RequestExtendedHeader struct {
	// Headers carries list of key-value headers
	Headers              []RequestExtendedHeader_KV `protobuf:"bytes,1,rep,name=Headers,proto3" json:"Headers"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RequestExtendedHeader) Reset()         { *m = RequestExtendedHeader{} }
func (m *RequestExtendedHeader) String() string { return proto.CompactTextString(m) }
func (*RequestExtendedHeader) ProtoMessage()    {}
func (*RequestExtendedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638867e7b43457c, []int{2}
}
func (m *RequestExtendedHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestExtendedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RequestExtendedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestExtendedHeader.Merge(m, src)
}
func (m *RequestExtendedHeader) XXX_Size() int {
	return m.Size()
}
func (m *RequestExtendedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestExtendedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestExtendedHeader proto.InternalMessageInfo

func (m *RequestExtendedHeader) GetHeaders() []RequestExtendedHeader_KV {
	if m != nil {
		return m.Headers
	}
	return nil
}

// KV contains string key-value pair
type RequestExtendedHeader_KV struct {
	// K carries extended header key
	K string `protobuf:"bytes,1,opt,name=K,proto3" json:"K,omitempty"`
	// V carries extended header value
	V                    string   `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestExtendedHeader_KV) Reset()         { *m = RequestExtendedHeader_KV{} }
func (m *RequestExtendedHeader_KV) String() string { return proto.CompactTextString(m) }
func (*RequestExtendedHeader_KV) ProtoMessage()    {}
func (*RequestExtendedHeader_KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638867e7b43457c, []int{2, 0}
}
func (m *RequestExtendedHeader_KV) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestExtendedHeader_KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RequestExtendedHeader_KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestExtendedHeader_KV.Merge(m, src)
}
func (m *RequestExtendedHeader_KV) XXX_Size() int {
	return m.Size()
}
func (m *RequestExtendedHeader_KV) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestExtendedHeader_KV.DiscardUnknown(m)
}

var xxx_messageInfo_RequestExtendedHeader_KV proto.InternalMessageInfo

func (m *RequestExtendedHeader_KV) GetK() string {
	if m != nil {
		return m.K
	}
	return ""
}

func (m *RequestExtendedHeader_KV) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestMetaHeader)(nil), "service.RequestMetaHeader")
	proto.RegisterType((*ResponseMetaHeader)(nil), "service.ResponseMetaHeader")
	proto.RegisterType((*RequestExtendedHeader)(nil), "service.RequestExtendedHeader")
	proto.RegisterType((*RequestExtendedHeader_KV)(nil), "service.RequestExtendedHeader.KV")
}

func init() { proto.RegisterFile("service/meta.proto", fileDescriptor_a638867e7b43457c) }

var fileDescriptor_a638867e7b43457c = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x8e, 0xda, 0x40,
	0x10, 0x86, 0x19, 0x03, 0x01, 0x96, 0x24, 0x0a, 0xab, 0x44, 0xb2, 0x28, 0x8c, 0x43, 0xe5, 0x14,
	0xb6, 0x25, 0xf2, 0x04, 0xa0, 0x10, 0x25, 0x72, 0x12, 0xa1, 0x05, 0xb9, 0x48, 0x67, 0xec, 0xc1,
	0xb8, 0xc0, 0xeb, 0x78, 0x0d, 0x49, 0x91, 0x07, 0xc9, 0x33, 0xe4, 0x0d, 0xee, 0x0d, 0x28, 0x29,
	0xaf, 0x42, 0x27, 0xdf, 0x8b, 0x9c, 0xbc, 0x36, 0x27, 0x38, 0xdd, 0x5d, 0xf7, 0x7f, 0xb3, 0x33,
	0x3b, 0x9f, 0x34, 0x84, 0x0a, 0x4c, 0x77, 0x91, 0x8f, 0xf6, 0x06, 0x33, 0xcf, 0x4a, 0x52, 0x9e,
	0x71, 0xda, 0xaa, 0x6a, 0x7d, 0x33, 0x8c, 0xb2, 0xf5, 0x76, 0x69, 0xf9, 0x7c, 0x63, 0x87, 0x3c,
	0xe4, 0xb6, 0x7c, 0x5f, 0x6e, 0x57, 0x92, 0x24, 0xc8, 0x54, 0xce, 0x0d, 0xaf, 0x80, 0xf4, 0x18,
	0xfe, 0xda, 0xa2, 0xc8, 0xbe, 0x63, 0xe6, 0x7d, 0x41, 0x2f, 0xc0, 0x94, 0xbe, 0x21, 0xf5, 0xc5,
	0xe2, 0x9b, 0x0a, 0x3a, 0x18, 0xaf, 0x58, 0x11, 0xe9, 0x5b, 0xd2, 0x9c, 0x26, 0xdc, 0x5f, 0xab,
	0x8a, 0x0e, 0x46, 0x83, 0x95, 0x40, 0x55, 0xd2, 0x72, 0x31, 0x15, 0x11, 0x8f, 0xd5, 0xba, 0xec,
	0x3d, 0x61, 0xf1, 0x03, 0xf3, 0x7e, 0xab, 0x0d, 0x1d, 0x8c, 0x36, 0x2b, 0x22, 0x9d, 0x91, 0xd7,
	0xd3, 0x3f, 0x19, 0xc6, 0x01, 0x06, 0xe5, 0x16, 0xb5, 0xa9, 0x83, 0xd1, 0x1d, 0x69, 0x56, 0xa5,
	0x6e, 0x55, 0x1e, 0x97, 0x5d, 0x93, 0xf6, 0xfe, 0x38, 0xa8, 0x1d, 0x8e, 0x03, 0x60, 0x0f, 0xe6,
	0x87, 0x9f, 0x08, 0x65, 0x28, 0x12, 0x1e, 0x0b, 0x3c, 0x73, 0xbf, 0x37, 0x85, 0x27, 0x4c, 0x95,
	0x0b, 0xd3, 0xe1, 0x5f, 0xf2, 0xee, 0xd1, 0xc5, 0x74, 0x4c, 0x5a, 0x65, 0x12, 0x2a, 0xe8, 0x75,
	0xa3, 0x3b, 0x7a, 0xff, 0xbc, 0xa9, 0xe5, 0xb8, 0x93, 0x46, 0x21, 0xcb, 0x4e, 0x73, 0x7d, 0x9d,
	0x28, 0x8e, 0x4b, 0x5f, 0x12, 0x70, 0xa4, 0x4d, 0x87, 0x81, 0x53, 0x90, 0x2b, 0x1d, 0x3a, 0x0c,
	0xdc, 0xc9, 0x7c, 0x9f, 0x6b, 0x70, 0xc8, 0x35, 0xb8, 0xce, 0x35, 0xb8, 0xc9, 0x35, 0xf8, 0x77,
	0xab, 0xd5, 0x7e, 0x7e, 0x38, 0x3b, 0x62, 0x2c, 0x12, 0xdf, 0x37, 0x03, 0xdc, 0xd9, 0x31, 0xf2,
	0x95, 0x30, 0xbd, 0x24, 0x32, 0x43, 0x6e, 0x57, 0x2a, 0xff, 0x95, 0xde, 0x0f, 0xe4, 0x9f, 0xe7,
	0xd6, 0x78, 0xf6, 0xd5, 0x9a, 0x97, 0xb5, 0xe5, 0x0b, 0x79, 0xdb, 0x8f, 0x77, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0xf7, 0x0a, 0xb8, 0x29, 0x02, 0x00, 0x00,
}

func (m *RequestMetaHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestMetaHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestMetaHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.RequestExtendedHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMeta(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Raw {
		i--
		if m.Raw {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if m.Epoch != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x10
	}
	if m.TTL != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseMetaHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseMetaHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseMetaHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Version != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if m.Epoch != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RequestExtendedHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestExtendedHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestExtendedHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMeta(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestExtendedHeader_KV) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestExtendedHeader_KV) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestExtendedHeader_KV) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.V) > 0 {
		i -= len(m.V)
		copy(dAtA[i:], m.V)
		i = encodeVarintMeta(dAtA, i, uint64(len(m.V)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.K) > 0 {
		i -= len(m.K)
		copy(dAtA[i:], m.K)
		i = encodeVarintMeta(dAtA, i, uint64(len(m.K)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestMetaHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TTL != 0 {
		n += 1 + sovMeta(uint64(m.TTL))
	}
	if m.Epoch != 0 {
		n += 1 + sovMeta(uint64(m.Epoch))
	}
	if m.Version != 0 {
		n += 1 + sovMeta(uint64(m.Version))
	}
	if m.Raw {
		n += 2
	}
	l = m.RequestExtendedHeader.Size()
	n += 1 + l + sovMeta(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseMetaHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovMeta(uint64(m.Epoch))
	}
	if m.Version != 0 {
		n += 1 + sovMeta(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestExtendedHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovMeta(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestExtendedHeader_KV) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.K)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	l = len(m.V)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeta(x uint64) (n int) {
	return sovMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestMetaHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: RequestMetaHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestMetaHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Raw = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestExtendedHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RequestExtendedHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseMetaHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: ResponseMetaHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseMetaHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestExtendedHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: RequestExtendedHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestExtendedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, RequestExtendedHeader_KV{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestExtendedHeader_KV) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: KV: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KV: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field K", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.K = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
				return 0, ErrInvalidLengthMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeta = fmt.Errorf("proto: unexpected end of group")
)
