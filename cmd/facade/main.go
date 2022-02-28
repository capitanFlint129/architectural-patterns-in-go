package main

// TODO отформатировать импорты
// TODO почитать и поправить модули убрать mod
// TODO почитать про линтер
import (
	"fmt"

	"go-facade/pkg/facade"
)

const (
	a = 15
	b = 5
	c = 5
)

func useConverter(converter converter) error {
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
