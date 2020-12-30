package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectToDB is a function used to connect to a mysql database.
func ConnectToDB(username string, password string, database string, host string,port_no string) *sql.DB {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port_no+")/"+database)
	if err != nil {
		db.Close()
		db = nil
		panic(err.Error())
	}

	fmt.Println("Connected to the database", database)

	return db
}
