package hub

import (
	"fmt"
)

type Route struct {
	None   bool
	All    bool
	One    string
	Only   []string
	Except []string
	Where  func(Endpoint) bool
}

func (r Route) String() string {
	if r.None {
		return "none"
	}
	if r.All {
		return "all"
	}
	if r.One != "" {
		return fmt.Sprintf("one [%s]", r.One)
	}
	if len(r.Except) != 0 {
		return fmt.Sprintf("except [%s]", r.Except)
	}
	if len(r.Only) != 0 {
		return fmt.Sprintf("only [%s]", r.Only)
	}
	return ""
}
