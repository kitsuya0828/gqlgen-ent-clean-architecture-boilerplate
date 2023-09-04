package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/graph/generated"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/adapter/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	})
}
