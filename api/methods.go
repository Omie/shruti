package api

import (
	"net/http"
	"reflect"

	"github.com/emicklei/go-restful"
)

func getFoo(request *restful.Request, response *restful.Response) {
	GetHandler(response, "{}", nil, http.StatusNoContent)
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
