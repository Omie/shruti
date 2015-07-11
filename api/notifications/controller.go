package notifications

import (
	"log"

	"github.com/omie/shruti/api"

	"github.com/emicklei/go-restful"
)

func init() {
	log.Println("--- initializing notifications resource")
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Path("/notifications").
		Doc("notifications api")

	initWebResource(ws)
	api.AddWebService(ws)
}

func initWebResource(ws *restful.WebService) {
	since := ws.PathParameter("since", "get all notifications since given timestamp").DataType("string")

	ws.Route(ws.GET("/{since:[0-9]+}").To(getNotificationsSince).
		Doc("get notifications").
		Param(since).
		Operation("getNotificationsSince"))
}
