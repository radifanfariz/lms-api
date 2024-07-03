package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func UserDataRouter(r *gin.Engine) {
	UserDataEndpoint := "/api/user/data"
	r.POST(UserDataEndpoint+"/create", controllers.UserDataCreate)
	r.POST(UserDataEndpoint+"/login", controllers.UserDataLogin)
	r.POST(UserDataEndpoint+"/login/through-sso", controllers.UserDataLoginThroughSSO)
	r.POST(UserDataEndpoint+"/login/through-portal", controllers.UserDataLoginThroughPortal)
	r.GET(UserDataEndpoint+"/paging", controllers.UserDataFindByPaging)
	r.GET(UserDataEndpoint+"/:id", controllers.UserDataFindById)
	r.POST(UserDataEndpoint+"/employeeId", controllers.UserDataFindByEmployeeIds)
	r.GET(UserDataEndpoint+"/all", controllers.UserDataFindAll)
	r.GET(UserDataEndpoint+"/params", controllers.UserDataFindByParams)
	r.PUT(UserDataEndpoint+"/update/:id", controllers.UserDataUpdate)
	r.PUT(UserDataEndpoint+"/upsert/:id", controllers.UserDataUpsert)
	r.PUT(UserDataEndpoint+"/upsert/bulk", controllers.UserDataBulkUpsert)
	r.PUT(UserDataEndpoint+"/upsert/resetall/bulk", controllers.UserDataBulkUpsertResetAll)
	r.DELETE(UserDataEndpoint+"/delete/:id", controllers.UserDataDelete)
}

func UserActionDataRouter(r *gin.Engine) {
	UserActionDataEndpoint := "/api/user/action"
	r.POST(UserActionDataEndpoint+"/create", controllers.UserActionDataCreate)
	r.GET(UserActionDataEndpoint+"/:id", controllers.UserActionDataFindById)
	r.GET(UserActionDataEndpoint+"/user/:user_id", controllers.UserActionDataFindByUserId)
	r.GET(UserActionDataEndpoint+"/:id/user/:user_id", controllers.UserActionDataFindByIdAndUserId)
	r.GET(UserActionDataEndpoint+"/all", controllers.UserActionDataFindAll)
	r.PUT(UserActionDataEndpoint+"/update/:id", controllers.UserActionDataUpdate)
	r.PUT(UserActionDataEndpoint+"/upsert/:id/user/:user_id", controllers.UserActionDataUpsert)
	r.DELETE(UserActionDataEndpoint+"/delete/:id", controllers.UserActionDataDelete)
}
