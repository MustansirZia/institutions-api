package provider

import (
	"encoding/json"
	"io/ioutil"
)

type StateProvider interface {
	Provide() ([]State, error)
}

type jsonStateProvider struct {
}

func NewJSONStateProvider() StateProvider {
	return &jsonStateProvider{}
}

func (p *jsonStateProvider) Provide() ([]State, error) {
	bytes, err := ioutil.ReadFile("./data/json/countries_states_cities.json")
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	statesSlice := make([]State, 0)
	for _, country := range data {
		statesMap, ok := country["states"].(map[string]interface{})
		if ok {
			for state := range statesMap {
				statesSlice = append(statesSlice, State{
					Country: country["name"].(string),
					Name:    state,
				})
			}
		}
	}
	return statesSlice, nil
}
