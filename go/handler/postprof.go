package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"minohen-card/entity"
)

func PostProf(c *gin.Context)  {
	var receive entity.MgramCard
	c.BindJSON(&receive)

	log.Println(receive)
	c.String(200,receive.Nickname)
}
