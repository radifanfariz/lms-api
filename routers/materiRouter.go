package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func MateriMetadataRouter(r *gin.Engine) {
	MateriMetadataEndpoint := "/api/materi/metadata"
	r.POST(MateriMetadataEndpoint+"/create", controllers.MateriMetadataCreate)
	r.GET(MateriMetadataEndpoint+"/:id", controllers.MateriMetadataFindById)
	r.GET(MateriMetadataEndpoint+"/all", controllers.MateriMetadataFindAll)
	r.PUT(MateriMetadataEndpoint+"/update/:id", controllers.MateriMetadataUpdate)
	r.PUT(MateriMetadataEndpoint+"/upsert/:id", controllers.MateriMetadataUpsert)
	r.DELETE(MateriMetadataEndpoint+"/delete/:id", controllers.MateriMetadataDelete)
}
func MateriDataRouter(r *gin.Engine) {
	MateriDataEndpoint := "/api/materi/data"
	r.POST(MateriDataEndpoint+"/create", controllers.MateriDataCreate)
	r.GET(MateriDataEndpoint+"/:id", controllers.MateriDataFindById)
	r.GET(MateriDataEndpoint+"/all", controllers.MateriDataFindAll)
	r.PUT(MateriDataEndpoint+"/update/:id", controllers.MateriDataUpdate)
	r.PUT(MateriDataEndpoint+"/upsert/:id", controllers.MateriDataUpsert)
	r.DELETE(MateriDataEndpoint+"/delete/:id", controllers.MateriDataDelete)
}
