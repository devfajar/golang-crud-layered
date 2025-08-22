package repository

import (
	"context"
	"database/sql"

	"github.com/devfajar/golang-rest-standlib/helper"
	"github.com/devfajar/golang-rest-standlib/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "DELETE FROM book WHERE id = $1"
	_, errExec := tx.ExecContext(ctx, query, bookId); 
	helper.PanicIfError(errExec)
	return nil
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) ([]*model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT id, title, author, published_date FROM book"
	result, errExec := tx.QueryContext(ctx, query)
	helper.PanicIfError(errExec)
	defer result.Close()

	var books []*model.Book
	for result.Next() {
		book := &model.Book{}
		errScan := result.Scan(&book.ID, &book.Title, &book.Author, &book.Published_Date)
		helper.PanicIfError(errScan)

		books = append(books, book)
	}

	return books, nil
}

// FindByID implements BookRepository.
func (b *BookRepositoryImpl) FindByID(ctx context.Context, id int) (*model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT id, title, author, published_date FROM book WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	book := &model.Book{}
	errScan := row.Scan(&book.ID, &book.Title, &book.Author, &book.Published_Date)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		}
		helper.PanicIfError(errScan)
	}

	return book, nil
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book *model.Book) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "INSERT INTO book (title, author, published_date) VALUES ($1, $2, $3) RETURNING id"
	_, errExec := tx.ExecContext(ctx, query, book.Title, book.Author, book.Published_Date)
	helper.PanicIfError(errExec)

	return nil
}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book *model.Book) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE book SET title = $1, author = $2, published_date = $3 WHERE id = $4"
	_, errExec := tx.ExecContext(ctx, query, book.Title, book.Author, book.Published_Date, book.ID)
	helper.PanicIfError(errExec)

	return nil
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: db}
}
