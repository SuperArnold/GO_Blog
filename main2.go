package main

import (
	"fmt"

	// _ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Tag struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     int    `form:"state,default=1"`
}

func (*Tag) TableName() string {
	return "tag"
}

func main() {
	connStr := fmt.Sprintf(
		"host=%s user=%s port=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		"postgres",
		"5432",
		"postgres",
		"blog",
	)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	tag := Tag{Name: "A1", State: 1, CreatedBy: "1111111"}

	// has := db.HasTable(&userPref)
	has := db.HasTable(tag)
	if !has {
		fmt.Println("table not exist")
	}

	fmt.Println(db.Create(&tag))
}
