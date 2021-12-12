package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "github.com/busraarsln/rest-api-with-go/models"
	utils "github.com/busraarsln/rest-api-with-go/utils"
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
