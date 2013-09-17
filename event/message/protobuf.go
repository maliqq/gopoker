package message

import (
	"gopoker/event/message/protobuf"
)

func (m AddBet) Proto() *protobuf.Message {
	return protobuf.NewAddBet(m.Pos, m.Bet.Proto())
}

func (m RequireBet) Proto() *protobuf.Message {
	return protobuf.NewRequireBet(m.Pos, m.Range.Proto())
}

func (m BettingComplete) Proto() *protobuf.Message {
	return protobuf.NewBettingComplete(m.Pot)
}

func (m DealCards) Proto() *protobuf.Message {
	return protobuf.NewDealCards(m.Pos, m.Cards.Proto(), m.Type)
}

func (m RequireDiscard) Proto() *protobuf.Message {
	return protobuf.NewRequireDiscard(m.Pos)
}

func (m Discarded) Proto() *protobuf.Message {
	return protobuf.NewDiscarded(m.Pos, m.Num)
}

func (m DiscardCards) Proto() *protobuf.Message {
	return protobuf.NewDiscardCards(m.Pos, m.Cards.Proto())
}

func (m PlayStart) Proto() *protobuf.Message {
	return protobuf.NewPlayStart(nil)
}

func (m StreetStart) Proto() *protobuf.Message {
	return protobuf.NewStreetStart(m.Name)
}

func (m PlayStop) Proto() *protobuf.Message {
	return protobuf.NewPlayStop() // FIXME
}

func (m ShowCards) Proto() *protobuf.Message {
	return protobuf.NewShowCards(m.Pos, m.Player.Proto(), m.Cards.Proto())
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

func (m Winner) Proto() *protobuf.Message {
	return protobuf.NewWinner(m.Pos, m.Player.Proto(), m.Amount)
}

func (m MoveButton) Proto() *protobuf.Message {
	return protobuf.NewMoveButton(m.Pos)
}

func (m JoinTable) Proto() *protobuf.Message {
	return protobuf.NewJoinTable(m.Player.Proto(), m.Pos, m.Amount)
}

func (m SitOut) Proto() *protobuf.Message {
	return protobuf.NewSitOut(m.Pos)
}

func (m ComeBack) Proto() *protobuf.Message {
	return protobuf.NewComeBack(m.Pos)
}

func (m LeaveTable) Proto() *protobuf.Message {
	return protobuf.NewLeaveTable(m.Player.Proto())
}

func (m ErrorMessage) Proto() *protobuf.Message {
	return protobuf.NewErrorMessage(m.Error)
}

func (m ChatMessage) Proto() *protobuf.Message {
	return protobuf.NewChatMessage(m.Body)
}
