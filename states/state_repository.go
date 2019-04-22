package states

import (
	"sort"
	"strings"
	"sync"

	"github.com/qazimusab/musalleen-apis/states/provider"
	"github.com/qazimusab/musalleen-apis/trie"
)

type StateRepository interface {
	GetStates(country string) ([]provider.State, error)
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
		r.trie.AddValue(state.String(), state)
	}
}

func (r *stateRepository) GetStates(country string) ([]provider.State, error) {

	var err error
	r.once.Do(func() {
		err = r.load()
	})
	if err != nil {
		return nil, err
	}

	results := r.trie.PrefixSearch(country, -1)
	states := make([]provider.State, 0, len(results))
	for _, result := range results {
		states = append(states, result.(provider.State))
	}
	sort.Slice(states, func(i, j int) bool {
		return strings.Compare(states[i].Name, states[j].Name) > 1
	})

	return states, nil
}
