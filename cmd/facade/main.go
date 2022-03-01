package main

import (
	"fmt"

	"architectural-patterns-in-go/pkg/facade"
)

const (
	a = 15
	b = 5
	c = 5
)

var (
	config = Config{
		a:     a,
		b:     b,
		c:     c,
		file:  facade.NewVideoFile(a, b),
		codec: facade.NewOggCompressionCodec(c),
	}
)

func useConverter(converter facade.Converter) error {
	er := converter.Convert()
	if er != nil {
		return fmt.Errorf("Error in facade")
	}
	return nil
}

func main() {
	videoConverterFacade := facade.NewVideoConverter(config.a, config.b, config.c, config.file, config.codec)
	er := useConverter(videoConverterFacade)
	if er != nil {
		fmt.Println("Error when use converter")
	}
}
