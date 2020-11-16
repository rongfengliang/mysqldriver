/*

this is a demo golang sql driver learning

package main

import (
	"database/sql"
	"log"

	_ "github.com/rongfengliang/mysqldriver"
)

func main() {
	db, err := sql.Open("mydb", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		log.Fatalf("some error %s", err.Error())
	}
	rows, err := db.Query("select * from demoapp")
	if err != nil {
		log.Println("some wrong for query", err.Error())
	}
	for rows.Next() {
		rows.Scan()
	}
}
*/

package mydb
