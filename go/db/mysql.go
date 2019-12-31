package db

import (
	"database/sql"
	"log"
	"minohen-card/entity"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1:3306)/card_database")
	if err != nil {
		log.Fatal("db error.")
	}
	return db
}

func FetchAllCrad() []entity.MgramCard{
	db := ConnectionDB()

	rows, err := db.Query("select * from mgram_cards")
	if err != nil {
		log.Fatal(err)
	}
	var response []entity.MgramCard
	for rows.Next() {
		var Id int
		res := entity.MgramCard{}
		if err := rows.Scan(
			&Id,
			&res.Facebook_id,
			&res.Nickname,
			&res.Year,
			&res.Month,
			&res.Twitter_id,
			&res.Team1,
			&res.Team2,
			&res.Word,
			&res.Mgram1,
			&res.Mgram2,
			&res.Mgram3,
			&res.Mgram4,
			&res.Mgram5,
			&res.Mgram6,
			&res.Mgram7,
			&res.Mgram8,
			&res.Mgram9,
			&res.Area,
			&res.Card_color); err != nil {
			log.Fatal(err)
		}
		response = append(response,res)
	}
	//fmt.Println(response)
	return response
}

