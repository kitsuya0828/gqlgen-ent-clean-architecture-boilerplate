package repository

import (
	"context"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent/user"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/entity/model"
	usecaseRepository "github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/usecase/repository"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) usecaseRepository.User {
	return &userRepository{client: client}
}

func (r *userRepository) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u, err := r.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) Update(ctx context.Context, id *model.ID, input model.UpdateUserInput) (*model.User, error) {
	u, err := r.client.User.UpdateOneID(*id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, orderBy []*model.UserOrder, where *model.UserWhereInput) (*model.UserConnection, error) {
	u, err := r.client.User.Query().Paginate(ctx, after, first, before, last,
		ent.WithUserOrder(orderBy),
		ent.WithUserFilter(where.Filter),
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}
