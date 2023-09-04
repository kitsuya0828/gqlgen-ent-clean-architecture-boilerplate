package testutil

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent/enttest"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/infrastructure/datastore"
)

// NewDBClient loads database for test.
func NewDBClient(t *testing.T) *ent.Client {
	d := datastore.NewDSN()
	return enttest.Open(t, dialect.MySQL, d)
}

// DropAll drops all data from database.
func DropAll(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	DropUser(t, client)
}

// DropUser drops data from users.
func DropUser(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.User.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
