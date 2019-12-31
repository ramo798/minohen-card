package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"minohen-card/db"
	"minohen-card/handler"
)

func main()  {
	db := db.ConnectionDB()
	defer db.Close()

	router := gin.Default()
	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})
	router.GET("/card/all",handler.GetAllCrad)
	router.POST("/register",handler.PostProf)

	router.Run(":80")


	//handler.GetAllCrad(db)

}
