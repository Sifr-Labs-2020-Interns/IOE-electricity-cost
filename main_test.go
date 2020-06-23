package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

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

func TestIsValidAdmin(t *testing.T) {
	// we register an sql driver named "txdb"
	txdb.Register("txdb", "mysql", "root@/test")

	// dsn serves as an unique identifier for connection pool
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if _, err := db.Exec("INSERT INTO admins (`ADMIN_NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`,`ADMIN_KEY`) VALUES(?,?,?,?,?)", "Jack Black", "jb12@gmail.com", "jbl", "jackbl", "jbl94"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added an admin")
	}
	result := isValid(db, "jbl94", "select count(admin_key) as admin from admins where admin_key=?")
	if result == false {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}

}
