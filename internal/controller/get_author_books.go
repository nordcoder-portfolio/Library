package controller

import (
	"github.com/project/library/generated/api/library"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *implementation) GetAuthorBooks(req *library.GetAuthorBooksRequest, stream library.Library_GetAuthorBooksServer) error {
	i.logger.Info("executing GetAuthorBooks")

	if err := req.ValidateAll(); err != nil {
		i.logger.Warn("invalid data", zap.Error(err))
		return status.Error(codes.InvalidArgument, err.Error())
	}

	books, err := i.authorUseCase.GetAuthorBooks(stream.Context(), req.GetAuthorId())
	if err != nil {
		i.logger.Warn("failed to get author books", zap.Error(err))
		return status.Error(codes.NotFound, "repository error")
	}

	for _, book := range books {
		i.logger.Info("SENT BOOK: " + book.ID)
		err = stream.Send(&library.Book{
			Id:        book.ID,
			Name:      book.Name,
			AuthorId:  book.AuthorsID,
			CreatedAt: timestamppb.New(book.CreatedAt),
			UpdatedAt: timestamppb.New(book.UpdatedAt),
		})
		if err != nil {
			i.logger.Warn("failed to insert record in stream", zap.Error(err))
			return i.convertErr(err)
		}
	}

	defer i.logger.Info("successfully finished GetAuthorBooks")

	return nil
}
