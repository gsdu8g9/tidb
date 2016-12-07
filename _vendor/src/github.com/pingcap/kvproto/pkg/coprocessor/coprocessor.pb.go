// Code generated by protoc-gen-gogo.
// source: coprocessor.proto
// DO NOT EDIT!

/*
	Package coprocessor is a generated protocol buffer package.

	It is generated from these files:
		coprocessor.proto

	It has these top-level messages:
		KeyRange
		Request
		Response
*/
package coprocessor

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	errorpb "github.com/pingcap/kvproto/pkg/errorpb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [start, end)
type KeyRange struct {
	Start            []byte `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End              []byte `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptorCoprocessor, []int{0} }

func (m *KeyRange) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *KeyRange) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

type Request struct {
	Context          *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Tp               int64            `protobuf:"varint,2,opt,name=tp" json:"tp"`
	Data             []byte           `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Ranges           []*KeyRange      `protobuf:"bytes,4,rep,name=ranges" json:"ranges,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorCoprocessor, []int{1} }

func (m *Request) GetContext() *kvrpcpb.Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Request) GetTp() int64 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetRanges() []*KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type Response struct {
	Data             []byte                   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	RegionError      *errorpb.Error           `protobuf:"bytes,2,opt,name=region_error,json=regionError" json:"region_error,omitempty"`
	Locked           *kvrpcpb.LockInfo        `protobuf:"bytes,3,opt,name=locked" json:"locked,omitempty"`
	OtherError       string                   `protobuf:"bytes,4,opt,name=other_error,json=otherError" json:"other_error"`
	UpdatedRegions   []*kvrpcpb.UpdatedRegion `protobuf:"bytes,5,rep,name=updated_regions,json=updatedRegions" json:"updated_regions,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorCoprocessor, []int{2} }

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetRegionError() *errorpb.Error {
	if m != nil {
		return m.RegionError
	}
	return nil
}

func (m *Response) GetLocked() *kvrpcpb.LockInfo {
	if m != nil {
		return m.Locked
	}
	return nil
}

func (m *Response) GetOtherError() string {
	if m != nil {
		return m.OtherError
	}
	return ""
}

func (m *Response) GetUpdatedRegions() []*kvrpcpb.UpdatedRegion {
	if m != nil {
		return m.UpdatedRegions
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "coprocessor.KeyRange")
	proto.RegisterType((*Request)(nil), "coprocessor.Request")
	proto.RegisterType((*Response)(nil), "coprocessor.Response")
}
func (m *KeyRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.Start)))
		i += copy(dAtA[i:], m.Start)
	}
	if m.End != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.End)))
		i += copy(dAtA[i:], m.End)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Context.Size()))
		n1, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	dAtA[i] = 0x10
	i++
	i = encodeVarintCoprocessor(dAtA, i, uint64(m.Tp))
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.Ranges) > 0 {
		for _, msg := range m.Ranges {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCoprocessor(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.RegionError != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.RegionError.Size()))
		n2, err := m.RegionError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Locked != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Locked.Size()))
		n3, err := m.Locked.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.OtherError)))
	i += copy(dAtA[i:], m.OtherError)
	if len(m.UpdatedRegions) > 0 {
		for _, msg := range m.UpdatedRegions {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintCoprocessor(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Coprocessor(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Coprocessor(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCoprocessor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *KeyRange) Size() (n int) {
	var l int
	_ = l
	if m.Start != nil {
		l = len(m.Start)
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.End != nil {
		l = len(m.End)
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	n += 1 + sovCoprocessor(uint64(m.Tp))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if len(m.Ranges) > 0 {
		for _, e := range m.Ranges {
			l = e.Size()
			n += 1 + l + sovCoprocessor(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.RegionError != nil {
		l = m.RegionError.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Locked != nil {
		l = m.Locked.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	l = len(m.OtherError)
	n += 1 + l + sovCoprocessor(uint64(l))
	if len(m.UpdatedRegions) > 0 {
		for _, e := range m.UpdatedRegions {
			l = e.Size()
			n += 1 + l + sovCoprocessor(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCoprocessor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCoprocessor(x uint64) (n int) {
	return sovCoprocessor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Start = append(m.Start[:0], dAtA[iNdEx:postIndex]...)
			if m.Start == nil {
				m.Start = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.End = append(m.End[:0], dAtA[iNdEx:postIndex]...)
			if m.End == nil {
				m.End = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &kvrpcpb.Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ranges = append(m.Ranges, &KeyRange{})
			if err := m.Ranges[len(m.Ranges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionError == nil {
				m.RegionError = &errorpb.Error{}
			}
			if err := m.RegionError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Locked == nil {
				m.Locked = &kvrpcpb.LockInfo{}
			}
			if err := m.Locked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OtherError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedRegions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedRegions = append(m.UpdatedRegions, &kvrpcpb.UpdatedRegion{})
			if err := m.UpdatedRegions[len(m.UpdatedRegions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
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
func skipCoprocessor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoprocessor
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
					return 0, ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoprocessor
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCoprocessor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCoprocessor
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCoprocessor(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCoprocessor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoprocessor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("coprocessor.proto", fileDescriptorCoprocessor) }

var fileDescriptorCoprocessor = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x50, 0xcd, 0x4a, 0xeb, 0x40,
	0x18, 0xed, 0x34, 0xe9, 0xcf, 0xfd, 0xd2, 0xf6, 0xb6, 0x43, 0xef, 0x25, 0x74, 0x91, 0x5b, 0x0a,
	0x17, 0xaa, 0x60, 0xc4, 0xbc, 0x80, 0x50, 0x71, 0x21, 0xba, 0x1a, 0x70, 0x5d, 0x62, 0x32, 0x46,
	0xa9, 0xe4, 0x1b, 0x67, 0xa6, 0xa2, 0x2f, 0xe1, 0xda, 0x47, 0xea, 0xd2, 0x27, 0x10, 0xa9, 0x3e,
	0x88, 0x64, 0x26, 0x09, 0x59, 0xe5, 0xfb, 0xce, 0x39, 0x39, 0xdf, 0x39, 0x03, 0x93, 0x04, 0x85,
	0xc4, 0x84, 0x2b, 0x85, 0x32, 0x14, 0x12, 0x35, 0x52, 0xaf, 0x01, 0xcd, 0x86, 0x5c, 0x4a, 0x94,
	0xe2, 0xc6, 0x72, 0xb3, 0xe1, 0xe6, 0x49, 0x8a, 0xa4, 0x5e, 0xa7, 0x19, 0x66, 0x68, 0xc6, 0xe3,
	0x62, 0xb2, 0xe8, 0x22, 0x82, 0xfe, 0x25, 0x7f, 0x61, 0x71, 0x9e, 0x71, 0x3a, 0x85, 0x8e, 0xd2,
	0xb1, 0xd4, 0x3e, 0x99, 0x93, 0xe5, 0x80, 0xd9, 0x85, 0x8e, 0xc1, 0xe1, 0x79, 0xea, 0xb7, 0x0d,
	0x56, 0x8c, 0x8b, 0x57, 0x02, 0x3d, 0xc6, 0x1f, 0xb7, 0x5c, 0x69, 0x7a, 0x08, 0xbd, 0x04, 0x73,
	0xcd, 0x9f, 0xed, 0x5f, 0x5e, 0x34, 0x0e, 0xab, 0xb3, 0x67, 0x16, 0x67, 0x95, 0x80, 0x4e, 0xa1,
	0xad, 0x85, 0x31, 0x72, 0x56, 0xee, 0xee, 0xe3, 0x5f, 0x8b, 0xb5, 0xb5, 0xa0, 0x14, 0xdc, 0x34,
	0xd6, 0xb1, 0xef, 0x98, 0x03, 0x66, 0xa6, 0x47, 0xd0, 0x95, 0x45, 0x24, 0xe5, 0xbb, 0x73, 0x67,
	0xe9, 0x45, 0x7f, 0xc2, 0x66, 0xf5, 0x2a, 0x30, 0x2b, 0x45, 0x8b, 0x6f, 0x02, 0x7d, 0xc6, 0x95,
	0xc0, 0x5c, 0xf1, 0xda, 0x8f, 0x34, 0xfc, 0x4e, 0x60, 0x20, 0x79, 0x76, 0x8f, 0xf9, 0xda, 0x3c,
	0x91, 0xc9, 0xe0, 0x45, 0xa3, 0xb0, 0x7a, 0xb0, 0xf3, 0xe2, 0xcb, 0x3c, 0xab, 0x31, 0x0b, 0x3d,
	0x80, 0xee, 0x03, 0x26, 0x1b, 0x9e, 0x9a, 0x60, 0x5e, 0x34, 0xa9, 0x7b, 0x5d, 0x61, 0xb2, 0xb9,
	0xc8, 0x6f, 0x91, 0x95, 0x02, 0xfa, 0x1f, 0x3c, 0xd4, 0x77, 0x5c, 0x96, 0xe6, 0xee, 0x9c, 0x2c,
	0x7f, 0x95, 0x05, 0xc1, 0x10, 0xd6, 0xf1, 0x14, 0x7e, 0x6f, 0x45, 0x1a, 0x6b, 0x9e, 0xae, 0xed,
	0x21, 0xe5, 0x77, 0x4c, 0xbb, 0xbf, 0xb5, 0xf5, 0xb5, 0xe5, 0x99, 0xa1, 0xd9, 0x68, 0xdb, 0x5c,
	0xd5, 0x6a, 0xbc, 0xdb, 0x07, 0xe4, 0x7d, 0x1f, 0x90, 0xcf, 0x7d, 0x40, 0xde, 0xbe, 0x82, 0xd6,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xfe, 0xaf, 0xd5, 0x12, 0x02, 0x00, 0x00,
}
