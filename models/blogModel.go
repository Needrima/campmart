package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	DatabaseID primitive.ObjectID `bson:"_id"`
	Commentor  string             `bson:"commentor"`
	Comment    string             `bson:"comment"`
}

type BlogPost struct {
	DatabaseID primitive.ObjectID `bson:"_id"`
	Id         string             `bson:"id"`
	Title      string             `bson:"title"`
	Content    string             `bson:"content"`
	CreatedAt  time.Time          `bson:"created_at"`
	ImageName  string             `bson:"image_name"`
	Comments   []Comment          `bson:"comments, omitempty"`
}

// AddComment adds a new comment to a blogpost
func (b *BlogPost) AddComment(c Comment) {
	b.Comments = append(b.Comments, c)
}

type BlogPage struct {
	BlogPosts  []BlogPost
	PageNumber int
}

type SingleBlogPage struct {
	BlogPost
	Suggestions []BlogPost
}
