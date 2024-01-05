package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import services "github.com/jepbura/go-server/pkg/usecase/usecase_interfaces"

type Resolver struct {
	Usecase services.UserUseCase
}

// type ResolverTarget struct {
// 	Logger         *zap.Logger
// 	userInteractor *usecase.UserInteractor
// }

// func NewResolver(target ResolverTarget) *Resolver {
// 	return &Resolver{userInteractor: target.userInteractor}
// }

// func NewResolver(userInteractor usecase.UserInteractor) *Resolver {
// 	return &Resolver{
// 		userInteractor: userInteractor,
// 	}
// }
