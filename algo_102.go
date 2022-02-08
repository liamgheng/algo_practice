package algo_practice

// 9:52 10:05
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		levelData := make([]int, 0)
		tmpLevel := make([]*TreeNode, 0)
		for _, node := range nextLevel {
			levelData = append(levelData, node.Val)
			if node.Left != nil {
				tmpLevel = append(tmpLevel, node.Left)
			}
			if node.Right != nil {
				tmpLevel = append(tmpLevel, node.Right)
			}
		}
		result = append(result, levelData)
		nextLevel = tmpLevel
	}
	return result
}
