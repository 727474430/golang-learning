package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

func ConnDB(host string, port string, username string, password string, database string) (db *gorm.DB, err error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username, password, host, port, database)
	fmt.Printf("Mysql conn info: %v\n", url)
	conn, err := gorm.Open("mysql", url)
	if err != nil {
		panic("Fail conn db info")
		return nil, errors.New("DB conn fail")
	}

	return conn, nil
}

func CloseDB(db *gorm.DB) {
	if db == nil {
		return
	}

	e := db.Close()
	if e != nil {
		panic("DB close fail")
	}
}
