package main

import "fmt"

type Trie struct {
	Children   []*Trie
	Val        string
	Terminator bool
}

func Constructor() Trie {
	return Trie{
		Val: "",
	}
}

func (this *Trie) Insert(word string) {
	curr := this

	for i := 1; i <= len(word); i++ {
		partial := word[0:i]
		inserted := false

		for _, child := range curr.Children {
			if child.Val == partial {
				curr = child
				inserted = true
				break
			}
		}

		if !inserted {
			tmp := &Trie{
				Val:        partial,
				Terminator: partial == word,
			}
			curr.Children = append(curr.Children, tmp)
			curr = tmp
		}
	}

	if curr.Val == word && !curr.Terminator {
		curr.Terminator = true
	}
}

func (this *Trie) Search(word string) bool {
	return this.InnerSearch(word, true)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.InnerSearch(prefix, false)
}

func (this *Trie) InnerSearch(word string, shouldBeTerminator bool) bool {
	curr := this

	for i := 1; i <= len(word); i++ {
		partial := word[0:i]

		for _, child := range curr.Children {
			if child.Val == partial {
				curr = child
			}
		}

		if curr.Val != partial {
			return false
		}
	}

	if !shouldBeTerminator {
		return true
	} else {
		return curr.Terminator
	}
}

func main() {
	obj := Constructor()
	obj.Insert("apple")

	fmt.Printf("search app: \x1b[31m%+v\x1b[0m (should be false)\n", obj.Search("app"))
	fmt.Printf("search apple: \x1b[32m%+v\x1b[0m (should be true)\n", obj.Search("apple"))
	fmt.Printf("starts with app: \x1b[32m%+v\x1b[0m (should be true)\n", obj.StartsWith("app"))

	obj.Insert("app")
	fmt.Printf("search app: \x1b[32m%+v\x1b[0m (should be true)\n", obj.StartsWith("app"))
}
