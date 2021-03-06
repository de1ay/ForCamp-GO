/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package handlers

import (
	"net/http"

	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/handlers/mainSite"

	"github.com/gorilla/mux"
)

func getFolder_Scripts() http.Dir {
	ScriptsFolder := http.Dir(conf.FOLDER_SCRIPTS)
	return ScriptsFolder
}

func getFolder_Css() http.Dir {
	CssFolder := http.Dir(conf.FOLDER_CSS)
	return CssFolder
}

func getFolder_Images() http.Dir {
	ImagesFolder := http.Dir(conf.FOLDER_IMAGES)
	return ImagesFolder
}

func folderHandler_Scripts() http.Handler {
	ScriptsFolder := getFolder_Scripts()
	ScriptsFileServer := http.FileServer(ScriptsFolder)
	ScriptsHandler := http.StripPrefix("/scripts", ScriptsFileServer)
	return ScriptsHandler
}

func folderHandler_css() http.Handler {
	CssFolder := getFolder_Css()
	CssFileServer := http.FileServer(CssFolder)
	CssHandler := http.StripPrefix("/css", CssFileServer)
	return CssHandler
}

func folderHandler_Images() http.Handler {
	ImagesFolder := getFolder_Images()
	ImagesFolderServer := http.FileServer(ImagesFolder)
	ImagesHandler := http.StripPrefix("/images", ImagesFolderServer)
	return ImagesHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_INDEX)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func HandleMainSite(router *mux.Router) {
	// Folders
	ScriptsHandler := folderHandler_Scripts()
	CssHandler := folderHandler_css()
	ImagesHandler := folderHandler_Images()
	router.PathPrefix("/scripts").Handler(ScriptsHandler)
	router.PathPrefix("/css").Handler(CssHandler)
	router.PathPrefix("/images").Handler(ImagesHandler)
	// Pages
	router.HandleFunc("/", mainSite.IndexHandler)
	router.HandleFunc("/orgset", mainSite.OrgSetHandler)
	router.HandleFunc("/marks", mainSite.MarksHandler)
	router.HandleFunc("/general", mainSite.GeneralHandler)
	router.HandleFunc("/profile", mainSite.ProfileHandler)
	router.HandleFunc("/apanel", mainSite.ApanelHandler)
}
