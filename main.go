package main // import "github.com/omie/shruti"

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/omie/shruti/api"
	_ "github.com/omie/shruti/api/notifications"
	_ "github.com/omie/shruti/api/providers"
	_ "github.com/omie/shruti/api/settings"
	"github.com/omie/shruti/lib/db"

	_ "github.com/lib/pq"
)

// var responseQueue *TTSResponseQueue

func main() {

	cString := os.Getenv("SHRUTI_CONN_STRING")
	if cString == "" {
		log.Println("main: database environment variables not set")
		return
	}
	if err := db.InitDB(cString); err != nil {
		log.Println("error initializing database, Abort", err)
		return
	}

	host := os.Getenv("SHRUTI_SERVER_HOST")
	port := os.Getenv("SHRUTI_SERVER_PORT")
	if host == "" || port == "" {
		log.Println("main: host or port not set")
		return
	}

	api.InitSwagger(fmt.Sprintf("http://%s:%s", host, port))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: api.Container,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server", err)
	}

}
