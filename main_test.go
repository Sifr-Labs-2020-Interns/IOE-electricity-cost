package main

import (
	"testing"

	"github.com/Sifr-Labs-2020-Interns/IOE-electricity-cost/connection"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// Connecting to the test database
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
		if result := convertToJSON(table.mappedMsg); result != table.jsonMsg {
			t.Errorf("result was incorrect, got: %s, want: %s.", result, table.jsonMsg)
		}
	}

}

func TestIsValidKey(t *testing.T) {
	key := "ETU3rSTYCnqM51lsgiZMXI5Y4B4sDNxBodsnaImKBgesNcbpf09JbwnFKurCL5zObqihGyDJEJrLXZxPXjvYM1Pe1psqbh4jpHADxjSZYZ8Pey2loQDByDBzdtYyDp8skkD7c3M5tVwWGSzu05zoJxOA8scQtwbgryFhErrGHKTAvUQ3hbRgnOEaj191mP4A7swVOQKqorU8OBTrlmj6W49IPzd0Cp85ZJKtXb4H1HVzR9v39wLFzBeaRGjOQ0EKIGdy3iiKzzLZeIKzy58PjgK2UF8aDw3YaRU9TJILy4q93xNBQJA9xh59HZ3mqJGUfyEUOC15sqEimxPflwrurewHBrc0GO1AjBYwYw4fLOmzgXUXrPBjCsxpTtHkDXzIdf9FqSG4q5BqmdqsDVU5FGcllHvKmhb9Gm2U9DHRniNJ9bLwLMNX1DpIQxBrgrT1Bnzrn1o80fDOqZwSc8KjRWzQpqxxbchlEQCqH8fz12KABRSPzs0k"
	if result := isValidKey(key); result == false {
		t.Errorf("result was incorrect, got %t want %t", result, true)
	}
}

func TestGetRandomString(t *testing.T) {
	key := getRandomString(500)
	if result := isValidKey(key); result == false {
		t.Errorf("result was incorrect, got %t want %t", result, true)
	}
}

func TestIsValidAdmin(t *testing.T) {

	adminKey := "ETU3rSTYCnqM51lsgiZMXI5Y4B4sDNxBodsnaImKBgesNcbpf09JbwnFKurCL5zObqihGyDJEJrLXZxPXjvYM1Pe1psqbh4jpHADxjSZYZ8Pey2loQDByDBzdtYyDp8skkD7c3M5tVwWGSzu05zoJxOA8scQtwbgryFhErrGHKTAvUQ3hbRgnOEaj191mP4A7swVOQKqorU8OBTrlmj6W49IPzd0Cp85ZJKtXb4H1HVzR9v39wLFzBeaRGjOQ0EKIGdy3iiKzzLZeIKzy58PjgK2UF8aDw3YaRU9TJILy4q93xNBQJA9xh59HZ3mqJGUfyEUOC15sqEimxPflwrurewHBrc0GO1AjBYwYw4fLOmzgXUXrPBjCsxpTtHkDXzIdf9FqSG4q5BqmdqsDVU5FGcllHvKmhb9Gm2U9DHRniNJ9bLwLMNX1DpIQxBrgrT1Bnzrn1o80fDOqZwSc8KjRWzQpqxxbchlEQCqH8fz12KABRSPzs0k"
	result := isValid(adminKey, "select count(admin_key) as admin from admins where admin_key=?")
	if result == false {
		t.Errorf("result was incorrect. The given admin key was not found in the database. isValid returned: %t, want: %t", result, true)
	}

}

func TestIsValidUser(t *testing.T) {

	result := isValid("chrisg1", "select count(username) as users from users where username=?")
	if result == false {
		t.Errorf("Result was incorrect. The given username was not found in the database. isValid returned: %t, want: %t", result, true)
	}
}

func TestAddUser(t *testing.T) {

	var newUser NewUser

	newUser.AdminKey = "ETU3rSTYCnqM51lsgiZMXI5Y4B4sDNxBodsnaImKBgesNcbpf09JbwnFKurCL5zObqihGyDJEJrLXZxPXjvYM1Pe1psqbh4jpHADxjSZYZ8Pey2loQDByDBzdtYyDp8skkD7c3M5tVwWGSzu05zoJxOA8scQtwbgryFhErrGHKTAvUQ3hbRgnOEaj191mP4A7swVOQKqorU8OBTrlmj6W49IPzd0Cp85ZJKtXb4H1HVzR9v39wLFzBeaRGjOQ0EKIGdy3iiKzzLZeIKzy58PjgK2UF8aDw3YaRU9TJILy4q93xNBQJA9xh59HZ3mqJGUfyEUOC15sqEimxPflwrurewHBrc0GO1AjBYwYw4fLOmzgXUXrPBjCsxpTtHkDXzIdf9FqSG4q5BqmdqsDVU5FGcllHvKmhb9Gm2U9DHRniNJ9bLwLMNX1DpIQxBrgrT1Bnzrn1o80fDOqZwSc8KjRWzQpqxxbchlEQCqH8fz12KABRSPzs0k"
	newUser.EmailID = "bb10@yahoo.com"
	newUser.Name = "Bruce Banner"
	newUser.Username = "hulk"
	newUser.Password = "hulk_smash"

	adduser(newUser)

	if result := isValid("hulk", "select count(username) as users from users where username=?"); result == false {
		t.Errorf("%s could not be added to the database", newUser.Username)
	}

}

func TestRemoveUser(t *testing.T) {
	var userToRemove RemoveUser

	userToRemove.AdminKey = "ETU3rSTYCnqM51lsgiZMXI5Y4B4sDNxBodsnaImKBgesNcbpf09JbwnFKurCL5zObqihGyDJEJrLXZxPXjvYM1Pe1psqbh4jpHADxjSZYZ8Pey2loQDByDBzdtYyDp8skkD7c3M5tVwWGSzu05zoJxOA8scQtwbgryFhErrGHKTAvUQ3hbRgnOEaj191mP4A7swVOQKqorU8OBTrlmj6W49IPzd0Cp85ZJKtXb4H1HVzR9v39wLFzBeaRGjOQ0EKIGdy3iiKzzLZeIKzy58PjgK2UF8aDw3YaRU9TJILy4q93xNBQJA9xh59HZ3mqJGUfyEUOC15sqEimxPflwrurewHBrc0GO1AjBYwYw4fLOmzgXUXrPBjCsxpTtHkDXzIdf9FqSG4q5BqmdqsDVU5FGcllHvKmhb9Gm2U9DHRniNJ9bLwLMNX1DpIQxBrgrT1Bnzrn1o80fDOqZwSc8KjRWzQpqxxbchlEQCqH8fz12KABRSPzs0k"
	userToRemove.Username = "hulk"

	removeuser(userToRemove)

	if result := isValid("hulk", "select count(username) as users from users where username=?"); result == true {
		t.Errorf("%s is still the database", userToRemove.Username)
	}
}
