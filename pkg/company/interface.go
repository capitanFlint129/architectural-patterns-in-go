package company

import "github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor"

// Company produces something
type Company interface {
	Accept(visitor visitor.Visitor)
}
