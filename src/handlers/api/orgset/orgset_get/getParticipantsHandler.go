/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/handlers"
	"nullteam.info/wplay/demo/src/api/orgset/participants"
)


func GetParticipantsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		participants.GetParticipants(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetParticipants(router *mux.Router)  {
	router.HandleFunc("/orgset.participants.get", GetParticipantsHandler)
}
