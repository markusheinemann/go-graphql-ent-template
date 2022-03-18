package repository

import (
	"context"

	"gitlab.com/trustify/core/pkg/entity/model"
)

type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
}
