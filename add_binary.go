package main

import "strconv"

func addBinary(a string, b string) string {
	resultStr := ""

	if len(a) > len(b) {
		a, b = b, a
	}

	bOffset := len(b) - len(a)
	previous := 0
	for i := len(a) - 1; i >= 0; i-- {
		aStr := string(a[i])
		aNum, _ := strconv.Atoi(aStr)

		bStr := string(b[i+bOffset])
		bNum, _ := strconv.Atoi(bStr)

		tmp := aNum + bNum + previous
		previous = tmp / 2
		remain := tmp % 2

		resultStr = strconv.Itoa(remain) + resultStr
	}

	for i := bOffset - 1; i >= 0; i-- {
		bStr := string(b[i])
		bNum, _ := strconv.Atoi(bStr)
		tmp := bNum + previous
		previous = tmp / 2
		remain := tmp % 2

		resultStr = strconv.Itoa(remain) + resultStr
	}

	if previous > 0 {
		resultStr = strconv.Itoa(previous) + resultStr
	}

	return resultStr
}
