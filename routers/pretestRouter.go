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
	r.PUT(PreTestMetadataEndpoint+"/upsert/:id", controllers.PreTestMetadataUpsert)
	r.DELETE(PreTestMetadataEndpoint+"/delete/:id", controllers.PreTestMetadataDelete)
}
func PreTestDataRouter(r *gin.Engine) {
	PreTestDataEndpoint := "/api/pretest/data"
	r.POST(PreTestDataEndpoint+"/create", controllers.PreTestDataCreate)
	r.GET(PreTestDataEndpoint+"/:id", controllers.PreTestDataFindById)
	r.GET(PreTestDataEndpoint+"/all", controllers.PreTestDataFindAll)
	r.PUT(PreTestDataEndpoint+"/update/:id", controllers.PreTestDataUpdate)
	r.PUT(PreTestDataEndpoint+"/upsert/:id", controllers.PreTestDataUpsert)
	r.DELETE(PreTestDataEndpoint+"/delete/:id", controllers.PreTestDataDelete)
}
func PreTestResultDataRouter(r *gin.Engine) {
	PreTestResultDataEndpoint := "/api/pretest/result"
	r.POST(PreTestResultDataEndpoint+"/create", controllers.PreTestResultDataCreate)
	r.GET(PreTestResultDataEndpoint+"/:id", controllers.PreTestResultDataFindById)
	r.GET(PreTestResultDataEndpoint+"/all", controllers.PreTestResultDataFindAll)
	r.PUT(PreTestResultDataEndpoint+"/update/:id", controllers.PreTestResultDataUpdate)
	r.PUT(PreTestResultDataEndpoint+"/upsert/:id", controllers.PreTestResultDataUpsert)
	r.DELETE(PreTestResultDataEndpoint+"/delete/:id", controllers.PreTestResultDataDelete)
}
