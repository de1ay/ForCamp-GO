package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"wplay/src/api/orgset/teams"
)

func getAddTeamPostValues(r *http.Request) (string, string){
	Token := r.PostFormValue("token")
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	return Name, Token
}

func AddTeamHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		Name, token := getAddTeamPostValues(r)
		teams.AddTeam(token, Name, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddTeam(router *mux.Router)  {
	router.HandleFunc("/orgset.team.add", AddTeamHandler)
}