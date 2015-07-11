package providers

import (
	"log"
	"net/http"

	"github.com/omie/shruti/api"

	"github.com/emicklei/go-restful"
)

func getProvider(request *restful.Request, response *restful.Response) {
	log.Println("--- getProvider")
	api.GetHandler(response, "{}", nil, http.StatusNoContent)
}
