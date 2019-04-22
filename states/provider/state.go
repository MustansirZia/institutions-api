package provider

import (
	"fmt"
)

type State struct {
	Country string
	Name    string
}

func (s State) String() string {
	return fmt.Sprintf("%s, %s", s.Name, s.Country)
}
