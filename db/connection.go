package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	//"time"
)
var Db *sql.DB
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "admin123"
    dbname   = "testdb2"
)
func InitDB()(err error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	Db,err = sql.Open("postgres",psqlInfo)
	if err!=nil {
		panic(err)
	}
	//defer Db.Close()
	err = Db.Ping()
	if err!=nil {
		fmt.Println("weer",err)
		panic(err)
	}

/* 	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxLifetime(time.Hour) */
	fmt.Println("Db Connected successfully!!")

	return nil
}
