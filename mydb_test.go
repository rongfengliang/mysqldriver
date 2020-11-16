package mydb_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/rongfengliang/mysqldriver"
	mydb "github.com/rongfengliang/mysqldriver"
)

func TestDb(t *testing.T) {
	db, err := sql.Open("mydb", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		t.Errorf("some error %s", err.Error())
	}
	rows, err := db.Query("select name,age,version from demoapp")
	if err != nil {
		log.Fatal("some wrong for query", err.Error())
	}
	for rows.Next() {
		var user mydb.MyUser
		if err := rows.Scan(&user.Name, &user.Age, &user.Version); err != nil {
			log.Println("scan value erro", err.Error())
		} else {
			log.Println(user)
		}
	}
}

func Example() {
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
