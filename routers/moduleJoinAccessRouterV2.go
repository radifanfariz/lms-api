package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func ModuleJoinAccessRouterV2(r *gin.Engine) {
	ModuleMetadataJoinAccessEndpoint := "/api/moduleMetadataJoinAccessData/v2"
	ModuleDataJoinAccessEndpoint := "/api/moduleDataJoinAccessData/v2"
	r.GET(ModuleMetadataJoinAccessEndpoint+"/paging", controllers.ModuleMetadataJoinAccessDataFindPagingV2)
	r.GET(ModuleMetadataJoinAccessEndpoint+"/:id", controllers.ModuleMetadataJoinAccessDataFindByIdWithParamsV2)
	r.GET(ModuleDataJoinAccessEndpoint+"/paging", controllers.ModuleDataJoinAccessDataFindPagingV2)
	r.GET(ModuleDataJoinAccessEndpoint+"/:id", controllers.ModuleDataJoinAccessDataFindByIdWithParamsV2)

}
