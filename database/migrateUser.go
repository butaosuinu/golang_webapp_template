package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"../model"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/goapp?charset=utf8&parseTime=True")
	fmt.Println(err)
	db.DropTable(&model.User{})
	db.CreateTable(&model.User{})
}
