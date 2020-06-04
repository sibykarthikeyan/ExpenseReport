package controllers

import (
	"github.com/sibykarthikeyan/ExpenseReport/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/sibykarthikeyan/ExpenseReport/db"
	"strconv"
)

func GetAllExpenseDetails(c *gin.Context){
	var expenses models.ExpenseDetailsModel
	var totalRepo []models.ExpenseDetailsModel
	sqlStatement:= `SELECT * FROM expensecollection`
	rows,err := db.Db.Query(sqlStatement)
	if err!=nil {
		panic(err)
	}
	for rows.Next() {
		
		err = rows.Scan(&expenses.ID,&expenses.Name,&expenses.ExpenseName,&expenses.ExpenseAmount,&expenses.Created_at) 
		if err!=nil {
			panic(err)
		}
		totalRepo = append(totalRepo,expenses)
	}
	c.JSON(http.StatusOK,totalRepo)
}

func GetExpense(c *gin.Context){
	var expenses models.ExpenseDetailsModel
	id := c.Param("id")
	sqlStatement:= `SELECT * FROM expensecollection WHERE id=$1`
	rows,err := db.Db.Query(sqlStatement,id)
	if err!=nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&expenses.ID,&expenses.Name,&expenses.ExpenseName,&expenses.ExpenseAmount,&expenses.Created_at) 
		if err!=nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK,expenses)
}

func CreateExpenseList(c *gin.Context){
	var input models.ExpenseDetailsModel
	err := c.ShouldBindJSON(&input)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	fmt.Println("### Inserting values")
	sqlStatement := `INSERT INTO expensecollection(name,expenseName,expenseAmount) 
	VALUES($1,$2,$3)
	RETURNING id`
	var id int
	err = db.Db.QueryRow(sqlStatement,input.Name,input.ExpenseName,input.ExpenseAmount).Scan(&id)
	if err!=nil{
		panic(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"id" : id,
		"name":input.Name,
		"expenseName":input.ExpenseName,
		"expenseAmount":input.ExpenseAmount,
	})
}

func UpdateExpenseList(c *gin.Context) {
	var expenses models.ExpenseDetailsModel
	id,_:= strconv.ParseInt(c.Param("id"),10,64)
	err:=c.ShouldBindJSON(&expenses)
	if err!=nil{
		panic(err)
	}
	expenses.ID = id
	fmt.Println("### Updating",expenses)
	sqlStatement := `UPDATE expenseCollection SET name=$2,expenseName=$3,expenseAmount=$4 WHERE id=$1`
	res,err := db.Db.Exec(sqlStatement,expenses.ID,expenses.Name,expenses.ExpenseName,expenses.ExpenseAmount)
	if err!=nil {
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}
	count,err := res.RowsAffected()
	if err!=nil {
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK,expenses)
	}else{
		c.JSON(400,gin.H{
			"error": "No records found",
		})
	}

}

func DeleteExpenseList(c *gin.Context){
	var expenses models.ExpenseDetailsModel
	id,_ := strconv.ParseInt(c.Param("id"),10,64)

	fmt.Println("### deleting")
	sqlStatement := `DELETE FROM expenseCollection WHERE id=$1 RETURNING id,name,expenseName,expenseAmount`
	err := db.Db.QueryRow(sqlStatement,id).Scan(&expenses.ID,&expenses.Name,&expenses.ExpenseName,&expenses.ExpenseAmount)
	if err!=nil {
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}
	if expenses.ID > 0 {
		c.JSON(http.StatusOK,gin.H{
			"data" : "Record deleted Successfully",
		})
	}else{
		c.JSON(400,gin.H{
			"error":"Record not found",
		})
	}
}	