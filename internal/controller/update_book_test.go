package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/project/library/generated/api/library"
	"github.com/project/library/mocks"

	"github.com/project/library/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateBookWrongId(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	_, err := i.UpdateBook(context.Background(), &library.UpdateBookRequest{Id: "1"})
	assert.Error(t, err)
}

func TestUpdateBook(t *testing.T) {
	t.Parallel()

	mockBooksUseCase := new(mocks.BooksUseCase)

	ctrl := &implementation{
		booksUseCase: mockBooksUseCase,
		logger:       zap.NewNop(),
	}

	id := uuid.NewString()

	req := &library.UpdateBookRequest{
		Id:       id,
		Name:     "Updated Book Name",
		AuthorId: []string{},
	}

	mockBooksUseCase.
		On("UpdateBook", mock.Anything, id, "Updated Book Name", []string{}).
		Return(entity.Book{ID: id, Name: "Updated Book Name", AuthorsID: []string{id}}, nil)

	resp, err := ctrl.UpdateBook(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, &library.UpdateBookResponse{}, resp)
}
