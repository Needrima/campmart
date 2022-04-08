package middlewares

import (
	"campmart/database"
	"campmart/models"
	"context"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetBlogposts gets 3 blogpost based on the pagenumber
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

// GetSinglePostAndSugestions gets a single blog posts by the provided id.
// It also returns three random blog posts as suggestions for other posts
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

// AddNewCommentToPost add a new comment to the blog post with the Id field matching id
// It return the Id Hex of the inserted comment and a nil error if no err occured
// or an empty string if an error occured
func AddNewCommentToPost(r *http.Request, id string) (string, error) {
	commentor := strings.TrimSpace(r.FormValue("commentor"))
	comment := strings.TrimSpace(r.FormValue("comment"))

	var postToCommentTo models.BlogPost
	blogPostsCollection := database.GetDatabaseCollection("blogposts")
	if err := blogPostsCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&postToCommentTo); err != nil {
		log.Println("Error geting post to comment to:", err)
		return "", err
	}

	newComment := models.Comment{
		DatabaseID: primitive.NewObjectID(),
		Commentor:  commentor,
		Comment:    comment,
	}

	postToCommentTo.AddComment(newComment)

	_, err := blogPostsCollection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": postToCommentTo})
	if err != nil {
		log.Println("Error updating result:", err)
		return "", err
	}

	return newComment.DatabaseID.Hex(), nil
}
