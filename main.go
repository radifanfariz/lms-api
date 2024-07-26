package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	/*----------------------------*/
	/* for production */
	// gin.SetMode(gin.ReleaseMode)
	/*----------------------------*/

	r := gin.Default()

	// CORS middleware with custom settings
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // Specify your allowed origin(s) here
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-Portal-Token")
	r.Use(cors.New(corsConfig))

	/* just initial testing */
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	/*------------------*/

	/* Module Metadata Endpoint */
	routers.ModuleMetadataRouter(r)
	/*----------------------------*/
	/* Module Data Endpoint */
	routers.ModuleDataRouter(r)
	/*----------------------------*/
	/* PreTest Metadata Endpoint */
	routers.PreTestMetadataRouter(r)
	/*----------------------------*/
	/* PreTest Data Endpoint */
	routers.PreTestDataRouter(r)
	/*----------------------------*/
	/* PreTest Result Data Endpoint */
	routers.PreTestResultDataRouter(r)
	/*----------------------------*/
	/* Materi Metadata Endpoint */
	routers.MateriMetadataRouter(r)
	/*----------------------------*/
	/* Materi Data Endpoint */
	routers.MateriDataRouter(r)
	/*----------------------------*/
	/* Materi Result Data Endpoint */
	routers.MateriResultDataRouter(r)
	/*----------------------------*/
	/* PostTest Metadata Endpoint */
	routers.PostTestMetadataRouter(r)
	/*----------------------------*/
	/* PostTest Metadata Endpoint */
	routers.PostTestDataRouter(r)
	/*----------------------------*/
	/* PostTest Result Data Endpoint */
	routers.PostTestResultDataRouter(r)
	/*----------------------------*/
	/* User Data Endpoint */
	routers.UserDataRouter(r)
	/*----------------------------*/
	/* User Action Data Endpoint */
	routers.UserActionDataRouter(r)
	/*----------------------------*/
	/* Access Data Endpoint */
	routers.AccessDataRouter(r)
	/*----------------------------*/
	/* Gallery Data Endpoint */
	routers.GalleryDataRouter(r)
	/*----------------------------*/
	/* Upload File Endpoint */
	routers.UploadFileRouter(r)
	/*----------------------------*/
	/* HRIS Data Endpoint */
	routers.HrisDataRouter(r)
	/*----------------------------*/
	/* ModuleMetadataJoinAccssedData Data Endpoint */
	routers.ModuleJoinAccessRouter(r)
	/*----------------------------*/
	/* ModuleMetadataJoinAccssedData V2 Data Endpoint */
	routers.ModuleJoinAccessRouterV2(r)
	/*----------------------------*/
	/* Category Data Endpoint */
	routers.CategoryDataRouter(r)
	/*----------------------------*/
	/* Duplicate Data Endpoint */
	routers.DuplicateDataRouter(r)
	/*----------------------------*/
	/* Portal Data Endpoint */
	routers.PortalDataRouter(r)
	/*----------------------------*/
	port := os.Getenv("PORT")
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
