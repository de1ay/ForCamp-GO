/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package teams

import (
	"net/http"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/orgset"
)

func DeleteTeam(token string, id int64, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndIdByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteTeam_Request(id)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		return conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteTeam_Request(id int64) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("DELETE FROM teams WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	APIerr := deleteTeam_Users(id)
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func deleteTeam_Users(id int64) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE users SET team='0' WHERE team=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}
