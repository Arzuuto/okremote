package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("FistName").NotEmpty(),
		field.String("LastName").NotEmpty(),
		field.String("Email").NotEmpty().Unique(),
		field.String("Password").NotEmpty().Sensitive(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("advertisement", Advertisement.Type),
	}
}
