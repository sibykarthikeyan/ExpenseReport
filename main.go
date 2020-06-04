package main
import (
	
	"github.com/gin-gonic/gin"
	"github.com/sibykarthikeyan/ExpenseReport/controllers"
	"github.com/sibykarthikeyan/ExpenseReport/db"
	"github.com/sibykarthikeyan/ExpenseReport/models"
)

func main (){
	r:=gin.Default()
	
	db.InitDB()
	models.Init()
	 
	r.GET("/expenses",controllers.GetAllExpenseDetails)
	r.POST("/createExpenseList",controllers.CreateExpenseList)
	r.PATCH("/updateExpenseList/:id",controllers.UpdateExpenseList)
	r.DELETE("/deleteExpenseList/:id",controllers.DeleteExpenseList)
	r.GET("/getExpense/:id",controllers.GetExpense)
	r.Run()
}


