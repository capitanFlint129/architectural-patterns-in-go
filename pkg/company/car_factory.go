package company

import (
	"fmt"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor"
)

//CarFactory produces cars
type CarFactory interface {
	Company
}

type carFactory struct{}

func (c *carFactory) Accept(visitor visitor.Visitor) {
	fmt.Println("Car factory: accept visitor")
	visitor.VisitCarFactory(c)
}

//NewCarFactory creates new car factory
func NewCarFactory() CarFactory {
	return &carFactory{}
}
