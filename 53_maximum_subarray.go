package main

import "math"

// 類型：Greedy / Dynamic Programming
// 算法：Kadane's Algorithm (算是 Greedy)
// 原因：
//    Kadane's 在每一步都做出當下最好的選擇，期望得到整體最好的解
//    他所做的選擇是「如果目前這段總和變負的，就放棄它，從下一個數字重新開始」
//    這是一種貪婪的選擇，因為判斷「目前這段字陣列」不可能為後續帶來正貢獻
// 核心思路：
//   - 採用 2 個 variables
//     1. globalSum 用來存放當前加總最大的 subarray 總和
//     2. currentSum 用來存放目前 subarray 的加總
//   - 遍歷 nums 每個元素，對 currentSum 加上 nums[i]
//   - 如果目前 currentSum > globalSum 就把 currentSum 的值存放到 globalSum
//   - 如果 currentSum 是負數，我們就 reset currentSum 為 0
// 為何貪婪能成功？
//    1. 最優子結構 optimal substructure：最終的最大子陣列必然來自某一段局部最佳的選擇
//    2. 無後效性  no aftereffect：一旦你決定丟掉某段，就不再依賴它，這不會影響之後的選擇
// 時間複雜度：O(n)

func maxSubArray(nums []int) int {
	globalSum := math.MinInt64
	curSum := 0

	for _, n := range nums {
		curSum += n
		globalSum = max(globalSum, curSum)

		if curSum < 0 {
			curSum = 0
		}
	}

	return globalSum
}

// Divide and Conquer 解法
// 時間複雜度：O(n log n)

func maxSubArray(nums []int) int {
	var divideAndConquer func(int, int) int
	var maxCrossingSubarray func(int, int, int) int

	// maxCrossingSubarray 是一個用來跨越中間，聯合左邊及右邊的最大子集合加總
	// 算是一個 divide and conquer 補丁做到這件事的方法
	// 本質上與 Kadane's 演算法是有差異的
	// O(2 * n/2) => O(n)
	maxCrossingSubarray = func(left, mid, right int) int {
		leftSum, sum := math.MinInt64, 0
		// 從中間往左邊找到最大加總值
		// O(n/2)
		for i := mid; i >= left; i-- {
			sum += nums[i]
			leftSum = max(leftSum, sum)
		}

		// 從中間+1 往右找最大加總值
		// O(n/2)
		rightSum, sum := math.MinInt64, 0
		for i := mid + 1; i <= right; i++ {
			sum += nums[i]
			rightSum = max(rightSum, sum)
		}

		// 把左邊最大加總值和右邊最大加總值直接加總
		return leftSum + rightSum
	}
	// 本質上是 O(log n) 但因為裡面有 maxCrossingSubarray
	// 所以是 O(n log n)
	divideAndConquer = func(left, right int) int {
		if left == right {
			return nums[left]
		}

		mid := left + (right-left)/2

		// 找到左邊最大
		lMax := divideAndConquer(left, mid)
		// 找到右邊最大
		rMax := divideAndConquer(mid+1, right)
		// 找到跨越中間 左邊+右邊最大
		maxCrossing := maxCrossingSubarray(left, mid, right)

		// 回傳三者最大
		return max(lMax, rMax, maxCrossing)
	}

	return divideAndConquer(0, len(nums)-1)
}
