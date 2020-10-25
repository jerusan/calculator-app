package models

// Expression Model to keep track of operations and operator of each request
type Expression struct {
	Operand1 float64 `form:"operand1" binding:"required"`
	Operand2 float64 `form:"operand2" binding:"required"`
	Operator string  `form:"operator" binding:"required"`
}
