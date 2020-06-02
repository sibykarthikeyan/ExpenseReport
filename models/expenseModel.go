package models
import (
	"context"
	"time"
	"github.com/sibykarthikeyan/ExpenseReport/api/db"
)

type ExpenseDetailsModel struct {
	ID int64 `json:"id"`
	Age int64 `json:"age" binding:"required"`
	Name string `json:"name" binding:"required"`
	ExpenseName string `json:"expenseName" binding:"required"`
	ExpenseAmount int64 `json:"expenseAmount" binding:"required"`
}

func CreateExpenseTable(){
	ctx,_ := context.WithTimeout(context.Background(),1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS expenseDetails(
		id SERIAL primary key NOT NULL,
		name varchar(45) NOT NULL,
		expenseName varchar(100) NOT NULL,
		expenseAmount numeric NOT NULL,
		created_at timestampz NOT NULL DEFAULT NOW()
	)`
	_,err := db.Db().ExecContext(ctx,sqlStatement)
	if err!=nil {
		fmt.Println("Error in creating expenseDetails table")
		fmt.Println(err.Error())
		panic(err)
	}
}