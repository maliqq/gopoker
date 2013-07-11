package poker

import (
	"sort"
)

type GroupedCards []Cards

func (g *GroupedCards) ArrangeByFirst(ord Ordering) *GroupedCards {
	groups := *g

	sort.Sort(ByFirst{groups, ord})

	return &groups
}

func (g *GroupedCards) ArrangeByMax(ord Ordering) *GroupedCards {
	groups := *g

	sort.Sort(ByMax{groups, ord})

	return &groups
}

func (groups *GroupedCards) Count() *map[int]GroupedCards {
	count := map[int]GroupedCards{}

	for _, group := range *groups {
		length := len(group)
		if _, present := count[length]; !present {
			count[length] = GroupedCards{}
		}

		count[length] = append(count[length], group)
	}

	return &count
}
