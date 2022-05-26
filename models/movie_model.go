package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id    		primitive.ObjectID	`json:"id,omitempty"`
	Name			string 							`json:"name,omitempty" validate:"required"`
	Director	string							`json:"director,omitempty" validate:"required"`
	Rating		int									`json:"rating,omitempty" validate:"required"`
	Comment		string							`json:"comment,omitempty" validate:"required"`
}
