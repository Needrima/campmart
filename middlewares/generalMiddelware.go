package middlewares

import (
	"bytes"
	"campmart/database"
	"campmart/helpers"
	"campmart/models"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	gomail "gopkg.in/gomail.v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewSubscriber return a new subscriber from email from form input.
// See type Subscriber
func CreateNewSubscriber(r *http.Request) (models.Subscriber, error) {

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading email from response body:", err)
		return models.Subscriber{}, errors.New("something went wrong try again later")
	}

	email, exp := string(bs), `^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	if !helpers.ValidFormInput(email, exp) {
		fmt.Println("Invaid email")
		return models.Subscriber{}, errors.New("invalid email, check and try again")
	}

	subscribersCollection := database.GetDatabaseCollection("subscribers")

	subscribersCursor, err := subscribersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("subscriber cursor error:", err)
		return models.Subscriber{}, errors.New("something went wrong, try again later")
	}

	for subscribersCursor.Next(context.TODO()) {
		var s models.Subscriber

		if err := subscribersCursor.Decode(&s); err != nil {
			fmt.Println("Error getting subscriber:", err)
			continue
		}

		if strings.EqualFold(s.Email, email) {
			fmt.Println("User already a subscriber")
			return models.Subscriber{}, errors.New("you are already a subscriber")
		}
	}

	databaseID := primitive.NewObjectID()
	id := databaseID.Hex()
	newSubcriber := models.Subscriber{
		DatabaseID: databaseID,
		Id:         id,
		Email:      email,
	}

	return newSubcriber, nil
}

// SendMail sends a new mail to the email specified with the templateName being the name of the html template
// that serves as the body of the mail with data being parsed into the template
func SendMail(email, templateName, subject string, data interface{}) error {
	mail := gomail.NewMessage()

	mail.SetHeader("From", mail.FormatAddress("emailservice@campmart.com", "The Campmart Team"))

	mail.SetHeaders(map[string][]string{
		"To":      {email},
		"Subject": {subject},
	})

	password := os.Getenv("emailPassword")

	tpl := helpers.LoadTemplate()

	msgBuffer := &bytes.Buffer{}
	if err := tpl.ExecuteTemplate(msgBuffer, templateName, data); err != nil {
		errBody := fmt.Sprintf("Error executing welcome mail template: %v", err.Error())
		return errors.New(errBody)
	}

	mail.SetBody("text/html", msgBuffer.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "oyebodeamirdeen@gmail.com", password)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println("Error sending mail:", err)
		return errors.New("sending welcome message failed")
	}

	return nil
}
