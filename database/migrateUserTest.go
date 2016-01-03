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
	db.NewRecord(&model.User{})
	user := model.User{
		ID:          1,
		Name:        "太郎",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Architecto quo, magni repellendus perspiciatis necessitatibus molestiae vero eos, omnis voluptates quisquam similique distinctio adipisci sit deleniti obcaecati provident soluta minus facere.",
	}

	db.Create(user)
}
