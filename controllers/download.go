package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
// Download godoc
// @Summary Download plugin
// @Tags sync
// @Description Get plugin binary file to override in application plugins path
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "plugin lib (.so, .dll or .app)"
// @Param --VERSION header string true "plugin version, ex. 1.0.1"
// @Param --OS header string true "platform - linux, windows or macos"
// @Param --PLUGIN_ID header string true "id - com.pulse.641.nfe"
// @Success 200 {object} Uploader
// @Router /upload [post]
func Download(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Plugin id :%v / version: %v / platform: %v", vars["pluginId"], vars["version"], vars["platform"])
}
