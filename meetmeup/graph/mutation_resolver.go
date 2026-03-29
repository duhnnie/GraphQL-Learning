package graph

import (
	"context"
	"fmt"
	"meetmeup/graph/model"
)

type mutationResolver struct{ *Resolver }

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetupInput) (*model.Meetup, error) {
	m := &model.Meetup{
		ID:          fmt.Sprintf("%d", len(meetups)+1),
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}

	meetups = append(meetups, m)
	return m, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u := &model.User{
		ID:       fmt.Sprintf("%d", len(users)+1),
		Username: input.Username,
		Email:    input.Email,
	}

	users = append(users, u)
	return u, nil
}
