package facade

import "fmt"

type Codec interface {
	notify()
	multiply(n int)
}

type oggCompressionCodec struct {
	x int
}

func (s *oggCompressionCodec) notify() {
	fmt.Println("Notification from codec")
	return
}

func (s *oggCompressionCodec) multiply(n int) {
	s.x *= n
}

func NewOggCompressionCodec(x int) Codec {
	return &oggCompressionCodec{
		x: x,
	}
}
