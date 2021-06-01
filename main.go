package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Uploader struct {
	Version string `json:"version"`
	Os string `json:"os"`
	PluginId string `json:"plugin_id"`
}

func summaryHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "summary handler")
}

func uploadHandler(w http.ResponseWriter, r *http.Request)  {

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
	w.Write(response)
	return
}

func upServer()  {
	http.HandleFunc("/summary", summaryHandler)
	http.HandleFunc("/upload", uploadHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	upServer()
}
