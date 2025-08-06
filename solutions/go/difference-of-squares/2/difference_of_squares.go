package diffsquares


//import "math"

func SquareOfSum(n int) int {
	var s int
	for i := 1; i <= n; i++ {
		s = s + i
	}
	//return int(math.Pow(float64(s), 2))
	return s * s
}

func SumOfSquares(n int) int {
	var s int
	for i := 1; i <= n; i++ {
		//s = s + int(math.Pow(float64(i), 2))
		s = s + i*i
	}
	return s
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
