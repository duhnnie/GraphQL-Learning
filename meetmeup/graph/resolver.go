package graph

import (
	"meetmeup/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

// var users = []*model.User{
// 	{ID: "1", Username: "pepiro", Email: "pepiro@gmail.com"},
// }

type Resolver struct {
	userRepo   *repository.UserRepository
	meetupRepo *repository.MeetupRepository
}

func NewResolver(userRepo *repository.UserRepository, meetupRepo *repository.MeetupRepository) *Resolver {
	return &Resolver{
		userRepo:   userRepo,
		meetupRepo: meetupRepo,
	}
}
