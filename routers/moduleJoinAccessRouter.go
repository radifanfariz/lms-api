package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func ModuleJoinAccessRouter(r *gin.Engine) {
	ModuleMetadataJoinAccessEndpoint := "/api/moduleMetadataJoinAccessData"
	ModuleDataJoinAccessEndpoint := "/api/moduleDataJoinAccessData"
	r.GET(ModuleMetadataJoinAccessEndpoint+"/paging", controllers.ModuleMetadataJoinAccessDataFindPaging)
	r.GET(ModuleMetadataJoinAccessEndpoint+"/first/:id", controllers.ModuleMetadataJoinAccessDataFindFirst)
	r.GET(ModuleDataJoinAccessEndpoint+"/first/:id", controllers.ModuleDataJoinAccessDataFindFirst)

}
