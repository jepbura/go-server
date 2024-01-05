package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import services "github.com/jepbura/go-server/pkg/usecase/usecase_interfaces"

type Resolver struct {
	Usecase services.UserUseCase
}
