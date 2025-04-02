package controller

import (
	"context"

	"go.uber.org/zap"

	"github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) RegisterAuthor(ctx context.Context, req *library.RegisterAuthorRequest) (*library.RegisterAuthorResponse, error) {
	i.logger.Info("executing RegisterAuthor")

	if err := req.ValidateAll(); err != nil {
		i.logger.Warn("invalid data", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	author, err := i.authorUseCase.RegisterAuthor(ctx, req.GetName())

	if err != nil {
		i.logger.Warn("failed to register author", zap.Error(err))
		return nil, i.convertErr(err)
	}

	defer i.logger.Info("successfully finished RegisterAuthor")

	return &library.RegisterAuthorResponse{
		Id: author.ID,
	}, nil
}
