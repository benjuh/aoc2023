package util

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func Order(x, x2 int) (int, int) {
	if x <= x2 {
		return x, x2
	}
	return x2, x

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
