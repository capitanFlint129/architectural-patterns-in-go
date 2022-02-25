package facade

import (
	"fmt"
)

type VideoConverter struct {
	parameter_1 *firstParameter
	parameter_2 *secondParameter
	parameter_3 int
}

func NewVideoConverter(a int, b int, c int) *VideoConverter {
	VideoConverter := &VideoConverter{
		parameter_1: newFirstParameter(a, b),
		parameter_2: newSecondParameter(c),
		parameter_3: a + b + c,
	}
	fmt.Println("New VideoConverter created")
	return VideoConverter
}

func (v *VideoConverter) Action1() error {
	fmt.Println("Action 1 started")
	er := v.parameter_1.check(10)
	if er != nil {
		return er
	}
	v.parameter_2.notify()
	fmt.Println("Action 1 finished")
	return nil
}

func (v *VideoConverter) Action2(some_int int) error {
	fmt.Println("Action 2 started")
	er := v.parameter_1.check(some_int)
	if er != nil {
		return er
	}
	v.parameter_3 = some_int
	v.parameter_2.multiply(some_int)
	fmt.Println("Action 1 finished")
	return nil
}

type firstParameter struct {
	x int
	y int
}

func newFirstParameter(x int, y int) *firstParameter {
	return &firstParameter{
		x: x,
		y: y,
	}
}

func (f *firstParameter) check(z int) error {
	if f.x+f.y < z {
		return fmt.Errorf("Error")
	}
	return nil
}

type secondParameter struct {
	x int
}

func newSecondParameter(x int) *secondParameter {
	return &secondParameter{
		x: x,
	}
}

func (s *secondParameter) notify() {
	fmt.Println("Message from second parameter")
	return
}

func (s *secondParameter) multiply(n int) {
	s.x *= n
}
