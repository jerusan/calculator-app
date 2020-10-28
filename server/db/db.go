package db

import (
	"fmt"
	"time"

	"github.com/jj/repo/calculator-app/server/models"
)

var maxmsgs int
var modifyIdx int

func Init(maxMsgs int) {
	maxmsgs = maxMsgs
	modifyIdx = 0
}

func AddToDb(expression models.Expression, mockDb []models.CalculationEntry) {
	var newCalculationEntity = new(models.CalculationEntry)
	newCalculationEntity.Expression = expression
	// Storing time in seconds
	newCalculationEntity.CreatedTime = time.Now().Unix()
	mockDb[modifyIdx] = *newCalculationEntity
	modifyIdx = getNextIdx(modifyIdx)
}

func getNextIdx(curIdx int) int {
	return (curIdx + 1) % maxmsgs
}

// GetDBValues fn Used only for testing
func GetDBValues(mockDb []models.CalculationEntry) {
	fmt.Println("Getting DB values")
	fmt.Println("2d: ", mockDb)
}
