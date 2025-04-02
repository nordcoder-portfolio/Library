package repository

import (
	"context"

	"github.com/project/library/internal/entity"
)

type (
	AuthorRepository interface {
		RegisterAuthor(ctx context.Context, author entity.Author) (entity.Author, error)
		GetAuthorByID(ctx context.Context, id string) (entity.Author, error)
		ChangeAuthorInfoByID(ctx context.Context, id, info string) (entity.Author, error)
		GetAuthorBooks(ctx context.Context, id string) ([]*entity.Book, error)
	}

	BooksRepository interface {
		AddBook(ctx context.Context, book entity.Book) (entity.Book, error)
		GetBookByID(ctx context.Context, bookID string) (entity.Book, error)
		UpdateBookInfoByID(ctx context.Context, bookID, info string, authorIDs []string) (entity.Book, error)
	}
)
