package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func CategoryDataRouter(r *gin.Engine) {
	CategoryDataEndpoint := "/api/category/data"
	r.POST(CategoryDataEndpoint+"/create", controllers.CategoryDataCreate)
	r.GET(CategoryDataEndpoint+"/:id", controllers.CategoryDataFindById)
	r.GET(CategoryDataEndpoint+"/domain/:domain", controllers.CategoryDataFindByDomain)
	r.GET(CategoryDataEndpoint+"/label/:label", controllers.CategoryDataFindByLabel)
	r.GET(CategoryDataEndpoint+"/all", controllers.CategoryDataFindAll)
	r.PUT(CategoryDataEndpoint+"/update/:id", controllers.CategoryDataUpdate)
	r.PUT(CategoryDataEndpoint+"/upsert/:id", controllers.CategoryDataUpsert)
	r.DELETE(CategoryDataEndpoint+"/delete/:id", controllers.CategoryDataDelete)
}
