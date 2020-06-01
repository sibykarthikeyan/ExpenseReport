package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "admin123"
    dbname   = "expenseReport"
)
func connectDB(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	database,err := sql.Open("postgres",psqlInfo)
	if err!=nil {
		panic(err)
	}
	defer database.Close()
	err = database.Ping()
	if err!=nil {
		panic(err)
	}
	//database.AutoMigrate(&expenseReport{})
	fmt.Println("Db Connected successfully!!")
}
