package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func HrisDataRouter(r *gin.Engine) {
	HrisDataEndpoint := "/api/hris/"
	r.GET(HrisDataEndpoint+"employee/", controllers.EmployeeDataFindByParams)
	r.GET(HrisDataEndpoint+"/grade/", controllers.GradeDataFindByParams)
}
