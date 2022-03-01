package main

import "architectural-patterns-in-go/pkg/facade"

type Config struct {
	a     int
	b     int
	c     int
	file  facade.File
	codec facade.Codec
}
