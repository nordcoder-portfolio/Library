package controller

import (
	"context"

	"go.uber.org/zap"

	"github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) ChangeAuthorInfo(ctx context.Context, req *library.ChangeAuthorInfoRequest) (*library.ChangeAuthorInfoResponse, error) {
	i.logger.Info("executing ChangeAuthorInfo")

	if err := req.ValidateAll(); err != nil {
		i.logger.Warn("invalid data", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := i.authorUseCase.ChangeAuthorInfo(ctx, req.GetId(), req.GetName())

	if err != nil {
		i.logger.Warn("failed to change author info", zap.Error(err))
		return nil, i.convertErr(err)
	}

	defer i.logger.Info("successfully finished ChangeAuthorInfo")

	return &library.ChangeAuthorInfoResponse{}, nil
}
