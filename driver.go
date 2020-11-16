// copyright to rongfengliang

// Package mydb mydb struct
package mydb

import (
	"database/sql/driver"
	"log"
)

// Driver mydb driver for implement database/sql/driver
type Driver struct {
}

func init() {
	log.Println("driver is call ")
}

// Open for implement driver interface
func (driver *Driver) Open(name string) (driver.Conn, error) {
	log.Println("exec open driver")
	return &Conn{}, nil
}
