package main

import (
	"log"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/config"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/adapter/controller"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/infrastructure/datastore"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/infrastructure/graphql"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/infrastructure/router"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/registry"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	e := router.New(srv)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed to open mysql client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
