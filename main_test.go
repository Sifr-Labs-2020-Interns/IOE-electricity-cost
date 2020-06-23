package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// we register an sql driver named "txdb"
	txdb.Register("txdb", "mysql", "root@/test")

	// dsn serves as an unique identifier for connection pool
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if _, err := db.Exec("INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`,`USER_KEY`) VALUES(?,?,?,?,?)", "Jack Black", "jb101@yahoo.com", "jb101", "password", "jb101"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}

func TestConvertToJSON(t *testing.T) {
	tables := []struct {
		mappedMsg map[string]string
		jsonMsg   string
	}{
		{map[string]string{"Message": "Hello World"}, `{"Message":"Hello World"}`},
		{map[string]string{"Name": "Tony", "Age": "30"}, `{"Age":"30","Name":"Tony"}`},
	}
	for _, table := range tables {
		result := convertToJSON(table.mappedMsg)
		if result != table.jsonMsg {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, table.jsonMsg)
		}
	}

}

// func TestIsValidAdmin(t *testing.T) {

// }
