package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func DuplicateDataRouter(r *gin.Engine) {
	DuplicateDataEndpoint := "/api/module/duplicate"
	r.POST(DuplicateDataEndpoint+"/:id", controllers.DuplicateLearningModuleAllData)
	r.POST(DuplicateDataEndpoint+"/v2/:id", controllers.DuplicateLearningModuleAllDataV2)
}
