package facade

import "fmt"

type VideoConverter struct {
	multiplier int
	videoFile  *videoFile
	codec      *oggCompressionCodec
}

func (v *VideoConverter) Convert() error {
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

func (v *VideoConverter) privateAction(someInt int) error {
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

func NewVideoConverter(a int, b int, c int) *VideoConverter {
	// Creates a fake video converter
	Converter := &VideoConverter{
		videoFile:  newVideoFile(a, b),
		codec:      newOggCompressionCodec(c),
		multiplier: a + b + c,
	}
	fmt.Println("New VideoConverter created")
	return Converter
}
