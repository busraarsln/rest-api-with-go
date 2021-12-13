package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "github.com/busraarsln/rest-api-with-go/models"
	utils "github.com/busraarsln/rest-api-with-go/utils"
	"github.com/gorilla/mux"
)

//GET - /categories
func GetCategories(w http.ResponseWriter, r *http.Request) {

	bytes, err := utils.ReadFile("../data/categories.json")
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

//GET - /categories/{id}
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var category Category
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, c := range utils.LoadCategoriesFromJson() {
		if c.ID == key {
			category = c
		}
	}

	data, err := json.Marshal(category)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
