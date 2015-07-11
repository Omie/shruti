package notifications

import (
	"log"
	"net/http"

	"github.com/omie/shruti/api"

	"github.com/emicklei/go-restful"
)

func getNotificationsSince(request *restful.Request, response *restful.Response) {
	log.Println("--- getNotificationsSince")
	api.GetHandler(response, "{}", nil, http.StatusNoContent)
}
