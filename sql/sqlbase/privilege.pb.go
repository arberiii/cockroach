// Code generated by protoc-gen-gogo.
// source: cockroach/sql/sqlbase/privilege.proto
// DO NOT EDIT!

package sqlbase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// UserPrivileges describes the list of privileges available for a given user.
type UserPrivileges struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user"`
	// privileges is a bitfield of 1<<Privilege values.
	Privileges uint32 `protobuf:"varint,2,opt,name=privileges" json:"privileges"`
}

func (m *UserPrivileges) Reset()                    { *m = UserPrivileges{} }
func (m *UserPrivileges) String() string            { return proto.CompactTextString(m) }
func (*UserPrivileges) ProtoMessage()               {}
func (*UserPrivileges) Descriptor() ([]byte, []int) { return fileDescriptorPrivilege, []int{0} }

// PrivilegeDescriptor describes a list of users and attached
// privileges. The list should be sorted by user for fast access.
type PrivilegeDescriptor struct {
	Users []UserPrivileges `protobuf:"bytes,1,rep,name=users" json:"users"`
}

func (m *PrivilegeDescriptor) Reset()                    { *m = PrivilegeDescriptor{} }
func (m *PrivilegeDescriptor) String() string            { return proto.CompactTextString(m) }
func (*PrivilegeDescriptor) ProtoMessage()               {}
func (*PrivilegeDescriptor) Descriptor() ([]byte, []int) { return fileDescriptorPrivilege, []int{1} }

func init() {
	proto.RegisterType((*UserPrivileges)(nil), "cockroach.sql.sqlbase.UserPrivileges")
	proto.RegisterType((*PrivilegeDescriptor)(nil), "cockroach.sql.sqlbase.PrivilegeDescriptor")
}
func (m *UserPrivileges) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UserPrivileges) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintPrivilege(data, i, uint64(len(m.User)))
	i += copy(data[i:], m.User)
	data[i] = 0x10
	i++
	i = encodeVarintPrivilege(data, i, uint64(m.Privileges))
	return i, nil
}

func (m *PrivilegeDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PrivilegeDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, msg := range m.Users {
			data[i] = 0xa
			i++
			i = encodeVarintPrivilege(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Privilege(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Privilege(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPrivilege(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *UserPrivileges) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	n += 1 + l + sovPrivilege(uint64(l))
	n += 1 + sovPrivilege(uint64(m.Privileges))
	return n
}

func (m *PrivilegeDescriptor) Size() (n int) {
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovPrivilege(uint64(l))
		}
	}
	return n
}

func sovPrivilege(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPrivilege(x uint64) (n int) {
	return sovPrivilege(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserPrivileges) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivilege
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserPrivileges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserPrivileges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivilege
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrivilege
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			m.Privileges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivilege
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Privileges |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrivilege(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivilege
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
func (m *PrivilegeDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivilege
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrivilegeDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivilegeDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivilege
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPrivilege
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, UserPrivileges{})
			if err := m.Users[len(m.Users)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivilege(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivilege
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
func skipPrivilege(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrivilege
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowPrivilege
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowPrivilege
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPrivilege
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPrivilege
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipPrivilege(data[start:])
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
	ErrInvalidLengthPrivilege = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrivilege   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/sql/sqlbase/privilege.proto", fileDescriptorPrivilege) }

var fileDescriptorPrivilege = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0x4f, 0xce,
	0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x2f, 0x2e, 0xcc, 0x01, 0xe1, 0xa4, 0xc4, 0xe2, 0x54, 0xfd,
	0x82, 0xa2, 0xcc, 0xb2, 0xcc, 0x9c, 0xd4, 0xf4, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21,
	0x51, 0xb8, 0x32, 0xbd, 0xe2, 0xc2, 0x1c, 0x3d, 0xa8, 0x32, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c,
	0xb0, 0x0a, 0x7d, 0x10, 0x0b, 0xa2, 0x58, 0x29, 0x80, 0x8b, 0x2f, 0xb4, 0x38, 0xb5, 0x28, 0x00,
	0x66, 0x46, 0xb1, 0x90, 0x04, 0x17, 0x4b, 0x69, 0x71, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xa7, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x60, 0x11, 0x21, 0x15, 0x2e, 0x2e, 0xb8, 0x5d,
	0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xbc, 0x50, 0x79, 0x24, 0x71, 0xa5, 0x08, 0x2e, 0x61, 0xb8,
	0x69, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x42, 0x8e, 0x5c, 0xac, 0x20,
	0x43, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x54, 0xf5, 0xb0, 0xba, 0x52, 0x0f, 0xd5,
	0x31, 0x50, 0xe3, 0x21, 0x3a, 0x9d, 0x14, 0x4f, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x1b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0x21, 0x8a, 0x1d, 0xaa, 0x1d, 0x10, 0x00, 0x00, 0xff, 0xff, 0x90, 0x14, 0xdb, 0x9b, 0x23, 0x01,
	0x00, 0x00,
}
