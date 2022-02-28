package facade

import "fmt"

// Converter converts files
type Converter interface {
	Convert() error
}

type videoConverter struct {
	multiplier int
	videoFile  *file
	codec      *codec
}

// Convert executes conversion
func (v *videoConverter) Convert() error {
	fmt.Println("Сonversion started")
	er := v.videoFile.check(10)
	if er != nil {
		return er
	}
	v.codec.notify()
	er = v.privateAction(1)
	if er != nil {
		return er
	}
	fmt.Println("Сonversion finished")
	return nil
}

func (v *videoConverter) privateAction(someInt int) error {
	fmt.Println("Private action started")
	er := v.videoFile.check(someInt)
	if er != nil {
		return er
	}
	v.multiplier = someInt
	v.codec.multiply(someInt)
	fmt.Println("Private action finished")
	return nil
}

// NewVideoConverter creates a fake video converter
func NewVideoConverter(a int, b int, c int) *videoConverter {
	Converter := &videoConverter{
		videoFile:  newVideoFile(a, b),
		codec:      newOggCompressionCodec(c),
		multiplier: a + b + c,
	}
	fmt.Println("New videoConverter created")
	return Converter
}
