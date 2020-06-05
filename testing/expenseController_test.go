package testing

import (
	"testing"
	"strconv"
	"fmt"
	"net/http"
	"io/ioutil"
)

func TestExpense(t *testing.T){

	Url := "http://localhost:8080/expenses/"
		res,err := http.Get(Url)
		if err!=nil{
			t.Errorf("Somme error-%v",err)
		}
		body,_ := ioutil.ReadAll(res.Body)
		s,_ := strconv.Unquote(string(body))
		fmt.Println("Datat=",s)
	

}
