package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/alisson/go-version-manager/docs"
	"github.com/alisson/go-version-manager/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Uploader struct {
	Version string `json:"version"`
	Os string `json:"os"`
	PluginId string `json:"plugin_id"`
	Path string `json:"path"`
}

type FileHandler func(string) string

func handler(r *http.Request) *Uploader {
	infos := &Uploader{
		Version: r.Header.Get("--VERSION"),
		Os: r.Header.Get("--OS"),
		PluginId: r.Header.Get("--PLUGIN_ID"),
	}

	if err := utils.MakeDir(fmt.Sprintf("./download/plugins/%s/%s/%s", infos.PluginId, infos.Version, infos.Os)); err != nil {
		panic(err)
	}

	infos.Path = fmt.Sprintf("./download/plugins/%s/%s/%s/", infos.PluginId, infos.Version, infos.Os)

	return infos
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
// @Router /upload/ [post]
func Upload(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(8 << 20)

	file, h, _ := r.FormFile("file")
	defer file.Close()

	dst, err := os.Create(handler(r).Path + h.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	response, _ := json.Marshal(handler(r))
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJsonFile(handler(r), h.Filename)
}

func writeJsonFile(u *Uploader, name string)  {

	var filePath = u.Path + "infos.json"
	u.Path = u.Path + name

	file, _ := json.MarshalIndent(u, "", " ")
	_ = ioutil.WriteFile(filePath, file, 0644)

}