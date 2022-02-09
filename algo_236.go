package algo_practice

// MARK
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 先遍历存储所有节点的父节点
	parentNodeKv := map[*TreeNode]*TreeNode{}
	parentNodeKv[root] = nil
	saveChild(root, parentNodeKv)

	// 从 p 往上遍历，记录访问的节点
	pParentNodeKv := map[*TreeNode]struct{}{}
	pCurrent := p
	for pCurrent != nil {
		// 记录访问的节点
		pParentNodeKv[pCurrent] = struct{}{}

		pCurrentParent := parentNodeKv[pCurrent]
		pCurrent = pCurrentParent
	}

	// q  往上遍历，如果在 p 中也访问过，则输出结果
	qCurrent := q
	for qCurrent != nil {
		if _, ok := pParentNodeKv[qCurrent]; ok {
			return qCurrent
		}
		qCurrentParent := parentNodeKv[qCurrent]
		qCurrent = qCurrentParent
	}
	return nil
}

func saveChild(parent *TreeNode, parentNodeKv map[*TreeNode]*TreeNode)  {
	if parent == nil {
		return
	}
	if parent.Left != nil {
		parentNodeKv[parent.Left] = parent
		saveChild(parent.Left, parentNodeKv)
	}
	if parent.Right != nil {
		parentNodeKv[parent.Right] = parent
		saveChild(parent.Right, parentNodeKv)
	}
}