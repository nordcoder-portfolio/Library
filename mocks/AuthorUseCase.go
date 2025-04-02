// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/project/library/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// AuthorUseCase is an autogenerated mock type for the AuthorUseCase type
type AuthorUseCase struct {
	mock.Mock
}

// ChangeAuthorInfo provides a mock function with given fields: ctx, authorID, authorName
func (_m *AuthorUseCase) ChangeAuthorInfo(ctx context.Context, authorID string, authorName string) (entity.Author, error) {
	ret := _m.Called(ctx, authorID, authorName)

	if len(ret) == 0 {
		panic("no return value specified for ChangeAuthorInfo")
	}

	var r0 entity.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (entity.Author, error)); ok {
		return rf(ctx, authorID, authorName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entity.Author); ok {
		r0 = rf(ctx, authorID, authorName)
	} else {
		r0 = ret.Get(0).(entity.Author)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, authorID, authorName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorBooks provides a mock function with given fields: ctx, authorID
func (_m *AuthorUseCase) GetAuthorBooks(ctx context.Context, authorID string) ([]*entity.Book, error) {
	ret := _m.Called(ctx, authorID)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorBooks")
	}

	var r0 []*entity.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entity.Book, error)); ok {
		return rf(ctx, authorID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.Book); ok {
		r0 = rf(ctx, authorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, authorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorInfo provides a mock function with given fields: ctx, authorID
func (_m *AuthorUseCase) GetAuthorInfo(ctx context.Context, authorID string) (entity.Author, error) {
	ret := _m.Called(ctx, authorID)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorInfo")
	}

	var r0 entity.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Author, error)); ok {
		return rf(ctx, authorID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Author); ok {
		r0 = rf(ctx, authorID)
	} else {
		r0 = ret.Get(0).(entity.Author)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, authorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterAuthor provides a mock function with given fields: ctx, authorName
func (_m *AuthorUseCase) RegisterAuthor(ctx context.Context, authorName string) (entity.Author, error) {
	ret := _m.Called(ctx, authorName)

	if len(ret) == 0 {
		panic("no return value specified for RegisterAuthor")
	}

	var r0 entity.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Author, error)); ok {
		return rf(ctx, authorName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Author); ok {
		r0 = rf(ctx, authorName)
	} else {
		r0 = ret.Get(0).(entity.Author)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, authorName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthorUseCase creates a new instance of AuthorUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthorUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthorUseCase {
	mock := &AuthorUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
