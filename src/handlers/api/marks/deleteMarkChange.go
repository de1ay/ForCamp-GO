package marks

import (
	"net/http"
	"forcamp/conf"
	"strings"
	"strconv"
	"log"
	"forcamp/src"
	"forcamp/src/marks"
	"github.com/gorilla/mux"
)

func getDeleteMarkChangePostData(r *http.Request) (string, int64, *conf.ApiError){
	token := strings.TrimSpace(r.PostFormValue("token"))
	id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("id")), 10, 64)
	if err != nil {
		log.Print(err)
		return "", 0, conf.ErrIDisNotINT
	}
	return token, id, nil
}

func deleteMarkChangeHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, id, APIerr := getDeleteMarkChangePostData(r)
		if APIerr != nil {
			conf.PrintError(APIerr, w)
		} else {
			marks.DeleteMarkChange(token, id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleDeleteMarkChange(router *mux.Router) {
	router.HandleFunc("/mark.change.delete", deleteMarkChangeHandler)
}