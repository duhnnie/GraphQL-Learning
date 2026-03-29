package graph

import (
	"context"
	"meetmeup/graph/model"
)

type meetupResolver struct{ *Resolver }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	for _, u := range users {
		if obj.UserID == u.ID {
			return u, nil
		}
	}

	return nil, nil
}
