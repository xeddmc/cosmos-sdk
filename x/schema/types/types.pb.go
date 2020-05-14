// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/schema/types/types.proto

package types

import (
	fmt "fmt"
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

type KeyDescriptor struct {
	Name        string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Store       string                `protobuf:"bytes,4,opt,name=store,proto3" json:"store,omitempty"`
	Prefix      []byte                `protobuf:"bytes,5,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Parts       []*KeyDescriptor_Part `protobuf:"bytes,6,rep,name=parts,proto3" json:"parts,omitempty"`
	ValueType   string                `protobuf:"bytes,7,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
}

func (m *KeyDescriptor) Reset()         { *m = KeyDescriptor{} }
func (m *KeyDescriptor) String() string { return proto.CompactTextString(m) }
func (*KeyDescriptor) ProtoMessage()    {}
func (*KeyDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_335ca074c797b8b4, []int{0}
}
func (m *KeyDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDescriptor.Merge(m, src)
}
func (m *KeyDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *KeyDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDescriptor proto.InternalMessageInfo

func (m *KeyDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *KeyDescriptor) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

func (m *KeyDescriptor) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *KeyDescriptor) GetParts() []*KeyDescriptor_Part {
	if m != nil {
		return m.Parts
	}
	return nil
}

func (m *KeyDescriptor) GetValueType() string {
	if m != nil {
		return m.ValueType
	}
	return ""
}

type KeyDescriptor_Part struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *KeyDescriptor_Part) Reset()         { *m = KeyDescriptor_Part{} }
func (m *KeyDescriptor_Part) String() string { return proto.CompactTextString(m) }
func (*KeyDescriptor_Part) ProtoMessage()    {}
func (*KeyDescriptor_Part) Descriptor() ([]byte, []int) {
	return fileDescriptor_335ca074c797b8b4, []int{0, 0}
}
func (m *KeyDescriptor_Part) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyDescriptor_Part) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyDescriptor_Part.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyDescriptor_Part) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDescriptor_Part.Merge(m, src)
}
func (m *KeyDescriptor_Part) XXX_Size() int {
	return m.Size()
}
func (m *KeyDescriptor_Part) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDescriptor_Part.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDescriptor_Part proto.InternalMessageInfo

func (m *KeyDescriptor_Part) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyDescriptor_Part) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *KeyDescriptor_Part) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyDescriptor)(nil), "cosmos_sdk.x.schema.v1.KeyDescriptor")
	proto.RegisterType((*KeyDescriptor_Part)(nil), "cosmos_sdk.x.schema.v1.KeyDescriptor.Part")
}

func init() { proto.RegisterFile("x/schema/types/types.proto", fileDescriptor_335ca074c797b8b4) }

var fileDescriptor_335ca074c797b8b4 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xaa, 0xd0, 0x2f, 0x4e,
	0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x90, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x62, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0xc5, 0x29, 0xd9, 0x7a, 0x15,
	0x7a, 0x10, 0x65, 0x7a, 0x65, 0x86, 0x4a, 0xcb, 0x98, 0xb8, 0x78, 0xbd, 0x53, 0x2b, 0x5d, 0x52,
	0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x05, 0x2e, 0xee, 0x14, 0xa8, 0x8a,
	0xcc, 0xfc, 0x3c, 0x09, 0x26, 0xb0, 0x14, 0xb2, 0x90, 0x90, 0x08, 0x17, 0x6b, 0x71, 0x49, 0x7e,
	0x51, 0xaa, 0x04, 0x0b, 0x58, 0x0e, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x28, 0x4a, 0x4d, 0xcb,
	0xac, 0x90, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0xf2, 0x84, 0x1c, 0xb8, 0x58, 0x0b, 0x12,
	0x8b, 0x4a, 0x8a, 0x25, 0xd8, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xb4, 0xf4, 0xb0, 0xbb, 0x4e, 0x0f,
	0xc5, 0x65, 0x7a, 0x01, 0x89, 0x45, 0x25, 0x41, 0x10, 0x8d, 0x42, 0xb2, 0x5c, 0x5c, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0xf1, 0x20, 0x4f, 0x4a, 0xb0, 0x83, 0x2d, 0xe5, 0x04, 0x8b, 0x84, 0x54, 0x16,
	0xa4, 0x4a, 0x05, 0x70, 0xb1, 0x80, 0x54, 0x93, 0xe9, 0x19, 0x21, 0x2e, 0x16, 0xb0, 0xb1, 0xcc,
	0x10, 0x5d, 0x20, 0xb6, 0x93, 0xeb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x69, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0xfc, 0x01, 0xa5,
	0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x51, 0xe3, 0x24, 0x89, 0x0d, 0x1c, 0x1d, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8a, 0x02, 0x93, 0x51, 0xac, 0x01, 0x00, 0x00,
}

func (m *KeyDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValueType) > 0 {
		i -= len(m.ValueType)
		copy(dAtA[i:], m.ValueType)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValueType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Parts) > 0 {
		for iNdEx := len(m.Parts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Parts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Store) > 0 {
		i -= len(m.Store)
		copy(dAtA[i:], m.Store)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Store)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeyDescriptor_Part) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyDescriptor_Part) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyDescriptor_Part) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Store)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Parts) > 0 {
		for _, e := range m.Parts {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.ValueType)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *KeyDescriptor_Part) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: KeyDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Store = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append(m.Prefix[:0], dAtA[iNdEx:postIndex]...)
			if m.Prefix == nil {
				m.Prefix = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parts = append(m.Parts, &KeyDescriptor_Part{})
			if err := m.Parts[len(m.Parts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *KeyDescriptor_Part) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Part: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Part: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)