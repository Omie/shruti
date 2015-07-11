package main

import (
	"fmt"
	"log"
	"net/http"
	//"os"

	"github.com/omie/shruti/api"
	_ "github.com/omie/shruti/api/notifications"
	_ "github.com/omie/shruti/api/providers"
	_ "github.com/omie/shruti/api/settings"
)

// var responseQueue *TTSResponseQueue

func main() {

	/*
		responseQueue = NewTTSResponseQueue()

		accessKey := os.Getenv("SHRUTI_ACCESSKEY")
		secretKey := os.Getenv("SHRUTI_SECRETKEY")
		if accessKey == "" || secretKey == "" {
			log.Println("main: environment variables not set")
			return
		}

		InitIvona(accessKey, secretKey)
	*/

	api.InitSwagger(fmt.Sprintf("http://%s:%d", "127.0.0.1", 9574))

	// go RunResponseWorker()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 9574),
		Handler: api.Container,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server", err)
	}

}
