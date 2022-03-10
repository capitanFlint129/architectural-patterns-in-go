package company

import (
	"fmt"
)

type steelMill = interface{}

type specificSteelMill struct{}

func (s *specificSteelMill) Accept(visitor visitor) {
	fmt.Println("Steel mill: accept visitor")
	visitor.VisitSteelMill(s)
}

//NewSteelMill creates new steel mill
func NewSteelMill() Company {
	return &specificSteelMill{}
}
