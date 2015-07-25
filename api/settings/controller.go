package settings

import (
	"log"

	"github.com/omie/shruti/api"
	"github.com/omie/shruti/models/setting"

	"github.com/emicklei/go-restful"
)

func init() {
	log.Println("--- initializing settings resource")
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Path("/settings").
		Doc("settings api")

	initWebResource(ws)
	api.AddWebService(ws)
}

func initWebResource(ws *restful.WebService) {
	key := ws.PathParameter("key", "setting key").DataType("string")

	ws.Route(ws.GET("/{key}").To(getSetting).
		Doc("get setting").
		Param(key).
		Operation("getSetting"))

	ws.Route(ws.POST("/{key}").To(createSetting).
		Doc("create setting").
		Param(key).
		Operation("createSetting").
		Reads(setting.Setting{}))

	ws.Route(ws.DELETE("/{key}").To(deleteSetting).
		Doc("delete setting").
		Param(key).
		Operation("deleteSetting"))

	ws.Route(ws.PUT("/{key}").To(updateSetting).
		Doc("update setting").
		Param(key).
		Operation("updateSetting").
		Reads(setting.Setting{}))

	ws.Route(ws.GET("/").To(getAllSettings).
		Doc("get all settings").
		Operation("getAllSettings"))

}
