package db

import (
	"database/sql"
	"fmt"
	"log"
	"minohen-card/handler"
)

var (
	Db  *sql.DB
	err error
)

func Testget() *model.MysqlResponse {
	var res handler.Card
	rows, err := Db.Query("select * from mgram_card")
	if err != nil {
		log.Println("クエリ検索のところ")
		log.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&res)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}


	return res
}
func init() {
	//q := config.MYSQL_USER + ":" + config.MYSQL_PASSWORD + "@tcp(" + config.MYSQL_HOSTNAME + ":" + config.MYSQL_PORT + ")/" + config.MYSQL_DATABASE
	//log.Println(q)
	//log.Println(config.MYSQL_ROOT_PASSWORD)
	//log.Println(config.MYSQL_HOSTNAME)
	//log.Println(config.MYSQL_PORT)
	//log.Println(config.MYSQL_DATABASE)

	q :=  "docker" + ":" + "docker" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "card_database"

	Db, err = sql.Open("mysql", q)
	if err != nil {
		log.Println("error: mysql init error.")
		log.Fatalln(err)
	}
}
