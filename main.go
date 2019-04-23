package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ken-aio/go-sqlboiler-sample/app/db"
)

func main() {
	initDB()
	//insert()
	//update()
	//delete()
	//insertTx()
	selectSample()
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

func update() {
	user := db.User{ID: 1}
	user.Email = null.StringFrom("update@example.com")
	user.UpdateGP(context.Background(), boil.Infer())
}

func delete() {
	user := db.User{ID: 1}
	user.DeleteGP(context.Background())
}

func insertTx() {
	ctx := context.Background()
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	user := db.User{Email: null.StringFrom("test@example.com"), PasswordDigest: null.StringFrom("digested-password")}
	fmt.Printf("before user = %+v\n", user)
	err = user.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	fmt.Printf("after user = %+v\n", user)
}

func selectSample() {
	//users := db.Users().AllGP(context.Background())
	//fmt.Printf("users = %+v\n", users)

	//user := db.Users().OneGP(context.Background())
	//fmt.Printf("user = %+v\n", user)

	//users := db.Users(qm.InnerJoin("group_members on group_members.user_id = users.id")).AllGP(context.Background())
	//fmt.Printf("users = %+v\n", users)

	//type userMember struct {
	//	db.User        `boil:",bind"`
	//	db.GroupMember `boil:",bind"`
	//}
	//var mem userMember
	//db.Users(qm.Select("users.*, group_members.*"), qm.InnerJoin("group_members on group_members.user_id = users.id")).BindG(context.Background(), &mem)
	//fmt.Printf("mem = %+v\n", mem)

	//user := db.Users(qm.Load("GroupMembers.Group")).OneGP(context.Background())
	//fmt.Printf("user = %+v\n", user)
	//fmt.Printf("user.R.GroupMembers = %+v\n", user.R.GroupMembers)
	//for _, mem := range user.R.GroupMembers {
	//	fmt.Printf("mem = %+v\n", mem)
	//	fmt.Printf("mem.R.Group = %+v\n", mem.R.Group)
	//}

	user := db.Users(qm.Load("GroupMembers", qm.Where("group_members.role = ?", "dummy")), qm.Load("GroupMembers.Group")).OneGP(context.Background())
	fmt.Printf("user = %+v\n", user)
	fmt.Printf("user.R.GroupMembers = %+v\n", user.R.GroupMembers)
	for _, mem := range user.R.GroupMembers {
		fmt.Printf("mem = %+v\n", mem)
		fmt.Printf("mem.R.Group = %+v\n", mem.R.Group)
	}
}
