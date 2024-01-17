package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/routers"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

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
	/* Upload File Endpoint */
	routers.UploadFileRouter(r)
	/*----------------------------*/
	r.Run(":1001") // listen and serve on 0.0.0.0:8080
}
