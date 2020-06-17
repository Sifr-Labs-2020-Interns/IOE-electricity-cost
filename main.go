package main

import (
	"unicode"

	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

//Post request parameters for route add user
type NewUser struct {
	Name      string `form:"name" binding:"Required"`
	Username  string `form: "username" binding:"Required"`
	Email_Id  string `form: "email_id" binding:"Required"`
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
   An admin key is valid only if it is alpha numeric and atleast 500 characters long
*/
func isValid(key string) bool {
     
     if len key < 500{
          return false
     }
     
     alphaFlag := false
	numFlag := false
	for _, c := range key {

		if unicode.IsLetter(c) {
			alphaFlag = true
		} else if unicode.IsDigit(c) {
			numFlag = true
		}

		if numFlag && alphaFlag {
			return true
		}
	}
	return false
}

// TODO: add a new user
func adduser(ctx *macaron.Context, newuser NewUser) {

	// Get user from post request - UNCOMMENT WHEN USING THE VARIABLES

	name := newuser.Name
	Username := newuser.Username
	Email_id := newuser.Email_Id
	Admin_key := newuser.Admin_key

	/* Check if the admin key is valid
	   |___ is valid
	        |___ Generate key for user and insert data to the database and return JSON success with
	             user key
	   |___ not valid
	        |___ Returns in json error admin key not valid

     */
     

	// Remove this when you have your JSON return statementa
	ctx.Resp.WriteHeader(200)
}

// TODO: remove user
func removeuser(ctx *macaron.Context, removeuser RemoveUser) {

	// Get user information from post request - UNCOMMENT WHEN USING THE VARIABLES
	/* Username := removeuser.Username
	Email_id := removeuser.Email_Id
	Admin_key := removeuser.Admin_key */

	/* Check if the admin key is valid
	   |___ is valid
	        |___ Remove user from database and return success
	   |___ not valid
	        |___ Returns in json error admin key not valid
	*/

	// Remove this when you have your JSON return statementa
	ctx.Resp.WriteHeader(200)
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
