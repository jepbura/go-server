package graph

import "github.com/jepbura/go-server/pkg/usecase/usecase_interfaces"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Usecase usecase_interfaces.UseCasesInterface
}
