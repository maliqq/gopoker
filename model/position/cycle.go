package position

// Cycle - cycle position
func Cycle(pos int, max int) int {
	if pos >= max {
		pos = 0
	}
	return pos
}
