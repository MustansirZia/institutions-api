package cities

import (
	"sort"
	"strings"
	"sync"

	"github.com/qazimusab/musalleen-apis/cities/provider"
	"github.com/qazimusab/musalleen-apis/set"
	"github.com/qazimusab/musalleen-apis/trie"
)

type CityRepository interface {
	GetCities(name, state string) ([]provider.City, error)
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
		r.trie.AddValue(city.String(), city)
	}
}

func (r *cityRepository) GetCities(name, state string) ([]provider.City, error) {

	var err error
	r.once.Do(func() {
		err = r.load()
	})
	if err != nil {
		return nil, err
	}

	citiesSet := set.NewSet()

	resultsForState := r.trie.PrefixSearch(state, -1)
	for _, result := range resultsForState {
		citiesSet.Add(result.(provider.City))
	}

	resultsForName := r.trie.PrefixSearch(name, -1)
	for _, result := range resultsForName {
		citiesSet.Add(result.(provider.City))
	}

	citiesGeneric := citiesSet.Values()
	cities := make([]provider.City, 0, len(citiesGeneric))

	for _, city := range citiesGeneric {
		cities = append(cities, city.(provider.City))
	}

	sort.Slice(cities, func(i, j int) bool {
		return strings.Compare(cities[i].Name, cities[j].Name) > 1
	})

	return cities, nil
}
