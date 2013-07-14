package gameplay

func (this *GamePlay) ResetSeats() {
  for _, s := range this.Table.Seats {
    switch s.State {
    case seat.Ready, seat.Play:
      s.Play()
    }
  }
}

func (this *GamePlay) SetButton(pos int) {
  this.Table.SetButton(pos)

  this.Broadcast.All <- protocol.NewMoveButton(pos)
}

func (this *GamePlay) MoveButton() {
  this.Table.MoveButton()

  this.Broadcast.All <- protocol.NewMoveButton(this.Table.Button)
}
