package apis

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "github.com/busraarsln/rest-api-with-go/models"
	utils "github.com/busraarsln/rest-api-with-go/utils"
	"github.com/gorilla/mux"
)

//GET - /products
func GetProducts(w http.ResponseWriter, r *http.Request) {

	bytes, err := utils.ReadFile("../data/products.json")
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

//GET - /products/{id}
func GetProductById(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, p := range utils.LoadProductsFromJson() {
		if p.ID == key {
			product = p
		}
	}

	data, err := json.Marshal(product)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

var id int = 3

//POST - /products
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	utils.CheckError(err)

	var createdProduct = make(map[string]Product)

	product.CreatedAt = time.Now()
	id++
	product.ID = id
	k := strconv.Itoa(id)
	createdProduct[k] = product

	data, err := json.Marshal(product)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

//PUT - /products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	utils.CheckError(err)

	for _, p := range utils.LoadProductsFromJson() {
		if p.ID == key {
			p.ID = key
			p.UpdatedAt = time.Now()
			//update db
		}
	}
	w.WriteHeader(http.StatusOK)

}

//DELETE - /products/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	utils.CheckError(err)

	for _, p := range utils.LoadProductsFromJson() {
		if p.ID == key {
			//delete from db
		}
	}
	w.WriteHeader(http.StatusOK)
}
