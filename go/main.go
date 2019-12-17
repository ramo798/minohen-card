package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"minohen-card/handler"
)


func main()  {
	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))



	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})
	router.GET("/test",handler.GetList)


	router.Run(":80")

}