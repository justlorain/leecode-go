package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, row := range matrix {
		res := binarySearch(row, target)
		if res != -1 {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	// m 行数 n 列数
	m, n := len(matrix), len(matrix[0])
	// 行的范围：[x, m-1]
	// 列的范围：[0, y]
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
