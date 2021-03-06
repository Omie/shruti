package notifications

import (
	"log"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/models/notification"

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
	ids := ws.PathParameter("ids", "ids of notifications").DataType("string")
	id := ws.PathParameter("id", "id of notification").DataType("int")

	ws.Route(ws.GET("/unheard").To(getNotificationsUnheard).
		Doc("get unheard notifications").
		Operation("getNotificationsUnheard"))

	ws.Route(ws.GET("/since/{since:[0-9]+}").To(getNotificationsSince).
		Doc("get notifications").
		Param(since).
		Operation("getNotificationsSince"))

	ws.Route(ws.GET("/{id:[0-9]+}").To(getNotification).
		Doc("get notification").
		Param(id).
		Operation("getNotification"))

	ws.Route(ws.POST("/").To(pushNotification).
		Doc("push notification").
		Operation("pushNotification").
		Reads(notification.Notification{}))

	ws.Route(ws.PUT("/mark-heard/{ids}").To(markAsHeard).
		Doc("mark notification as heard").
		Param(ids).
		Operation("markAsHeard"))
}
