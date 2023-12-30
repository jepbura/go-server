package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jepbura/go-server/feature/domain"
	"github.com/jepbura/go-server/feature/usecase"
)

type BookController struct {
	bookInteractor usecase.BookInteractor
}

func NewBookController(bookInteractor usecase.BookInteractor) *BookController {
	return &BookController{bookInteractor}
}

func (controller *BookController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book domain.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.bookInteractor.CreateBook(book)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *BookController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.bookInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
