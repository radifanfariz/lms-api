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
	r.DELETE(UserMetadataEndpoint+"/delete/:id", controllers.UserDataDelete)
}
