package facade

import "fmt"

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

func newOggCompressionCodec(x int) *oggCompressionCodec {
	return &oggCompressionCodec{
		x: x,
	}
}
