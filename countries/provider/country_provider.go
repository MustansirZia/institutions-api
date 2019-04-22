package provider

import (
	"encoding/json"
	"io/ioutil"
)

type CountryProvider interface {
	Provide() ([]string, error)
}

type jsonCountryProvider struct{}

func NewJSONCountryProvider() CountryProvider {
	return &jsonCountryProvider{}
}

func (p *jsonCountryProvider) Provide() ([]string, error) {
	bytes, err := ioutil.ReadFile("./data/json/countries_states_cities.json")
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	countries := make([]string, 0, len(data))
	for _, c := range data {
		countries = append(countries, c["name"].(string))
	}
	return countries, nil
}
