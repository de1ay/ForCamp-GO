/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package reasons

import (
	"nullteam.info/wplay/demo/src/api/orgset"
	"net/http"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/conf"
)

func DeleteReason(token string, id int64, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndIdByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteReason_Request(id)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteReason_Request(id int64) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("DELETE FROM reasons WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
