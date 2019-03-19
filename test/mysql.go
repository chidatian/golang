package main

import(
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("github.com")
	db, err := sql.Open("mysql", "root:root@/test")	
	if err != nil {
	    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	}
	// stmtIn, err := db.Prepare("insert into goods (name) values(?)")
	// defer stmtIn.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM goods WHERE id = ?")
	if err != nil {
		fmt.Println("something wrong")
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// stmtIn.Exec("dell")
	var name string
	var id int
	var time int
	stmtOut.QueryRow(1).Scan(&id,&name,&time)
	fmt.Printf("%d,%s,%d",id,name,time)
}