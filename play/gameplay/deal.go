package gameplay

func (this *GamePlay) StartNewDeal() {
  this.Deal = model.NewDeal()
  this.Betting = NewBetting()
  if this.Game.Discards {
    this.Discarding = NewDiscarding(this.Deal)
  }
}

func (this *GamePlay) ScheduleNextDeal() {
  this.Control <- command.NextDeal
}
