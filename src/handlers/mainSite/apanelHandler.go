/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package mainSite

import (
	"html/template"
	"net/http"
	"net/url"

	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/orgset"
	"nullteam.info/wplay/demo/src/tools"
)

type apanelTemplateData struct {
	Token string
}

func ApanelHandler(responseWriter http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(responseWriter)
		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(responseWriter, r, "https://"+conf.MAIN_SITE_DOMAIN+"/exit", http.StatusTemporaryRedirect)
		}
		token.Value, err = url.QueryUnescape(token.Value)
		if err == nil && tools.CheckToken(token.Value) {
			isAdmin := isUserAdmin(token.Value)
			if isAdmin != nil {
				http.Redirect(responseWriter, r, "https://"+conf.MAIN_SITE_DOMAIN+"/profile", http.StatusTemporaryRedirect)
				return
			}
			apanelHTML, err := template.New(conf.FILE_APANEL).ParseFiles(conf.FILE_APANEL)
			if err != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)
				return
			}
			atd := getApanelTemplateData(token.Value, r)
			apanelHTML.ExecuteTemplate(responseWriter, "apanel", atd)
		} else {
			http.Redirect(responseWriter, r, "https://"+conf.MAIN_SITE_DOMAIN+"/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(responseWriter, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getApanelTemplateData(token string, r *http.Request) apanelTemplateData {
	var atd apanelTemplateData
	atd.Token = token
	return atd
}

func isUserAdmin(token string) *conf.ApiResponse {
	user_id, apiErr := orgset.GetUserIdByToken(token)
	if apiErr != nil {
		return apiErr
	}
	var admin_status bool
	err := src.Connection.QueryRow("SELECT admin FROM users WHERE id=?", user_id).Scan(&admin_status)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if admin_status {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}
