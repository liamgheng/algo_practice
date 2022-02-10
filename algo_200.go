package algo_practice

import (
	"fmt"
	"math"
)

type Dot struct {
	x int
	y int
}

// start 10:43
func numIslands(grid [][]byte) int {
	islandCounts := 0

	// 获取搜索 1 的节点
	var allIslandDots []*Dot
	for yIndex, dots := range grid {
		for xIndex, val := range dots {
			if val == 1 {
				allIslandDots = append(allIslandDots, &Dot{xIndex, yIndex})
			}
		}
	}

	for len(allIslandDots) > 0 {
		islandCounts += 1

		var islandDots []*Dot
		// 从每个 1 开始迭代，如果找到相邻的则组成岛屿
		for _, dot := range allIslandDots {
			if inIsland(dot, islandDots) {
				islandDots = append(islandDots, dot)
			} else {
				fmt.Println("warning")
			}
		}

		// 已经组成岛屿的，从 allIslandDots 中删除
		for _, willDeleteDot := range islandDots {
			index := -1
			for willDeleteIndex, dot := range allIslandDots {
				if dot == willDeleteDot {
					index = willDeleteIndex
				}
			}
			allIslandDots = append(allIslandDots[:index], allIslandDots[index+1:]...)
		}
	}

	return islandCounts
}

func inIsland(dot *Dot, islandDots []*Dot) bool {
	// 空的情况
	if len(islandDots) == 0 {
		return true
	}
	for _, islandDot := range islandDots {
		if dot.x == islandDot.x && math.Abs(float64(dot.y - islandDot.y)) == 1 {
			return true
		}
		if dot.y == islandDot.y && math.Abs(float64(dot.x - islandDot.x)) == 1 {
			return true
		}
	}
	return false
}
