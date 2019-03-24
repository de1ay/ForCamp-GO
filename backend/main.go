/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package main

import (
	"wplay/conf"
	"wplay/src"
	"wplay/src/handlers"
	"wplay/src/handlers/api/apanel"
	"wplay/src/handlers/api/emotional_marks"
	"wplay/src/handlers/api/events"
	"wplay/src/handlers/api/marks"
	"wplay/src/handlers/api/orgset/orgset_add"
	"wplay/src/handlers/api/orgset/orgset_delete"
	"wplay/src/handlers/api/orgset/orgset_edit"
	"wplay/src/handlers/api/orgset/orgset_get"
	"wplay/src/handlers/api/users/users_edit"
	"wplay/src/handlers/api/users/users_get"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Domains routing
	Router := mux.NewRouter()
	// WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	// MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	APISite := Router.NewRoute().Subrouter();

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
	// handlers.HandleMainSite(MainSite)
	// handlers.HandleExit(WWWSite)
	// handlers.HandleExit(MainSite)

	// Database: "wplay"
	src.Connection = src.Connect()

	// Server
	go handlers.HandleTLS(Router)
	http.ListenAndServe(conf.SERVER_PORT, Router)
}
