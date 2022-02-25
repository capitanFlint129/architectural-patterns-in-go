package main

import (
	"fmt"
	"go-facade/cmd/facade/pkg/facade"
)

func main() {
	videoConverterFacade := facade.NewVideoConverter(15, 5, 5)
	er := videoConverterFacade.Action1()
	if er != nil {
		fmt.Println("Error at action 1")
	}
	er = videoConverterFacade.Action2(1)
	if er != nil {
		fmt.Println("Error at action 2")
	}
}
