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
	ErrStringToInt = errors.New("could not convert string to int")
	ErrFormFile    = errors.New("formfile error")
)

// CreateNewProduct creates new product from form inputs in new-product.html
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
	newProduct.Types = strings.Split(strings.TrimSpace(r.FormValue("productType")), ",")
	newProduct.Description = r.FormValue("productDescription")
	newProduct.Properties = strings.Split(strings.TrimSpace(r.FormValue("productProperties")), ",")
	newProduct.Category = r.FormValue("category")
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		fmt.Printf("%v: for rating", ErrStringToInt.Error())
		return models.Product{}, ErrStringToInt
	}
	newProduct.Rating = rating
	newProduct.Brand = r.FormValue("brandName")
	newProduct.Date_added = time.Now().Format(time.ANSIC)
	newProduct.Seller = seller

	// get form files
	if err := r.ParseMultipartForm(2 << 10); err != nil {
		log.Println("Max memory err:", err)
	}

	imgForm := r.MultipartForm

	formFiles := imgForm.File["img_files"]

	// range over form file and get productimages for each file in form
	var productImages []models.ProductImage

	for _, file := range formFiles {
		f, _ := file.Open()
		productImg := models.ProductImage{
			File:      f,
			Name:      file.Filename,
			Extension: filepath.Ext(file.Filename),
		}

		productImages = append(productImages, productImg)
	}

	imgNames, err := helpers.ProcessImageAndReturnNames(productImages, newProduct.Id)
	newProduct.Image_names = imgNames

	return newProduct, err
}
