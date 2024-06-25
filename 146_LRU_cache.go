package main

type MyNode struct {
	Key  int
	Val  int
	Prev *MyNode
	Next *MyNode
}

type LRUCache struct {
	Head     *MyNode
	Tail     *MyNode
	Hash     map[int]*MyNode
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Capacity: capacity, Hash: map[int]*MyNode{}}
}

func (this *LRUCache) Get(key int) int {
	keyNode, keyExist := this.Hash[key]

	if !keyExist {
		return -1
	}

	moveToHead(this, keyNode)
	return keyNode.Val
}

func moveToHead(cache *LRUCache, nodeShouldMove *MyNode) {
	originalHead := cache.Head
	if originalHead == nil || nodeShouldMove == originalHead {
		return
	}

	orgPrevOfNodeShouldMove := nodeShouldMove.Prev
	orgNextOfNodeShouldMove := nodeShouldMove.Next
	if orgPrevOfNodeShouldMove != nil {
		orgPrevOfNodeShouldMove.Next = orgNextOfNodeShouldMove
	}
	if orgNextOfNodeShouldMove != nil {
		orgNextOfNodeShouldMove.Prev = orgPrevOfNodeShouldMove
	}
	nodeShouldMove.Prev = nil
	nodeShouldMove.Next = originalHead
	originalHead.Prev = nodeShouldMove
	cache.Head = nodeShouldMove

	if cache.Tail == nodeShouldMove {
		cache.Tail = orgPrevOfNodeShouldMove
	}
}

func (this *LRUCache) Put(key int, value int) {
	keyNode, keyExist := this.Hash[key]

	if keyExist {
		keyNode.Val = value
		moveToHead(this, keyNode)
		return
	}

	newNode := &MyNode{Key: key, Val: value}
	this.Hash[key] = newNode
	originalTail := this.Tail

	if originalTail == nil {
		this.Head = newNode
		this.Tail = newNode
		return
	}

	moveToHead(this, newNode)

	if len(this.Hash) > this.Capacity {
		delete(this.Hash, originalTail.Key)

		if originalTail.Prev != nil {
			originalTail.Prev.Next = nil
		}

		this.Tail = originalTail.Prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
