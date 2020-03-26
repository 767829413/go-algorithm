package recursive

func Seat(n int) int {
	if n == 1 {
		return 1
	}
	if n < 1 {
		return 0
	}
	return Seat(n-1) + 1
}

func StepCount(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return StepCount(n-1) + StepCount(n-2)
}
