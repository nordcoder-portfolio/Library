package controller

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
	"github.com/project/library/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBookBadAuthor(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	_, err := i.AddBook(context.Background(), &library.AddBookRequest{AuthorId: []string{"1"}})
	assert.Error(t, err)
}

func TestAddBook(t *testing.T) {
	t.Parallel()

	mockBooksUseCase := new(mocks.BooksUseCase)

	ctrl := &implementation{
		booksUseCase: mockBooksUseCase,
		logger:       zap.NewNop(),
	}

	authorID := uuid.NewString()

	req := &library.AddBookRequest{
		Name:     "New Book",
		AuthorId: []string{authorID},
	}

	mockBooksUseCase.
		On("AddBook", mock.Anything, "New Book", []string{authorID}).
		Return(entity.Book{ID: "1", Name: "New Book", AuthorsID: []string{authorID}}, nil)

	resp, err := ctrl.AddBook(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, &library.AddBookResponse{
		Book: &library.Book{
			Id:        "1",
			Name:      "New Book",
			AuthorId:  []string{authorID},
			CreatedAt: timestamppb.New(time.Time{}),
			UpdatedAt: timestamppb.New(time.Time{}),
		},
	}, resp)
}
