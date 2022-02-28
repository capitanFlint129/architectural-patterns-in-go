package facade

import "fmt"

type codec interface {
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

func newOggCompressionCodec(x int) codec {
	codec := codec(&oggCompressionCodec{
		x: x,
	})
	return codec
}
