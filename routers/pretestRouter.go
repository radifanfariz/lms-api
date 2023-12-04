package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func PreTestMetadataRouter(r *gin.Engine) {
	PreTestMetadataEndpoint := "/api/pretest/metadata"
	r.POST(PreTestMetadataEndpoint+"/create", controllers.PreTestMetadataCreate)
	r.GET(PreTestMetadataEndpoint+"/:id", controllers.PreTestMetadataFindById)
	r.GET(PreTestMetadataEndpoint+"/all", controllers.PreTestMetadataFindAll)
	r.PUT(PreTestMetadataEndpoint+"/update/:id", controllers.PreTestMetadataUpdate)
	r.DELETE(PreTestMetadataEndpoint+"/delete/:id", controllers.PreTestMetadataDelete)
}
func PreTestDataRouter(r *gin.Engine) {
	PreTestDataEndpoint := "/api/pretest/data"
	r.POST(PreTestDataEndpoint+"/create", controllers.PreTestDataCreate)
	r.GET(PreTestDataEndpoint+"/:id", controllers.PreTestDataFindById)
	r.GET(PreTestDataEndpoint+"/all", controllers.PreTestDataFindAll)
	r.PUT(PreTestDataEndpoint+"/update/:id", controllers.PreTestDataUpdate)
	r.DELETE(PreTestDataEndpoint+"/delete/:id", controllers.PreTestDataDelete)
}
func PreTestResultDataRouter(r *gin.Engine) {
	PreTestDataEndpoint := "/api/pretest/result"
	r.POST(PreTestDataEndpoint+"/create", controllers.PreTestResultDataCreate)
	r.GET(PreTestDataEndpoint+"/:id", controllers.PreTestResultDataFindById)
	r.GET(PreTestDataEndpoint+"/all", controllers.PreTestResultDataFindAll)
	r.PUT(PreTestDataEndpoint+"/update/:id", controllers.PreTestResultDataUpdate)
	r.DELETE(PreTestDataEndpoint+"/delete/:id", controllers.PreTestResultDataDelete)
}
