package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.1.86:3308)/test");
	if err != nil {
		panic(err)
	}

	id := 1
	rows, err := db.Query("SELECT name FROM weather_tmp WHERE id = ?", id);
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
	    var name string
	    if err := rows.Scan(&name); err != nil {
	        log.Fatal(err)
	    }
	    fmt.Printf("%s is %d\n", name, id)
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
}