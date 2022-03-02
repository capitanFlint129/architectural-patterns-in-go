package main

import (
	"fmt"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/facade"
)

const (
	a = 15
	b = 5
	c = 5
)

var (
	config = facade.Config{a, b, c, facade.NewVideoFile(a, b), facade.NewOggCompressionCodec(c)}
)

func useConverter(converter facade.Converter) error {
	er := converter.Convert()
	if er != nil {
		return fmt.Errorf("Error in facade")
	}
	return nil
}

func main() {
	videoConverterFacade := facade.NewVideoConverter(config.A, config.B, config.C, config.File, config.Codec)
	er := useConverter(videoConverterFacade)
	if er != nil {
		fmt.Println("Error when use converter")
	}
}
