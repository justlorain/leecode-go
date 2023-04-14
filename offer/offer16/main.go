package main

func main() {

}

// 快速幂
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMulti(x, n)
	}
	return 1.0 / quickMulti(x, n)
}

func quickMulti(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMulti(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
