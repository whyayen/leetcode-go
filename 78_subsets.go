package main

// [1, 2, 3]
// 結果：[[], [1], [2], [3], [1, 2], [1, 3], [1, 2, 3]]

// 採用 Backtracking 方法其實就是對每一個數字做包括、不包括的決策樹
// 每一個數字的 include 與 exclude
// 例如樹的 root 是 []
// root:
//   include 1 => [1]
//   exclude 1 => []
// Left [1]:
//   include 2 => [1, 2]
//       Left [1, 2]:
//         include 3 => [1, 2, 3]
//         exclude 3 => [1, 2]
//   exclude 2 => [1]
//       Right [1]:
//         include 3 => [1, 3]
//         exclude 3 => [1]
// Right []:
//   include 2 => [2]
//       Left [2]:
//         include 3 => [2, 3]
//         exclude 3 => [2]
//   exclude 2 => []
//		 Right []:
//         include 3 => [3]
//         exclude 3 => []
// 至此所有葉子都走到了 len(nums) 所以葉子就是所有組合
// [[1, 2, 3], [1, 2], [1, 3], [1], [2, 3], [2], [3], []]

// 時間複制度為 O(2^N * N)
// * N 是因為 copy

// 空間複雜度為 O(2^N * N)
// 2^N 子集合數量
// N 是每個子集合可能的最大長度

func subsets(nums []int) [][]int {
	var result [][]int
	var curr []int
	var explore func(int)

	explore = func(idx int) {
		if idx == len(nums) {
			subset := make([]int, len(curr))
			copy(subset, curr)
			result = append(result, subset)
			return
		}

		// include
		curr = append(curr, nums[idx])
		explore(idx + 1)
		curr = curr[:len(curr)-1] // backtrack

		// exclude
		explore(idx + 1)
	}

	explore(0)
	return result
}
