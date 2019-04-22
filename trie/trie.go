package trie

import (
	"strings"

	originalTrie "github.com/derekparker/trie"
	"github.com/qazimusab/musalleen-apis/set"
)

type Trie interface {
	PrefixSearch(pre string, count int) []interface{}
	AddValue(key string, value interface{})
	GetAllValues() []interface{}
}

type trie struct {
	originalTrie *originalTrie.Trie
}

func NewTrie() Trie {
	return &trie{
		originalTrie: originalTrie.New(),
	}
}

func (t *trie) PrefixSearch(pre string, count int) []interface{} {
	results := make([]interface{}, 0)
	nodes := t.originalTrie.PrefixSearch(strings.ToLower(pre))
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
	node, found := t.originalTrie.Find(strings.ToLower(key))
	if !found {
		return nil, false
	}
	return node.Meta(), true
}

func (t *trie) AddValue(key string, value interface{}) {
	t.originalTrie.Add(strings.ToLower(key), value)
	wordsInKey := strings.Split(key, " ")
	if len(wordsInKey) > 1 {
		for _, word := range wordsInKey[1:] {
			t.originalTrie.Add(strings.ToLower(word), value)
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
