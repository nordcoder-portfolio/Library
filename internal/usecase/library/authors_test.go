package library

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/project/library/mocks"

	"github.com/project/library/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterAuthor(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)

	authorName := "Test"
	author := entity.Author{
		ID:   "1",
		Name: authorName,
	}

	mockAuthorRepo.
		On("RegisterAuthor", mock.Anything, mock.AnythingOfType("entity.Author")).
		Return(author, nil)

	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	result, err := libraryService.RegisterAuthor(context.Background(), authorName)

	require.NoError(t, err)
	assert.Equal(t, author.Name, result.Name)

	mockAuthorRepo.AssertExpectations(t)
}

func TestRegisterAuthorBad(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)
	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	mockAuthorRepo.
		On("RegisterAuthor", mock.Anything, mock.AnythingOfType("entity.Author")).
		Return(entity.Author{}, errors.New("couldn't add :("))
	_, err := libraryService.RegisterAuthor(context.Background(), "authorName")

	assert.Error(t, err)
}

func TestChangeAuthorInfo(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)

	authorID := "1"
	authorName := "Test Updated"

	updatedAuthor := entity.Author{
		ID:   authorID,
		Name: authorName,
	}
	mockAuthorRepo.
		On("ChangeAuthorInfoByID", mock.Anything, authorID, authorName).
		Return(updatedAuthor, nil)

	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	result, err := libraryService.ChangeAuthorInfo(context.Background(), authorID, authorName)

	require.NoError(t, err)
	assert.Equal(t, updatedAuthor.Name, result.Name)

	mockAuthorRepo.AssertExpectations(t)
}

func TestChangeAuthorInfoBad(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)
	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	mockAuthorRepo.
		On("ChangeAuthorInfoByID", mock.Anything, "authorID", "authorName").
		Return(entity.Author{}, errors.New("oops"))
	_, err := libraryService.ChangeAuthorInfo(context.Background(), "authorID", "authorName")

	assert.Error(t, err)
}

func TestGetAuthorBooks(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)
	mockBooksRepo := new(mocks.BooksRepository)

	authorID := "1"
	books := []*entity.Book{
		{ID: "book1", Name: "Book One", AuthorsID: []string{authorID}},
		{ID: "book2", Name: "Book Two", AuthorsID: []string{authorID}},
	}

	mockAuthorRepo.
		On("GetAuthorBooks", mock.Anything, authorID).
		Return(books, nil)

	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	result, err := libraryService.GetAuthorBooks(context.Background(), authorID)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, books, result)

	mockAuthorRepo.AssertExpectations(t)
	mockBooksRepo.AssertExpectations(t)
}

func TestGetAuthorBooksBad(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)
	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	mockAuthorRepo.
		On("GetAuthorBooks", mock.Anything, "authorID").
		Return(nil, errors.New("oops"))

	_, err := libraryService.GetAuthorBooks(context.Background(), "authorID")

	assert.Error(t, err)
}

func TestGetAuthorInfo(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)

	authorID := "1"
	author := entity.Author{
		ID:   authorID,
		Name: "Test",
	}

	mockAuthorRepo.
		On("GetAuthorByID", mock.Anything, authorID).
		Return(author, nil)

	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	result, err := libraryService.GetAuthorInfo(context.Background(), authorID)

	require.NoError(t, err)
	assert.Equal(t, author.Name, result.Name)

	mockAuthorRepo.AssertExpectations(t)
}

func TestGetAuthorInfoBad(t *testing.T) {
	t.Parallel()

	mockAuthorRepo := new(mocks.AuthorRepository)
	libraryService := &libraryImpl{
		authorRepository: mockAuthorRepo,
	}

	mockAuthorRepo.
		On("GetAuthorByID", mock.Anything, "authorID").
		Return(entity.Author{}, errors.New("oops"))
	_, err := libraryService.GetAuthorInfo(context.Background(), "authorID")

	assert.Error(t, err)
}
