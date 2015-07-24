package notifications

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/lib/db"
	"github.com/omie/shruti/lib/pusher"
	"github.com/omie/shruti/models/notification"

	"github.com/emicklei/go-restful"
)

func getNotification(request *restful.Request, response *restful.Response) {
	log.Println("--- getNotification")
	_id := request.PathParameter("id")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		api.Error(response, err)
		return
	}
	n, err := notification.GetSingle(db.DB, int(id))

	api.GetHandler(response, n, err, http.StatusNoContent)
}

func getNotificationsSince(request *restful.Request, response *restful.Response) {
	log.Println("--- getNotificationsSince")
	_since := request.PathParameter("since")
	since, err := strconv.ParseInt(_since, 10, 64)
	if err != nil {
		api.Error(response, err)
		return
	}
	tm := time.Unix(since, 0).In(time.UTC)
	n, err := notification.GetSince(db.DB, tm)

	api.GetHandler(response, n, err, http.StatusNoContent)
}

func getNotificationsUnheard(request *restful.Request, response *restful.Response) {
	log.Println("--- getNotificationsUnheard")
	n, err := notification.GetUnheard(db.DB)

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

	// response cares about err during creation, do not modify it
	// for supplimentary action
	// Let it die on its own
	n, err2 := notification.GetSingle(db.DB, newNotification.Id)
	if err2 == nil {
		newNotification = n
		go pusher.Push(n)
	}

	api.CreateHandler(response, newNotification, err)
}

func markAsHeard(request *restful.Request, response *restful.Response) {
	log.Println("--- markAsHeard")

	ids := request.PathParameter("ids")
	var unheardIds []int

	err := json.Unmarshal([]byte(ids), &unheardIds)
	if err != nil {
		api.Error(response, err)
		return
	}
	err = notification.MarkHeard(db.DB, unheardIds)
	api.PutHandler(response, unheardIds, err)
}
