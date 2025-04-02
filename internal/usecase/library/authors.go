package library

import (
	"context"

	"github.com/project/library/internal/entity"
)

func (l *libraryImpl) RegisterAuthor(ctx context.Context, authorName string) (entity.Author, error) {
	author, err := l.authorRepository.RegisterAuthor(ctx, entity.Author{
		Name: authorName,
	})

	if err != nil {
		return entity.Author{}, err
	}

	return author, nil
}

func (l *libraryImpl) ChangeAuthorInfo(ctx context.Context, authorID, authorName string) (entity.Author, error) {
	author, err := l.authorRepository.ChangeAuthorInfoByID(ctx, authorID, authorName)

	if err != nil {
		return entity.Author{}, err
	}

	return author, nil
}

func (l *libraryImpl) GetAuthorBooks(ctx context.Context, authorID string) ([]*entity.Book, error) {
	books, err := l.authorRepository.GetAuthorBooks(ctx, authorID)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (l *libraryImpl) GetAuthorInfo(ctx context.Context, authorID string) (entity.Author, error) {
	author, err := l.authorRepository.GetAuthorByID(ctx, authorID)

	if err != nil {
		return entity.Author{}, err
	}

	return author, nil
}
