package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func ModuleMetadataRouter(r *gin.Engine) {
	ModuleMetadataEndpoint := "/api/module/metadata"
	r.POST(ModuleMetadataEndpoint+"/create", controllers.ModuleMetadataCreate)
	r.GET(ModuleMetadataEndpoint+"/:id", controllers.ModuleMetadataFindById)
	r.GET(ModuleMetadataEndpoint+"/all", controllers.ModuleMetadataFindAll)
	r.PUT(ModuleMetadataEndpoint+"/update/:id", controllers.ModuleMetadataUpdate)
	r.DELETE(ModuleMetadataEndpoint+"/delete/:id", controllers.ModuleMetadataDelete)
}
func ModuleDataRouter(r *gin.Engine) {
	ModuleDataEndpoint := "/api/module/data"
	r.POST(ModuleDataEndpoint+"/create", controllers.ModuleDataCreate)
	r.GET(ModuleDataEndpoint+"/:id", controllers.ModuleDataFindById)
	r.GET(ModuleDataEndpoint+"/all", controllers.ModuleDataFindAll)
	r.PUT(ModuleDataEndpoint+"/update/:id", controllers.ModuleDataUpdate)
	r.DELETE(ModuleDataEndpoint+"/delete/:id", controllers.ModuleDataDelete)
}
