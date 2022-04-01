package middlewares

import (
	"campmart/helpers"
	"campmart/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrStringToInt = errors.New("could not convert string to int")
	ErrFormFile    = errors.New("formfile error")
)

// CreateNewProduct creates new product from form inputs in new-product.html by admin.
// Error occurs if a string text is given to a number field or a formfile err
func CreateNewProduct(r *http.Request) (models.Product, error) {
	wantedPassword := os.Getenv("campmartAdminPassword")
	gottenPassword := strings.TrimSpace(r.FormValue("adminPassword"))
	if err := bcrypt.CompareHashAndPassword([]byte(wantedPassword), []byte(gottenPassword)); err != nil {
		log.Println("Admin password error:", err.Error())
		return models.Product{}, errors.New("YOU ARE NOT AN ADMIN")
	}

	var newProduct models.Product
	var seller models.Seller

	seller.DatabaseID = primitive.NewObjectID()
	seller.Id = seller.DatabaseID.Hex()
	seller.Seller_name = strings.TrimSpace(r.FormValue("sellerName"))
	seller.Seller_email = strings.TrimSpace(r.FormValue("sellerEmail"))
	seller.Seller_phone = strings.TrimSpace(r.FormValue("sellerPhone"))

	newProduct.DatabaseID = primitive.NewObjectID()
	newProduct.Id = newProduct.DatabaseID.Hex()
	newProduct.Name = strings.TrimSpace(r.FormValue("productName"))
	price, err := strconv.Atoi(strings.TrimSpace(r.FormValue("productPrice")))
	if err != nil {
		fmt.Printf("%v: for price", ErrStringToInt.Error())
		return models.Product{}, ErrStringToInt
	}
	newProduct.Price = price
	newProduct.Types = strings.Split(strings.TrimSpace(r.FormValue("productType")), ",")
	newProduct.Description = strings.TrimSpace(r.FormValue("productDescription"))
	newProduct.Properties = strings.Split(strings.TrimSpace(r.FormValue("productProperties")), ",")
	newProduct.Category = strings.TrimSpace(r.FormValue("category"))
	rating, err := strconv.Atoi(strings.TrimSpace(r.FormValue("rating")))
	if err != nil {
		fmt.Printf("%v: for rating", ErrStringToInt.Error())
		return models.Product{}, ErrStringToInt
	}
	newProduct.Rating = rating
	newProduct.Brand = strings.TrimSpace(r.FormValue("brandName"))
	newProduct.Date_added = time.Now().Format(time.ANSIC)
	newProduct.Seller = seller

	// get form files
	if err := r.ParseMultipartForm(2 << 10); err != nil {
		log.Println("Max memory err:", err)
	}

	imgForm := r.MultipartForm

	formFiles := imgForm.File["img_files"]

	// range over form file and get productimages for each file in form
	var productImages []models.Image

	for _, fileHeader := range formFiles {
		f, _ := fileHeader.Open()
		productImg := models.Image{
			File:      f,
			Name:      fileHeader.Filename,
			Extension: filepath.Ext(fileHeader.Filename),
		}

		productImages = append(productImages, productImg)
	}

	imgNames, err := helpers.ProcessImageAndReturnNames(productImages, newProduct.Id, "website-pub/images/products")
	newProduct.Image_names = imgNames

	return newProduct, err
}

func CreateNewBlog(r *http.Request) (models.BlogPost, error) {
	wantedPassword := os.Getenv("campmartAdminPassword")
	gottenPassword := strings.TrimSpace(r.FormValue("adminPassword"))
	if err := bcrypt.CompareHashAndPassword([]byte(wantedPassword), []byte(gottenPassword)); err != nil {
		log.Println("Admin password error:", err.Error())
		return models.BlogPost{}, errors.New("YOU ARE NOT AN ADMIN")
	}

	var newBlog models.BlogPost

	newBlog.DatabaseID = primitive.NewObjectID()
	newBlog.Id = newBlog.DatabaseID.Hex()
	newBlog.Title = strings.TrimSpace(r.FormValue("blog_title"))
	newBlog.Content = strings.TrimSpace(r.FormValue("blog_content"))
	newBlog.CreatedAt = time.Now()

	file, header, err := r.FormFile("blog_img")
	if err != nil {
		log.Println("Error formfile for blog image:", err)
		return models.BlogPost{}, err
	}

	blogImg := models.Image{
		File:      file,
		Name:      header.Filename,
		Extension: filepath.Ext(header.Filename),
	}

	imageNames, err := helpers.ProcessImageAndReturnNames([]models.Image{blogImg}, newBlog.Id, "website-pub/images/blog")
	newBlog.ImageName = imageNames[0]

	return newBlog, err
}
