package settings

import (
	"log"
	"net/http"

	"github.com/omie/shruti/api"

	"github.com/emicklei/go-restful"
)

func getAllSettings(request *restful.Request, response *restful.Response) {
	log.Println("--- getAllSettings")
	api.GetHandler(response, "{}", nil, http.StatusNoContent)
}
