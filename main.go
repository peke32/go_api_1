package main

import (
	"./db"
	"fmt"
)

func main() {
	fmt.Println(db.Read())
}
