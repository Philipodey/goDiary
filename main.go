package main

// import "fmt"
import "github.com/gin-gonic/gin"
func main(){
	// fmt.Println("Hello Nation")
	router:=gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"Hello world",
		})
	})
	
}

func add(numberOne, numberTwo int) int{
	// total int
	total := numberOne + numberTwo
	return total
}