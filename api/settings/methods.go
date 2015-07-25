package settings

import (
	"log"
	"net/http"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/lib/db"
	"github.com/omie/shruti/models/setting"

	"github.com/emicklei/go-restful"
)

func getSetting(request *restful.Request, response *restful.Response) {
	log.Println("--- getSetting")
	key := request.PathParameter("key")

	s, err := setting.Get(db.DB, key)

	api.GetHandler(response, s, err, http.StatusNoContent)
}

func getAllSettings(request *restful.Request, response *restful.Response) {
	log.Println("--- getAllSettings")

	s, err := setting.GetAll(db.DB)

	api.GetHandler(response, s, err, http.StatusNoContent)
}

func createSetting(request *restful.Request, response *restful.Response) {
	log.Println("--- createSetting")
	key := request.PathParameter("key")

	newSetting := new(setting.Setting)
	err := request.ReadEntity(newSetting)
	if err != nil {
		api.Error(response, err)
		return
	}
	newSetting.Key = key
	err = newSetting.Insert(db.DB)
	api.CreateHandler(response, newSetting, err)
}

func deleteSetting(request *restful.Request, response *restful.Response) {
	log.Println("--- deleteSetting")
	key := request.PathParameter("key")

	err := setting.Delete(db.DB, key)
	api.DeleteHandler(response, err)
}

func updateSetting(request *restful.Request, response *restful.Response) {
	log.Println("--- updateSetting")
	key := request.PathParameter("key")

	newSetting := new(setting.Setting)
	err := request.ReadEntity(newSetting)
	if err != nil {
		api.Error(response, err)
		return
	}

	newSetting.Key = key
	err = newSetting.Update(db.DB)
	api.PutHandler(response, newSetting, err)
}
