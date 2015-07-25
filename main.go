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
	"github.com/omie/shruti/lib/pusher"

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

	pusherAppId := os.Getenv("SHRUTI_PUSHER_APPID")
	pusherKey := os.Getenv("SHRUTI_PUSHER_KEY")
	pusherSecret := os.Getenv("SHRUTI_PUSHER_SECRET")
	pusherChannel := os.Getenv("SHRUTI_PUSHER_CHANNEL")
	pusherEvent := os.Getenv("SHRUTI_PUSHER_EVENT")
	usePusher := true
	if pusherAppId == "" || pusherKey == "" || pusherSecret == "" ||
		pusherChannel == "" || pusherEvent == "" {
		log.Println("main: pusher variables not set, continuing without")
		usePusher = false
	}
	pusher.InitPusher(pusherAppId, pusherKey, pusherSecret, pusherChannel,
		pusherEvent, usePusher)

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
