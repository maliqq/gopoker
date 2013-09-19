package play

func (play *Play) Start() {
  go play.Run()
  play.NextDeal <- time.After(5 * time.Second)
  play.Broadcast.Notify(&message.NextDealAnnounce{
    After: 5,
  })
}

func (play *Play) Stop() {
}

func (play *Play) Pause() {
}

func (play *Play) Resume() {
}

func (play *Play) Close() {
}

func (play *Play) AddBet() {
}

func (play *Play) JoinTable(msg *message.Player) {
    player, pos, amount := msg.Player, msg.Pos, msg.Amount
    
    _, err := play.Table.AddPlayer(player, pos, amount)
    
    if err != nil {
      play.Broadcast.Notify(&message.ErrorMessage{Error: err}).One(player)
    } else {
      play.Broadcast.Notify(msg).All()
    }
}

func (play *Play) LeaveTable(msg *message.Player) {
    player := msg.Player

    play.Table.RemovePlayer(player)
    play.Broadcast.Notify(msg).All()
}

func (play *Play) SitOut(msg *message.Position) {
    pos := msg.Pos
    play.Table.Seat(pos).State = seat.Idle
    play.Broadcast.Notify(&message.SeatState{
      State: seat.Idle,
    }).All()
}

func (play *Play) ComeBack(msg *message.Position) {
    pos := msg.Pos
    play.Table.Seat(pos).State = seat.Ready
    play.Broadcast.Notify(&message.SeatState{
      State: seat.Ready,
    }).All()
}

func (play *Play) AddChatMessage(msg *message.ChatMessage) {
    play.Broadcast.Notify(msg).All()
}

func (play *Play) AddBet(msg *message.Bet) {
    if !play.Betting.IsActive() {
      return
    }

    pos := msg.Pos
    if pos != play.Betting.Pos {
      return
    }
    //seat := play.Table.Seat(pos)

    play.Betting.Bet <- msg.Bet
    play.Broadcast.Pass(event).All()
}
