package utils

import (
	"encoding/json"

	. "github.com/busraarsln/rest-api-with-go/models"
)

//get users data from a json file
func LoadUsersFromJson() []User {
	bytes, _ := ReadFile("../data/users.json")
	var users []User
	err := json.Unmarshal(bytes, &users)
	CheckError(err)
	return users
}

//get products data from a json file
func LoadProductsFromJson() []Product {
	bytes, _ := ReadFile("../data/products.json")
	var products []Product
	err := json.Unmarshal(bytes, &products)
	CheckError(err)
	return products
}

//get categories data from a json file
func LoadCategoriesFromJson() []Category {
	bytes, _ := ReadFile("../data/categories.json")
	var categories []Category
	err := json.Unmarshal(bytes, &categories)
	CheckError(err)
	return categories
}
