package institutions

import (
	"github.com/qazimusab/musalleen-apis/institutions/providers"
	"github.com/qazimusab/musalleen-apis/trie"
	"golang.org/x/sync/errgroup"
)

type InstitutionRepository interface {
	Load() error
	GetInstitutions(name string, count int) []string
}

type institutionRepository struct {
	providers []providers.InstitutionProvider
	trie      trie.Trie
}

func NewInstitutionRepository(providers ...providers.InstitutionProvider) InstitutionRepository {
	return &institutionRepository{
		providers: providers,
		trie:      trie.NewTrie(),
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
		r.trie.AddValue(institution, institution)
	}
}

func (r *institutionRepository) GetInstitutions(name string, count int) []string {
	results := r.trie.PrefixSearch(name, count)
	institutions := make([]string, 0, len(results))
	for _, result := range results {
		institutions = append(institutions, result.(string))
	}
	return institutions
}
