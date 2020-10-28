package main

import (
	"flag"
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jj/repo/calculator-app/server/db"
	"github.com/jj/repo/calculator-app/server/models"
	"github.com/jj/repo/calculator-app/server/websocket"
)

var maxmsgs int = 10
var mockDb models.CalculationEntries

func getRecentOperations(c *gin.Context) {
	list := getUpdatedCalculationList()

	// Allow CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, list)
}

func handleCalculation(c *gin.Context) {
	expression := new(models.Expression)

	// bind post request body to Expression format
	c.Bind(&expression) // To do: error handling

	db.AddToDb(*expression, mockDb)

	calculatedResult := calculate(*expression)

	// Allow CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, calculatedResult)

	// send latest operations to active clients
	list := getUpdatedCalculationList()

	websocket.BroadCastLatestCacluations(list)

	// Used only for testing
	// db.GetDBValues(mockDb)
}

func cors(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	flag.IntVar(&maxmsgs, "list-size", 10, "You can adjust the no. of recent operations, default is 10")
	flag.Parse()

	fmt.Println("Size of the list: ", maxmsgs)
	go websocket.SetUpWS()

	initDb()

	initRouters()

}

func initDb() {
	mockDb = make([]models.CalculationEntry, maxmsgs)
	db.Init(maxmsgs)
}

func initRouters() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(cors)

	// Serve frontend static files
	router.StaticFile("./", "../client/calculator.html")

	router.POST("/", handleCalculation)
	router.GET("/recentOperations", getRecentOperations)

	// Start and run the server
	router.Run(":3000")
}

func calculateInOutputFormat(expression models.Expression) string {
	res := fmt.Sprintf("%.2f %s %.2f = %.2f", expression.Operand1, expression.Operator, expression.Operand2, calculate(expression))
	return res
}

func calculate(expression models.Expression) float64 {
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
	return result
}

func getUpdatedCalculationList() []string {
	list := make([]string, maxmsgs)

	// Copy mockDb temporarily to tempSlice to sort them
	tempDb := make(models.CalculationEntries, maxmsgs)
	copy(tempDb, mockDb)
	sort.Sort(sort.Reverse(tempDb))

	for i, entry := range tempDb {
		if entry.CreatedTime != 0 {
			list[i] = calculateInOutputFormat(entry.Expression)
		}
	}
	return list
}
