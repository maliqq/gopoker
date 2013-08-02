package poker

// Flags
const (
	HighCardFlag = 0x100000 // HighCardFlag - high card flag

	OnePairFlag = 0x200000 // OnePairFlag - one pair flag

	TwoPairFlag = 0x300000 // TwoPairFlag - two pair flag

	ThreeKindFlag = 0x400000 // ThreeKindFlag - three kind flag

	StraightFlag = 0x500000 // StraightFlag - straight flag

	FlushFlag = 0x600000 // FlushFlag - flush flag

	FullHouseFlag = 0x700000 // FullHouseFlag - full house flag

	FourKindFlag = 0x800000 // FourKindFlag - four of a kind flag

	StraightFlushFlag = 0x900000 // StraightFlushFlag - straight flush flag
)

// Indexes
var (
	Flush = make([]uint, 8129) // Flush index

	Straight = make([]uint, 8129) // Straight index

	Top1Of16 = make([]uint, 8129) // Top1Of16 index

	Top1Of12 = make([]uint, 8129) // Top1Of12 index

	Top1Of8 = make([]uint, 8129) // Top1Of8 index

	Top2Of12 = make([]uint, 8129) // Top2Of12 index

	Top2Of8 = make([]uint, 8129) // Top2Of8 index

	Top3Of4 = make([]uint, 8129) // Top3Of4 index

	Top5 = make([]uint, 8129) // Top5 index

	Bit1 = make([]uint, 8129) // Bit1 index

	Bit2 = make([]uint, 8129) // Bit2 index
)

func doRank(hand uint64) uint {
	var c, h, d, s uint
	var p1, p2, p3, p4 uint

	s = uint(hand & 0x1fff)
	h = uint((hand >> 16) & 0x1fff)
	d = uint((hand >> 32) & 0x1fff)
	c = uint((hand >> 48) & 0x1fff)

	if Flush[s]|Flush[h]|Flush[d]|Flush[c] != 0 {
		return Flush[s] | Flush[h] | Flush[d] | Flush[c]
	}

	p1 = s
	p2 = p1 & h
	p1 = p1 | h
	p3 = p2 & d
	p2 = p2 | (p1 & d)
	p1 = p1 | d
	p4 = p3 & c
	p3 = p3 | (p2 & c)
	p2 = p2 | (p1 & c)
	p1 = p1 | c

	if Straight[p1] != 0 {
		return Straight[p1]
	}

	if p2 == 0 { // There are no pairs
		return HighCardFlag | uint(Top5[p1])
	}

	if p3 == 0 { // There are pairs but no triplets
		if Bit2[p2] == 0 {
			return OnePairFlag | uint(Top1Of16[p2]) | uint(Top3Of4[p1^Bit1[p2]])
		}
		return TwoPairFlag | uint(Top2Of12[p2]) | uint(Top1Of8[p1^Bit2[p2]])
	}

	if p4 == 0 { // Deal with trips/sets/boats
		if (p2 > p3) || (p3&(p3-1) != 0) {
			return FullHouseFlag | uint(Top1Of16[p3]) | uint(Top1Of12[p2^Bit1[p3]])
		}
		return ThreeKindFlag | uint(Top1Of16[p3]) | uint(Top2Of8[p1^Bit1[p3]])
	}

	return FourKindFlag | uint(Top1Of16[p4]) | uint(Top1Of12[p1^p4])
}

// InitFast - create indexes
func InitFast() {
	var i, c1, c2, c3, c4, c5, c6, c7 uint

	for c5 = 14; c5 > 4; c5-- {
		c4 = c5 - 1
		c3 = c4 - 1
		c2 = c3 - 1
		c1 = c2 - 1
		if c1 == 1 {
			c1 = 14
		}

		for c6 = 14; c6 > 1; c6-- {
			if c6 != c5+1 {
				for c7 = c6 - 1; c7 > 1; c7-- {
					if c7 != c5+1 {
						i = (1 << c1) | (1 << c2) | (1 << c3) | (1 << c4) | (1 << c5) | (1 << c6) | (1 << c7)
						Flush[i>>2] = StraightFlushFlag | (c1 << 16) | (c2 << 12) | (c3 << 8) | (c4 << 4) | c5
						Straight[i>>2] = StraightFlag | (c1 << 16) | (c2 << 12) | (c3 << 8) | (c4 << 4) | c5
					}
				}
			}
		}
	}

	for c1 = 14; c1 > 5; c1-- {
		for c2 = c1 - 1; c2 > 4; c2-- {
			for c3 = c2 - 1; c3 > 3; c3-- {
				for c4 = c3 - 1; c4 > 2; c4-- {
					for c5 = c4 - 1; c5 > 1; c5-- {
						for c6 = c5; c6 > 1; c6-- {
							for c7 = c6; c7 > 1; c7-- {
								i = (1 << c1) | (1 << c2) | (1 << c3) | (1 << c4) | (1 << c5) | (1 << c6) | (1 << c7)
								if Flush[i>>2] == 0 {
									Flush[i>>2] = FlushFlag | (c1 << 16) | (c2 << 12) | (c3 << 8) | (c4 << 4) | c5
								}
								Top5[i>>2] = HighCardFlag | (c1 << 16) | (c2 << 12) | (c3 << 8) | (c4 << 4) | c5
							}
						}
					}
				}
			}
		}
	}

	for c1 = 14; c1 > 3; c1-- {
		for c2 = c1 - 1; c2 > 2; c2-- {
			for c3 = c2 - 1; c3 > 1; c3-- {
				for c4 = c3; c4 > 1; c4-- {
					for c5 = c4; c5 > 1; c5-- {
						for c6 = c5; c6 > 1; c6-- {
							for c7 = c6; c7 > 1; c7-- {
								i = (1 << c1) | (1 << c2) | (1 << c3) | (1 << c4) | (1 << c5) | (1 << c6) | (1 << c7)
								Top3Of4[i>>2] = (c1 << 12) | (c2 << 8) | (c3 << 4)
							}
						}
					}
				}
			}
		}
	}

	for c1 = 14; c1 > 2; c1-- {
		for c2 = c1 - 1; c2 > 1; c2-- {
			for c3 = c2; c3 > 1; c3-- {
				for c4 = c3; c4 > 1; c4-- {
					for c5 = c4; c5 > 1; c5-- {
						for c6 = c5; c6 > 1; c6-- {
							for c7 = c6; c7 > 1; c7-- {
								i = (1 << c1) | (1 << c2) | (1 << c3) | (1 << c4) | (1 << c5) | (1 << c6) | (1 << c7)
								Top2Of12[i>>2] = (c1 << 16) | (c2 << 12)
								Top2Of8[i>>2] = (c1 << 12) | (c2 << 8)
								Bit2[i>>2] = (1 << (c1 - 2)) | (1 << (c2 - 2))
							}
						}
					}
				}
			}
		}
	}

	for c1 = 14; c1 > 1; c1-- {
		for c2 = c1; c2 > 1; c2-- {
			for c3 = c2; c3 > 1; c3-- {
				for c4 = c3; c4 > 1; c4-- {
					for c5 = c4; c5 > 1; c5-- {
						for c6 = c5; c6 > 1; c6-- {
							for c7 = c6; c7 > 1; c7-- {
								i = (1 << c1) | (1 << c2) | (1 << c3) | (1 << c4) | (1 << c5) | (1 << c6) | (1 << c7)
								Top1Of16[i>>2] = (c1 << 16)
								Top1Of12[i>>2] = (c1 << 12)
								Top1Of8[i>>2] = (c1 << 8)
								Bit1[i>>2] = (1 << (c1 - 2))
							}
						}
					}
				}
			}
		}
	}
}
