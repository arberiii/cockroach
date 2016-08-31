// Code generated by protoc-gen-gogo.
// source: cockroach/sql/sqlbase/encoded_datum.proto
// DO NOT EDIT!

package sqlbase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DatumEncoding identifies the encoding used for an EncDatum.
type DatumEncoding int32

const (
	// Indicates that the datum is encoded using the order-preserving encoding
	// used for keys (ascending order).
	DatumEncoding_ASCENDING_KEY DatumEncoding = 0
	// Indicates that the datum is encoded using the order-preserving encoding
	// used for keys (descending order).
	DatumEncoding_DESCENDING_KEY DatumEncoding = 1
	// Indicates that the datum is encoded using the encoding used for values.
	DatumEncoding_VALUE DatumEncoding = 2
)

var DatumEncoding_name = map[int32]string{
	0: "ASCENDING_KEY",
	1: "DESCENDING_KEY",
	2: "VALUE",
}
var DatumEncoding_value = map[string]int32{
	"ASCENDING_KEY":  0,
	"DESCENDING_KEY": 1,
	"VALUE":          2,
}

func (x DatumEncoding) Enum() *DatumEncoding {
	p := new(DatumEncoding)
	*p = x
	return p
}
func (x DatumEncoding) String() string {
	return proto.EnumName(DatumEncoding_name, int32(x))
}
func (x *DatumEncoding) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatumEncoding_value, data, "DatumEncoding")
	if err != nil {
		return err
	}
	*x = DatumEncoding(value)
	return nil
}
func (DatumEncoding) EnumDescriptor() ([]byte, []int) { return fileDescriptorEncodedDatum, []int{0} }

func init() {
	proto.RegisterEnum("cockroach.sql.sqlbase.DatumEncoding", DatumEncoding_name, DatumEncoding_value)
}

func init() {
	proto.RegisterFile("cockroach/sql/sqlbase/encoded_datum.proto", fileDescriptorEncodedDatum)
}

var fileDescriptorEncodedDatum = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xce, 0x4f, 0xce,
	0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x2f, 0x2e, 0xcc, 0x01, 0xe1, 0xa4, 0xc4, 0xe2, 0x54, 0xfd,
	0xd4, 0xbc, 0xe4, 0xfc, 0x94, 0xd4, 0x94, 0xf8, 0x94, 0xc4, 0x92, 0xd2, 0x5c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x51, 0xb8, 0x52, 0xbd, 0xe2, 0xc2, 0x1c, 0x3d, 0xa8, 0x52, 0x29, 0x91,
	0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x0a, 0x7d, 0x10, 0x0b, 0xa2, 0x58, 0xcb, 0x91, 0x8b, 0xd7, 0x05,
	0xa4, 0xd7, 0x15, 0x64, 0x50, 0x66, 0x5e, 0xba, 0x90, 0x20, 0x17, 0xaf, 0x63, 0xb0, 0xb3, 0xab,
	0x9f, 0x8b, 0xa7, 0x9f, 0x7b, 0xbc, 0xb7, 0x6b, 0xa4, 0x00, 0x83, 0x90, 0x10, 0x17, 0x9f, 0x8b,
	0x2b, 0x8a, 0x18, 0xa3, 0x10, 0x27, 0x17, 0x6b, 0x98, 0xa3, 0x4f, 0xa8, 0xab, 0x00, 0x93, 0x93,
	0xe2, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0x43, 0xed, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x6f, 0x78, 0x01, 0xa7, 0xbe, 0x00, 0x00, 0x00,
}
