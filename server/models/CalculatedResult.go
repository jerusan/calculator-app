package models

// CalculatedResult returns CalculationEntry with Result
type CalculatedResult struct {
	Expression Expression
	Result     float64
}

// CalculatedResultList returns list
type CalculatedResultList struct {
	Operand1 float64
	Operand2 float64
	Operator string
	Result   float64
}
