package countries

import (
	"sort"
	"strings"
	"sync"

	"github.com/qazimusab/musalleen-apis/countries/provider"
	"github.com/qazimusab/musalleen-apis/trie"
)

type CountryRepository interface {
	GetAllCountries() ([]string, error)
}

type countryRepository struct {
	provider provider.CountryProvider
	trie     trie.Trie
	once     sync.Once
}

func NewCountryRepository() CountryRepository {
	return &countryRepository{
		provider: provider.NewJSONCountryProvider(),
		trie:     trie.NewTrie(),
		once:     sync.Once{},
	}
}

func (r *countryRepository) load() error {
	countries, err := r.provider.Provide()
	if err != nil {
		return err
	}
	r.addContriesToTrie(countries)
	return nil
}

func (r *countryRepository) addContriesToTrie(countries []string) {
	for _, country := range countries {
		r.trie.AddValue(country, country)
	}
}

func (r *countryRepository) GetAllCountries() ([]string, error) {
	var err error
	r.once.Do(func() {
		err = r.load()
	})
	if err != nil {
		return nil, err
	}

	results := r.trie.GetAllValues()
	countries := make([]string, 0, len(results))
	for _, result := range results {
		countries = append(countries, result.(string))
	}

	sort.SliceStable(countries, func(i, j int) bool {
		return strings.Compare(countries[i], countries[j]) < 1
	})

	return countries, nil
}
