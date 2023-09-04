package registry

import (
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/adapter/controller"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/adapter/repository"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
