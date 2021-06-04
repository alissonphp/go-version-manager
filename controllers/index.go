package controllers

import (
	"fmt"
	"net/http"
)
// Index godoc
// @Summary List all plugins
// @Tags sync
// @Description Retrieve list plugin with all versions
// @Produce json
// @Success 200 {array} Uploader
// @Router / [get]
func Index(w http.ResponseWriter, r *http.Request)  {
	_, err := fmt.Fprintf(w, "summary handler")
	if err != nil {
		return
	}
}
