/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"database/sql"
	"strconv"
)

type Permission struct {
	Id    int64 `json:"id"`
	Value string `json:"value"`
}

type Employee struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Middlename  string `json:"middlename"`
	Sex         int `json:"sex"`
	Team        int64 `json:"team"`
	Post        string `json:"post"`
	Permissions []Permission `json:"permissions"`
}

type getEmployees_Success struct {
	Employees []Employee `json:"employees"`
}

func GetEmployees(token string, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			rawResp, APIerr := getEmployees_Request()
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func getEmployees_Request() (getEmployees_Success, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT login,name,surname,middlename,sex,team FROM users WHERE access='1'")
	if err != nil {
		log.Print(err)
		return getEmployees_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getEmployeesFromResponse(Query)
}

func getEmployeesFromResponse(rows *sql.Rows) (getEmployees_Success, *conf.ApiResponse) {
	defer rows.Close()
	Permissions, Posts, APIerr := getPermissionsAndPosts()
	if APIerr != nil {
		return getEmployees_Success{}, APIerr
	}
	var (
		employees []Employee
		employee Employee
	)
	for rows.Next() {
		err := rows.Scan(&employee.Login, &employee.Name, &employee.Surname, &employee.Middlename, &employee.Sex, &employee.Team)
		if err != nil {
			log.Print(err)
			return getEmployees_Success{}, conf.ErrDatabaseQueryFailed
		}
		employee.Permissions = Permissions[employee.Login]
		employee.Post = Posts[employee.Login]
		employees = append(employees, Employee{employee.Login, employee.Name, employee.Surname, employee.Middlename, employee.Sex, employee.Team, employee.Post, employee.Permissions})
	}
	if employees == nil {
		return getEmployees_Success{make([]Employee, 0)}, nil
	}
	return getEmployees_Success{employees}, nil
}

func getPermissionsAndPosts() (map[string][]Permission, map[string]string, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT * FROM employees")
	if err != nil {
		log.Print(err)
		return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		log.Print(err)
		return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getPermissionsIfNoCategories(Query)
	}
	return getPermissionsIfCategories(Query, CategoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) (map[string][]Permission, map[string]string, *conf.ApiResponse) {
	defer rows.Close()
	var (
		login string
		post string
		Permissions = make(map[string][]Permission)
		Posts = make(map[string]string)
	)
	for rows.Next() {
		err := rows.Scan(&login, &post)
		if err != nil {
			log.Print(err)
			return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
		}
		Permissions[login] = make([]Permission, 0)
		Posts[login] = post
	}
	return Permissions, Posts, nil
}

func getPermissionsIfCategories(rows *sql.Rows, CategoriesIDs []string) (map[string][]Permission, map[string]string,*conf.ApiResponse) {
	CategoriesIDs = CategoriesIDs[2:]
	var (
		rawResult = make([][]byte, len(CategoriesIDs) + 2)
		Result = make([]interface{}, len(CategoriesIDs) + 2)
		Permissions = make(map[string][]Permission)
		Values = make([]string, len(CategoriesIDs) + 2)
		Posts = make(map[string]string)
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(Result...)
		if err != nil {
			return make(map[string][]Permission), make(map[string]string),conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				Result[i] = "\\N"
			} else {
				Values[i] = string(raw)
			}
		}
		Login := Values[0]
		Post := Values[1]
		Posts[Login] = Post
		Values = Values[2:]
		for i := 0; i < len(Values); i++ {
			id, err := strconv.ParseInt(CategoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return make(map[string][]Permission), make(map[string]string),conf.ErrConvertStringToInt
			}
			Permissions[Login] = append(Permissions[Login], Permission{id, Values[i]})
		}
		Values = make([]string, len(CategoriesIDs) + 2)
	}
	return Permissions, Posts, nil
}