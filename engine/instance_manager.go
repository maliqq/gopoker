package engine

import (
	"gopoker/message"
	"gopoker/model"
	"gopoker/model/seat"
)

func (i *Instance) JoinTable(msg *message.Join) {
	player, pos, amount := msg.Player, msg.Pos, msg.Amount

	_, err := i.Table.AddPlayer(player, pos, amount)

	if err != nil {
		i.e.Notify(&message.ErrorMessage{Error: err}).One(player)
	} else {
		i.e.Notify(msg).All()
	}
}

func (i *Instance) LeaveTable(msg *message.Leave) {
	player := msg.Player

	i.Table.RemovePlayer(player)
	i.e.Notify(msg).All()
}

func (i *Instance) SitOut(msg *message.Seat) {
	pos := msg.Pos
	i.Table.Seat(pos).State = seat.Idle
	i.e.Notify(&message.SeatState{
		Pos: pos,
		State: seat.Idle,
	}).All()
}

func (i *Instance) ComeBack(msg *message.Seat) {
	pos := msg.Pos
	i.Table.Seat(pos).State = seat.Ready
	i.e.Notify(&message.SeatState{
		Pos: pos,
		State: seat.Ready,
	}).All()
}

func (i *Instance) AddChatMessage(msg *message.ChatMessage) {
	i.e.Notify(msg).All()
}

func (i *Instance) AddBet(bet *model.Bet) {
	
/*	if !i.b.IsActive() {
		return
	}

	pos := msg.Pos
	if pos != i.Betting.Pos {
		return
	}
	//seat := i.Table.Seat(pos)

	i.b.Bet <- msg.Bet
	i.e.Pass(event).All()
	*/
}
