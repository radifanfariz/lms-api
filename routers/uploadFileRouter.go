package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func UploadFileRouter(r *gin.Engine) {
	UploadFileEndpoint := "/api/upload"
	r.POST(UploadFileEndpoint+"/create", controllers.UploadFileCreate)
	// r.POST(UserDataEndpoint+"/login", controllers.UserDataLogin)
	// r.POST(UserDataEndpoint+"/login/through-portal", controllers.UserDataLoginThroughPortal)
	// r.GET(UserDataEndpoint+"/:id", controllers.UserDataFindById)
	// r.GET(UserDataEndpoint+"/all", controllers.UserDataFindAll)
	// r.PUT(UserDataEndpoint+"/update/:id", controllers.UserDataUpdate)
	// r.PUT(UserDataEndpoint+"/upsert/:id", controllers.UserDataUpsert)
	// r.DELETE(UserDataEndpoint+"/delete/:id", controllers.UserDataDelete)
}
