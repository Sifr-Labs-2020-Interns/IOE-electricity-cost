package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"unicode"
	"unsafe"

	"github.com/Sifr-Labs-2020-Interns/IOE-electricity-cost/connection"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-macaron/binding"
	macaron "gopkg.in/macaron.v1"
)

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
	User_key string `form: "admin_key" binding:"Required"`
	Watts    string `form: "Watts" binding:"Required"`
	Type     string `form: "Type" binding:"Required"`
}

func main() {

	m := macaron.Classic()
	// Public files
	m.Use(macaron.Static("public"))

	// All routes
	m.Post("/adduser", binding.Bind(NewUser{}), adduser)
	m.Post("/removeuser", binding.Bind(RemoveUser{}), removeuser)
	m.Post("/addtransaction", binding.Bind(AddTransaction{}), addtransaction)

	m.Run()
}

/* A function that checks if the provided admin key is valid or not.
   An admin key is valid only if it is alpha numeric (no special characters) and atleast 500 characters long.

	Used https://stackoverflow.com/questions/38554353/how-to-check-if-a-string-only-contains-alphabetic-characters-in-go (17/06/2020)
	for reference
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

	db := connection.ConnectToDB("root", "", "ioe")

	if db == nil {
		panic("Database Connection Failed")
	}

	defer db.Close()

	/* Check if the admin key is valid
	   |___ is valid
	        |___ Generate key for user and insert data to the database and return JSON success with
	             user key
	   |___ not valid
	        |___ Returns in json error admin key not valid

	*/

	if isValid(adminKey) {

		userKey = getRandomString(500)

		mResponse = map[string]string{"Generated Key": userKey}
		jResponse, _ = json.Marshal(mResponse)

		// Taken from https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html (Accessed 19/06/2020)
		query, err := db.Prepare("INSERT INTO users (`NAME`,`EMAIL_ID`, `USERNAME`, `PASSWORD`,`KEY`) VALUES(?,?,?,?,?)")
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

	db := connection.ConnectToDB("root", "", "ioe")

	if db == nil {
		panic("Database Connection Failed")
	}

	defer db.Close()

	/* Check if the admin key is valid
	   |___ is valid
	        |___ Remove user from database and return success
	   |___ not valid
	        |___ Returns in json error admin key not valid
	*/

	if isValid(Admin_key) {

		// Taken from https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html (Accessed 19/06/2020)
		query, err := db.Prepare("DELETE FROM users WHERE username = ? and email_id = ?")
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
func addtransaction(ctx *macaron.Context, addtransaction AddTransaction) {

	// Get information to add transaaction - UNCOMMENT WHEN USING THE VARIABLES
	/* User_key := addtransaction.User_key
	Watts := addtransaction.Watts
	Type := addtransaction.Type */

	/* Check if the user key is valid
	   |___ is valid
	        |___ Add the transaction information in the database and return JSON success
	   |___ not valid
	        |___ Returns in json error user key not valid
	*/

	// Remove this when you have your JSON return statementa
	ctx.Resp.WriteHeader(200)

}
