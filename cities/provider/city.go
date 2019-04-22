package provider

import (
	"fmt"
)

type City struct {
	Country string
	State   string
	Name    string
}

func (s City) String() string {
	return fmt.Sprintf("%s, %s, %s", s.Name, s.State, s.Country)
}
