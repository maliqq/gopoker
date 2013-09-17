// Code generated by protoc-gen-go.
// source: event/message/format/protobuf/message_deal.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// DealType - deal type
type DealType int32

const (
	DealType_Board DealType = 0
	DealType_Hole  DealType = 1
	DealType_Door  DealType = 2
)

var DealType_name = map[int32]string{
	0: "Board",
	1: "Hole",
	2: "Door",
}
var DealType_value = map[string]int32{
	"Board": 0,
	"Hole":  1,
	"Door":  2,
}

func (x DealType) Enum() *DealType {
	p := new(DealType)
	*p = x
	return p
}
func (x DealType) String() string {
	return proto.EnumName(DealType_name, int32(x))
}
func (x DealType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DealType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DealType_value, data, "DealType")
	if err != nil {
		return err
	}
	*x = DealType(value)
	return nil
}

// DealCards - deal cards
type DealCards struct {
	Pos              *int32    `protobuf:"varint,1,opt" json:"Pos,omitempty"`
	Cards            []byte    `protobuf:"bytes,2,req" json:"Cards,omitempty"`
	Type             *DealType `protobuf:"varint,3,req,enum=protobuf.DealType" json:"Type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *DealCards) Reset()         { *m = DealCards{} }
func (m *DealCards) String() string { return proto.CompactTextString(m) }
func (*DealCards) ProtoMessage()    {}

func (m *DealCards) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *DealCards) GetCards() []byte {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DealCards) GetType() DealType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

// RequireDiscard - require discard
type RequireDiscard struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequireDiscard) Reset()         { *m = RequireDiscard{} }
func (m *RequireDiscard) String() string { return proto.CompactTextString(m) }
func (*RequireDiscard) ProtoMessage()    {}

func (m *RequireDiscard) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

// Discarded - discarded n cards
type Discarded struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Num              *int32 `protobuf:"varint,2,req" json:"Num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Discarded) Reset()         { *m = Discarded{} }
func (m *Discarded) String() string { return proto.CompactTextString(m) }
func (*Discarded) ProtoMessage()    {}

func (m *Discarded) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *Discarded) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// DiscardCards - deal discard cards
type DiscardCards struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Cards            []byte `protobuf:"bytes,2,req" json:"Cards,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DiscardCards) Reset()         { *m = DiscardCards{} }
func (m *DiscardCards) String() string { return proto.CompactTextString(m) }
func (*DiscardCards) ProtoMessage()    {}

func (m *DiscardCards) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *DiscardCards) GetCards() []byte {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.DealType", DealType_name, DealType_value)
}
