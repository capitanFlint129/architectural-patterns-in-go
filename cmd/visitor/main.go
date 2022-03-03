package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/company"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor"
)

func main() {
	steelMill := company.NewSteelMill()
	chemicalFactory := company.NewChemicalFactory()
	carFactory := company.NewCarFactory()

	companies := [3]company.Company{steelMill, chemicalFactory, carFactory}

	auditor := visitor.NewVisitor()
	for _, auditedCompany := range companies {
		auditedCompany.Accept(auditor)
	}
}
