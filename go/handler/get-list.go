package handler

import 	"github.com/gin-gonic/gin"


type MgramCard struct{
	facebook_id string
	nickname   string
	year       int
	month      int
	twitter_id  string
	team1      string
	team2      string
	word       string
	mgram1     string
	mgram2     string
	mgram3     string
	mgram4     string
	mgram5     string
	mgram6     string
	mgram7     string
	mgram8     string
	mgram9     string
	area       string
	card_color  int
}


func GetList(c *gin.Context) {
	c.JSON(200,"")
}
