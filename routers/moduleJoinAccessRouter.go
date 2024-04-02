package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func ModuleJoinAccessRouter(r *gin.Engine) {
	ModuleJoinAccssedEndpoint := "/api/moduleJoinAccess"
	r.GET(ModuleJoinAccssedEndpoint+"/paging", controllers.ModuleMeatadataJoinAccessDataFindPaging)

}
