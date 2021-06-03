package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "summary handler")
}
