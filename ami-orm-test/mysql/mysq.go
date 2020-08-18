package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ami-orm?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = db.Close()
	}()

	// 设置数据库最大连接数
	db.SetConnMaxLifetime(100)

	// 设置数据库最大空闲连接数
	db.SetMaxIdleConns(100)

	if err := db.Ping(); err != nil {
		fmt.Println("Open database fail")
		return
	}

	// Exec一般用于执行sql语句
	_, err = db.Exec("drop table if exists  user;")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("create table user(name varchar(100), age int);")
	if err != nil {
		log.Println(err)
	}
	result, err := db.Exec("insert into user(`name`, `age`) values (?, ?),(?, ?)", "Tom", 12, "Sam", 13)
	if err != nil {
		log.Println(err)
		return
	}

	// 影响的行数 2行
	//
	affected, _ := result.RowsAffected()
	log.Println(affected)

	// 查询语句一般用Query或者QueryRow
	row := db.QueryRow("select name from user limit 1")
	var name string

	// row.Scan接收一个或多个指针作为参数，可以获取对应列（column）值[
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
	fmt.Println("Open mysql success")

}
