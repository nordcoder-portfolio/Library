package library

import (
	"context"

	"github.com/project/library/internal/entity"
	"github.com/project/library/internal/usecase/repository"
	"go.uber.org/zap"
)

type (
	AuthorUseCase interface {
		RegisterAuthor(ctx context.Context, authorName string) (entity.Author, error)
		ChangeAuthorInfo(ctx context.Context, authorID string, authorName string) (entity.Author, error)
		GetAuthorBooks(ctx context.Context, authorID string) ([]*entity.Book, error)
		GetAuthorInfo(ctx context.Context, authorID string) (entity.Author, error)
	}

	BooksUseCase interface {
		AddBook(ctx context.Context, name string, authorIDs []string) (entity.Book, error)
		GetBook(ctx context.Context, bookID string) (entity.Book, error)
		UpdateBook(ctx context.Context, bookID string, name string, authorIDs []string) (entity.Book, error)
	}
)

var _ AuthorUseCase = (*libraryImpl)(nil)
var _ BooksUseCase = (*libraryImpl)(nil)

type libraryImpl struct {
	logger           *zap.Logger
	authorRepository repository.AuthorRepository
	booksRepository  repository.BooksRepository
}

func New(
	logger *zap.Logger,
	authorRepository repository.AuthorRepository,
	booksRepository repository.BooksRepository,
) *libraryImpl {
	return &libraryImpl{
		logger:           logger,
		authorRepository: authorRepository,
		booksRepository:  booksRepository,
	}
}
