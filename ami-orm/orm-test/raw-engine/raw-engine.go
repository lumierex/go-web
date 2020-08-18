package main

import (
	amiorm "ami-orm"
	"ami-orm/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine, err := amiorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/ami-orm?charset=utf8mb4")
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table user(name varchar(100),age int(10));").Exec()

	// 测试Log
	_, _ = s.Raw("create table user(name varchar(100),age int(10));").Exec()
	result, _ := s.Raw("insert into user(`name`, `age`) values (?, ?), (?, ?)", "tome", 12, "sam", 20).Exec()
	count, _ := result.RowsAffected()

	fmt.Printf("Exec success %d affected", count)

}
