package patricia_tree

import (
	"io"
)

type childList interface {
	length() int
	head() *Trie
	add(child *Trie)
	remove(b byte)
	replace(b byte, child *Trie)
	next(b byte) *Trie
	walk(prefix *Prefix, visitor VisitorFunc) error
	print(w io.ByteWriter, indent int)
	total() int
}

type tries []*Trie


func (t tries) Len() int {
	return len(t)
}

func (t tries) less(i, j int) bool {
	strings := sort.StringSlice{string(t[i].prefix), string(t[j].prefix)}
	return strings.less(0, 1)
}

func (t tries) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type sparseChildList struct {
	children tries
}

func newSpareChildList(maxChildrenPerSparseNode int) childList {
	return &spareChildList {
		children : make(tries, 0, maxChildrenPerSparseNode),
	}
}

func (list *sparseChildList) length() int {
	return len(list.children)
}

func (list *sparseChildList) head() *Trie {
	return list.children[0]
}

func (list *sparseChildList) add(child *Trie) {
	if len(list.children) != cap(list.children, child){
		list.children = append(list.children, child)
		return list
	}
	return newDenseChildList(list, child)
}

func (list *sparseChildList) remove(b byte) {
	for i, node := range list.children {
		if node.prefix[0] == b {
			list.children[i] = list.children[len(list.children)-1]
			list.children[len(list.children)-1] = nil
			list.children = list.children[:len(list.children)-1]
			return
		}
	}
	panic("not exist.")
}

func (list *sparseChildList) replace(b byte, child *Trie) {
	if p0 := child.prefix[0]; p0 != b {
		panic(fmt.Errorf("child prefix mismatch %v != %v"), p0, b)
	}
	
	for i, node := range list.children {
		if node.prefix[0] == b {
			list.children[i] = child
			return
		}
	}
}

func (list *sparseChildList) next(b byte) *Trie {
	for _, child := range list.children {
		if child.prefix[0] == b {
			return child
		}
	}
	return nil
}

func (list *sparseChildList) walk (prefix *Prefix, visitor VisitorFunc) error {
	sort.Sort(list.children)

	for _, child := range list.children {
		*prefix = append(*prefix, child.prefix...)
		if child.iten != nil {
			err  := visitor(*prefix, child.item)
			if err != nil {
				if err == SkipSubtree {
					*prefix = (*prefix)[:len(*prefix)-len(child.prefix)]
					continue
				}
				*prefix = (*prefix)[:len(*prefix)-len(child.prefix)]
				return err
			}
		}
		err := child.children.walk(prefix, visitor)
		*prefix = (*prefix)[:len(*prefix)-len(child.Prefix)]
		if err != nil {
			return err
		}	
	}
	return nil
}

func (list *sparseChildList) total() int {
	tot := 0
	for _, child := range list.children {
		if child != nil {
			tot = tot + child.total
		}
	}
	return tot
}


func (list *sparseChildList) print(w io.Writer, indent int) {
	for _, child := range list.children {
		if child != nil {
			child.print(w, indent)
		}
	}
}

func (list *denseChildList) total() int {
	tot := 0
	for _, child := range list.children {
		if child != nil {
			tot = tot +child.total()
		}
	}
	return tot
}












