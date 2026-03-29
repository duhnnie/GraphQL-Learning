package graph

import "meetmeup/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

var users = []*model.User{
	{ID: "1", Username: "pepiro", Email: "pepiro@gmail.com"},
}
var meetups = []*model.Meetup{
	{ID: "1", Name: "meetup #1", Description: "description 1", UserID: "1"},
}

type Resolver struct {
}
