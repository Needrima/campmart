package middlewares

import "campmart/models"

// GetCartItemFomProduct create new cart item from a product
// See type CartItem
func GetCartItemFomProduct(product models.Product, qty int, selectedType string) models.CartItem {
	cartItem := models.CartItem{
		Id:           product.Id,
		Image_name:   product.Image_names[0],
		Name:         product.Name,
		Price:        product.Price,
		Quantity:     qty,
		Types:        product.Types,
		SelectedType: selectedType,
	}

	return cartItem
}
