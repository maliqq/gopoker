// Code generated by protoc-gen-go.
// source: event/message/format/protobuf/message_betting.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// BetType - bet type
type BetType int32

const (
	BetType_Ante       BetType = 1
	BetType_BringIn    BetType = 2
	BetType_SmallBlind BetType = 3
	BetType_BigBlind   BetType = 4
	BetType_GuestBlind BetType = 5
	BetType_Straddle   BetType = 6
	BetType_Raise      BetType = 7
	BetType_Call       BetType = 8
	BetType_Check      BetType = 9
	BetType_Fold       BetType = 10
	BetType_Discard    BetType = 11
	BetType_StandPat   BetType = 12
	BetType_Show       BetType = 13
	BetType_Muck       BetType = 14
)

var BetType_name = map[int32]string{
	1:  "Ante",
	2:  "BringIn",
	3:  "SmallBlind",
	4:  "BigBlind",
	5:  "GuestBlind",
	6:  "Straddle",
	7:  "Raise",
	8:  "Call",
	9:  "Check",
	10: "Fold",
	11: "Discard",
	12: "StandPat",
	13: "Show",
	14: "Muck",
}
var BetType_value = map[string]int32{
	"Ante":       1,
	"BringIn":    2,
	"SmallBlind": 3,
	"BigBlind":   4,
	"GuestBlind": 5,
	"Straddle":   6,
	"Raise":      7,
	"Call":       8,
	"Check":      9,
	"Fold":       10,
	"Discard":    11,
	"StandPat":   12,
	"Show":       13,
	"Muck":       14,
}

func (x BetType) Enum() *BetType {
	p := new(BetType)
	*p = x
	return p
}
func (x BetType) String() string {
	return proto.EnumName(BetType_name, int32(x))
}
func (x BetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *BetType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BetType_value, data, "BetType")
	if err != nil {
		return err
	}
	*x = BetType(value)
	return nil
}

// Bet - bet
type Bet struct {
	Type             *BetType `protobuf:"varint,1,req,enum=protobuf.BetType" json:"Type,omitempty"`
	Amount           *float64 `protobuf:"fixed64,2,opt" json:"Amount,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Bet) Reset()         { *m = Bet{} }
func (m *Bet) String() string { return proto.CompactTextString(m) }
func (*Bet) ProtoMessage()    {}

func (m *Bet) GetType() BetType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Bet) GetAmount() float64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// BetRange - bet range
type BetRange struct {
	Call             *float64 `protobuf:"fixed64,1,req" json:"Call,omitempty"`
	Min              *float64 `protobuf:"fixed64,2,req" json:"Min,omitempty"`
	Max              *float64 `protobuf:"fixed64,3,req" json:"Max,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BetRange) Reset()         { *m = BetRange{} }
func (m *BetRange) String() string { return proto.CompactTextString(m) }
func (*BetRange) ProtoMessage()    {}

func (m *BetRange) GetCall() float64 {
	if m != nil && m.Call != nil {
		return *m.Call
	}
	return 0
}

func (m *BetRange) GetMin() float64 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *BetRange) GetMax() float64 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

// RequireBet - require bet
type RequireBet struct {
	Pos              *int32    `protobuf:"varint,1,req" json:"Pos,omitempty"`
	BetRange         *BetRange `protobuf:"bytes,2,req" json:"BetRange,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *RequireBet) Reset()         { *m = RequireBet{} }
func (m *RequireBet) String() string { return proto.CompactTextString(m) }
func (*RequireBet) ProtoMessage()    {}

func (m *RequireBet) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *RequireBet) GetBetRange() *BetRange {
	if m != nil {
		return m.BetRange
	}
	return nil
}

// AddBet - add bet
type AddBet struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Bet              *Bet   `protobuf:"bytes,2,req" json:"Bet,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AddBet) Reset()         { *m = AddBet{} }
func (m *AddBet) String() string { return proto.CompactTextString(m) }
func (*AddBet) ProtoMessage()    {}

func (m *AddBet) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *AddBet) GetBet() *Bet {
	if m != nil {
		return m.Bet
	}
	return nil
}

// BettingComplete - betting complete
type BettingComplete struct {
	Pot              *float64 `protobuf:"fixed64,1,req" json:"Pot,omitempty"`
	Rake             *float64 `protobuf:"fixed64,2,opt" json:"Rake,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BettingComplete) Reset()         { *m = BettingComplete{} }
func (m *BettingComplete) String() string { return proto.CompactTextString(m) }
func (*BettingComplete) ProtoMessage()    {}

func (m *BettingComplete) GetPot() float64 {
	if m != nil && m.Pot != nil {
		return *m.Pot
	}
	return 0
}

func (m *BettingComplete) GetRake() float64 {
	if m != nil && m.Rake != nil {
		return *m.Rake
	}
	return 0
}

func init() {
	proto.RegisterEnum("protobuf.BetType", BetType_name, BetType_value)
}
