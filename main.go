package main
import (
	
	"github.com/gin-gonic/gin"
	"github.com/sibykarthikeyan/ExpenseReport/models"
	"github.com/sibykarthikeyan/ExpenseReport/controllers"
	"github.com/sibykarthikeyan/ExpenseReport/db"
)

func main (){
	r:=gin.Default()
	
	db.InitB()

	/* r.Use(func(c *gin.Context){
		c.Set("db",db)
		c.Next()
	}) */
	r.GET("/expenses",controllers.GetExpenseDetails)
	r.POST("/CreateExpenseList",controllers.CreateExpenseList)
	r.Run()
}


