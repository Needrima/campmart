package database

import "campmart/models"

// TemporaryCartDB stores cart items remporarily till user places an order or clears out cart items
var TemporaryCartDB = map[string]map[string]models.CartItem{}
