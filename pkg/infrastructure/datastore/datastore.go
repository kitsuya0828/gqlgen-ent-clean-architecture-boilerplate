package datastore

import (
	"entgo.io/ent/dialect"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/config"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/go-sql-driver/mysql"
)

// New returns data source name
func NewDSN() string {
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

// NewClient returns an orm client
func NewClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	d := NewDSN()

	return ent.Open(dialect.MySQL, d, entOptions...)
}
