package main

/* 主要技巧：Sliding Window
 * 核心思路：
 *     1. 維護一個 map 紀錄字母的頻率
 *     2. 設定一個 window (i,j) 紀錄 window 大小
 *     3. 計算字母出現最多次的頻率 maxF
 *     4. 如果 window 大小（j - i + 1） - 最頻繁出現的字母數量（maxF）超過 k 的時候，i++ 左邊邊界向右移縮小 window 大小，並且把該字母的頻率減一 freqMap[s[i]] - 1
 *     5. 計算最大長度 maxLength = max(maxLength, j - i + 1)
 *     6. window 右邊邊界擴展 1（j++）
 * 例如這筆資料：ABBB, k = 2
 *     1. window { i:0, j:0, A }    freqMap { A: 1 }        maxF 1 => 1 - 1 <= 2
 *     2. window { i:0, j:1, AB }   freqMap { A: 1, B: 1 }  maxF 1 => 2 - 1 <= 2
 *     3. window { i:0, j:2, ABB }  freqMap { A: 1, B: 2 }  maxF 2 => 3 - 2 <= 2
 *     4. window { i:0, j:3, ABBB } freqMap { A: 1, B: 3 }  maxF 3 => 4 - 3 <= 2
 * 範例資料 2：AABABBA, k = 1
 *     1. window { i:0, j:0, A }     freqMap { A: 1 }       maxF 1 => 1 - 1 <= 1  maxLength 1
 *     2. window { i:0, j:1, AA }    freqMap { A: 2 }       maxF 2 => 2 - 2 <= 1  maxLength 2
 *     3. window { i:0, j:2, AAB }   freqMap { A: 2, B: 1 } maxF 2 => 3 - 2 <= 1  maxLength 3
 *     4. window { i:0, j:3, AABA }  freqMap { A: 3, B: 1 } maxF 3 => 4 - 3 <= 1  maxLength 4
 *     5. window { i:0, j:4, AABAB } freqMap { A: 3, B: 2 } maxF 3 => 5 - 3 > 1   maxLength 4 (因為 i++ => 4 - 1 + 1)
 *     6. window { i:1, j:5, ABABB } freqMap { A: 2, B: 3 } maxF 3 => 5 - 3 > 1   maxLength 4 (因為 i++ => 5 - 2 + 1)
 *     7. window { i:2, j:6, BABBA } freqMap { A: 2, B: 3 } maxF 3 => 5 - 3 > 1   maxLength 4 (因為 i++ => 6 - 3 + 1)
 */

func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	maxF, i, j := 0, 0, 0
	maxLength := 0

	for j < len(s) {
		freqMap[s[j]] += 1

		maxF = max(maxF, freqMap[s[j]])

		if (j-i+1)-maxF > k {
			freqMap[s[i]] -= 1
			i++
		}

		maxLength = max(maxLength, j-i+1)
		j++
	}

	return maxLength
}
