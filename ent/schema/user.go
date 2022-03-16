package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password"),
		field.Time("created_at").Default(time.Now()).SchemaType(map[string]string{
			dialect.Postgres: "timestamp DEFAULt now()",
		}).Immutable(),
		field.Time("updated_at").Default(time.Now()).SchemaType(map[string]string{
			dialect.Postgres: "timestamp DEFAULT now()",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
