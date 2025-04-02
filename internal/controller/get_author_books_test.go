package controller

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"github.com/project/library/generated/api/library"
	"github.com/project/library/mocks"

	"github.com/project/library/internal/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type mockGetAuthorBooksStream struct {
	grpc.ServerStream
	sentBooks []*library.Book
}

func (m *mockGetAuthorBooksStream) Send(book *library.Book) error {
	m.sentBooks = append(m.sentBooks, book)
	return nil
}

func (m *mockGetAuthorBooksStream) Context() context.Context {
	return context.Background()
}

func TestGetAuthorBooksWrongId(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	err := i.GetAuthorBooks(&library.GetAuthorBooksRequest{AuthorId: "1"}, &mockGetAuthorBooksStream{})
	assert.Error(t, err)
}

func TestGetAuthorBooks(t *testing.T) {
	t.Parallel()

	authorUseCase := new(mocks.AuthorUseCase)

	id := uuid.NewString()

	ctrl := &implementation{
		authorUseCase: authorUseCase,
		logger:        zap.NewNop(),
	}

	mockStream := &mockGetAuthorBooksStream{}

	books := []*entity.Book{
		{ID: "book-1", Name: "test1", AuthorsID: []string{id}},
		{ID: "book-2", Name: "test2", AuthorsID: []string{id}},
	}

	authorUseCase.
		On("GetAuthorBooks", mock.Anything, id).
		Return(books, nil)

	err := ctrl.GetAuthorBooks(&library.GetAuthorBooksRequest{AuthorId: id}, mockStream)

	expected := []*library.Book{
		{
			Id:        "book-1",
			Name:      "test1",
			AuthorId:  []string{id},
			CreatedAt: timestamppb.New(time.Time{}),
			UpdatedAt: timestamppb.New(time.Time{}),
		},
		{
			Id:        "book-2",
			Name:      "test2",
			AuthorId:  []string{id},
			CreatedAt: timestamppb.New(time.Time{}),
			UpdatedAt: timestamppb.New(time.Time{}),
		},
	}

	require.NoError(t, err)
	require.Equal(t, expected, mockStream.sentBooks)
}
