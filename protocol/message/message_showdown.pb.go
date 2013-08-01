// Code generated by protoc-gen-go.
// source: protocol/message/message_showdown.proto
// DO NOT EDIT!

package message

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Rank int32

const (
	Rank_StraightFlush Rank = 0
	Rank_FourKind      Rank = 1
	Rank_FullHouse     Rank = 2
	Rank_Flush         Rank = 3
	Rank_Straight      Rank = 4
	Rank_ThreeKind     Rank = 5
	Rank_TwoPair       Rank = 6
	Rank_OnePair       Rank = 7
	Rank_HighCard      Rank = 8
	Rank_BadugiFour    Rank = 9
	Rank_BadugiThree   Rank = 10
	Rank_BadugiTwo     Rank = 11
	Rank_BadugiOne     Rank = 12
	Rank_CompleteLow   Rank = 13
	Rank_IncompleteLow Rank = 14
)

var Rank_name = map[int32]string{
	0:  "StraightFlush",
	1:  "FourKind",
	2:  "FullHouse",
	3:  "Flush",
	4:  "Straight",
	5:  "ThreeKind",
	6:  "TwoPair",
	7:  "OnePair",
	8:  "HighCard",
	9:  "BadugiFour",
	10: "BadugiThree",
	11: "BadugiTwo",
	12: "BadugiOne",
	13: "CompleteLow",
	14: "IncompleteLow",
}
var Rank_value = map[string]int32{
	"StraightFlush": 0,
	"FourKind":      1,
	"FullHouse":     2,
	"Flush":         3,
	"Straight":      4,
	"ThreeKind":     5,
	"TwoPair":       6,
	"OnePair":       7,
	"HighCard":      8,
	"BadugiFour":    9,
	"BadugiThree":   10,
	"BadugiTwo":     11,
	"BadugiOne":     12,
	"CompleteLow":   13,
	"IncompleteLow": 14,
}

func (x Rank) Enum() *Rank {
	p := new(Rank)
	*p = x
	return p
}
func (x Rank) String() string {
	return proto.EnumName(Rank_name, int32(x))
}
func (x Rank) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Rank) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Rank_value, data, "Rank")
	if err != nil {
		return err
	}
	*x = Rank(value)
	return nil
}

type Hand struct {
	Rank             *Rank  `protobuf:"varint,1,req,enum=message.Rank" json:"Rank,omitempty"`
	Value            []byte `protobuf:"bytes,2,req" json:"Value,omitempty"`
	High             []byte `protobuf:"bytes,3,req" json:"High,omitempty"`
	Kicker           []byte `protobuf:"bytes,4,req" json:"Kicker,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *Hand) Reset()         { *this = Hand{} }
func (this *Hand) String() string { return proto.CompactTextString(this) }
func (*Hand) ProtoMessage()       {}

func (this *Hand) GetRank() Rank {
	if this != nil && this.Rank != nil {
		return *this.Rank
	}
	return 0
}

func (this *Hand) GetValue() []byte {
	if this != nil {
		return this.Value
	}
	return nil
}

func (this *Hand) GetHigh() []byte {
	if this != nil {
		return this.High
	}
	return nil
}

func (this *Hand) GetKicker() []byte {
	if this != nil {
		return this.Kicker
	}
	return nil
}

type ShowHand struct {
	Pos              *int32  `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Player           *string `protobuf:"bytes,2,req" json:"Player,omitempty"`
	Cards            []byte  `protobuf:"bytes,3,req" json:"Cards,omitempty"`
	Hand             *Hand   `protobuf:"bytes,4,req" json:"Hand,omitempty"`
	HandString       *string `protobuf:"bytes,5,req" json:"HandString,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ShowHand) Reset()         { *this = ShowHand{} }
func (this *ShowHand) String() string { return proto.CompactTextString(this) }
func (*ShowHand) ProtoMessage()       {}

func (this *ShowHand) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *ShowHand) GetPlayer() string {
	if this != nil && this.Player != nil {
		return *this.Player
	}
	return ""
}

func (this *ShowHand) GetCards() []byte {
	if this != nil {
		return this.Cards
	}
	return nil
}

func (this *ShowHand) GetHand() *Hand {
	if this != nil {
		return this.Hand
	}
	return nil
}

func (this *ShowHand) GetHandString() string {
	if this != nil && this.HandString != nil {
		return *this.HandString
	}
	return ""
}

type ShowCards struct {
	Pos              *int32  `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Player           *string `protobuf:"bytes,2,req" json:"Player,omitempty"`
	Cards            []byte  `protobuf:"bytes,3,req" json:"Cards,omitempty"`
	Muck             *bool   `protobuf:"varint,4,req,def=0" json:"Muck,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ShowCards) Reset()         { *this = ShowCards{} }
func (this *ShowCards) String() string { return proto.CompactTextString(this) }
func (*ShowCards) ProtoMessage()       {}

const Default_ShowCards_Muck bool = false

func (this *ShowCards) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *ShowCards) GetPlayer() string {
	if this != nil && this.Player != nil {
		return *this.Player
	}
	return ""
}

func (this *ShowCards) GetCards() []byte {
	if this != nil {
		return this.Cards
	}
	return nil
}

func (this *ShowCards) GetMuck() bool {
	if this != nil && this.Muck != nil {
		return *this.Muck
	}
	return Default_ShowCards_Muck
}

type Winner struct {
	Pos              *int32   `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Player           *string  `protobuf:"bytes,2,req" json:"Player,omitempty"`
	Amount           *float64 `protobuf:"fixed64,3,req" json:"Amount,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Winner) Reset()         { *this = Winner{} }
func (this *Winner) String() string { return proto.CompactTextString(this) }
func (*Winner) ProtoMessage()       {}

func (this *Winner) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *Winner) GetPlayer() string {
	if this != nil && this.Player != nil {
		return *this.Player
	}
	return ""
}

func (this *Winner) GetAmount() float64 {
	if this != nil && this.Amount != nil {
		return *this.Amount
	}
	return 0
}

func init() {
	proto.RegisterEnum("message.Rank", Rank_name, Rank_value)
}
