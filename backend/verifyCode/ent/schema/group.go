package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"regexp"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Match(regexp.MustCompile("[a-zA-Z_]+$")), //组名的Regexp验证。
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
