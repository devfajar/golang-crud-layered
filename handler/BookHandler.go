package handler

import (
	"net/http"
	"strconv"

	"github.com/devfajar/golang-rest-standlib/data/request"
	"github.com/devfajar/golang-rest-standlib/data/response"
	"github.com/devfajar/golang-rest-standlib/helper"
	"github.com/devfajar/golang-rest-standlib/service"
	"github.com/gorilla/mux"
)

type BookHandler struct {
	// You can add fields here if needed, such as a service or repository
	BookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		BookService: bookService,
	}
}


func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Logic to create a book
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(r, &bookCreateRequest)

	h.BookService.Create(r.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:  http.StatusCreated,
		Message: "Book created successfully",
		Data:    nil,
	}

	helper.WriteResponse(w, webResponse)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Logic to update a book
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(r, &bookUpdateRequest)

	
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.ID = id
	h.BookService.Update(r.Context(), bookUpdateRequest)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Book updated successfully",
		Data:    nil,
	}
	helper.WriteResponse(w, webResponse)	
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Logic to delete a book
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	h.BookService.Delete(r.Context(), id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Book deleted successfully",
		Data:    nil,
	}

	helper.WriteResponse(w, webResponse)

}

func (h *BookHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// Logic to find all books
	result := h.BookService.FindAll(r.Context())
	webResponse := response.WebResponse {
		Code:    http.StatusOK,
		Message: "Books retrieved successfully",
		Data:    result,
	}

	helper.WriteResponse(w, webResponse)
} 

func (h *BookHandler) FindById(w http.ResponseWriter, r *http.Request) {
	// Logic to find a book by ID
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	result := h.BookService.FindById(r.Context(), id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Book retrieved successfully",
		Data:    result,
	}
	helper.WriteResponse(w, webResponse)

}	