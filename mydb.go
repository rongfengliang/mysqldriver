// copyright to rongfengliang

// this is a demo golang driver for test

package mydb

import (
	"database/sql"
	"log"
)

func init() {
	log.Println("register mydb driver")
	sql.Register("mydb", &Driver{})
}
