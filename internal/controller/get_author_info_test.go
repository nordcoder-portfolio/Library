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

func TestGetAuthorInfoWrongId(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	_, err := i.GetAuthorInfo(context.Background(), &library.GetAuthorInfoRequest{Id: "1"})
	assert.Error(t, err)
}

func TestGetAuthorInfo(t *testing.T) {
	t.Parallel()

	mockAuthorUseCase := new(mocks.AuthorUseCase)

	ctrl := &implementation{
		authorUseCase: mockAuthorUseCase,
		logger:        zap.NewNop(),
	}

	id := uuid.NewString()

	req := &library.GetAuthorInfoRequest{
		Id: id,
	}

	mockAuthorUseCase.
		On("GetAuthorInfo", mock.Anything, id).
		Return(entity.Author{ID: id, Name: "Test Author"}, nil)

	resp, err := ctrl.GetAuthorInfo(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, &library.GetAuthorInfoResponse{
		Id:   id,
		Name: "Test Author",
	}, resp)
}
