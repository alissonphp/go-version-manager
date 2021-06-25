package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/alisson/go-version-manager/utilities"
	version2 "github.com/hashicorp/go-version"
	"net/http"
	"sort"
	"strings"
)

type Summary struct {
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	ID string `json:"id"`
	Latest Version `json:"latest"`
	Versions []Version `json:"versions"`
}

type Version struct {
	Number string `json:"number"`
	Os []Os `json:"os"`
}

type Os struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Index godoc
// @Summary List all plugins
// @Tags sync
// @Description Retrieve list plugin with all versions
// @Produce json
// @Success 200 {array} Summary
// @Router /summary [get]
func Index(w http.ResponseWriter, r *http.Request)  {

	sum := Summary{}
	for _, p := range utilities.ReadListDir("./download/plugins") {
		var versions = getSortedPluginVersions(p.Name())
		sum.Plugins = append(sum.Plugins, Plugin{
			ID: p.Name(),
			Latest: getLatestPluginVersion(versions),
			Versions: versions,
		})
	}
	res, err := json.Marshal(sum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.Write(res)

}

func getOsList(pluginId string, version string) []Os  {
	var os []Os
	var path = "./download/plugins/" + pluginId + "/" + version
	for _, o := range utilities.ReadListDir(path) {
	   os = append(os, Os {
		Name: o.Name(),
		Path: getBinaryPathFile(path, o.Name()),
	   })
	}
	return os
}

func getBinaryPathFile(path string, os string) string {
	 file := utilities.ReadListDir(path + "/" + os)[0]
	 return fmt.Sprintf("http://localhost:8000%s/%s/%s", strings.Replace(path, "./download", "/download", 1), os, file.Name())
}

func getSortedPluginVersions(pluginId string) []Version {
	var path = "./download/plugins/" + pluginId
	versionsRaw := utilities.ReadListDir(path)
	versions := make([]*version2.Version, len(versionsRaw))

	for i, raw := range versionsRaw {
		v, _ := version2.NewVersion(raw.Name())
		versions[i] = v
	}

	sort.Sort(version2.Collection(versions))
	return getVersionsObject(pluginId, versions)
}

func getVersionsObject(pluginId string, sorted []*version2.Version) []Version  {
	var versions []Version

	for _, v := range sorted {
		versions = append(versions, Version{
			Number: v.String(),
			Os: getOsList(pluginId, v.String()),
		})
	}

	return versions
}

func getLatestPluginVersion(versions []Version) Version  {
	return versions[len(versions)-1]
}