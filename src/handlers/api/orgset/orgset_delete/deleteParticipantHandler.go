/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"wplay/src/api/orgset/participants"
	"strconv"
)

func getDeleteParticipantPostValues(r *http.Request) (string, int64, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	employee_id, err := strconv.ParseInt(strings.TrimSpace(
		strings.ToLower(r.PostFormValue("participant_id"))), 10, 64)
	if err != nil {
		return "", 0, conf.ErrIdIsNotINT
	}
	return token, employee_id, nil
}

func DeleteParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, employee_id, apiErr := getDeleteParticipantPostValues(r); if apiErr != nil {
			apiErr.Print(w)
		} else {
			participants.DeleteParticipant(token, employee_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleDeleteParticipant(router *mux.Router)  {
	router.HandleFunc("/orgset.participant.delete", DeleteParticipantHandler)
}
