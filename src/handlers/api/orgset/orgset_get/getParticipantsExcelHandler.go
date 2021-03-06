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
	"nullteam.info/wplay/demo/src/api/orgset"
	"io/ioutil"
	"time"
	"bytes"
)


func GetParticipantsExcelHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		token := handlers.GetToken(r)
		if orgset.IsUserAdmin(token, w){
			organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token);
			if apiErr != nil {
				src.SetHeaders_API_GET(w)
				apiErr.Print(w)
			} else {
				file, err := ioutil.ReadFile(conf.FOLDER_PARTICIPANTS+"/"+ organizationName +".xlsx")
				if err != nil {
					src.SetHeaders_API_GET(w)
					conf.ErrOpenExcelFile.Print(w)
				} else {
					src.SetHeaders_API_Download(w, "участники.xlsx", r)
					http.ServeContent(w, r, "участники.xlsx", time.Now(), bytes.NewReader(file))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetParticipantsExcel(router *mux.Router)  {
	router.HandleFunc("/orgset.participants.password.get", GetParticipantsExcelHandler)
}
