package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"gitlab.com/trustify/core/ent"
	"gitlab.com/trustify/core/graph/generated"
	"gitlab.com/trustify/core/pkg/adapter/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	})
}
