package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/routers"
)

func init() {
	initializers.ConnectToDB()
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

	/* PreTest Metadata Endpoint */
	routers.ModuleMetadataRouter(r)
	/*----------------------------*/
	/* PreTest Data Endpoint */
	routers.ModuleDataRouter(r)
	/*----------------------------*/
	/* PreTest Metadata Endpoint */
	routers.PreTestMetadataRouter(r)
	/*----------------------------*/
	/* PreTest Metadata Endpoint */
	routers.PreTestDataRouter(r)
	/*----------------------------*/
	/* PreTest Metadata Endpoint */
	routers.PreTestResultDataRouter(r)
	/*----------------------------*/
	/* Materi Metadata Endpoint */
	routers.MateriMetadataRouter(r)
	/*----------------------------*/
	/* Materi Metadata Endpoint */
	routers.MateriDataRouter(r)
	/*----------------------------*/
	/* PostTest Metadata Endpoint */
	routers.PostTestMetadataRouter(r)
	/*----------------------------*/
	/* PostTest Metadata Endpoint */
	routers.PostTestDataRouter(r)
	/*----------------------------*/
	/* PostTest Metadata Endpoint */
	routers.PostTestResultDataRouter(r)
	/*----------------------------*/
	/* User Data Endpoint */
	routers.UserDataRouter(r)
	/*----------------------------*/
	/* Access Data Endpoint */
	routers.AccessDataRouter(r)
	/*----------------------------*/
	r.Run(":1001") // listen and serve on 0.0.0.0:8080
}
