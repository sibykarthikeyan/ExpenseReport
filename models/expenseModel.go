package models

type ExpenseModel struct {
	ID int `json:"id" gorm:"primary_key"`
	ExpenseName string `json:"expenseName"`
	ExpenseAmount float64 `json:"expenseAmount"`
}