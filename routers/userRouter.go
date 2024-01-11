package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func UserDataRouter(r *gin.Engine) {
	UserDataEndpoint := "/api/user/data"
	r.POST(UserDataEndpoint+"/create", controllers.UserDataCreate)
	r.POST(UserDataEndpoint+"/login", controllers.UserDataLogin)
	r.GET(UserDataEndpoint+"/:id", controllers.UserDataFindById)
	r.GET(UserDataEndpoint+"/all", controllers.UserDataFindAll)
	r.PUT(UserDataEndpoint+"/update/:id", controllers.UserDataUpdate)
	r.PUT(UserDataEndpoint+"/upsert/:id", controllers.UserDataUpsert)
	r.DELETE(UserDataEndpoint+"/delete/:id", controllers.UserDataDelete)
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
