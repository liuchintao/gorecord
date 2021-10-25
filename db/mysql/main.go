package main

import (

	// mysql driver
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Test struct {
	NAME   string `db:"NAME"`
	ID     string `db:"ID"`
	STATUS string `db:"STATUS"`
}

func main() {
	d := Test{}
	db := sqlx.MustConnect("mysql", "root:111111@tcp(127.0.0.1:3306)/study?parseTime=true")
	if err := db.GetContext(context.Background(), &d, "SELECT NAME, ID, STATUS FROM attempt"); err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", d)
}
