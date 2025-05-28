package main

// 2 維 DP 題
// 思路：每個格子可以從 grid[x-1][y] & grid[x][y-1] 來的
// 所以轉化成公式就是：
// f(x, y) = f(x-1, y) + f(x, y-1)
// f(0, 0) => 1
// f(0, y) => 1 因為只能從上面往下走一條路
// f(x, 0) => 1 因為只能從左邊往右走一條路

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for x := range m {
		for y := range n {
			if x == 0 || y == 0 {
				grid[x][y] = 1
			} else {
				grid[x][y] = grid[x-1][y] + grid[x][y-1]
			}
		}
	}

	return grid[m-1][n-1]
}
