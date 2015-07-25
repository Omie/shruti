package providers

import (
	"log"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/models/provider"

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
	provider_name := ws.PathParameter("provider_name", "provider name").DataType("string")

	ws.Route(ws.GET("/{provider_name}").To(getProvider).
		Doc("get provider").
		Param(provider_name).
		Operation("getProvider"))

	ws.Route(ws.POST("/{provider_name}").To(createProvider).
		Doc("create provider").
		Param(provider_name).
		Operation("createProvider").
		Reads(provider.Provider{}))

	ws.Route(ws.DELETE("/{provider_name}").To(deleteProvider).
		Doc("delete provider").
		Param(provider_name).
		Operation("deleteProvider"))

	ws.Route(ws.PUT("/{provider_name}").To(updateProvider).
		Doc("update provider").
		Param(provider_name).
		Operation("updateProvider").
		Reads(provider.Provider{}))

	ws.Route(ws.GET("/").To(getAllProviders).
		Doc("get all providers").
		Operation("getAllProviders"))
}
