package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func InsightDataRouter(r *gin.Engine) {
	InsightDataEndPoint := "/api/insights/data"
	r.GET(InsightDataEndPoint+"/total", controllers.TotalInsightsDataFindAll)
	r.GET(InsightDataEndPoint+"/total/enrolledpermonth", controllers.TotalEnrolledPerMonthInsightsDataFindAll)
	r.GET(InsightDataEndPoint+"/enrolled", controllers.EnrolledInsightsDataFindAll)
	r.GET(InsightDataEndPoint+"/user", controllers.UserInsightsDataFindAll)
}
