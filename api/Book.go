package api

import "go.mongodb.org/mongo-driver/bson/primitive"

//Book Resource
type Book struct {
    ID     *primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Title  string             	`json:"title"  bson:"title" binding:"required"`
	Author string             	`json:"author" bson:"author" binding:"required"`
	Pages  int                	`json:"pages"  bson:"pages" binding:"required,min=4,max=500"`
}
