package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ken-aio/go-sqlboiler-sample/app/db"
)

func main() {
	initDB()
	insert()
}

func initDB() {
	dns := "user=postgres dbname=sampledb host=localhost sslmode=disable connect_timeout=10"
	con, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}
	// connection pool settings
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

	// global connection setting
	boil.SetDB(con)
	boil.DebugMode = true
}

func insert() {
	user := db.User{Email: null.StringFrom("test@example.com"), PasswordDigest: null.StringFrom("digested-password")}
	fmt.Printf("before user = %+v\n", user)
	user.InsertGP(context.Background(), boil.Infer())
	fmt.Printf("after user = %+v\n", user)
}
