package library

import (
	"context"

	"github.com/project/library/internal/entity"
)

func (l *libraryImpl) AddBook(ctx context.Context, name string, authorIDs []string) (entity.Book, error) {
	book, err := l.booksRepository.AddBook(ctx, entity.Book{
		Name:      name,
		AuthorsID: authorIDs,
	})

	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (l *libraryImpl) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	book, err := l.booksRepository.GetBookByID(ctx, bookID)

	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (l *libraryImpl) UpdateBook(ctx context.Context, bookID, newName string, authorIDs []string) (entity.Book, error) {
	book, err := l.booksRepository.UpdateBookInfoByID(ctx, bookID, newName, authorIDs)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}
