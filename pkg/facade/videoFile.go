package facade

import "fmt"

// TODO добавить интерфейс

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

// TODO возвращаем интерфейс а не класс
func newVideoFile(x int, y int) *videoFile {
	return &videoFile{
		x: x,
		y: y,
	}
}
