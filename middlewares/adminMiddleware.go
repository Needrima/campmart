package middlewares

import (
	"campmart/helpers"
	"campmart/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrStringToInt = errors.New("Could not convert string to int")
	ErrFormFile    = errors.New("FormFile error")
)

func CreateNewProduct(r *http.Request) (models.Product, error) {
	var newProduct models.Product
	var seller models.Seller

	seller.DatabaseID = primitive.NewObjectID()
	seller.Id = seller.DatabaseID.Hex()
	seller.Seller_name = r.FormValue("sellerName")
	seller.Seller_email = r.FormValue("sellerEmail")
	seller.Seller_phone = r.FormValue("sellerPhone")

	newProduct.DatabaseID = primitive.NewObjectID()
	newProduct.Id = newProduct.DatabaseID.Hex()
	newProduct.Name = r.FormValue("productName")
	price, err := strconv.Atoi(r.FormValue("productPrice"))
	if err != nil {
		fmt.Printf("%v: for price", ErrStringToInt.Error())
		return models.Product{}, ErrStringToInt
	}
	newProduct.Price = price
	newProduct.Types = strings.Split(r.FormValue("productType"), ",")
	newProduct.Description = r.FormValue("productDescription")
	newProduct.Properties = strings.Split(r.FormValue("productProperties"), ",")
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		fmt.Printf("%v: for rating", ErrStringToInt.Error())
		return models.Product{}, ErrStringToInt
	}
	newProduct.Rating = rating
	newProduct.Brand = r.FormValue("brandName")
	newProduct.Date_added = time.Now().Format(time.ANSIC)
	newProduct.Seller = seller

	f1, h1, err := r.FormFile("f1")
	if err != nil {
		log.Printf("%v: %v", ErrFormFile, err)
		return models.Product{}, ErrFormFile
	}

	f2, h2, err := r.FormFile("f2")
	if err != nil {
		log.Printf("%v: %v", ErrFormFile, err)
		return models.Product{}, ErrFormFile
	}

	f3, h3, err := r.FormFile("f3")
	if err != nil {
		log.Printf("%v: %v", ErrFormFile, err)
		return models.Product{}, ErrFormFile
	}

	f4, h4, err := r.FormFile("f4")
	if err != nil {
		log.Printf("%v: %v", ErrFormFile, err)
		return models.Product{}, ErrFormFile
	}

	img_names, err := helpers.ProcessImageAndReturnNames([]models.ProductImage{
		{File: f1, Name: h1.Filename, Extension: filepath.Ext(h1.Filename)},
		{File: f2, Name: h2.Filename, Extension: filepath.Ext(h2.Filename)},
		{File: f3, Name: h3.Filename, Extension: filepath.Ext(h3.Filename)},
		{File: f4, Name: h4.Filename, Extension: filepath.Ext(h4.Filename)},
	}, newProduct.Id)
	newProduct.Image_names = img_names

	return newProduct, err
}
