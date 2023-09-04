package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/config"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent/migrate"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	d := newDSN()
	client, err := ent.Open(dialect.MySQL, d)
	if err != nil {
		log.Fatalf("failed to open mysql client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	var migrateOptions []schema.MigrateOption
	if config.IsDev() {
		migrateOptions = append(migrateOptions, migrate.WithDropIndex(true))
		migrateOptions = append(migrateOptions, migrate.WithDropColumn(true))
		log.Println("dropping columns and indexes mode")
	} else if config.IsProd() {
		log.Println("append-only mode")
	}

	if err := client.Debug().Schema.Create(ctx, migrateOptions...); err != nil {
		log.Fatalf("failed to create schema resources: %v", err)
	}
}

func newDSN() string {
	mc := mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
			"charset":   config.C.Database.Params.Charset,
			"loc":       config.C.Database.Params.Loc,
		},
	}
	return mc.FormatDSN()
}
