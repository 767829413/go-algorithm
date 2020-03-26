package recursive

import "testing"

func TestSeat(t *testing.T) {
	t.Log(Seat(0))
	t.Log(Seat(1))
	t.Log(Seat(2))
	t.Log(Seat(10))
}

func TestStepCount(t *testing.T) {
	t.Log(StepCount(0))
	t.Log(StepCount(1))
	t.Log(StepCount(2))
	t.Log(StepCount(10))
}
