package graph

import "github.com/jepbura/go-server/feature/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userInteractor usecase.UserInteractor
}
