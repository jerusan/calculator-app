package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jj/repo/calculator-app/server/models"

	"github.com/gin-gonic/gin"
)

var maxmsgs int = 10
var modifyIdx int
var mockDb []CalculationEntry

func handleCalculation(c *gin.Context) {
	var expression Expression

	// bind post request body to Expression format
	c.Bind(&expression) // To do: need to add error handling

	// To do: error handling
	c.JSON(http.StatusOK, gin.H{"result": CalculateInOutputFormat(expression)})
	GetDBValues()

	// if operator, err := strconv.ParseFloat(c.Param("operator"), 32); err == nil {
	//   fmt.Printf( "operand1 = %f\n", operator)
	//   c.JSON(http.StatusOK, operator)
	// } else {
	//   c.AbortWithStatus(http.StatusNotFound)
	// }
}

func main() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	mockDb = make([]CalculationEntry, MAX_MSGS)

	// Serve frontend static files
	router.StaticFile("./", "./client/calculator.html")

	router.POST("/", handleCalculation)

	// Start and run the server
	router.Run(":3000")

}

func AddToDb(expression models.Expression) {
	var newCalculationEntity = new(CalculationEntry)
	newCalculationEntity.Expression = expression
	// Storing time in seconds
	newCalculationEntity.CreatedTime = time.Now().Unix()
	mockDb[modifyIdx] = *newCalculationEntity
	modifyIdx = GetNextIdx(modifyIdx)
}

func GetNextIdx(curIdx int) int {
	return (curIdx + 1) % MAX_MSGS
}

func GetDBValues() {
	fmt.Println("Getting DB values")
	fmt.Println("2d: ", mockDb)
}

func CalculateInOutputFormat(expression Expression) string {
	return fmt.Sprintf("%.2f %s %.2f = %.2f", expression.Operand1, expression.Operator, expression.Operand2, Calculate(expression))
}

func Calculate(expression Expression) float64 {
	var result float64

	operand1 := expression.Operand1
	operand2 := expression.Operand2
	operator := expression.Operator

	switch {
	case operator == "+":
		result = operand1 + operand2
	case operator == "-":
		result = operand1 - operand2
	case operator == "/":
		result = operand1 / operand2
	case operator == "*":
		result = operand1 * operand2
	default:
		// To Do: need to handle non valid operator
	}

	AddToDb(expression)

	return result
}
