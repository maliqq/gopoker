// Code generated by protoc-gen-go.
// source: protocol/message/message.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type BetType int32

const (
	BetType_CHECK BetType = 0
	BetType_CALL  BetType = 1
	BetType_FOLD  BetType = 2
	BetType_RAISE BetType = 3
)

var BetType_name = map[int32]string{
	0: "CHECK",
	1: "CALL",
	2: "FOLD",
	3: "RAISE",
}
var BetType_value = map[string]int32{
	"CHECK": 0,
	"CALL":  1,
	"FOLD":  2,
	"RAISE": 3,
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

type DealType int32

const (
	DealType_BOARD   DealType = 0
	DealType_HOLE    DealType = 1
	DealType_DOOR    DealType = 2
	DealType_DISCARD DealType = 3
)

var DealType_name = map[int32]string{
	0: "BOARD",
	1: "HOLE",
	2: "DOOR",
	3: "DISCARD",
}
var DealType_value = map[string]int32{
	"BOARD":   0,
	"HOLE":    1,
	"DOOR":    2,
	"DISCARD": 3,
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

type GameType int32

const (
	GameType_TEXAS    GameType = 0
	GameType_OMAHA    GameType = 1
	GameType_OMAHA8   GameType = 2
	GameType_STUD     GameType = 3
	GameType_STUD8    GameType = 4
	GameType_RAZZ     GameType = 9
	GameType_LONDON   GameType = 10
	GameType_FIVECARD GameType = 11
	GameType_SINGLE27 GameType = 12
	GameType_TRIPLE27 GameType = 13
	GameType_BADUGI   GameType = 14
)

var GameType_name = map[int32]string{
	0:  "TEXAS",
	1:  "OMAHA",
	2:  "OMAHA8",
	3:  "STUD",
	4:  "STUD8",
	9:  "RAZZ",
	10: "LONDON",
	11: "FIVECARD",
	12: "SINGLE27",
	13: "TRIPLE27",
	14: "BADUGI",
}
var GameType_value = map[string]int32{
	"TEXAS":    0,
	"OMAHA":    1,
	"OMAHA8":   2,
	"STUD":     3,
	"STUD8":    4,
	"RAZZ":     9,
	"LONDON":   10,
	"FIVECARD": 11,
	"SINGLE27": 12,
	"TRIPLE27": 13,
	"BADUGI":   14,
}

func (x GameType) Enum() *GameType {
	p := new(GameType)
	*p = x
	return p
}
func (x GameType) String() string {
	return proto.EnumName(GameType_name, int32(x))
}
func (x GameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *GameType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GameType_value, data, "GameType")
	if err != nil {
		return err
	}
	*x = GameType(value)
	return nil
}

type GameLimit int32

const (
	GameLimit_NO_LIMIT    GameLimit = 0
	GameLimit_POT_LIMIT   GameLimit = 1
	GameLimit_FIXED_LIMIT GameLimit = 2
)

var GameLimit_name = map[int32]string{
	0: "NO_LIMIT",
	1: "POT_LIMIT",
	2: "FIXED_LIMIT",
}
var GameLimit_value = map[string]int32{
	"NO_LIMIT":    0,
	"POT_LIMIT":   1,
	"FIXED_LIMIT": 2,
}

func (x GameLimit) Enum() *GameLimit {
	p := new(GameLimit)
	*p = x
	return p
}
func (x GameLimit) String() string {
	return proto.EnumName(GameLimit_name, int32(x))
}
func (x GameLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *GameLimit) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GameLimit_value, data, "GameLimit")
	if err != nil {
		return err
	}
	*x = GameLimit(value)
	return nil
}

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

type ErrorMessage struct {
	Message          *string `protobuf:"bytes,1,req" json:"Message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ErrorMessage) Reset()         { *this = ErrorMessage{} }
func (this *ErrorMessage) String() string { return proto.CompactTextString(this) }
func (*ErrorMessage) ProtoMessage()       {}

func (this *ErrorMessage) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

type ChatMessage struct {
	Pos              *int32  `protobuf:"varint,1,req,name=pos" json:"pos,omitempty"`
	Message          *string `protobuf:"bytes,2,req" json:"Message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ChatMessage) Reset()         { *this = ChatMessage{} }
func (this *ChatMessage) String() string { return proto.CompactTextString(this) }
func (*ChatMessage) ProtoMessage()       {}

func (this *ChatMessage) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *ChatMessage) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

type DealerMessage struct {
	Message          *string `protobuf:"bytes,1,req" json:"Message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *DealerMessage) Reset()         { *this = DealerMessage{} }
func (this *DealerMessage) String() string { return proto.CompactTextString(this) }
func (*DealerMessage) ProtoMessage()       {}

func (this *DealerMessage) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

type BetRange struct {
	Call             *float32 `protobuf:"fixed32,1,req" json:"Call,omitempty"`
	Min              *float32 `protobuf:"fixed32,2,req" json:"Min,omitempty"`
	Max              *float32 `protobuf:"fixed32,3,req" json:"Max,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *BetRange) Reset()         { *this = BetRange{} }
func (this *BetRange) String() string { return proto.CompactTextString(this) }
func (*BetRange) ProtoMessage()       {}

func (this *BetRange) GetCall() float32 {
	if this != nil && this.Call != nil {
		return *this.Call
	}
	return 0
}

func (this *BetRange) GetMin() float32 {
	if this != nil && this.Min != nil {
		return *this.Min
	}
	return 0
}

func (this *BetRange) GetMax() float32 {
	if this != nil && this.Max != nil {
		return *this.Max
	}
	return 0
}

type RequireBet struct {
	Pos              *int32    `protobuf:"varint,1,req,name=pos" json:"pos,omitempty"`
	BetRange         *BetRange `protobuf:"bytes,2,req" json:"BetRange,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (this *RequireBet) Reset()         { *this = RequireBet{} }
func (this *RequireBet) String() string { return proto.CompactTextString(this) }
func (*RequireBet) ProtoMessage()       {}

func (this *RequireBet) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *RequireBet) GetBetRange() *BetRange {
	if this != nil {
		return this.BetRange
	}
	return nil
}

type Bet struct {
	Type             *BetType `protobuf:"varint,1,req,enum=protocol.BetType" json:"Type,omitempty"`
	Amount           *float32 `protobuf:"fixed32,2,opt" json:"Amount,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Bet) Reset()         { *this = Bet{} }
func (this *Bet) String() string { return proto.CompactTextString(this) }
func (*Bet) ProtoMessage()       {}

func (this *Bet) GetType() BetType {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Bet) GetAmount() float32 {
	if this != nil && this.Amount != nil {
		return *this.Amount
	}
	return 0
}

type AddBet struct {
	Pos              *int32 `protobuf:"varint,1,req,name=pos" json:"pos,omitempty"`
	Bet              *Bet   `protobuf:"bytes,2,req" json:"Bet,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *AddBet) Reset()         { *this = AddBet{} }
func (this *AddBet) String() string { return proto.CompactTextString(this) }
func (*AddBet) ProtoMessage()       {}

func (this *AddBet) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *AddBet) GetBet() *Bet {
	if this != nil {
		return this.Bet
	}
	return nil
}

type BettingComplete struct {
	Pot              *float32 `protobuf:"fixed32,1,req" json:"Pot,omitempty"`
	Rake             *float32 `protobuf:"fixed32,2,req" json:"Rake,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *BettingComplete) Reset()         { *this = BettingComplete{} }
func (this *BettingComplete) String() string { return proto.CompactTextString(this) }
func (*BettingComplete) ProtoMessage()       {}

func (this *BettingComplete) GetPot() float32 {
	if this != nil && this.Pot != nil {
		return *this.Pot
	}
	return 0
}

func (this *BettingComplete) GetRake() float32 {
	if this != nil && this.Rake != nil {
		return *this.Rake
	}
	return 0
}

type DealCards struct {
	Pos              *int32    `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Cards            []int64   `protobuf:"varint,2,rep" json:"Cards,omitempty"`
	Type             *DealType `protobuf:"varint,3,req,enum=protocol.DealType" json:"Type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (this *DealCards) Reset()         { *this = DealCards{} }
func (this *DealCards) String() string { return proto.CompactTextString(this) }
func (*DealCards) ProtoMessage()       {}

func (this *DealCards) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *DealCards) GetCards() []int64 {
	if this != nil {
		return this.Cards
	}
	return nil
}

func (this *DealCards) GetType() DealType {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

type RequireDiscard struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RequireDiscard) Reset()         { *this = RequireDiscard{} }
func (this *RequireDiscard) String() string { return proto.CompactTextString(this) }
func (*RequireDiscard) ProtoMessage()       {}

func (this *RequireDiscard) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

type Discarded struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Num              *int32 `protobuf:"varint,2,req" json:"Num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *Discarded) Reset()         { *this = Discarded{} }
func (this *Discarded) String() string { return proto.CompactTextString(this) }
func (*Discarded) ProtoMessage()       {}

func (this *Discarded) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *Discarded) GetNum() int32 {
	if this != nil && this.Num != nil {
		return *this.Num
	}
	return 0
}

type DiscardCards struct {
	Pos              *int32  `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Cards            []int32 `protobuf:"varint,2,rep" json:"Cards,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *DiscardCards) Reset()         { *this = DiscardCards{} }
func (this *DiscardCards) String() string { return proto.CompactTextString(this) }
func (*DiscardCards) ProtoMessage()       {}

func (this *DiscardCards) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *DiscardCards) GetCards() []int32 {
	if this != nil {
		return this.Cards
	}
	return nil
}

type PlayStart struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *PlayStart) Reset()         { *this = PlayStart{} }
func (this *PlayStart) String() string { return proto.CompactTextString(this) }
func (*PlayStart) ProtoMessage()       {}

type StreetStart struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *StreetStart) Reset()         { *this = StreetStart{} }
func (this *StreetStart) String() string { return proto.CompactTextString(this) }
func (*StreetStart) ProtoMessage()       {}

func (this *StreetStart) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

type ChangeGame struct {
	Type             *GameType  `protobuf:"varint,1,req,enum=protocol.GameType" json:"Type,omitempty"`
	Limit            *GameLimit `protobuf:"varint,2,req,enum=protocol.GameLimit" json:"Limit,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *ChangeGame) Reset()         { *this = ChangeGame{} }
func (this *ChangeGame) String() string { return proto.CompactTextString(this) }
func (*ChangeGame) ProtoMessage()       {}

func (this *ChangeGame) GetType() GameType {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *ChangeGame) GetLimit() GameLimit {
	if this != nil && this.Limit != nil {
		return *this.Limit
	}
	return 0
}

type Hand struct {
	Rank             *Rank   `protobuf:"varint,1,req,enum=protocol.Rank" json:"Rank,omitempty"`
	Value            []int32 `protobuf:"varint,2,rep" json:"Value,omitempty"`
	High             []int32 `protobuf:"varint,3,rep" json:"High,omitempty"`
	Kicker           []int32 `protobuf:"varint,4,rep" json:"Kicker,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
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

func (this *Hand) GetValue() []int32 {
	if this != nil {
		return this.Value
	}
	return nil
}

func (this *Hand) GetHigh() []int32 {
	if this != nil {
		return this.High
	}
	return nil
}

func (this *Hand) GetKicker() []int32 {
	if this != nil {
		return this.Kicker
	}
	return nil
}

type ShowHand struct {
	Pos              *int32  `protobuf:"varint,1,req" json:"Pos,omitempty"`
	Cards            []int32 `protobuf:"varint,2,rep" json:"Cards,omitempty"`
	Hand             *Hand   `protobuf:"bytes,3,req" json:"Hand,omitempty"`
	HandString       *string `protobuf:"bytes,4,req" json:"HandString,omitempty"`
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

func (this *ShowHand) GetCards() []int32 {
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
	Cards            []int32 `protobuf:"varint,2,rep" json:"Cards,omitempty"`
	Muck             *bool   `protobuf:"varint,3,req,def=0" json:"Muck,omitempty"`
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

func (this *ShowCards) GetCards() []int32 {
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
	Amount           *float32 `protobuf:"fixed32,2,req" json:"Amount,omitempty"`
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

func (this *Winner) GetAmount() float32 {
	if this != nil && this.Amount != nil {
		return *this.Amount
	}
	return 0
}

type MoveButton struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *MoveButton) Reset()         { *this = MoveButton{} }
func (this *MoveButton) String() string { return proto.CompactTextString(this) }
func (*MoveButton) ProtoMessage()       {}

func (this *MoveButton) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

type JoinTable struct {
	Player           *string  `protobuf:"bytes,1,req" json:"Player,omitempty"`
	Pos              *int32   `protobuf:"varint,2,req" json:"Pos,omitempty"`
	Amount           *float32 `protobuf:"fixed32,3,req" json:"Amount,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *JoinTable) Reset()         { *this = JoinTable{} }
func (this *JoinTable) String() string { return proto.CompactTextString(this) }
func (*JoinTable) ProtoMessage()       {}

func (this *JoinTable) GetPlayer() string {
	if this != nil && this.Player != nil {
		return *this.Player
	}
	return ""
}

func (this *JoinTable) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

func (this *JoinTable) GetAmount() float32 {
	if this != nil && this.Amount != nil {
		return *this.Amount
	}
	return 0
}

type LeaveTable struct {
	Player           *string `protobuf:"bytes,1,req" json:"Player,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *LeaveTable) Reset()         { *this = LeaveTable{} }
func (this *LeaveTable) String() string { return proto.CompactTextString(this) }
func (*LeaveTable) ProtoMessage()       {}

func (this *LeaveTable) GetPlayer() string {
	if this != nil && this.Player != nil {
		return *this.Player
	}
	return ""
}

type SitOut struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *SitOut) Reset()         { *this = SitOut{} }
func (this *SitOut) String() string { return proto.CompactTextString(this) }
func (*SitOut) ProtoMessage()       {}

func (this *SitOut) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

type ComeBack struct {
	Pos              *int32 `protobuf:"varint,1,req" json:"Pos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *ComeBack) Reset()         { *this = ComeBack{} }
func (this *ComeBack) String() string { return proto.CompactTextString(this) }
func (*ComeBack) ProtoMessage()       {}

func (this *ComeBack) GetPos() int32 {
	if this != nil && this.Pos != nil {
		return *this.Pos
	}
	return 0
}

type Envelope struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *Envelope) Reset()         { *this = Envelope{} }
func (this *Envelope) String() string { return proto.CompactTextString(this) }
func (*Envelope) ProtoMessage()       {}

func init() {
	proto.RegisterEnum("protocol.BetType", BetType_name, BetType_value)
	proto.RegisterEnum("protocol.DealType", DealType_name, DealType_value)
	proto.RegisterEnum("protocol.GameType", GameType_name, GameType_value)
	proto.RegisterEnum("protocol.GameLimit", GameLimit_name, GameLimit_value)
	proto.RegisterEnum("protocol.Rank", Rank_name, Rank_value)
}
