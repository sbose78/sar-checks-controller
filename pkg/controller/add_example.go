package controller

import (
	"github.com/sbose78/sar-checks-operator/pkg/controller/example"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, example.Add)
}
