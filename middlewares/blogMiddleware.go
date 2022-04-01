package middlewares

import (
	"campmart/database"
	"campmart/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlogposts(pageNumber int) []models.BlogPost {
	blogPostsCollection := database.GetDatabaseCollection("blogposts")
	limit, skip, sort := int64(3), int64(pageNumber*3), bson.M{"created_at": -1}
	findOptions := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  sort,
	}

	var blogPosts []models.BlogPost

	blogPostsCursor, err := blogPostsCollection.Find(context.TODO(), bson.M{}, &findOptions)
	if err != nil {
		log.Println("Error getting blogposts cursor:", err)
		return []models.BlogPost{}
	}
	defer blogPostsCursor.Close(context.TODO())

	for blogPostsCursor.Next(context.TODO()) {
		var b models.BlogPost
		if err := blogPostsCursor.Decode(&b); err != nil {
			log.Println("Error decoding cursor into blogpost:", err)
			continue
		}

		blogPosts = append(blogPosts, b)
	}

	return blogPosts
}
