package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func ModuleMetadataJoinAccssedDataRouter(r *gin.Engine) {
	ModuleMetadataJoinAccssedDataEndpoint := "/api/moduleJoinAccess/data"
	r.GET(ModuleMetadataJoinAccssedDataEndpoint+"/paging", controllers.ModuleMeatadataJoinAccssedDataFindPaging)

}
