package api

import (
	"../db"
  "fmt"
  "net/http"
  "github.com/ant0ine/go-json-rest/rest"
  "log"
)

// type User struct{
//   name string
//   pass string
// }

func Read(){

  // users := make([]User, 2)

  i := 0
  for _, user := range db.Read(){
    // users[i] = *user
      fmt.Println(user)
    i++
  }

  // fmt.Println(users)

  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(db.Read())
  }))
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
