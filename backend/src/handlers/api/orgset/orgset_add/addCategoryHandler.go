package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"wplay/src/api/orgset/categories"
)

func getAddCategoryPostValues(r *http.Request) (categories.Category, string){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	NegativeMarks := strings.TrimSpace(strings.ToLower(r.PostFormValue("negative_marks")))
	return categories.Category{ID: 0, Name: Name, NegativeMarks: NegativeMarks}, Token
}

func AddCategoryHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		category, token := getAddCategoryPostValues(r)
		categories.AddCategory(token, category, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddCategory(router *mux.Router)  {
	router.HandleFunc("/orgset.category.add", AddCategoryHandler)
}