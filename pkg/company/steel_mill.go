package company

import (
	"fmt"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor"
)

// SteelMill produces steel products
type SteelMill interface {
	Company
}

type steelMill struct{}

func (s *steelMill) Accept(visitor visitor.Visitor) {
	fmt.Println("Steel mill: accept visitor")
	visitor.VisitSteelMill(s)
}

//NewSteelMill creates new steel mill
func NewSteelMill() SteelMill {
	return &steelMill{}
}
