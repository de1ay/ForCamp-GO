package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
	"forcamp/src/orgset/employees"
)


func GetEmployeesHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		employees.GetEmployees(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetEmployees(router *mux.Router)  {
	router.HandleFunc("/orgset.employees.get", GetEmployeesHandler)
}