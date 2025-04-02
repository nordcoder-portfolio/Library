package library

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
	"github.com/project/library/mocks"

	"github.com/project/library/internal/entity"
)

func TestAddBook(t *testing.T) {
	t.Parallel()

	mockBookRepo := new(mocks.BooksRepository)

	bookID := uuid.NewString()
	authors := []string{"author-1", "author-2"}
	book := entity.Book{ID: bookID, Name: "Test Book", AuthorsID: authors}

	mockBookRepo.
		On("AddBook", mock.Anything, mock.AnythingOfType("entity.Book")).
		Return(book, nil)

	libraryService := &libraryImpl{
		booksRepository: mockBookRepo,
	}

	createdBook, err := libraryService.AddBook(context.Background(), book.Name, authors)

	require.NoError(t, err)
	assert.Equal(t, book, createdBook)

	mockBookRepo.AssertExpectations(t)
}

func TestAddBookBad(t *testing.T) {
	t.Parallel()

	mockBookRepo := new(mocks.BooksRepository)
	libraryService := &libraryImpl{booksRepository: mockBookRepo}

	mockBookRepo.
		On("AddBook", mock.Anything, mock.AnythingOfType("entity.Book")).
		Return(entity.Book{}, errors.New("oops"))

	_, err := libraryService.AddBook(context.Background(), uuid.NewString(), []string{})

	assert.Error(t, err)
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	mockBookRepo := new(mocks.BooksRepository)

	ctx := context.Background()
	bookID := "book-123"
	book := entity.Book{ID: bookID, Name: "Existing Book", AuthorsID: []string{"author-1"}}

	mockBookRepo.
		On("GetBookByID", ctx, bookID).
		Return(book, nil)

	libraryService := &libraryImpl{
		booksRepository: mockBookRepo,
	}

	foundBook, err := libraryService.GetBook(ctx, bookID)

	require.NoError(t, err)
	assert.Equal(t, book, foundBook)
	mockBookRepo.AssertExpectations(t)
}

func TestGetBookBad(t *testing.T) {
	t.Parallel()

	mockBookRepo := new(mocks.BooksRepository)
	libraryService := &libraryImpl{booksRepository: mockBookRepo}

	mockBookRepo.
		On("GetBookByID", mock.Anything, "bookID").
		Return(entity.Book{}, errors.New("oops"))

	_, err := libraryService.GetBook(context.Background(), "bookID")

	assert.Error(t, err)
}

func TestUpdateBook(t *testing.T) {
	t.Parallel()

	mockBookRepo := new(mocks.BooksRepository)

	ctx := context.Background()
	bookID := "book-456"
	authorID := "author1"
	updatedBook := entity.Book{ID: bookID, Name: "Updated Name", AuthorsID: []string{authorID}}

	mockBookRepo.
		On("UpdateBookInfoByID", ctx, bookID, "Updated Name", []string{authorID}).
		Return(updatedBook, nil)

	libraryService := &libraryImpl{
		booksRepository: mockBookRepo,
	}

	resultBook, err := libraryService.UpdateBook(ctx, bookID, "Updated Name", []string{authorID})

	require.NoError(t, err)
	assert.Equal(t, updatedBook, resultBook)
	mockBookRepo.AssertExpectations(t)
}
func TestUpdateBookBad(t *testing.T) {
	t.Parallel()

	mockBookRepo := new(mocks.BooksRepository)
	libraryService := &libraryImpl{booksRepository: mockBookRepo}

	mockBookRepo.
		On("UpdateBookInfoByID", mock.Anything, "bookID", "info", []string{"boys"}).
		Return(entity.Book{}, errors.New("oops"))

	_, err := libraryService.UpdateBook(context.Background(), "bookID", "info", []string{"boys"})

	assert.Error(t, err)
}
