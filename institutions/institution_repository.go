package institutions

import (
	"strings"

	"github.com/derekparker/trie"
	"github.com/qazimusab/musalleen-apis/institutions/providers"
	"golang.org/x/sync/errgroup"
)

type InstitutionRepository interface {
	Load() error
	GetInstitutions(name string, count int) []string
}

type institutionRepository struct {
	providers []providers.InstitutionProvider
	trie      *trie.Trie
}

func NewInstitutionRepository(providers ...providers.InstitutionProvider) InstitutionRepository {
	return &institutionRepository{
		providers: providers,
		trie:      trie.New(),
	}
}

func (r *institutionRepository) Load() error {
	routineGroup := errgroup.Group{}
	for _, provider := range r.providers {
		// Extracting institutions from each provider concurrently.
		currentProvider := provider
		routineGroup.Go(func() error {
			institutions, err := currentProvider.Provide()
			if err != nil {
				return err
			}
			r.addInstitutionsToTrie(institutions)
			return nil
		})
	}
	// Waiting for the extraction to finish.
	return routineGroup.Wait()
}

func (r *institutionRepository) addInstitutionsToTrie(institutions []string) {
	for _, institution := range institutions {
		r.trie.Add(strings.ToLower(institution), institution)
		split := strings.Split(institution, " ")
		if len(split) > 1 {
			for _, split := range split[1:] {
				r.trie.Add(strings.ToLower(split), institution)
			}
		}
	}
}

func (r *institutionRepository) GetInstitutions(name string, count int) []string {
	institutions := make([]string, 0)
	nodes := r.trie.PrefixSearch(strings.ToLower(name))
	for i, k := range nodes {
		node, found := r.trie.Find(k)
		if found {
			institutions = append(institutions, node.Meta().(string))
		}
		if i+1 == count {
			break
		}
	}
	return institutions
}
