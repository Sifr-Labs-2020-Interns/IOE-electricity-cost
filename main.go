package main


import (
  "gopkg.in/macaron.v1"
  "fmt"
  "github.com/go-macaron/binding"
)

//Post request parameters for route add user
type NewUser struct {
    Name  string `form:"name" binding:"Required"`,
    Username string `form: "username" binding:"Required"`,
    Email_Id string `form: "email_id" binding:"Required"`,
    Admid_key string `form: "admin_key" binding:"Required"`
}



func main() {
    // All routes
    m := macaron.Classic()
    // Public files
    m.Use(macaron.Static("public"))
    m.Post("/adduser",binding.Bind(NewUser{}),adduser)
    m.Run()
}

func adduser(ctx *macaron.Context,newuser NewUser){
   fmt.Printf("%v\n", newuser.Name)
   ctx.Resp.WriteHeader(200)
}
