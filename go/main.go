package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


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



func main()  {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1:3306)/card_database")
	if err != nil{
		log.Fatal("db error.")
	}
	defer db.Close()

	rows, err := db.Query("select * from mgram_cards")
	if err != nil {
		log.Fatal(err)
	}



	for rows.Next() {
		var Id int
		res := MgramCard{}
		if err := rows.Scan(
			&Id,
			&res.facebook_id,
			&res.nickname,
			&res.year,
			&res.month,
			&res.twitter_id,
			&res.team1,
			&res.team2,
			&res.word,
			&res.mgram1,
			&res.mgram2,
			&res.mgram3,
			&res.mgram4,
			&res.mgram5,
			&res.mgram6,
			&res.mgram7,
			&res.mgram8,
			&res.mgram9,
			&res.area,
			&res.card_color); err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)

	}



	//router := gin.Default()
	//// CORS 対応
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"*"}
	//router.Use(cors.New(config))
	//
	//router.GET("/", func(c *gin.Context) {
	//	c.String(200, "hello, API")
	//})
	//router.GET("/test",handler.GetList)
	//
	//router.Run(":80")
}
