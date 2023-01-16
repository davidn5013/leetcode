package ds

// implementation of trie
// Leetcode 208
// Grind75 number 35

// Trie node store word for easy find
type Trie struct {
	words []string
}

// NewTrie pointer to new trie struct
func NewTrie() Trie {
	return Trie{
		words: []string{},
	}
}

// normal idiomatic go is (t *Trie) but this for compliance to leetcode specs

// Insert new order
func (t *Trie) Insert(word string) {
	t.words = append(t.words, word)
}

// Search for word in trie
func (t *Trie) Search(word string) bool {
	for _, data := range t.words {
		if data == word {
			return true
		}
	}

	return false
}

// StartsWith for word starting with in trie
func (t *Trie) StartsWith(prefix string) bool {
	lenPrefix := len(prefix)

	for _, word := range t.words {
		if len(word) < lenPrefix {
			continue
		}

		if prefix == word[:lenPrefix] {
			return true
		}
	}

	return false
}
