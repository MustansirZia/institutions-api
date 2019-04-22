package provider

import (
	"encoding/json"
	"io/ioutil"
)

type CityProvider interface {
	Provide() ([]City, error)
}

type jsonCityProvider struct {
}

func NewJSONCityProvider() CityProvider {
	return &jsonCityProvider{}
}

func (p *jsonCityProvider) Provide() ([]City, error) {
	bytes, err := ioutil.ReadFile("./data/json/countries_states_cities.json")
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	citiesSlice := make([]City, 0)
	for _, country := range data {
		statesMap, ok := country["states"].(map[string]interface{})
		if ok {
			for state, citiesOfAState := range statesMap {
				for _, city := range citiesOfAState.([]string) {
					citiesSlice = append(citiesSlice, City{
						Country: country["name"].(string),
						State:   state,
						Name:    city,
					})
				}
			}
		}
	}
	return citiesSlice, nil
}
