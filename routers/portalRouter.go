package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func PortalDataRouter(r *gin.Engine) {
	PortalDataEndpoint := "/api/portal"
	r.GET(PortalDataEndpoint+"/company/all", controllers.CompanyDataFindAll)
}
