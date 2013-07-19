package poker

import (
	"sort"
	"strings"
)

type GroupedCards []Cards

func (g GroupedCards) String() string {
	s := "["

	strs := make([]string, len(g))
	for i, group := range g {
		strs[i] = "{" + group.String() + "}"
	}
	s += strings.Join(strs, ", ")

	return s + "]"
}

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
