package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	simpleJson "github.com/bitly/go-simplejson"
	"io/ioutil"
)

// 構造体の定義
type Apival struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func Read() []Apival {
	// json読み込み
	file, err := ioutil.ReadFile("/var/www/go/db/config.json")
	if err != nil {
			panic(err)
	}

	// json取得
	js, err := simpleJson.NewJson([]byte(file))
	if err != nil {
		panic(err)
	}

	// 設定の構造体
	type Dbconf struct {
	    host string `json:host`
	}

	var dbconfig Dbconf

  dbconfig.host, err = js.Get("host").String()
  if err != nil {
    panic(err)
  }

	db, err := sql.Open("mysql", dbconfig.host)
	if err != nil {
		panic(err.Error())
	}

	// 件数取得
	count_rows, err := db.Query("SELECT count(*) as count FROM users")
	if err != nil {
		panic(err.Error())
	}

	var count int
	for count_rows.Next() {
		err = count_rows.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}

	users := make([]Apival, count)

	// レコード数取得
	rows, err := db.Query("SELECT name, pass FROM users")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	i := 0
	for rows.Next() {

		var name string
		var pass string

		err = rows.Scan(&name, &pass)

		if err != nil {
			panic(err.Error())
		}
		users[i].Name = name
		users[i].Pass = pass
		i++
	}

	return users
}
