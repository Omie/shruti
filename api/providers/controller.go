package providers

import (
	"log"

	"github.com/omie/shruti/api"

	"github.com/emicklei/go-restful"
)

func init() {
	log.Println("--- initializing providers resource")
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Path("/providers").
		Doc("providers api")

	initWebResource(ws)
	api.AddWebService(ws)
}

func initWebResource(ws *restful.WebService) {
	provider_name := ws.PathParameter("provider_name", "get provider").DataType("string")

	ws.Route(ws.GET("/{provider_name:[a-z]+}").To(getProvider).
		Doc("get provider").
		Param(provider_name).
		Operation("getProvider"))
}
