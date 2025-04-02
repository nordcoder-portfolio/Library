package controller

import (
	"context"

	"go.uber.org/zap"

	"github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) UpdateBook(ctx context.Context, req *library.UpdateBookRequest) (*library.UpdateBookResponse, error) {
	i.logger.Info("executing UpdateBook")

	if err := req.ValidateAll(); err != nil {
		i.logger.Warn("invalid data", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := i.booksUseCase.UpdateBook(ctx, req.GetId(), req.GetName(), req.GetAuthorId())

	if err != nil {
		i.logger.Warn("failed to update book", zap.Error(err))
		return nil, i.convertErr(err)
	}

	defer i.logger.Info("successfully finished UpdateBook")

	return &library.UpdateBookResponse{}, nil
}
