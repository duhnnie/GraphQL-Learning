package repository

import (
	"context"
	"fmt"
	"meetmeup/graph/model"
)

type MeetupRepository struct {
	meetups []*model.Meetup
}

func NewMeetupRepo() *MeetupRepository {
	return &MeetupRepository{
		meetups: []*model.Meetup{
			{ID: "1", Name: "meetup #1", Description: "description 1", UserID: "1"},
		},
	}
}

func (r *MeetupRepository) CreateMeetup(ctx context.Context, input model.NewMeetupInput) (*model.Meetup, error) {
	m := &model.Meetup{
		ID:          fmt.Sprintf("%d", len(r.meetups)+1),
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}

	r.meetups = append(r.meetups, m)
	return m, nil
}

func (r *MeetupRepository) GetAll(ctx context.Context) ([]*model.Meetup, error) {
	return r.meetups, nil
}
