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

func GetSinglePostAndSugestions(id string) models.SingleBlogPage {
	blogPostsCollection := database.GetDatabaseCollection("blogposts")

	var blogPost models.BlogPost
	if err := blogPostsCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&blogPost); err != nil {
		log.Println("Error getting specified blogpost:", err)
		return models.SingleBlogPage{}
	}

	sampleStage := bson.M{"$sample": bson.M{"size": 3}}
	suggestionsCursor, err := blogPostsCollection.Aggregate(context.TODO(), []bson.M{sampleStage})
	if err != nil {
		log.Println("error getting suggestions cursor:", err)
		return models.SingleBlogPage{}
	}

	var suggestions []models.BlogPost
	for suggestionsCursor.Next(context.TODO()) {
		var b models.BlogPost

		if err := suggestionsCursor.Decode(&b); err != nil {
			log.Println("Suggestions cursor ranging error:", err)
			continue
		}

		suggestions = append(suggestions, b)
	}

	singleBlogPage := models.SingleBlogPage{
		BlogPost:    blogPost,
		Suggestions: suggestions,
	}

	return singleBlogPage
}
