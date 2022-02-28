package facade

import "fmt"

type file interface {
	check(z int) error
}

type videoFile struct {
	x int
	y int
}

func (f *videoFile) check(z int) error {
	if f.x+f.y < z {
		return fmt.Errorf("Error")
	}
	return nil
}

func newVideoFile(x int, y int) *file {
	return &videoFile{
		x: x,
		y: y,
	}
}
