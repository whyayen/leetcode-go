package main

func numIslands(grid [][]byte) int {
	water := byte('0')
	rows := len(grid)
	cols := len(grid[0])

	var visit func(int, int) int

	// 走訪函數
	//   主要功能為：
	//       1. 如果是 island 就把它變 water (視為已走訪），並把 res 設為 1 表示有 island 存在
	//       2. 走訪其他四個方向，把其回傳值相加，屆時回傳的 res > 0 則表示至少有一個 node 是 island
	//       3. 如果超過陣列大小或是 water（表示已走訪）則回傳 0
	visit = func(x0 int, y0 int) int {
		if x0 < 0 || y0 < 0 || x0 >= rows || y0 >= cols || grid[x0][y0] == water {
			return 0
		}

		// 是 island 就把它變 water 當做 visited
		// 把 water 視為 visited 就不需要額外的 M x N 的 visited 二維陣列
		grid[x0][y0] = water
		res := 1 // 有 island 就 res 設為 1

		// 走訪四個方向，把 return 結果相加
		res += visit(x0-1, y0)
		res += visit(x0, y0-1)
		res += visit(x0+1, y0)
		res += visit(x0, y0+1)

		return res
	}

	islandCount := 0
	x, y := 0, 0 // 從 0, 0 走訪到 rows - 1, cols - 1
	for x < rows && y < cols {
		// 如果 res > 0 就表示這次走訪存在 island
		if visit(x, y) > 0 {
			islandCount++
		}

		// 往右一格走訪
		if y+1 < cols {
			y++
		} else {
			// y + 1 >= cols 表示已經超過陣列大小，往下一列第一欄走訪
			x++
			y = 0
		}
	}

	return islandCount
}
