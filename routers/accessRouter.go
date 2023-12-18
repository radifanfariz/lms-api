package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func AccessDataRouter(r *gin.Engine) {
	AccessDataEndpoint := "/api/access/data"
	r.POST(AccessDataEndpoint+"/create", controllers.AccessDataCreate)
	r.GET(AccessDataEndpoint+"/:id", controllers.AccessDataFindById)
	r.GET(AccessDataEndpoint+"/all", controllers.AccessDataFindAll)
	r.PUT(AccessDataEndpoint+"/update/:id", controllers.AccessDataUpdate)
	r.PUT(AccessDataEndpoint+"/upsert/:id", controllers.AccessDataUpsert)
	r.DELETE(AccessDataEndpoint+"/delete/:id", controllers.AccessDataDelete)
}
