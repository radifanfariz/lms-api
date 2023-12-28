package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func UserDataRouter(r *gin.Engine) {
	UserMetadataEndpoint := "/api/user/data"
	r.POST(UserMetadataEndpoint+"/create", controllers.UserDataCreate)
	r.GET(UserMetadataEndpoint+"/:id", controllers.UserDataFindById)
	r.GET(UserMetadataEndpoint+"/all", controllers.UserDataFindAll)
	r.PUT(UserMetadataEndpoint+"/update/:id", controllers.UserDataUpdate)
	r.PUT(UserMetadataEndpoint+"/upsert/:id", controllers.UserDataUpsert)
	r.DELETE(UserMetadataEndpoint+"/delete/:id", controllers.UserDataDelete)
}

func UserActionDataRouter(r *gin.Engine) {
	UserActionDataEndpoint := "/api/user/action"
	r.POST(UserActionDataEndpoint+"/create", controllers.UserActionDataCreate)
	r.GET(UserActionDataEndpoint+"/:id", controllers.UserActionDataFindById)
	r.GET(UserActionDataEndpoint+"/all", controllers.UserActionDataFindAll)
	r.PUT(UserActionDataEndpoint+"/update/:id", controllers.UserActionDataUpdate)
	r.PUT(UserActionDataEndpoint+"/upsert/:id", controllers.UserActionDataUpsert)
	r.DELETE(UserActionDataEndpoint+"/delete/:id", controllers.UserActionDataDelete)
}
