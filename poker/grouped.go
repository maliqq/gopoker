package poker

import (
	"sort"
)

type GroupedCards []Cards

func (g *GroupedCards) ArrangeByFirst(ord Ordering) *GroupedCards {
	// copy
	groups := *g

	sort.Sort(ByFirst{groups, ord})

	return &groups
}

func (g *GroupedCards) ArrangeByMax(ord Ordering) *GroupedCards {
	// copy
	groups := *g

	sort.Sort(ByMax{groups, ord})

	return &groups
}

func (g *GroupedCards) Count() *map[int]GroupedCards {
	count := map[int]GroupedCards{}

	for _, group := range *g {
		length := len(group)
		if _, present := count[length]; !present {
			count[length] = GroupedCards{}
		}

		count[length] = append(count[length], group)
	}

	return &count
}
