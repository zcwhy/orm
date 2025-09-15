package main

import (
	"fmt"
	"orm/engine"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e, _ := engine.Open("sqlite3", "gee.db")
	defer e.Close()
	e.Raw("DROP TABLE IF EXISTS User;").Exec()
	e.Raw("CREATE TABLE User(Name text);").Exec()
	e.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := e.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
