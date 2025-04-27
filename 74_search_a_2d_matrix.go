package main

// 用 2 次二分搜尋
func searchMatrix(matrix [][]int, target int) bool {
	lowRows := 0
	highRows := len(matrix) - 1

	for lowRows <= highRows {
		midRows := lowRows + (highRows-lowRows)/2
		n := len(matrix[midRows])

		if target >= matrix[midRows][0] && target <= matrix[midRows][n-1] {
			lowCols := 0
			highCols := n - 1

			for lowCols <= highCols {
				midCols := lowCols + (highCols-lowCols)/2

				if target == matrix[midRows][midCols] {
					return true
				} else if target < matrix[midRows][midCols] {
					highCols = midCols - 1
				} else {
					lowCols = midCols + 1
				}
			}

			return false
		} else if target < matrix[midRows][0] {
			highRows = midRows - 1
		} else {
			lowRows = midRows + 1
		}
	}

	return false
}
