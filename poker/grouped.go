package poker

import (
	"sort"
	"strings"
)

// GroupedCards - list of list of cards
type GroupedCards []Cards

// String - grouped cards to string
func (g GroupedCards) String() string {
	s := "["

	strs := make([]string, len(g))
	for i, group := range g {
		strs[i] = "{" + group.String() + "}"
	}
	s += strings.Join(strs, ", ")

	return s + "]"
}

// ArrangeByFirst - take first card of each group and reorder
func (g *GroupedCards) ArrangeByFirst(ord Ordering) GroupedCards {
	// copy
	groups := *g

	sort.Sort(ByFirst{groups, ord})

	return groups
}

// ArrangeByMax - take largest card of each group and reorder
func (g *GroupedCards) ArrangeByMax(ord Ordering) GroupedCards {
	// copy
	groups := *g

	sort.Sort(ByMax{groups, ord})

	return groups
}

// Count - map with lengths of each group
func (g *GroupedCards) Count() map[int]GroupedCards {
	count := map[int]GroupedCards{}

	for _, group := range *g {
		length := len(group)
		if _, present := count[length]; !present {
			count[length] = GroupedCards{}
		}

		count[length] = append(count[length], group)
	}

	return count
}
