package usecase

import (
	"context"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/entity/model"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/usecase/repository"
)

type user struct {
	userRepository repository.User
}

// User of usecase.
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, id *model.ID, input model.UpdateUserInput) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, orderBy []*model.UserOrder, where *model.UserWhereInput) (*model.UserConnection, error)
}

// NewUserUsecase returns user usecase
func NewUserUsecase(r repository.User) User {
	return &user{userRepository: r}
}

func (u *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u *user) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return u.userRepository.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, id *model.ID, input model.UpdateUserInput) (*model.User, error) {
	return u.userRepository.Update(ctx, id, input)
}

func (u *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, orderBy []*model.UserOrder, where *model.UserWhereInput) (*model.UserConnection, error) {
	return u.userRepository.List(ctx, after, first, before, last, orderBy, where)
}
