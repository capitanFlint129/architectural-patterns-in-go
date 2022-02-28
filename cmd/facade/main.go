package main

import (
	"fmt"
	"go-facade/cmd/facade/pkg/facade"
)

type converter interface {
	Convert() error
}

func useConverter(converter converter) {
	er := converter.Convert()
	if er != nil {
		fmt.Println("Error in facade")
	}
}

func main() {
	videoConverterFacade := facade.NewVideoConverter(15, 5, 5)
	useConverter(videoConverterFacade)
}
