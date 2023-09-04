package repository

import (
	"context"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/entity/model"
)

// User is an interface of repository
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, id *model.ID, input model.UpdateUserInput) (*model.User, error)
}
