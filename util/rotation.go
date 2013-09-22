package util

type Rotatable interface {
  RotateLen() int
}

// Rotation - game rotation context
type Rotation struct {
  Rotatable
  every   int
  index   int
  counter int
}

// NewRotation - create rotation context
func NewRotation(rotatable Rotatable, every int) *Rotation {
  if every == 0 {
    every = rotatable.RotateLen()
  }

  return &Rotation{
    Rotatable:   rotatable,
    every: every,
  }
}

// Current - get current game in rotation
func (rotation *Rotation) Current() int {
  return rotation.index
}

// Next - get next rotated game
func (rotation *Rotation) Move() bool {
  hasNext := rotation.counter >= rotation.every
  
  if hasNext {
    rotation.counter = 0
    rotation.rotate()
  } else {
    rotation.counter++
  }

  return hasNext
}

func (rotation *Rotation) rotate() {
  rotation.index++
  if rotation.index >= rotation.Rotatable.RotateLen() {
    rotation.index = 0
  }
}
