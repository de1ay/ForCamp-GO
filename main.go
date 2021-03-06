/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package main

import (
	"net/http"

	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/handlers"
	"nullteam.info/wplay/demo/src/handlers/api/apanel"
	"nullteam.info/wplay/demo/src/handlers/api/emotional_marks"
	"nullteam.info/wplay/demo/src/handlers/api/events"
	"nullteam.info/wplay/demo/src/handlers/api/marks"
	"nullteam.info/wplay/demo/src/handlers/api/orgset/orgset_add"
	"nullteam.info/wplay/demo/src/handlers/api/orgset/orgset_delete"
	"nullteam.info/wplay/demo/src/handlers/api/orgset/orgset_edit"
	"nullteam.info/wplay/demo/src/handlers/api/orgset/orgset_get"
	"nullteam.info/wplay/demo/src/handlers/api/users/users_edit"
	"nullteam.info/wplay/demo/src/handlers/api/users/users_get"

	"github.com/gorilla/mux"
)

func main() {
	conf.GetEnvVars()
	// Domains routing
	Router := mux.NewRouter()
	// WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	APISite := Router.Host(conf.API_SITE_DOMAIN).Subrouter()

	// Handlers: API site
	// Authorization
	handlers.HandleAuthorizationByLoginAndPassword(APISite)
	handlers.HandleTokenVerification(APISite)
	// Users: GET
	users_get.HandleGetUserLoginByToken(APISite)
	users_get.HandleGetUserData(APISite)
	// Users: EDIT
	users_edit.HandleChangeUserPassword(APISite)
	users_edit.HandleChangeUserAvatar(APISite)
	// OrgSet: GET
	orgset_get.HandleGetTeams(APISite)
	orgset_get.HandleGetOrgSettings(APISite)
	orgset_get.HandleGetCategories(APISite)
	orgset_get.HandleGetParticipants(APISite)
	orgset_get.HandleGetEmployees(APISite)
	orgset_get.HandleGetParticipantsExcel(APISite)
	orgset_get.HandleGetReasons(APISite)
	orgset_get.HandleGetEmployeesExcel(APISite)
	// OrgSet: ADD
	orgset_add.HandleAddTeam(APISite)
	orgset_add.HandleAddCategory(APISite)
	orgset_add.HandleAddParticipant(APISite)
	orgset_add.HandleAddEmployee(APISite)
	orgset_add.HandleAddReason(APISite)
	// OrgSet: EDIT
	orgset_edit.HandleSetOrgSettingsValue(APISite)
	orgset_edit.HandleEditCategory(APISite)
	orgset_edit.HandleEditTeam(APISite)
	orgset_edit.HandleResetParticipantPassword(APISite)
	orgset_edit.HandleEditParticipant(APISite)
	orgset_edit.HandleResetEmployeePassword(APISite)
	orgset_edit.HandleEditEmployee(APISite)
	orgset_edit.HandleEditEmployeePermission(APISite)
	orgset_edit.HandleEditReason(APISite)
	// OrgSet: DELETE
	orgset_delete.HandleDeleteCategory(APISite)
	orgset_delete.HandleDeleteTeam(APISite)
	orgset_delete.HandleDeleteParticipant(APISite)
	orgset_delete.HandleDeleteEmployee(APISite)
	orgset_delete.HandleDeleteReason(APISite)
	// Marks
	marks.HandleEditMark(APISite)
	// Emotional marks
	emotional_marks.HandleSetEmotionalMark(APISite)
	// Events
	events.HandleGetEvents(APISite)
	events.HandleDeleteEvent(APISite)
	// Apanel
	apanel.HandleAddOrganization(APISite)

	// Handlers: Main site
	// handlers.HandleMainSite(WWWSite)
	handlers.HandleMainSite(MainSite)
	// handlers.HandleExit(WWWSite)
	handlers.HandleExit(MainSite)

	// Database: "wplay"
	src.Connection = src.Connect()

	// Server
	go handlers.HandleTLS(Router)
	http.ListenAndServe(conf.SERVER_PORT, Router)
}
