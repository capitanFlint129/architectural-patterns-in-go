package company

import (
	"fmt"
)

type chemicalFactory = interface{}

type specificChemicalFactory struct{}

func (s *specificChemicalFactory) Accept(visitor visitor) {
	fmt.Println("Chemical factory: accept visitor")
	visitor.VisitChemicalFactory(s)
}

//NewChemicalFactory creates new chemical factory
func NewChemicalFactory() Company {
	return &specificChemicalFactory{}
}
