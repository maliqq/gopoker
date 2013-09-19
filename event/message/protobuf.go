package message

import (
	"fmt"
)

import (
	"gopoker/event/message/protobuf"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/poker"
)

func (m AddBet) Proto() *protobuf.Message {
	return protobuf.NewAddBet(m.Pos, m.Bet.Proto())
}
func (m *AddBet) Unproto(msg *protobuf.Message) {
	p := msg.AddBet
	bet := &model.Bet{}
	bet.Unproto(p.Bet)
	*m = AddBet{
		int(p.GetPos()),
		bet,
	}
}

func (m RequireBet) Proto() *protobuf.Message {
	return protobuf.NewRequireBet(m.Pos, m.Range.Proto())
}
func (m *RequireBet) Unproto(msg *protobuf.Message) {
	p := msg.RequireBet
	betRange := &bet.Range{}
	betRange.Unproto(p.BetRange)
	*m = RequireBet{
		int(p.GetPos()),
		betRange,
	}
}

func (m BettingComplete) Proto() *protobuf.Message {
	return protobuf.NewBettingComplete(m.Pot)
}
func (m *BettingComplete) Unproto(msg *protobuf.Message) {
	p := msg.BettingComplete
	*m = BettingComplete{
		p.GetPot(),
	}
}

func (m DealCards) Proto() *protobuf.Message {
	return protobuf.NewDealCards(m.Pos, m.Cards.Proto(), m.Type)
}
func (m *DealCards) Unproto(msg *protobuf.Message) {
	p := msg.DealCards
	dealType := deal.Type(p.GetType().String())
	*m = DealCards{
		int(p.GetPos()),
		poker.BinaryCards(p.Cards),
		dealType,
	}
}

func (m RequireDiscard) Proto() *protobuf.Message {
	return protobuf.NewRequireDiscard(m.Pos)
}
func (m *RequireDiscard) Unproto(msg *protobuf.Message) {
	p := msg.RequireDiscard
	*m = RequireDiscard{
		int(p.GetPos()),
	}
}

func (m Discarded) Proto() *protobuf.Message {
	return protobuf.NewDiscarded(m.Pos, m.Num)
}
func (m *Discarded) Unproto(msg *protobuf.Message) {
	p := msg.Discarded
	*m = Discarded{
		int(p.GetPos()),
		int(p.GetNum()),
	}
}

func (m DiscardCards) Proto() *protobuf.Message {
	return protobuf.NewDiscardCards(m.Pos, m.Cards.Proto())
}
func (m *DiscardCards) Unproto(msg *protobuf.Message) {
	p := msg.DiscardCards
	*m = DiscardCards{
		int(p.GetPos()),
		poker.BinaryCards(p.Cards),
	}
}

func (m PlayStart) Proto() *protobuf.Message {
	return protobuf.NewPlayStart(m.Game.Proto(), m.Stake.Proto(), m.Table.Proto())
}
func (m *PlayStart) Unproto(msg *protobuf.Message) {
	p := msg.PlayStart

	game := &model.Game{}
	game.Unproto(p.Game)

	stake := &model.Stake{}
	stake.Unproto(p.Stake)

	table := &model.Table{}
	table.Unproto(p.Table)

	*m = PlayStart{
		Game:  game,
		Stake: stake,
		Table: table,
	}
}

func (m StreetStart) Proto() *protobuf.Message {
	return protobuf.NewStreetStart(m.Name)
}
func (m *StreetStart) Unproto(msg *protobuf.Message) {
	p := msg.StreetStart
	*m = StreetStart{
		p.GetName(),
	}
}

func (m PlayStop) Proto() *protobuf.Message {
	return protobuf.NewPlayStop() // FIXME
}
func (m *PlayStop) Unproto(msg *protobuf.Message) {
	//p := msg.PlayStop
	*m = PlayStop{}
}

func (m ShowCards) Proto() *protobuf.Message {
	return protobuf.NewShowCards(m.Pos, m.Player.Proto(), m.Cards.Proto())
}
func (m *ShowCards) Unproto(msg *protobuf.Message) {
	p := msg.ShowCards
	*m = ShowCards{
		int(p.GetPos()),
		p.GetMuck(),
		poker.BinaryCards(p.Cards),
		model.Player(p.GetPlayer()),
	}
}

func (m ShowHand) Proto() *protobuf.Message {
	return protobuf.NewShowHand(
		m.Pos,
		m.Player.Proto(),
		m.Cards.Proto(),
		m.Hand.Proto(),
		m.HandName,
	)
}
func (m *ShowHand) Unproto(msg *protobuf.Message) {
	p := msg.ShowHand
	hand := &poker.Hand{}
	hand.Unproto(p.Hand)

	*m = ShowHand{
		int(p.GetPos()),
		model.Player(p.GetPlayer()),
		poker.BinaryCards(p.Cards),
		*hand,
		p.GetHandString(),
	}
}

func (m Winner) Proto() *protobuf.Message {
	return protobuf.NewWinner(m.Pos, m.Player.Proto(), m.Amount)
}
func (m *Winner) Unproto(msg *protobuf.Message) {
	p := msg.Winner
	*m = Winner{
		int(p.GetPos()),
		model.Player(p.GetPlayer()),
		p.GetAmount(),
	}
}

func (m MoveButton) Proto() *protobuf.Message {
	return protobuf.NewMoveButton(m.Pos)
}
func (m *MoveButton) Unproto(msg *protobuf.Message) {
	p := msg.MoveButton
	*m = MoveButton{
		int(p.GetPos()),
	}
}

func (m JoinTable) Proto() *protobuf.Message {
	return protobuf.NewJoinTable(m.Player.Proto(), m.Pos, m.Amount)
}
func (m *JoinTable) Unproto(msg *protobuf.Message) {
	p := msg.JoinTable
	*m = JoinTable{
		model.Player(p.GetPlayer()),
		int(p.GetPos()),
		p.GetAmount(),
	}
}

func (m SitOut) Proto() *protobuf.Message {
	return protobuf.NewSitOut(m.Pos)
}
func (m *SitOut) Unproto(msg *protobuf.Message) {
	p := msg.SitOut
	*m = SitOut{
		int(p.GetPos()),
	}
}

func (m ComeBack) Proto() *protobuf.Message {
	return protobuf.NewComeBack(m.Pos)
}
func (m *ComeBack) Unproto(msg *protobuf.Message) {
	p := msg.ComeBack
	*m = ComeBack{
		int(p.GetPos()),
	}
}

func (m LeaveTable) Proto() *protobuf.Message {
	return protobuf.NewLeaveTable(m.Player.Proto())
}
func (m *LeaveTable) Unproto(msg *protobuf.Message) {
	p := msg.LeaveTable
	*m = LeaveTable{
		model.Player(p.GetPlayer()),
	}
}

func (m ErrorMessage) Proto() *protobuf.Message {
	return protobuf.NewErrorMessage(m.Error)
}
func (m *ErrorMessage) Unproto(msg *protobuf.Message) {
	p := msg.ErrorMessage
	*m = ErrorMessage{
		fmt.Errorf(p.GetError()),
	}
}

func (m ChatMessage) Proto() *protobuf.Message {
	return protobuf.NewChatMessage(m.Body)
}
func (m *ChatMessage) Unproto(msg *protobuf.Message) {
	p := msg.ChatMessage
	*m = ChatMessage{
		p.GetBody(),
	}
}
