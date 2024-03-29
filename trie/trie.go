package trie

import (
	"math/rand"
	"strings"

	originalTrie "github.com/derekparker/trie"
	"github.com/mustansirzia/institutions-api/set"
	uuid "github.com/satori/go.uuid"
)

type Trie interface {
	PrefixSearch(pre string, count int) []interface{}
	AddValue(key string, value interface{})
	GetAllValues() []interface{}
}

type trie struct {
	originalTrie *originalTrie.Trie
	uuidPool     []string
}

func NewTrie() Trie {
	return &trie{
		originalTrie: originalTrie.New(),
		uuidPool:     generateUUIDS(20, 3),
	}
}

func (t *trie) PrefixSearch(pre string, count int) []interface{} {
	results := make([]interface{}, 0)
	nodes := t.originalTrie.PrefixSearch(normalizeKey(pre))
	for i, k := range nodes {
		value, found := t.findValue(k)
		if found {
			results = append(results, value)
		}
		if i+1 == count {
			break
		}
	}
	return results
}

func (t *trie) findValue(key string) (interface{}, bool) {
	node, found := t.originalTrie.Find(key)
	if !found {
		return nil, false
	}
	return node.Meta(), true
}

func (t *trie) AddValue(key string, value interface{}) {
	t.originalTrie.Add(normalizeKey(key), value)
	wordsInKey := strings.Split(key, " ")
	if len(wordsInKey) > 1 {
		for _, word := range wordsInKey[1:] {
			word += t.getRandomUUID()
			t.originalTrie.Add(normalizeKey(word), value)
		}
	}
}

func (t *trie) GetAllValues() []interface{} {
	set := set.NewSet()
	keys := t.originalTrie.Keys()
	for _, key := range keys {
		value, found := t.findValue(key)
		if found {
			set.Add(value)
		}
	}
	return set.Values()
}

func (t *trie) getRandomUUID() string {
	index := rand.Intn(len(t.uuidPool) - 1)
	return t.uuidPool[index]
}

func generateUUIDS(count int, length int) []string {
	uuids := make([]string, 0, count)
	for i := 0; i < count; i++ {
		uuids = append(uuids, uuid.NewV1().String()[:length])
	}
	return uuids
}

func normalizeKey(key string) string {
	return strings.ToLower(strings.Trim(key, " "))
}
