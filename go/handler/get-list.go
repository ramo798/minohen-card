package handler

import 	"github.com/gin-gonic/gin"


type Card struct{
	Name string
	Nickname string
	Mgram []int
	TwitterId string
	BelongTeam string
	Where string
	Word string
}


func GetList(c *gin.Context) {
	card1 := Card{
		Name:"goma",
		Nickname:"goman",
		Mgram:[]int{1,2,3,4,5},
		TwitterId:"hogehoge",
		BelongTeam:"gakusei",
		Where:"osaka",
		Word:"jocker",
	}
	card2 := Card{
		Name:"goma",
		Nickname:"goman",
		Mgram:[]int{1,2,3,4,5},
		TwitterId:"hogehoge",
		BelongTeam:"gakusei",
		Where:"osaka",
		Word:"jocker",
	}
	card3 := Card{
		Name:"goma",
		Nickname:"goman",
		Mgram:[]int{1,2,3,4,5},
		TwitterId:"hogehoge",
		BelongTeam:"gakusei",
		Where:"osaka",
		Word:"jocker",
	}
	card4 := Card{
		Name:"goma",
		Nickname:"goman",
		Mgram:[]int{1,2,3,4,5},
		TwitterId:"hogehoge",
		BelongTeam:"gakusei",
		Where:"osaka",
		Word:"jocker",
	}

	cards := make([]Card, 0)
	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)

	c.JSON(200,cards)
}
