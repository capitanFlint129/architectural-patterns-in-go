package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor/company"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor/visitor"
)

func main() {
	auditor := visitor.NewVisitor()
	for _, auditedCompany := range [...]company.Company{
		company.NewSteelMill(),
		company.NewChemicalFactory(),
		company.NewCarFactory(),
	} {
		auditedCompany.Accept(auditor)
	}
}
