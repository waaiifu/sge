// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/house/ticket.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sge-network/sge/types"
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

// DepositTicketPayload indicates data of the deposit ticket.
type DepositTicketPayload struct {
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,1,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
}

func (m *DepositTicketPayload) Reset()         { *m = DepositTicketPayload{} }
func (m *DepositTicketPayload) String() string { return proto.CompactTextString(m) }
func (*DepositTicketPayload) ProtoMessage()    {}
func (*DepositTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f686c28436675f2, []int{0}
}
func (m *DepositTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositTicketPayload.Merge(m, src)
}
func (m *DepositTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *DepositTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_DepositTicketPayload proto.InternalMessageInfo

func (m *DepositTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

// WithdrawTicketPayload indicates data of the withdrawal ticket.
type WithdrawTicketPayload struct {
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,1,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
}

func (m *WithdrawTicketPayload) Reset()         { *m = WithdrawTicketPayload{} }
func (m *WithdrawTicketPayload) String() string { return proto.CompactTextString(m) }
func (*WithdrawTicketPayload) ProtoMessage()    {}
func (*WithdrawTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f686c28436675f2, []int{1}
}
func (m *WithdrawTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WithdrawTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WithdrawTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WithdrawTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawTicketPayload.Merge(m, src)
}
func (m *WithdrawTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *WithdrawTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawTicketPayload proto.InternalMessageInfo

func (m *WithdrawTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

func init() {
	proto.RegisterType((*DepositTicketPayload)(nil), "sgenetwork.sge.house.DepositTicketPayload")
	proto.RegisterType((*WithdrawTicketPayload)(nil), "sgenetwork.sge.house.WithdrawTicketPayload")
}

func init() { proto.RegisterFile("sge/house/ticket.proto", fileDescriptor_1f686c28436675f2) }

var fileDescriptor_1f686c28436675f2 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x4f, 0xd5,
	0xcf, 0xc8, 0x2f, 0x2d, 0x4e, 0xd5, 0x2f, 0xc9, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x29, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b,
	0x4e, 0x4f, 0xd5, 0x03, 0x2b, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd0, 0x07, 0xb1,
	0x20, 0x6a, 0xa5, 0x84, 0x40, 0x66, 0x94, 0x54, 0x16, 0xa4, 0xea, 0x67, 0x57, 0x26, 0x43, 0xc4,
	0x94, 0x62, 0xb8, 0x44, 0x5c, 0x52, 0x0b, 0xf2, 0x8b, 0x33, 0x4b, 0x42, 0xc0, 0xc6, 0x06, 0x24,
	0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x08, 0xb9, 0x70, 0x71, 0x64, 0x57, 0x26, 0xc7, 0xa7, 0x24, 0x96,
	0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0xeb, 0xa1, 0x59, 0x05, 0x32, 0x49, 0xcf,
	0xbb, 0x32, 0xd9, 0x25, 0xb1, 0x24, 0x11, 0xaa, 0xcd, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20,
	0xf6, 0x6c, 0x88, 0xa8, 0x52, 0x2c, 0x97, 0x68, 0x78, 0x66, 0x49, 0x46, 0x4a, 0x51, 0x62, 0x39,
	0x0d, 0x8c, 0x77, 0x72, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x8d,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0xf4, 0x54, 0x5d, 0xa8,
	0xc1, 0x20, 0xb6, 0x7e, 0x05, 0x2c, 0x1c, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xe1, 0x60,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x24, 0xd0, 0xa4, 0x7a, 0x61, 0x01, 0x00, 0x00,
}

func (m *DepositTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *WithdrawTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WithdrawTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WithdrawTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.KycData.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func (m *WithdrawTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.KycData.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: DepositTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *WithdrawTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: WithdrawTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WithdrawTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
