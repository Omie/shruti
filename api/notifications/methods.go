package notifications

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/lib/db"
	"github.com/omie/shruti/models/notification"

	"github.com/emicklei/go-restful"
)

func getNotificationsSince(request *restful.Request, response *restful.Response) {
	log.Println("--- getNotificationsSince")
	_since := request.PathParameter("since")
	since, err := strconv.ParseInt(_since, 10, 64)
	if err != nil {
		api.Error(response, err)
		return
	}
	tm := time.Unix(since, 0)
	n, err := notification.GetSince(db.DB, tm)

	api.GetHandler(response, n, err, http.StatusNoContent)
}

func pushNotification(request *restful.Request, response *restful.Response) {
	log.Println("--- pushNotification")

	newNotification := new(notification.Notification)
	err := request.ReadEntity(newNotification)
	if err != nil {
		api.Error(response, err)
		return
	}
	err = newNotification.Insert(db.DB)
	api.CreateHandler(response, newNotification, err)
}
