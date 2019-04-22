package states

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/qazimusab/musalleen-apis/states/provider"
	"github.com/qazimusab/musalleen-apis/trie"
)

type StateRepository interface {
	GetStates(country string) ([]string, error)
}

type stateRepository struct {
	provider provider.StateProvider
	trie     trie.Trie
	once     sync.Once
}

func NewStateRepository() StateRepository {
	return &stateRepository{
		provider: provider.NewJSONStateProvider(),
		trie:     trie.NewTrie(),
		once:     sync.Once{},
	}
}

func (r *stateRepository) load() error {
	states, err := r.provider.Provide()
	if err != nil {
		return err
	}
	r.addStatesToTrie(states)
	return nil
}

func (r *stateRepository) addStatesToTrie(states []provider.State) {
	for _, state := range states {
		r.trie.AddValue(fmt.Sprintf("%s, %s", state.Country, state.Name), state)
	}
}

func (r *stateRepository) GetStates(country string) ([]string, error) {

	var err error
	r.once.Do(func() {
		err = r.load()
	})
	if err != nil {
		return nil, err
	}

	results := r.trie.PrefixSearch(fmt.Sprintf("%s,", country), -1)
	states := make([]string, 0, len(results))
	for _, result := range results {
		states = append(states, result.(provider.State).Name)
	}
	sort.SliceStable(states, func(i, j int) bool {
		return strings.Compare(states[i], states[j]) < 0
	})

	return states, nil
}
