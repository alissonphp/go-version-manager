package main

import (
	"github.com/alisson/go-version-manager/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
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
	http.HandleFunc("/summary", controllers.Index)
	http.HandleFunc("/upload", controllers.Upload)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
