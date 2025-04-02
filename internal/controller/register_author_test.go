package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/project/library/generated/api/library"
	"github.com/project/library/mocks"

	"github.com/project/library/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterAuthorWrongName(t *testing.T) {
	t.Parallel()

	i := &implementation{logger: zap.NewNop()}
	_, err := i.RegisterAuthor(context.Background(), &library.RegisterAuthorRequest{Name: ""})
	assert.Error(t, err)
}

func TestRegisterAuthor(t *testing.T) {
	t.Parallel()

	mockAuthorUseCase := new(mocks.AuthorUseCase)

	ctrl := &implementation{
		authorUseCase: mockAuthorUseCase,
		logger:        zap.NewNop(),
	}

	req := &library.RegisterAuthorRequest{
		Name: "New Author",
	}

	mockAuthorUseCase.
		On("RegisterAuthor", mock.Anything, "New Author").
		Return(entity.Author{ID: "1", Name: "New Author"}, nil)

	resp, err := ctrl.RegisterAuthor(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, &library.RegisterAuthorResponse{
		Id: "1",
	}, resp)
}
