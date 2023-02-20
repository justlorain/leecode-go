package main

func main() {

}

func isPerfectSquare(num int) bool {
	ans := func(num int) int {
		left, right := 0, num
		res := -1
		for left <= right {
			mid := left + ((right - left) >> 1)
			if mid*mid <= num {
				res = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return res
	}(num)
	if ans*ans == num {
		return true
	}
	return false
}

func isPerfectSquare2(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := left + ((right - left) >> 1)
		square := mid * mid
		if square < num {
			left = mid + 1
		} else if square > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
