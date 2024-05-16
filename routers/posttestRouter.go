package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func PostTestMetadataRouter(r *gin.Engine) {
	PostTestMetadataEndpoint := "/api/posttest/metadata"
	r.POST(PostTestMetadataEndpoint+"/create", controllers.PostTestMetadataCreate)
	r.GET(PostTestMetadataEndpoint+"/:id", controllers.PostTestMetadataFindById)
	r.GET(PostTestMetadataEndpoint+"/all", controllers.PostTestMetadataFindAll)
	r.PUT(PostTestMetadataEndpoint+"/update/:id", controllers.PostTestMetadataUpdate)
	r.PUT(PostTestMetadataEndpoint+"/upsert/:id", controllers.PostTestMetadataUpsert)
	r.DELETE(PostTestMetadataEndpoint+"/delete/:id", controllers.PostTestMetadataDelete)
}
func PostTestDataRouter(r *gin.Engine) {
	PostTestDataEndpoint := "/api/posttest/data"
	r.POST(PostTestDataEndpoint+"/create", controllers.PostTestDataCreate)
	r.GET(PostTestDataEndpoint+"/:id", controllers.PostTestDataFindById)
	r.GET(PostTestDataEndpoint+"/all", controllers.PostTestDataFindAll)
	r.PUT(PostTestDataEndpoint+"/update/:id", controllers.PostTestDataUpdate)
	r.PUT(PostTestDataEndpoint+"/upsert/:id", controllers.PostTestDataUpsert)
	r.DELETE(PostTestDataEndpoint+"/delete/:id", controllers.PostTestDataDelete)
}
func PostTestResultDataRouter(r *gin.Engine) {
	PostTestResultDataEndpoint := "/api/posttest/result"
	r.POST(PostTestResultDataEndpoint+"/create", controllers.PostTestResultDataCreate)
	r.POST(PostTestResultDataEndpoint+"/create/autotime/:tracked_part", controllers.PostTestResultDataAutotimeCreate)
	r.GET(PostTestResultDataEndpoint+"/:id", controllers.PostTestResultDataFindById)
	r.GET(PostTestResultDataEndpoint+"/user/:user_id", controllers.PostTestResultDataFindByUserId)
	r.GET(PostTestResultDataEndpoint+"/:id/user/:user_id", controllers.PostTestResultDataFindByIdAndUserId)
	r.GET(PostTestResultDataEndpoint+"/all", controllers.PostTestResultDataFindAll)
	r.PUT(PostTestResultDataEndpoint+"/update/:id", controllers.PostTestResultDataUpdate)
	r.PUT(PostTestResultDataEndpoint+"/update/autotime/:id/:tracked_part", controllers.PostTestResultDataAutotimeUpdate)
	r.PUT(PostTestResultDataEndpoint+"/upsert/:id", controllers.PostTestResultDataUpsert)
	r.PUT(PostTestResultDataEndpoint+"/upsert/autotime/:id/:tracked_part", controllers.PostTestResultDataAutotimeUpsert)
	r.DELETE(PostTestResultDataEndpoint+"/delete/:id", controllers.PostTestResultDataDelete)
}
