package marks

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"strconv"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/marks"
)

func getEditMarkPostValues(r *http.Request) (string, string, int64, int64, *conf.ApiError){
	token := strings.ToLower(r.PostFormValue("token"))
	category_id, err := strconv.ParseInt(r.PostFormValue("category_id"), 10, 64)
	if err != nil {
		return "", "", 0, 0, conf.ErrCategoryIdNotINT
	}
	reason_id, err := strconv.ParseInt(r.PostFormValue("reason_id"), 10, 64)
	if err != nil {
		return "", "", 0, 0, conf.ErrReasonIncorrect
	}
	participant_login := strings.ToLower(r.PostFormValue("login"))
	return token, participant_login, category_id, reason_id, nil
}

func editMarkHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, participant_login, category_id, reason_id, APIerr := getEditMarkPostValues(r)
		if APIerr != nil {
			conf.PrintError(APIerr, w)
		} else {
			marks.EditMark(token, participant_login, category_id, reason_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditMark(router *mux.Router) {
	router.HandleFunc("/mark.edit", editMarkHandler)
}