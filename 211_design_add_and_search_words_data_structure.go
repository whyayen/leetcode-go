package main

type WordDictionary struct {
	charMap   map[rune]*WordDictionary
	endOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		charMap:   make(map[rune]*WordDictionary),
		endOfWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	current := this

	for _, r := range word {
		nodeInMap, isExist := current.charMap[r]

		if isExist {
			current = nodeInMap
			continue
		}

		newNode := &WordDictionary{
			charMap: make(map[rune]*WordDictionary),
		}
		current.charMap[r] = newNode
		current = newNode
	}
	current.endOfWord = true
}

// 可以用 DFS 下去找，比較簡單
func (this *WordDictionary) Search(word string) bool {
	current := this

	for i, r := range word {
		if _, found := current.charMap[r]; !found {
			if r != '.' {
				return false
			}

			for _, node := range current.charMap {
				if node.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}

		current = current.charMap[r]
	}

	return current.endOfWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
