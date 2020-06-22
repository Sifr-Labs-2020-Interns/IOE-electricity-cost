package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
	"unsafe"

	"./connection"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-macaron/binding"
	macaron "gopkg.in/macaron.v1"
)

//Connection object
var conn *sql.DB

/*	Taken from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	Accessed 17/06/2020. Line 16 - 24 and function getRandomString has been taken from the source specified above.
*/

var src = rand.NewSource(time.Now().UnixNano())

// This is the character set we will be using to build our key.
const charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const (
	charIndexBits = 6                    // 6 bits to represent a character index (since we have 62 characters)
	charIndexMask = 1<<charIndexBits - 1 // All 1-bits, as many as charIndexBits
	charIndexMax  = 63 / charIndexBits   // Number of character indices fitting in 63 bits
)

/*----------------------------------------- Struct for Post request ------------------------------------------- */
//Post request parameters for route add user
type NewUser struct {
	Name      string `form:"name" binding:"Required"`
	Username  string `form: "username" binding:"Required"`
	Password  string `form: "password" binding:"Required"`
	Email_ID  string `form: "email_id" binding:"Required"`
	Admin_key string `form: "admin_key" binding:"Required"`
}

//Post request parameters for route to remove user
type RemoveUser struct {
	Username  string `form: "username" binding:"Required"`
	Email_Id  string `form: "email_id" binding:"Required"`
	Admin_key string `form: "admin_key" binding:"Required"`
}

//Post request parameters for route to Add transaction
type AddTransaction struct {
	User_key string `form: "user_key" binding:"Required"`
	Watts    string `form: "watts" binding:"Required"`
	Type     string `form: "type" binding:"Required"`
}

/*-----------------------------------------------------------------------------------------------------------------*/

/* ----------------------------------------------------------------------------------------
------------------------------------ MAIN FUNCTION ---------------------------------------
------------------------------------------------------------------------------------------*/

func main() {

	argsWithoutProg := os.Args[1:]

	// Getting database information from arguments
	db_username := argsWithoutProg[0]
	db_password := argsWithoutProg[1]
	db := argsWithoutProg[2]

	// database connection
	conn = connection.ConnectToDB(db_username, db_password, db)

	if conn == nil {
		panic("Database Connection Failed")
	}

	m := macaron.Classic()
	// Public files
	m.Use(macaron.Static("public"))

	// All routes
	m.Post("/adduser", binding.Bind(NewUser{}), adduser)
	m.Post("/removeuser", binding.Bind(RemoveUser{}), removeuser)
	m.Post("/addtransaction", binding.Bind(AddTransaction{}), addtransaction)

	m.Run()
}

/*-----------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------*/

/*---------------------------------------------------------------------------------------------
------------------------------------- HELPER FUNCTIONS ----------------------------------------
---------------------------------------------------------------------------------------------*/

/* A function that checks if the provided admin key is valid or not.
   An admin key is valid only if it is alpha numeric (no special characters) and atleast
	 500 characters long.
	Used https://stackoverflow.com/questions/38554353/how-to-check-if-a-string-only-contains-alphabetic-characters-in-go (17/06/2020)
	for reference
	NOT USED : To be used later
*/
func isValid(key string) bool {

	// if len(key) < 500 {
	// 	return false
	// }

	// Flags to check if the string that's read contains alphabets, letters and no special characters.
	alphaFlag := false
	numFlag := false
	specialCharFlag := false
	for _, c := range key {

		if c >= 48 && c <= 57 {
			numFlag = true
		} else if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			alphaFlag = true
		} else {
			specialCharFlag = true
		}

		if specialCharFlag {
			return false
		}
	}
	if !numFlag || !alphaFlag {
		return false
	}

	return true
}

/* Checks if the admin key is valid or not  */
func isValidAdmin(key string) bool {

	result, err := conn.Query("select count(admin_key) as admin from admins where admin_key=?", key)

	if err != nil {
		panic(err.Error())
	}

	if result.Next() {

		var count int
		// for each row, scan the result into our tag composite object
		err = result.Scan(&count)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if count == 1 {
			return true
		}

	}

	return false

}

//Function to check if user key is valid

func isValidUser(key string) bool {

	result, err := conn.Query("select count(user_key) as user from users where user_key=?", key)

	if err != nil {
		panic(err.Error())
	}

	if result.Next() {

		var count int
		// for each row, scan the result into our tag composite object
		err = result.Scan(&count)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if count == 1 {
			return true
		}

	}

	return false

}

// A function that return a random string containing alpha numeric characters
func getRandomString(n int) string {

	numFlag := false // A flag to keep track if the string we're creating has a number in it.
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), charIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), charIndexMax
		}
		if index := int(cache & charIndexMask); index < len(charSet) {
			b[i] = charSet[index]

		} else {
			b[i] = charSet[index-2] // In case index gets a value of 63 or 62, decrement it by 2
		}

		if !numFlag { // If numFlag is false
			// If the rune of this byte is a digit, set numFlag to true
			if unicode.IsDigit(rune(b[i])) {
				numFlag = true
			}
		}

		i--
		cache >>= charIndexBits
		remain--
	}

	// In the event that the string we created doesn't have any number in it
	if numFlag == false {
		b[5] = charSet[1] // set the 10th character in the string to the number 1
	}

	return *(*string)(unsafe.Pointer(&b))
}

/*-----------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------*/

/*---------------------------------------------------------------------------------------------
------------------------------- FUNCTION TO HANDLE ROUTES  ------------------------------------
---------------------------------------------------------------------------------------------*/

// TODO: add a new user
func adduser(ctx *macaron.Context, newuser NewUser) string {

	// Get user from post request
	name := newuser.Name
	username := newuser.Username
	password := newuser.Password
	email_id := newuser.Email_ID
	adminKey := newuser.Admin_key

	//Hashing the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		panic(err.Error())
	}

	password = string(bytes)

	userKey := "null"                // the auto generated key
	mResponse := map[string]string{} // the JSON response in a map[string] string
	jResponse := []byte{}            // the JSON response as a JSON object

	/* Check if the admin key is valid
	   |___ is valid
	        |___ Generate key for user and insert data to the database and return JSON success with
	             user key
	   |___ not valid
	        |___ Returns in json error admin key not valid
	*/
	if isValidAdmin(adminKey) {

		userKey = getRandomString(500)

		mResponse = map[string]string{"Generated Key": userKey}
		jResponse, _ = json.Marshal(mResponse)

		// Taken from https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html (Accessed 19/06/2020)
		query, err := conn.Prepare("INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`,`USER_KEY`) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(name, email_id, username, password, userKey)

		fmt.Println("Entered value")
		if err != nil {
			panic(err.Error())
		}

	} else {
		mResponse = map[string]string{"Error": "Admin key not valid"}
		jResponse, _ = json.Marshal(mResponse)
	}

	return string(jResponse)

}

// TODO: remove user
func removeuser(ctx *macaron.Context, removeuser RemoveUser) string {

	// Get user information from post request - UNCOMMENT WHEN USING THE VARIABLES
	Username := removeuser.Username
	Email_id := removeuser.Email_Id
	Admin_key := removeuser.Admin_key

	mResponse := map[string]string{} // the JSON response in a map[string] string
	jResponse := []byte{}            // the JSON response as a JSON object

	/* Check if the admin key is valid
	   |___ is valid
	        |___ Remove user from database and return success
	   |___ not valid
	        |___ Returns in json error admin key not valid
	*/

	if isValidAdmin(Admin_key) {

		// Taken from https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html (Accessed 19/06/2020)
		query, err := conn.Prepare("DELETE FROM users WHERE username = ? and email_id = ?")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(Username, Email_id)

		mResponse = map[string]string{"Status": "Successfully removed " + Username}
		jResponse, _ = json.Marshal(mResponse)
	} else {
		mResponse = map[string]string{"Error": "Admin key not valid"}
		jResponse, _ = json.Marshal(mResponse)
	}

	return string(jResponse)

}

// TODO: Add transaction
func addtransaction(ctx *macaron.Context, addtransaction AddTransaction) string {

	// Get information to add transaaction - UNCOMMENT WHEN USING THE VARIABLES
	/* User_key := addtransaction.User_key
	Watts := addtransaction.Watts
	Type := addtransaction.Type */

	User_key := addtransaction.User_key
	Watts := addtransaction.Watts
	Type := addtransaction.Type

	mResponse := map[string]string{}
	jResponse := []byte{}

	/* Check if the user key is valid
	   |___ is valid
	        |___ Add the transaction information in the database and return JSON success
	   |___ not valid
	        |___ Returns in json error user key not valid
	*/

	if isValid(User_key) {

		query, err := conn.Prepare("INSERT INTO transactions (`USER_ID`, `WATT/SECOND`,`TYPE`) SELECT user_id,?,? FROM users WHERE user_key=?")

		if err != nil {
			panic(err.Error())
		}

		// query.Exec(User_key, Watts, Type)

		query.Exec(Watts, Type, User_key)

		fmt.Println("Transaction done")
		if err != nil {
			panic(err.Error())
		}

	} else {

		mResponse = map[string]string{"Error": "User key does not exist"}
		jResponse, _ = json.Marshal(mResponse)
	}

	return string(jResponse)

}

/*-----------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------*/
