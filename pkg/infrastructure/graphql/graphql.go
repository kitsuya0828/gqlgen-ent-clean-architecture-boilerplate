package graphql

import (
	"entgo.io/contrib/entgql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/adapter/controller"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/adapter/resolver"
)

// NewServer generates graphql server
func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}
