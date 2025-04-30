package main

/* Trie
 * 一種 Prefix Tree 的資料結構，也叫作 Digital Tree，透過樹把每個字元串起來
 * 大致會有這三個欄位，分別儲存對應的資料
 *   char: 儲存此節點的字元，root 為 nil
 *   isKey: 是否為一個詞的最後一個字元，例如 app 則最後一個 p 節點的 isKey 為 true，但 a 節點和接在 a 後面的 p 節點其 isKey 都是 false
 *   charMap: 用來儲存此節點下面的其他字元節點，例如一個 trie 有儲存兩個詞 tree & try，則 r 節點後面會有一個這樣的 map
 *   	{
 *          e: &TrieNode{ char: e, isKey: false }
 *          y: &TrieNode{ char: y, isKey: true  }
 * 		}
 * 基本支援 3 個方法
 *    Insert: 插入新的詞進入 Trie
 *    Search: 找尋 Trie 是否有包含這個詞
 *    StartsWith: 找尋是否有這個 prefix 的詞
 * Insert 方法：
 *   1. Iterate 每個字元，同時從 root 開始往下找，如果該 char 能在 charMap 找到，則跳到 charMap 該 char 的 node，並跳到下一個字元
 *   2. 如果目前 iterate 的該字元不存在當前節點 charMap，則建立一個 TrieNode 並存放到 charMap 對應字元的位置，並跳到下一個字元
 * Search 方法：
 *   1. 由於 search 是要找到該詞是否存在在 Trie 內，而不是 Prefix
 *   2. 所以 Iterate 每個字元從 root 開始往下找時，如果有任何一個字元在 charMap 找不到時，就直接 return false
 *   3. Iterate 到最後一個字元時，如果該 TrieNode 的 isKey 不是 true，則表示它存在在 Trie 內部並不是以 key 存在，而是一個 prefix，因此回傳 false
 * StartsWith 方法：
 *   1. 基本邏輯同 Search，但不用檢查最後一個字元 isKey 是否為 true，只要節點都存在，則表示有 StartsWith 這段 prefix 的詞
 */

type Trie struct {
	char    rune
	isKey   bool
	charMap map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		charMap: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	current := this

	for _, r := range word {
		nodeInMap, isExist := current.charMap[r]

		if isExist {
			current = nodeInMap
			continue
		}

		newNode := &Trie{
			char:    r,
			charMap: make(map[rune]*Trie),
		}
		current.charMap[r] = newNode
		current = newNode
	}

	current.isKey = true
}

func (this *Trie) Search(word string) bool {
	current := this

	for _, r := range word {
		nodeInMap, isExist := current.charMap[r]

		if !isExist {
			return false
		}

		current = nodeInMap
	}

	return current.isKey
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this

	for _, r := range prefix {
		nodeInMap, isExist := current.charMap[r]

		if !isExist {
			return false
		}

		current = nodeInMap
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
