package patricia_tree

const (
	DefaultMaxPrefixPerNode = 10
	DefaultMaxChildrenPerSparseNode = 8
)

type (
	Prefix []byte
	Item interface{}
	VisitorFunc func(prefix Prefix, item Item) error
)

type Trie struct {
	prefix Prefix
	item Item
	maxPrefixPerNode int
	maxChildrenPerSparseNode int
	children childList
}

func NewTrie (options ...Option) *Trie {
	trie := &Trie{}
	for _, opt := range options {
		opt(trie)
	}
	if trie.maxPrefixPerNode <= 0 {
		trie.maxChildrenPerSparseNode = DefaultMaxChildrenPerSparseNode
	}
	trie.children = newSpareChildList(trie.maxChildrenPerSparseNode)
	return trie
}

func MaxPrefixPerNode(value int) Option {
	return func(trie *Trie) {
		trie.maxPrefixPerNode = value
	}
}

func MaxChildrenPerSpareNode(value int) Option {
	return func(trie *Trie) {
		trie.maxChildrenPerSparseNode = value
	}
}

func (trie *Trie) Item() Item {
	return trie.item
}


func (trie *Trie) Insert(key Prefix, item Item) (inserted bool) {
	return trie.put(key, item, false)
}


