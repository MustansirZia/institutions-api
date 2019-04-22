package cities

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/qazimusab/musalleen-apis/cities/provider"
	"github.com/qazimusab/musalleen-apis/set"
	"github.com/qazimusab/musalleen-apis/trie"
)

type CityRepository interface {
	GetCitiesByName(name string) ([]string, error)
	GetCitiesByState(state string) ([]string, error)
}

type cityRepository struct {
	provider provider.CityProvider
	trie     trie.Trie
	once     sync.Once
}

func NewCityRepository() CityRepository {
	return &cityRepository{
		provider: provider.NewJSONCityProvider(),
		trie:     trie.NewTrie(),
		once:     sync.Once{},
	}
}

func (r *cityRepository) load() error {
	cities, err := r.provider.Provide()
	if err != nil {
		return err
	}
	r.addCitiesToTrie(cities)
	return nil
}

func (r *cityRepository) addCitiesToTrie(cities []provider.City) {
	for _, city := range cities {
		r.trie.AddValue(fmt.Sprintf("%s, %s", city.State, city.Name), city)
	}
}

func (r *cityRepository) GetCitiesByName(name string) ([]string, error) {
	cities, err := r.getCities(&name, nil)
	if err != nil {
		return nil, err
	}
	citiesAsStrings := make([]string, 0, len(cities))
	for _, city := range cities {
		citiesAsStrings = append(citiesAsStrings, city.String())
	}
	return citiesAsStrings, nil
}

func (r *cityRepository) GetCitiesByState(state string) ([]string, error) {
	cities, err := r.getCities(nil, &state)
	if err != nil {
		return nil, err
	}
	citiesAsStrings := make([]string, 0, len(cities))
	for _, city := range cities {
		citiesAsStrings = append(citiesAsStrings, city.Name)
	}
	return citiesAsStrings, nil
}

func (r *cityRepository) getCities(name, state *string) ([]provider.City, error) {

	var err error
	r.once.Do(func() {
		err = r.load()
	})
	if err != nil {
		return nil, err
	}

	citiesSet := set.NewSet()

	if state != nil {
		resultsForState := r.trie.PrefixSearch(*state, -1)
		for _, result := range resultsForState {
			citiesSet.Add(result.(provider.City))
		}
	}

	if name != nil {
		resultsForName := r.trie.PrefixSearch(*name, -1)
		for _, result := range resultsForName {
			citiesSet.Add(result.(provider.City))
		}
	}

	citiesGeneric := citiesSet.Values()
	cities := make([]provider.City, 0, len(citiesGeneric))

	for _, city := range citiesGeneric {
		cities = append(cities, city.(provider.City))
	}

	sort.SliceStable(cities, func(i, j int) bool {
		return strings.Compare(cities[i].String(), cities[j].String()) < 0
	})

	return cities, nil
}
