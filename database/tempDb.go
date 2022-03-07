package database

import "campmart/models"

var cartSession = map[string]string{}
var TemporaryCartDatabase = map[string]models.CartItem{}
