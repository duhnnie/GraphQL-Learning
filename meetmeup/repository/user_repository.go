package repository

import (
	"context"
	"fmt"
	"meetmeup/graph/model"
	"slices"
	"strings"
)

type UserRepository struct {
	users []*model.User
}

func NewUserRepo() *UserRepository {
	return &UserRepository{
		users: []*model.User{
			{ID: "1", Username: "pepiro", Email: "pepiro@gmail.com"},
		},
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *UserRepository) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	fmt.Printf("SELECT * from Users where id IN (%s)\n", strings.Join(userIDs, ", "))
	var u []*model.User

	for _, user := range r.users {
		if slices.Contains(userIDs, user.ID) {
			u = append(u, user)
		}
	}

	return u, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u := &model.User{
		ID:       fmt.Sprintf("%d", len(r.users)+1),
		Username: input.Username,
		Email:    input.Email,
	}

	r.users = append(r.users, u)
	return u, nil
}
