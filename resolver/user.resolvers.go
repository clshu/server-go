package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"example.com/server-go/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data *model.CreateUserInput) (*model.UserView, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LogIn(ctx context.Context, data *model.LogInInput) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Profile(ctx context.Context) (*model.UserView, error) {
	panic(fmt.Errorf("not implemented"))
}
