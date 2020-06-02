package controllers

import (
	"github.com/sibykarthikeyan/ExpenseReport/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExpenseDetails(c *gin.Context){
	var exp models.ExpenseModel
	err := c.ShouldBindJSON(&exp)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,exp)
}

func CreateExpenseList(c *gin.Context){
	var input models.CreateExpenseInput
	err := c.ShouldBindJSON(&input)
	fmt.Println("input:-",input)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	

	/* sqlStatement := `INSERT INTO expenseDetails(age,name,expenseName,expenseAmount) 
	VALUES($1,$2,$3,$4,$5)
	RETURNING id`
	id=0
	err = db.QueryRow(sqlStatement,)
	
	c.JSON(http.StatusOK,gin.H{"data":expense}) */
}