package main

import (
	"amiorm"
	"amiorm/log"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int `amiorm:"primary key"`
	Name string
	Age  int
}

func main() {
	engine, err := amiorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/ami-orm?charset=utf8mb4")
	if err != nil {
		log.Error(err)
	}

	s := engine.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		log.Error("failed to create table user")
	}
}
