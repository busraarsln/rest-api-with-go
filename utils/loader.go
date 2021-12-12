package utils

import (
	"encoding/json"

	. "github.com/busraarsln/rest-api-with-go/models"
)

//get users data from a json file
func LoadUsersFromJson() []User {
	bytes, _ := ReadFile("../json/users.json")
	var users []User
	json.Unmarshal(bytes, &users)
	return users
}

//get products data from a json file
func LoadProductsFromJson() []Product {
	bytes, _ := ReadFile("../json/products.json")
	var products []Product
	json.Unmarshal(bytes, &products)
	return products
}

//get categories data from a json file
func LoadCategoriesFromJson() []Category {
	bytes, _ := ReadFile("../json/categories.json")
	var categories []Category
	json.Unmarshal(bytes, &categories)
	return categories
}
