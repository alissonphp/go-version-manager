package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type LatestVersionByKernelType struct {
	ID string `json:"id"`
	Version string `json:"version"`
	OS string `json:"os"`
	BinaryPath string `json:"binary_path"`
}

// Plugins godoc
// @Summary Get plugin metadata from latest version
// @Tags sync
// @Param id path string true "Plugin id"
// @Param os path string true "Operation system kernel type"
// @Description Retrieve plugin infos
// @Produce json
// @Success 200 string LatestVersionByKernelType
// @Router /plugin/{id}/{os} [get]
func Plugins(w http.ResponseWriter, r *http. Request) {
	id, os := mux.Vars(r)["id"], mux.Vars(r)["os"]
	response := LatestVersionByKernelType{ID: id, OS: os}
	var versions = getSortedPluginVersions(id)
	latest := getLatestPluginVersion(versions)

	response.Version = latest.Number

	for _, i := range latest.Os {
		if i.Name == os {
			response.OS = i.Name
			response.BinaryPath = i.Path
		}
	}
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.Write(res)
}