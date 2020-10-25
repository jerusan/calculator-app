package models

// CalculationEntry Model helps with adding to mockDb
type CalculationEntry struct {
	Username    string
	CreatedTime int64
	Expression  Expression
}
