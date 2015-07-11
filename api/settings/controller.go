package settings

import (
	"log"

	"github.com/omie/shruti/api"

	"github.com/emicklei/go-restful"
)

func init() {
	log.Println("--- initializing settings resource")
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Path("/settings").
		Doc("settings api")

	initWebResource(ws)
	api.AddWebService(ws)
}

func initWebResource(ws *restful.WebService) {

	ws.Route(ws.GET("/").To(getAllSettings).
		Doc("get all settings").
		Operation("getAllSettings"))
}
