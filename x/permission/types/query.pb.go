// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/permission/v1beta1/query.proto

package types

import (
	fmt "fmt"
	multisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// QueryGovernanceKeyRequest is the request type for the
// Query/GovernanceKey RPC method
type QueryGovernanceKeyRequest struct {
}

func (m *QueryGovernanceKeyRequest) Reset()         { *m = QueryGovernanceKeyRequest{} }
func (m *QueryGovernanceKeyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGovernanceKeyRequest) ProtoMessage()    {}
func (*QueryGovernanceKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8183035d07a1606, []int{0}
}
func (m *QueryGovernanceKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGovernanceKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGovernanceKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGovernanceKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGovernanceKeyRequest.Merge(m, src)
}
func (m *QueryGovernanceKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGovernanceKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGovernanceKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGovernanceKeyRequest proto.InternalMessageInfo

// QueryGovernanceKeyResponse is the response type for the
// Query/GovernanceKey RPC method
type QueryGovernanceKeyResponse struct {
	GovernanceKey multisig.LegacyAminoPubKey `protobuf:"bytes,1,opt,name=governance_key,json=governanceKey,proto3" json:"governance_key"`
}

func (m *QueryGovernanceKeyResponse) Reset()         { *m = QueryGovernanceKeyResponse{} }
func (m *QueryGovernanceKeyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGovernanceKeyResponse) ProtoMessage()    {}
func (*QueryGovernanceKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8183035d07a1606, []int{1}
}
func (m *QueryGovernanceKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGovernanceKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGovernanceKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGovernanceKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGovernanceKeyResponse.Merge(m, src)
}
func (m *QueryGovernanceKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGovernanceKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGovernanceKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGovernanceKeyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryGovernanceKeyRequest)(nil), "axelar.permission.v1beta1.QueryGovernanceKeyRequest")
	proto.RegisterType((*QueryGovernanceKeyResponse)(nil), "axelar.permission.v1beta1.QueryGovernanceKeyResponse")
}

func init() {
	proto.RegisterFile("axelar/permission/v1beta1/query.proto", fileDescriptor_e8183035d07a1606)
}

var fileDescriptor_e8183035d07a1606 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0x33, 0x31,
	0x10, 0xc4, 0xef, 0xa4, 0x4f, 0x5f, 0x71, 0x08, 0x8a, 0x88, 0x82, 0x04, 0xc9, 0x81, 0x48, 0x48,
	0x50, 0x60, 0x2b, 0x50, 0x50, 0x93, 0x86, 0x22, 0x14, 0x10, 0x09, 0x0a, 0x1a, 0x74, 0x67, 0x2d,
	0xc6, 0x4a, 0xec, 0x75, 0xfc, 0x27, 0xc4, 0x6f, 0xc1, 0x63, 0xa5, 0x4c, 0x49, 0x85, 0x20, 0x79,
	0x11, 0x94, 0xf8, 0x80, 0x14, 0x74, 0x1e, 0xcd, 0xcf, 0x3b, 0xbb, 0x53, 0x1c, 0x95, 0x53, 0x18,
	0x95, 0x96, 0x19, 0xb0, 0x4a, 0x3a, 0x27, 0x51, 0xb3, 0x49, 0xb7, 0x02, 0x5f, 0x76, 0xd9, 0x38,
	0x80, 0x8d, 0xd4, 0x58, 0xf4, 0xd8, 0x68, 0x26, 0x8c, 0xfe, 0x62, 0xb4, 0xc6, 0x5a, 0xbb, 0x02,
	0x05, 0xae, 0x29, 0xb6, 0x7a, 0xa5, 0x0f, 0xad, 0xb6, 0x40, 0x14, 0x23, 0x60, 0x6b, 0x55, 0x85,
	0x27, 0xe6, 0xa5, 0x02, 0xe7, 0x4b, 0x65, 0x6a, 0xe0, 0x90, 0xa3, 0x53, 0xe8, 0x18, 0xb7, 0xd1,
	0x78, 0x64, 0x2a, 0x8c, 0xbc, 0x74, 0x52, 0xb0, 0x21, 0x44, 0x97, 0x90, 0xce, 0x7e, 0xd1, 0xbc,
	0x5d, 0xed, 0x70, 0x85, 0x13, 0xb0, 0xba, 0xd4, 0x1c, 0xfa, 0x10, 0x07, 0x30, 0x0e, 0xe0, 0x7c,
	0xc7, 0x17, 0xad, 0xbf, 0x4c, 0x67, 0x50, 0x3b, 0x68, 0xdc, 0x17, 0x3b, 0xe2, 0xc7, 0x78, 0x1c,
	0x42, 0xdc, 0xcb, 0x0f, 0xf2, 0xe3, 0xad, 0xb3, 0x13, 0x9a, 0x62, 0x69, 0x8a, 0xa5, 0xdf, 0xb1,
	0xf4, 0x1a, 0x44, 0xc9, 0xe3, 0xa5, 0x92, 0x1a, 0x6f, 0x42, 0xd5, 0x87, 0xd8, 0xfb, 0x37, 0x7b,
	0x6f, 0x67, 0x83, 0x6d, 0xb1, 0x39, 0xbf, 0x77, 0x37, 0xfb, 0x24, 0xd9, 0x6c, 0x41, 0xf2, 0xf9,
	0x82, 0xe4, 0x1f, 0x0b, 0x92, 0xbf, 0x2e, 0x49, 0x36, 0x5f, 0x92, 0xec, 0x6d, 0x49, 0xb2, 0x87,
	0x0b, 0x21, 0xfd, 0x73, 0xa8, 0x28, 0x47, 0xc5, 0x52, 0x61, 0x1a, 0xfc, 0x0b, 0xda, 0x61, 0xad,
	0x4e, 0x39, 0x5a, 0x60, 0xd3, 0xcd, 0xb2, 0x7d, 0x34, 0xe0, 0xaa, 0xff, 0xeb, 0x83, 0xcf, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x79, 0x06, 0xe3, 0xae, 0x8e, 0x01, 0x00, 0x00,
}

func (m *QueryGovernanceKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGovernanceKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGovernanceKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGovernanceKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGovernanceKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGovernanceKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GovernanceKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGovernanceKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGovernanceKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GovernanceKey.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGovernanceKeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGovernanceKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGovernanceKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGovernanceKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGovernanceKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGovernanceKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovernanceKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GovernanceKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
