package repository

import (
	"context"
	"time"

	"gitlab.com/trustify/core/ent"
	"gitlab.com/trustify/core/ent/user"
	"gitlab.com/trustify/core/pkg/entity/model"
	usecaseRepository "gitlab.com/trustify/core/pkg/usercase/repository"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) usecaseRepository.User {
	return &userRepository{client: client}
}

func (r *userRepository) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
}

func (r *userRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

func (r *userRepository) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return r.client.User.UpdateOneID(input.ID).SetInput(input).SetUpdatedAt(time.Now()).Save(ctx)
}
