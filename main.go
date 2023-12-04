package main

import (
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/routers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	/* just initial testing */
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	/*------------------*/

	/* PreTest Metadata Endpoint */
	routers.PreTestMetadataRouter(r)
	/*----------------------------*/
	/* PreTest Metadata Endpoint */
	routers.PreTestDataRouter(r)
	/*----------------------------*/
	/* PreTest Metadata Endpoint */
	routers.PreTestResultDataRouter(r)
	/*----------------------------*/
	/* User Data Endpoint */
	routers.UserDataRouter(r)
	/*----------------------------*/
	r.Run(":1001") // listen and serve on 0.0.0.0:8080
}
