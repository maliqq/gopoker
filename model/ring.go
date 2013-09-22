package model

import (
  "container/ring"
)

import (
  "gopoker/model/seat"
)

type Ring struct {
  *ring.Ring
}

type Box struct {
  Pos int
  Seat *Seat
}

func (r *Ring) Where(f func(seat *Seat) bool) []Box {
  index := []Box{}
  r.Do(func(value interface{}) {
    box := value.(Box)
    if f(box.Seat) {
      index = append(index, box)
    }
  })

  return index
}

func (r *Ring) Active() []Box {
  return r.Where(func(s *Seat) bool {
    return s.State == seat.Play || s.State == seat.PostBB
  })
}

func (r *Ring) Waiting() []Box {
  return r.Where(func(s *Seat) bool {
    return s.State == seat.WaitBB
  })
}

func (r *Ring) Playing() []Box {
  return r.Where(func(s *Seat) bool {
    return s.State == seat.Play
  })
}

func (r *Ring) InPlay() []Box {
  return r.Where(func(s *Seat) bool {
    return s.State == seat.Play || s.State == seat.Bet
  })
}

// InPot - get all seats in pot
func (r *Ring) InPot() []Box {
  return r.Where(func(s *Seat) bool {
    return s.State == seat.Play || s.State == seat.Bet || s.State == seat.AllIn
  })
}
