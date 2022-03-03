package company

import (
	"fmt"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor"
)

//ChemicalFactory produces chemicals
type ChemicalFactory interface {
	Company
}

type chemicalFactory struct{}

func (c *chemicalFactory) Accept(visitor visitor.Visitor) {
	fmt.Println("Chemical factory: accept visitor")
	visitor.VisitChemicalFactory(c)
}

//NewChemicalFactory creates new chemical factory
func NewChemicalFactory() ChemicalFactory {
	return &chemicalFactory{}
}
