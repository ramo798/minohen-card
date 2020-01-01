package handler

import (
	"github.com/gin-gonic/gin"
	"minohen-card/db"
	"minohen-card/entity"
)

func PostProf(c *gin.Context)  {
	var receive entity.MgramCard
	c.BindJSON(&receive)

	//log.Println(receive)

	db.RegisterCard(receive)

	c.Status(200)
}
