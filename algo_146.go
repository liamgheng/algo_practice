package algo_practice

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	Capacity int
	Len int
	Kv map[int]*DLinkedNode
	DLinkedHeadNode *DLinkedNode  // 双向链表头部哨兵节点
	// NOTE 这里一开始想用循环链表做，想麻烦了，单独记录尾部节点哨兵即可
	DLinkedTrailNode *DLinkedNode // 双向链表尾部哨兵节点
}


func Constructor(capacity int) LRUCache {
	headNode := &DLinkedNode{0, 0, nil, nil}
	tailNode := &DLinkedNode{0, 0, nil, nil}
	headNode.next = tailNode
	tailNode.prev = headNode
	return LRUCache{
		Capacity: capacity,
		Len: 0,
		Kv: make(map[int]*DLinkedNode),
		DLinkedHeadNode: headNode,
		DLinkedTrailNode: tailNode,
	}
}


func (this *LRUCache) Get(key int) int {
	if node, ok := this.Kv[key]; ok {
		this.topNode(node)
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.Kv[key]; ok {
		node.value = value
		this.topNode(node)
	} else {
		node = &DLinkedNode{key, value, nil, nil}
		this.Kv[key] = node
		// 头部插入一个节点
		this.insertHead(node)
		if this.Len > this.Capacity {
			this.removeLast()
		}
	}
}

func (this *LRUCache) topNode(node *DLinkedNode) {
	// 删除节点
	preNode := node.prev
	nextNode := node.next
	preNode.next = nextNode
	nextNode.prev = preNode

	preHeadNode := this.DLinkedHeadNode.next

	// 在头部双边插入节点
	this.DLinkedHeadNode.next = node
	node.prev = this.DLinkedHeadNode

	node.next = preHeadNode
	preHeadNode.prev = node
}

func (this *LRUCache) insertHead(newHeadNode *DLinkedNode) {
	this.Len += 1

	preHeadNode := this.DLinkedHeadNode.next

	this.DLinkedHeadNode.next = newHeadNode
	newHeadNode.prev = this.DLinkedHeadNode

	newHeadNode.next = preHeadNode
	preHeadNode.prev = newHeadNode
}

func (this *LRUCache) removeLast() {
	willDeleteNode := this.DLinkedTrailNode.prev
	willDeletePreNode := willDeleteNode.prev

	willDeletePreNode.next = this.DLinkedTrailNode
	this.DLinkedTrailNode.prev = willDeletePreNode

	delete(this.Kv, willDeleteNode.key)
}