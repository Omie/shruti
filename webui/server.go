package webui

import (
	//"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Serves the homepage with the question form
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html", "./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

// Serves the /assets/ directory for js, html, css assets
func assetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	static := vars["path"]
	http.ServeFile(w, r, filepath.Join("./assets", static))
}

// Handles the actions such as
// mute, unmute, force speak etc
func actionHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	action := vars["action"]

	action = strings.TrimSpace(action)
	if action == "" {
		w.Write([]byte("Invalid action"))
	}
	w.Write([]byte(action))
}

func StartHTTPServer(host string, port string) error {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", indexHandler)

	r.HandleFunc(`/assets/{path:[a-zA-Z0-9=\-\/\.\_]+}`, assetsHandler)

	r.HandleFunc("/action/{action}", actionHandler).Methods("GET")

	http.Handle("/", r)

	bind := fmt.Sprintf("%s:%s", host, port)
	log.Println("listening on: ", bind)
	return http.ListenAndServe(bind, nil)
}
