package main

import (
	"github.com/gin-gonic/gin"
	"minohen-card/handler"
)


func main()  {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})
	router.GET("/test",handler.GetList)


	router.Run(":80")

}