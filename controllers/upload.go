package controllers

import (
	"encoding/json"
	_ "github.com/alisson/go-version-manager/docs"
	"net/http"
)

type Uploader struct {
	Version string `json:"version"`
	Os string `json:"os"`
	PluginId string `json:"plugin_id"`
}

// Upload godoc
// @Summary Upload plugin
// @Tags sync
// @Description Recieve plugins's binary and metadata from Gitlab CI
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "plugin lib (.so, .dll or .app)"
// @Param --VERSION header string true "plugin version, ex. 1.0.1"
// @Param --OS header string true "platform - linux, windows or macos"
// @Param --PLUGIN_ID header string true "id - com.pulse.641.nfe"
// @Success 200 {object} Uploader
// @Router /upload [post]
func Upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		http.Error(w, "METHOD NOT ALOWED", http.StatusMethodNotAllowed)
		return
	}

	infos := &Uploader{
		Version: r.Header.Get("--VERSION"),
		Os: r.Header.Get("--OS"),
		PluginId: r.Header.Get("--PLUGIN_ID"),
	}

	w.Header().Add("Content-type", "application/json")
	response, _ := json.Marshal(infos)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}
