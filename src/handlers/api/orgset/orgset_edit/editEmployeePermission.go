package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"log"
	"forcamp/src/orgset/employees"
)

func getEditEmployeePermissionPostValues(r *http.Request) (string, int64, string, string, *conf.ApiError){
	Token := r.PostFormValue("token")
	Login := strings.ToLower(r.PostFormValue("login"))
	Value := strings.ToLower(r.PostFormValue("value"))
	CatId, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		log.Print(err)
		return "", 0, "", "", conf.ErrCategoryIdNotINT
	}
	return Login, CatId, Value, Token, nil
}

func EditEmployeePermissionHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		login, catId, value, token, APIerr := getEditEmployeePermissionPostValues(r)
		if APIerr != nil {
			conf.PrintError(APIerr, w)
		} else {
			employees.EditEmployeePermission(token, login, catId, value, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditEmployeePermission(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.permission.edit", EditEmployeePermissionHandler)
}

