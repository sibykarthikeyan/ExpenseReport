package models
import (
	"context"
	"time"
	"fmt"
	"github.com/sibykarthikeyan/ExpenseReport/db"
)

type ExpenseDetailsModel struct {
	ID int64 `json:"id"`
	Name string `json:"name" binding:"required"`
	ExpenseName string `json:"expenseName" binding:"required"`
	ExpenseAmount int64 `json:"expenseAmount" binding:"required"`
	Created_at time.Time `json:"created_at"` 
}



func CreateExpenseTable(){
	ctx,_ := context.WithTimeout(context.Background(),1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS expenseCollection(
		id SERIAL primary key NOT NULL,
		name varchar(45) NOT NULL,
		expenseName varchar(100) NOT NULL,
		expenseAmount numeric NOT NULL,
		created_at DATE NOT NULL DEFAULT NOW()
	)`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err!=nil {
		fmt.Println("Error in creating expenseCollection table")
		fmt.Println(err.Error())
		panic(err)
	}
}