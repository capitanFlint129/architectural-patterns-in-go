package facade

import "fmt"

type File interface {
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

func NewVideoFile(x int, y int) File {
	return &videoFile{
		x: x,
		y: y,
	}
}
