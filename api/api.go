package api

import (
	"database/sql"
	"log"
	"net/http"
	"reflect"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

var Container *restful.Container
var Cors restful.CrossOriginResourceSharing

func init() {
	log.Println("--- initializing api")
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Path("/api").
		Doc("shruti api")

	Cors = restful.CrossOriginResourceSharing{
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "HEAD"},
		ExposeHeaders:  []string{restful.HEADER_AccessControlAllowOrigin, restful.HEADER_AccessControlAllowMethods},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		CookiesAllowed: false,
		Container:      Container,
	}

	initWebResource(ws)

	Container = restful.NewContainer()
	Container.Filter(Cors.Filter)
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
		Error(response, err)
		return
	}

	if json == nil || (reflect.TypeOf(json).Kind() == reflect.Ptr && !reflect.ValueOf(json).Elem().IsValid()) {
		response.WriteErrorString(notFoundCode, "")
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteAsJson(json)
}

func Error(response *restful.Response, err error) {
	response.AddHeader("Content-Type", "text/plain")
	if err == nil {
		log.Println("invalid controller.Error() call: error must not be nil!")
		response.WriteHeader(http.StatusInternalServerError)
	} else if err == sql.ErrNoRows {
		log.Println("no rows returned")
		response.WriteErrorString(http.StatusNotFound, "not found")
	} else {
		log.Println("internal server error:", err)
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func CreateHandler(response *restful.Response, respData interface{}, err error) {
	if err != nil {
		Error(response, err)
		return
	}

	if reflect.TypeOf(respData).Kind() == reflect.Ptr && !reflect.ValueOf(respData).Elem().IsValid() {
		response.WriteErrorString(http.StatusNotFound, "not found")
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteAsJson(respData)
}

func DeleteHandler(response *restful.Response, err error) {
	if err != nil {
		Error(response, err)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func PutHandler(response *restful.Response, respData interface{}, err error) {
	if err != nil {
		Error(response, err)
		return
	}

	if reflect.TypeOf(respData).Kind() == reflect.Ptr && !reflect.ValueOf(respData).Elem().IsValid() {
		response.WriteErrorString(http.StatusNotFound, "not found")
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteAsJson(respData)
}
