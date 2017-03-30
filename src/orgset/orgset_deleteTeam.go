package orgset

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

func DeleteTeam(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	if checkUserAccess(token, ResponseWriter){
		Organization, _, APIerr := getUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection = src.Connect_Custom(Organization)
		APIerr = deleteTeam_Request(id)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteTeam_Request(id int64) *conf.ApiError{
	Query, err := NewConnection.Prepare("DELETE FROM teams WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	APIerr := deleteTeam_Users(id)
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func deleteTeam_Users(id int64) *conf.ApiError{
	Query, err := NewConnection.Prepare("UPDATE users SET team='0' WHERE team=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}
