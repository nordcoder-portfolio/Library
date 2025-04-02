package controller

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"github.com/project/library/generated/api/library"
	"github.com/project/library/mocks"

	"github.com/project/library/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookInfoWrongId(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	_, err := i.GetBookInfo(context.Background(), &library.GetBookInfoRequest{Id: "1"})
	assert.Error(t, err)
}

func TestGetBookInfo(t *testing.T) {
	t.Parallel()

	mockBooksUseCase := new(mocks.BooksUseCase)

	ctrl := &implementation{
		booksUseCase: mockBooksUseCase,
		logger:       zap.NewNop(),
	}

	id := uuid.NewString()

	req := &library.GetBookInfoRequest{
		Id: id,
	}

	mockBooksUseCase.
		On("GetBook", mock.Anything, id).
		Return(entity.Book{ID: id, Name: "Test Book", AuthorsID: []string{id}}, nil)

	resp, err := ctrl.GetBookInfo(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, &library.GetBookInfoResponse{
		Book: &library.Book{
			Id:        id,
			Name:      "Test Book",
			AuthorId:  []string{id},
			CreatedAt: timestamppb.New(time.Time{}),
			UpdatedAt: timestamppb.New(time.Time{}),
		},
	}, resp)
}
