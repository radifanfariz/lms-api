package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func GalleryDataRouter(r *gin.Engine) {
	GalleryDataEndpoint := "/api/gallery/data"
	r.POST(GalleryDataEndpoint+"/create", controllers.GalleryDataCreate)
	r.GET(GalleryDataEndpoint+"/:id", controllers.GalleryDataFindById)
	r.GET(GalleryDataEndpoint+"/user/:user_id", controllers.GalleryDataFindByUserId)
	r.GET(GalleryDataEndpoint+"/all", controllers.GalleryDataFindAll)
	r.POST(GalleryDataEndpoint+"/params", controllers.GalleryDataFindByParams)
	r.PUT(GalleryDataEndpoint+"/update/:id", controllers.GalleryDataUpdate)
	r.PUT(GalleryDataEndpoint+"/upsert/:id", controllers.GalleryDataUpsert)
	r.DELETE(GalleryDataEndpoint+"/delete/:id", controllers.GalleryDataDelete)
}
