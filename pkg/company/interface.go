package company

type visitor = interface {
	VisitSteelMill(steelMill steelMill)
	VisitChemicalFactory(chemicalFactory chemicalFactory)
	VisitCarFactory(carFactory carFactory)
}

// Company produces something
type Company interface {
	Accept(visitor visitor)
}
