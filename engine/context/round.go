package context

import (
  "container/ring"
)

import (
  "gopoker/model"
)

type Round struct {
  *ring.Ring
}

func (r *Round) Box() model.Box {
  return r.Ring.Value.(model.Box)
}

func (r *Round) Pos() int {
  return r.Box().Pos
}

func (r *Round) Current() *model.Seat {
  return r.Box().Seat
}

func (r *Round) Move() {
  r.Ring = r.Ring.Move(1)
}
