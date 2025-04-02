package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
	"github.com/project/library/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeAuthorInfoWrongId(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	_, err := i.ChangeAuthorInfo(context.Background(), &library.ChangeAuthorInfoRequest{Id: "1"})
	assert.Error(t, err)
}

func TestChangeAuthorInfo(t *testing.T) {
	t.Parallel()

	mockAuthorUseCase := new(mocks.AuthorUseCase)

	ctrl := &implementation{
		authorUseCase: mockAuthorUseCase,
		logger:        zap.NewNop(),
	}

	id := uuid.NewString()

	req := &library.ChangeAuthorInfoRequest{
		Id:   id,
		Name: "Updated Name",
	}

	mockAuthorUseCase.
		On("ChangeAuthorInfo", mock.Anything, id, "Updated Name").
		Return(entity.Author{ID: id, Name: "Updated Name"}, nil)

	resp, err := ctrl.ChangeAuthorInfo(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, &library.ChangeAuthorInfoResponse{}, resp)
}
