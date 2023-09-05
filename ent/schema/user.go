package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent/schema/ulid"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/const/globalid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(globalid.New().User.Prefix)
			}).
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("name").
			NotEmpty().
			MaxLen(255).
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Int("age").
			Positive().
			Annotations(
				entgql.OrderField("AGE"),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", User.Type).
			Unique().
			From("children"),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
    return []schema.Annotation{
		entgql.MultiOrder(),
		entgql.RelayConnection(),
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
}
