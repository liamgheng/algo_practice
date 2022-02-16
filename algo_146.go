package algo_practice

type DLinkListNode struct {
	key int
	val int
	prev *DLinkListNode
	next *DLinkListNode
}

type LRUCache struct {
	cap int
	kv map[int]*DLinkListNode
	dLinkListHead *DLinkListNode  // 哨兵头节点
	dLinkListTail *DLinkListNode
}

func Constructor(capacity int) LRUCache {
	// 通过哨兵头尾节点，来减少判断条件
	dLinkListHead := &DLinkListNode{-1, -1, nil, nil}
	dLinkListTail := &DLinkListNode{-1, -1, nil, nil}
	dLinkListHead.next = dLinkListTail
	dLinkListTail.prev = dLinkListHead
	return LRUCache{
		cap: capacity,
		kv: make(map[int]*DLinkListNode),
		dLinkListHead: dLinkListHead,
		dLinkListTail: dLinkListTail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.kv[key]; ok {
		this.removeNodeFromLinkList(node)
		this.putNodeHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.kv[key]; ok {
		node.val = value
		this.removeNodeFromLinkList(node)
		this.putNodeHead(node)
		return
	}

	// 创建新的元素
	node := &DLinkListNode{key, value, nil, nil}
	this.kv[key] = node
	this.putNodeHead(node)

	// 如果超过 cap，则删除队尾元素
	if len(this.kv) > this.cap {
		node := this.removeTailNode()
		// MARK 一开始这里删除有问题，删除的是 node.value 找不到
		// 改为在 node 中存储 key 就好了
		delete(this.kv, node.key)
	}
}

func (this *LRUCache) putNodeHead(node *DLinkListNode) {
	// 把节点放入头部
	originalHeadNext := this.dLinkListHead.next
	this.dLinkListHead.next = node
	node.prev = this.dLinkListHead

	node.next = originalHeadNext
	originalHeadNext.prev = node
}

func (this *LRUCache) removeNodeFromLinkList(node *DLinkListNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) removeTailNode() *DLinkListNode {
	willDeleteNode := this.dLinkListTail.prev
	this.removeNodeFromLinkList(willDeleteNode)
	return willDeleteNode
}
