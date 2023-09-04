package controller

import (
	"context"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/entity/model"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/usecase/usecase"
)

type user struct {
	userUsecase usecase.User
}

// User of interface
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, id *model.ID, input model.UpdateUserInput) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, orderBy []*model.UserOrder, where *model.UserWhereInput) (*model.UserConnection, error)
}

// NewUserController returns user controller
func NewUserController(uu usecase.User) User {
	return &user{userUsecase: uu}
}

func (u *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return u.userUsecase.Get(ctx, id)
}

func (u *user) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return u.userUsecase.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, id *model.ID, input model.UpdateUserInput) (*model.User, error) {
	return u.userUsecase.Update(ctx, id, input)
}

func (u *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, orderBy []*model.UserOrder, where *model.UserWhereInput) (*model.UserConnection, error) {
	return u.userUsecase.List(ctx, after, first, before, last, orderBy, where)
}
