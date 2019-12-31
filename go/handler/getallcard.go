package handler

import (
	"github.com/gin-gonic/gin"
	"minohen-card/db"
)


func GetAllCrad(c *gin.Context){
	response := db.FetchAllCrad()
	//log.Println(response)
	//log.Println(reflect.TypeOf(response))
	c.JSON(200, response)
}
