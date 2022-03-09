package company

import (
	"fmt"
)

type carFactory = interface{}

type specificCarFactory struct{}

func (s *specificCarFactory) Accept(visitor visitor) {
	fmt.Println("Car factory: accept visitor")
	visitor.VisitCarFactory(s)
}

// NewCarFactory creates new car factory
func NewCarFactory() Company {
	return &specificCarFactory{}
}
