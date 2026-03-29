package graph

import (
	"context"
	"meetmeup/graph/model"
)

type userResolver struct{ *Resolver }

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var m = []*model.Meetup{}

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}
