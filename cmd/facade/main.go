package main

import (
	"fmt"

	"go-facade/pkg/facade"
)

const (
	a = 15
	b = 5
	c = 5
)

func useConverter(converter facade.Converter) error {
	er := converter.Convert()
	if er != nil {
		return fmt.Errorf("Error in facade")
	}
	return nil
}

func main() {
	videoConverterFacade := facade.NewVideoConverter(a, b, c)
	er := useConverter(videoConverterFacade)
	if er != nil {
		fmt.Println("Error when use converter")
	}
}
