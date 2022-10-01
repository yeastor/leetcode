package main

type LRUCache struct {
	Values   map[int]int
	capacity int
	Weight   *WeightNode
	LastEl   *WeightNode
}

type WeightNode struct {
	Val  int
	Next *WeightNode
}

func findKey(root *WeightNode, prevNode *WeightNode, findVal int) (*WeightNode, *WeightNode) {
	if root.Val == findVal {
		return prevNode, root
	}
	return findKey(root.Next, root, findVal)
}

func FindLast(root *WeightNode) *WeightNode {
	if root.Next == nil {
		return root
	}
	return FindLast(root.Next)
}

func (this *LRUCache) MoveListNodeToEnd(key int, latNode *WeightNode) {
	prevNode, currentNode := findKey(this.Weight, nil, key)
	nextNode := currentNode.Next
	if nextNode == nil {
		// last elemetn return
		return
	}
	if prevNode != nil {
		prevNode.Next = nextNode
	} else {
		this.Weight = nextNode
	}

	currentNode.Next = nil
	this.LastEl.Next = currentNode
	this.LastEl = currentNode

	return
}

func RemoveFirstElement(root *WeightNode) *WeightNode {
	return root.Next
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Values:   make(map[int]int, capacity),
		capacity: capacity,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.Values[key]; ok {
		this.MoveListNodeToEnd(key, this.LastEl)
		return val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {

	weightList := &WeightNode{
		Val:  key,
		Next: nil,
	}

	isElemExist := this.Get(key)

	if len(this.Values) == this.capacity {
		if isElemExist == -1 {
			keyForRemove := this.Weight.Val
			delete(this.Values, keyForRemove)
			this.Weight = RemoveFirstElement(this.Weight)
		} else {
			this.Values[key] = value
			return
		}
		this.Values[key] = value
	} else {
		this.Values[key] = value
	}

	if this.Weight == nil {
		this.Weight = weightList
		this.LastEl = weightList
		return
	}

	if isElemExist == -1 {
		this.LastEl.Next = weightList
		this.LastEl = weightList
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

}
