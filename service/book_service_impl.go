package service

import (
	"context"

	"github.com/devfajar/golang-rest-standlib/data/request"
	"github.com/devfajar/golang-rest-standlib/data/response"
	"github.com/devfajar/golang-rest-standlib/helper"
	"github.com/devfajar/golang-rest-standlib/model"
	"github.com/devfajar/golang-rest-standlib/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

// Create implements BookService.
func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := model.Book{
		Title:          request.Title,
		Author:         request.Author,
		Published_Date: helper.ParseDate(request.Published_Date),
	}
	b.BookRepository.Save(ctx, &book)
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindByID(ctx, bookId)
	helper.PanicIfError(err)
	b.BookRepository.Delete(ctx, book.ID)
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books, err := b.BookRepository.FindAll(ctx)
	helper.PanicIfError(err)
	var bookResponses []response.BookResponse
	for _, value := range books {
		book := response.BookResponse{
			ID:             value.ID,
			Title:          value.Title,
			Author:         value.Author,
			Published_Date: value.Published_Date.Format("2006-01-02"),
		}
		bookResponses = append(bookResponses, book)
	}

	return bookResponses
}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindByID(ctx, bookId)
	helper.PanicIfError(err)
	bookResponse := response.BookResponse{
		ID:             book.ID,
		Title:          book.Title,
		Author:         book.Author,
		Published_Date: book.Published_Date.Format("2006-01-02"),
	}
	return bookResponse
}

// Update implements BookService.
func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindByID(ctx, request.ID)
	helper.PanicIfError(err)

	book.Title = request.Title
	book.Author = request.Author
	book.Published_Date = helper.ParseDate(request.Published_Date)
	b.BookRepository.Update(ctx, book)
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}
