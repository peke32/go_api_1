package api

import (
	"../db"
  "net/http"
  "github.com/ant0ine/go-json-rest/rest"
  "log"
	"encoding/json"
)

func Read(){

	bs, err := json.Marshal(db.Read())
	if err != nil {
		panic(err.Error())
	}

  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(string(bs))
  }))
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
