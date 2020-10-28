package models

// Expression Model to keep track of operations and operator of each request
type Expression struct {
	Operand1 float64 `json:"operand1" binding:"required"`
	Operand2 float64 `json:"operand2" binding:"required"`
	Operator string  `json:"operator" binding:"required"`
}
