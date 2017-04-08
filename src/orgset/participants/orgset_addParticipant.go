package participants

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/tealeg/xlsx"
	"database/sql"
)

type AddParticipant_Success struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Login string `json:"login"`
}

func AddParticipant(token string, participant Participant, ResponseWriter http.ResponseWriter) bool {
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		if checkAddParticipantData(participant, ResponseWriter, NewConnection) {
			Resp, APIerr := addParticipantRequest(participant, Organization, Connection, NewConnection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		}
	}
	return true
}

func addParticipantRequest(participant Participant, organization string, Connection *sql.DB, NewConnection *sql.DB) (AddParticipant_Success, *conf.ApiError){
	Password, Hash := orgset.GeneratePassword()
	login, APIerr := addParticipant_Main(organization, Connection, Hash)
	if APIerr != nil {
		return AddParticipant_Success{}, APIerr
	}
	participant.Login = login
	APIerr = addParticipant_Organization(participant, NewConnection)
	if APIerr != nil {
		return AddParticipant_Success{}, APIerr
	}
	APIerr = addParticipant_Excel(participant, organization, Password, NewConnection)
	if APIerr != nil {
		return AddParticipant_Success{}, APIerr
	}
	return AddParticipant_Success{200, "success", login}, nil
}

func addParticipant_Main(organization string, Connection *sql.DB, hash string) (string, *conf.ApiError){
	Query, err := Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(hash, organization)
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	ID, err := Resp.LastInsertId()
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = Connection.Prepare("UPDATE users SET login=? WHERE id=?")
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	login := "participant_"+strconv.FormatInt(ID, 10)
	_, err = Query.Exec(login, ID)
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return login, nil
}

func addParticipant_Organization(participant Participant, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("INSERT INTO users VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(participant.Login, participant.Name, participant.Surname, participant.Middlename, participant.Team, 0, participant.Sex, "default.jpg")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = Connection.Prepare("INSERT INTO participants(login) VALUES(?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(participant.Login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}

func addParticipant_Excel(participant Participant, organization string, password string, Connection *sql.DB) *conf.ApiError{
	teamName, APIerr := getTeamNameById(participant.Team, Connection)
	if APIerr != nil {
		return APIerr
	}
	excelFilePath := conf.FOLDER_PARTICIPANTS+"/"+organization+".xlsx"
	xlFile, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		log.Print(err)
		return conf.ErrOpenExcelFile
	}
	sheet := xlFile.Sheets[0]
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = participant.Name
	cell = row.AddCell()
	cell.Value = participant.Surname
	cell = row.AddCell()
	cell.Value = participant.Middlename
	cell = row.AddCell()
	cell.Value = teamName
	cell = row.AddCell()
	cell.Value = participant.Login
	cell = row.AddCell()
	cell.Value = password
	err = xlFile.Save(excelFilePath)
	if err != nil {
		log.Print(err)
		return conf.ErrSaveExcelFile
	}
	return nil
}

func getTeamNameById(id int64, Connection *sql.DB) (string, *conf.ApiError){
	if id == 0{
		return "отуствует", nil
	} else {
		Query, err := Connection.Query("SELECT name FROM teams WHERE id=?", id)
		if err != nil {
			log.Print(err)
			return "", conf.ErrDatabaseQueryFailed
		}
		defer Query.Close()
		var name string
		for Query.Next(){
			err = Query.Scan(&name)
			if err != nil {
				log.Print(err)
				return "", conf.ErrDatabaseQueryFailed
			}
		}
		return name, nil
	}

}

func checkAddParticipantData(participant Participant, w http.ResponseWriter, Connection *sql.DB) bool {
	if len(participant.Name) > 0 {
		if len(participant.Surname) > 0 {
			if len(participant.Middlename) > 0 {
				if participant.Sex == 0 || participant.Sex == 1 {
					if orgset.CheckTeamID(participant.Team, w, Connection) {
						return true
					} else {
						return false
					}
				} else {
					return conf.PrintError(conf.ErrParticipantSexIncorrect, w)
				}
			} else {
				return conf.PrintError(conf.ErrParticipantMiddlenameEmpty, w)
			}
		} else {
			return conf.PrintError(conf.ErrParticipantSurnameEmpty, w)
		}
	} else {
		return conf.PrintError(conf.ErrParticipantNameEmpty, w)
	}
}
