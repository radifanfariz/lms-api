package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func ModuleJoinAccessRouter(r *gin.Engine) {
	ModuleMetadataJoinAccessEndpoint := "/api/moduleMetadataJoinAccessData"
	ModuleDataJoinAccessEndpoint := "/api/moduleDataJoinAccessData"
	r.GET(ModuleMetadataJoinAccessEndpoint+"/paging", controllers.ModuleMetadataJoinAccessDataFindPaging)
	r.GET(ModuleMetadataJoinAccessEndpoint+"/:id", controllers.ModuleMetadataJoinAccessDataFindByIdWithParams)
	r.GET(ModuleDataJoinAccessEndpoint+"/paging", controllers.ModuleDataJoinAccessDataFindPaging)
	r.GET(ModuleDataJoinAccessEndpoint+"/:id", controllers.ModuleDataJoinAccessDataFindByIdWithParams)

}
