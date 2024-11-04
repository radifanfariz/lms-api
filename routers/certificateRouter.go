package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/controllers"
)

func CertificateMasterDataRouter(r *gin.Engine) {
	CertificateMasterDataEndpoint := "/api/certificate/master/data"
	r.POST(CertificateMasterDataEndpoint+"/create", controllers.CertificateMasterDataCreate)
	r.GET(CertificateMasterDataEndpoint+"/:id", controllers.CertificateMasterDataFindById)
	r.GET(CertificateMasterDataEndpoint+"/all", controllers.CertificateMasterDataFindAll)
	r.GET(CertificateMasterDataEndpoint+"/isactive", controllers.CertificateMasterDataFindByIsActive)
	r.PUT(CertificateMasterDataEndpoint+"/update/:id", controllers.CertificateMasterDataUpdate)
	r.PUT(CertificateMasterDataEndpoint+"/upsert/:id", controllers.CertificateMasterDataUpsert)
	r.DELETE(CertificateMasterDataEndpoint+"/delete/:id", controllers.CertificateMasterDataDelete)
}
func CertificateUserDataRouter(r *gin.Engine) {
	CertificateUserDataEndpoint := "/api/certificate/user/data"
	r.POST(CertificateUserDataEndpoint+"/create", controllers.CertificateUserDataCreate)
	r.GET(CertificateUserDataEndpoint+"/:id", controllers.CertificateUserDataFindById)
	r.GET(CertificateUserDataEndpoint+"/user/:user_id", controllers.CertificateUserDataFindByUserId)
	r.GET(CertificateUserDataEndpoint+"/:id/user/:user_id", controllers.CertificateUserDataFindByIdAndUserId)
	r.GET(CertificateUserDataEndpoint+"/all", controllers.CertificateUserDataFindAll)
	r.PUT(CertificateUserDataEndpoint+"/update/:id", controllers.CertificateUserDataUpdate)
	r.PUT(CertificateUserDataEndpoint+"/upsert/:id", controllers.CertificateUserDataUpsert)
	r.DELETE(CertificateUserDataEndpoint+"/delete/:id", controllers.CertificateUserDataDelete)
}
