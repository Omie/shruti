package providers

import (
	"log"
	"net/http"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/lib/db"
	"github.com/omie/shruti/models/provider"

	"github.com/emicklei/go-restful"
)

func getProvider(request *restful.Request, response *restful.Response) {
	log.Println("--- getProvider")
	provider_name := request.PathParameter("provider_name")

	p, err := provider.Get(db.DB, provider_name)

	api.GetHandler(response, p, err, http.StatusNoContent)
}

func createProvider(request *restful.Request, response *restful.Response) {
	log.Println("--- createProvider")
	provider_name := request.PathParameter("provider_name")

	newProvider := new(provider.Provider)
	err := request.ReadEntity(newProvider)
	if err != nil {
		api.Error(response, err)
	}
	newProvider.Name = provider_name
	err = newProvider.Insert(db.DB)
	api.CreateHandler(response, newProvider, err)
}
