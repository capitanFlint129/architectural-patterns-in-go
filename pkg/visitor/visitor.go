package visitor

import (
	"fmt"
)

type steelMill = interface {
	// Здесь пишутся методы, нужные для посещения этой комапнии
}

type chemicalFactory = interface {
	// Здесь пишутся методы, нужные для посещения этой комапнии
}

type carFactory = interface {
	// Здесь пишутся методы, нужные для посещения этой комапнии
}

// Visitor audit different companies
type Visitor = interface {
	VisitSteelMill(steelMill steelMill)
	VisitChemicalFactory(chemicalFactory chemicalFactory)
	VisitCarFactory(carFactory carFactory)
}

type visitor struct{}

// VisitSteelMill audits steel mill
func (v *visitor) VisitSteelMill(steelMill steelMill) {
	fmt.Println("Visitor: visit steel mill")
}

// VisitChemicalFactory audits chemical factory
func (v *visitor) VisitChemicalFactory(chemicalFactory chemicalFactory) {
	fmt.Println("Visitor: visit chemical factory")

}

// VisitCarFactory audits car factory
func (v *visitor) VisitCarFactory(carFactory carFactory) {
	fmt.Println("Visitor: visit car factory")

}

// NewVisitor creates new auditor
func NewVisitor() Visitor {
	return &visitor{}
}
