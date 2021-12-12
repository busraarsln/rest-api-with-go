package apis

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "../models"
	utils "../utils"
)

//GET - /categories
func GetCategories(w http.ResponseWriter, r *http.Request) {

	bytes, err := utils.ReadFile("../json/categories.json")
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

//GET - /categories/{id}
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var category Category
	id := r.URL.Path[2:]
	key, _ := strconv.Atoi(id)
	for _, p := range utils.LoadCategoriesFromJson() {
		if p.ID == key {
			category = p
		}
	}

	data, err := json.Marshal(category)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//POST - /categories
func CreateCategory(w http.ResponseWriter, r *http.Request) {

	var id int = 0
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	utils.CheckError(err)

	var createdCategory = make(map[string]Category)

	category.CreatedAt = time.Now()
	id++
	category.ID = id
	k := strconv.Itoa(id)
	createdCategory[k] = category

	data, err := json.Marshal(category)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

//PUT - /categories/{id}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[2:]
	key, _ := strconv.Atoi(id)
	var p Category
	err := json.NewDecoder(r.Body).Decode(&p)
	utils.CheckError(err)

	for _, p := range utils.LoadCategoriesFromJson() {
		if p.ID == key {
			p.ID = key
			p.UpdatedAt = time.Now()
			//update db
		}
	}
	w.WriteHeader(http.StatusOK)

}

//DELETE - /categories/{id}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[2:]
	key, _ := strconv.Atoi(id)
	var p Category
	err := json.NewDecoder(r.Body).Decode(&p)
	utils.CheckError(err)

	for _, p := range utils.LoadCategoriesFromJson() {
		if p.ID == key {
			//delete from db
		}
	}
	w.WriteHeader(http.StatusOK)
}
