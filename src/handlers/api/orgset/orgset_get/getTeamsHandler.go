package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/orgset/teams"
	"forcamp/src/handlers"
)

func GetTeamsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		teams.GetTeams(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetTeams(router *mux.Router)  {
	router.HandleFunc("/orgset.teams.get", GetTeamsHandler)
}