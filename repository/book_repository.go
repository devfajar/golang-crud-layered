package repository

import (
	"context"

	"github.com/devfajar/golang-rest-standlib/model"
)

type BookRepository interface {
	Save(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*model.Book, error)
	FindAll(ctx context.Context) ([]*model.Book, error)
}