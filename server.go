package main

import (
	"flag"
	"fmt"
	"github.com/alisson/go-version-manager/controllers"
	"github.com/alisson/go-version-manager/utilities"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"time"
)

func UpServer()  {
	var dir string
	flag.StringVar(&dir, "dir", "./download", "the directory to serve plugins files")
	flag.Parse()


	r := mux.NewRouter()
	r.HandleFunc("/summary", controllers.Index).Methods("GET")
	r.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir(dir))))
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	p := r.PathPrefix("/plugin").Subrouter()
		p.HandleFunc("/{id}/{os}", controllers.Plugins).Methods("GET")

	s := r.PathPrefix("/upload").Subrouter()
		s.HandleFunc("/", controllers.Upload).Methods("POST")
		s.Use(utilities.MimeTypeChecker)

	srv := &http.Server{
		Handler:      r,
		Addr:         checkEnv(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func checkEnv() string {
	var server = ""
	if os.Getenv("HOST") != "" && os.Getenv("PORT") != "" {
		server = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
		return server
	}

	err := os.Setenv("HOST", "0.0.0.0")
	if err != nil {
		return ""
	}
	err = os.Setenv("PORT", "8000")
	if err != nil {
		return ""
	}

	return checkEnv()
}
