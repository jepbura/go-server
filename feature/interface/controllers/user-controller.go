package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
	"github.com/jepbura/go-server/feature/usecase"
)

type UserController struct {
	userInteractor usecase.UserInteractor
}

func NewUserController(userInteractor usecase.UserInteractor) *UserController {
	return &UserController{userInteractor}
}

func (controller *UserController) SaveUser_Controller(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user model.NewUser
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.userInteractor.SaveUser_Usecase(user)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *UserController) FindAllUser_Controller(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.userInteractor.FindAllUsers_Usecase()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
