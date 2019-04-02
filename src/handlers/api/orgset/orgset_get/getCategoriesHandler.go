/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/orgset/categories"
	"wplay/src/handlers"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		categories.GetCategories(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetCategories(router *mux.Router)  {
	router.HandleFunc("/orgset.categories.get", GetCategoriesHandler)
}
