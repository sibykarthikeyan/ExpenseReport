package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
//	"github.com/sibykarthikeyan/ExpenseReport/api/models"
)

func main (){
	r:=gin.Default()

	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data":"Hello expense"})
	})
	//db := models.connection()
	r.Run()
}
