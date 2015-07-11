package api

import (
	"log"
	"net/http"
	"reflect"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
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

	ws.Route(ws.GET("/{foo}").To(getFoo).
		Doc("get a foo").
		Param(foo).
		Operation("getFoo"))
}

func InitSwagger(path string) {
	config := swagger.Config{
		WebServices:     Container.RegisteredWebServices(),
		WebServicesUrl:  path,
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "./swagger",
	}
	swagger.RegisterSwaggerService(config, Container)
}

func getFoo(request *restful.Request, response *restful.Response) {
	foo := request.PathParameter("foo")
	log.Println("getFoo: you sent: ", foo)
	GetHandler(response, "{}", nil, http.StatusNoContent)
}

func AddWebService(ws *restful.WebService) {
	Container.Add(ws)
}

func GetHandler(response *restful.Response, json interface{}, err error, notFoundCode int) {
	if err != nil {
		return
	}

	if json == nil || (reflect.TypeOf(json).Kind() == reflect.Ptr && !reflect.ValueOf(json).Elem().IsValid()) {
		response.WriteErrorString(notFoundCode, "")
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteAsJson(json)
}
