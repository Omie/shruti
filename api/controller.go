package api

import (
	"log"

	"github.com/emicklei/go-restful"
)

var Container *restful.Container

func init() {
	log.Println("--- initializing api")
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Path("/api").
		Doc("shruti api")

	initWebResource(ws)
	Container = restful.NewContainer()
	Container.Add(ws)
}

func initWebResource(ws *restful.WebService) {
	foo := ws.PathParameter("foo", "foo bar").DataType("string")

	ws.Route(ws.GET("/{foo:[a-z]{2}").To(getFoo).
		Doc("get a foo").
		Param(foo).
		Operation("getFoo"))
}
