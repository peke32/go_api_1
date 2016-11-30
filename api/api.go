package api

import (
	"../db"
  "fmt"
  "net/http"
  "github.com/ant0ine/go-json-rest/rest"
  "log"
	"encoding/json"
)

func Read(){

  for _, user := range db.Read(){
			bs, _ := json.Marshal(user)
			fmt.Println(string(bs))
  }

  // fmt.Println(users)

  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(db.Read())
  }))
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
