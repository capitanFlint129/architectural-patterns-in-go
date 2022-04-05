package strategy

import "os"

type Strategy interface {
	Convert(file *os.File)
}
