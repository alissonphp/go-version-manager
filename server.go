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

// @title gmkernel plugins sync
// @version 0.1
// @description Service to upload and syncronize plugins to gmkernel flow
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @contact.name Pulse | Software Engineer
// @contact.url https://engenharia.pulse.io
// @contact.email engenharia@grupomateus.com.br

// UpServer @host localhost:8000
// @BasePath /
func UpServer()  {
	var dir string
	flag.StringVar(&dir, "dir", "./download", "the directory to serve plugins files")
	flag.Parse()


	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir(dir))))
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

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

	err := os.Setenv("HOST", "localhost")
	if err != nil {
		return ""
	}
	err = os.Setenv("PORT", "8000")
	if err != nil {
		return ""
	}

	return checkEnv()
}
