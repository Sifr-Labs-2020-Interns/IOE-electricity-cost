package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wgb-10/IOE-electricity-cost/connection"
)

func init() {
	// we register an sql driver named "txdb"
	//txdb.Register("txdb", "mysql", "root@/test")
	conn = connection.ConnectToDB("root", "", "test")
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

func TestIsValidAdmin(t *testing.T) {

	// making the conn variable point to the driver returned by sql open here
	//conn, err := sql.Open("txdb", "identifier")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	defer conn.Close()

	// if _, err := conn.Exec("INSERT INTO admins (`ADMIN_NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`,`ADMIN_KEY`) VALUES(?,?,?,?,?)", "Jack Black", "jb12@gmail.com", "jbl", "jackbl", "jbl94"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added an admin")
	// }
	result := isValid("jbl94", "select count(admin_key) as admin from admins where admin_key=?")
	if result == false {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}

}

/*func TestAddUserToDB(t *testing.T) {
	// making the conn variable point to the driver returned by sql open here
	conn, err := sql.Open("txdb", "identifier")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	var newUser NewUser

	newUser.Admin_key = "jbl94"
	newUser.Email_ID = "blah@yahoo.com"
	newUser.Name = "Bruce Banner"
	newUser.Username = "hulk"
	newUser.Password = "punny god"

	userKey := "Abc123"

	_, errString := addUserToDB(conn, newUser, userKey)

	if errString != "null" {
		t.Errorf("Could not add %s to the database", newUser.Username)
	}
}
*/
