package models

// CalculationEntry Model helps with adding to mockDb
type CalculationEntry struct {
	Username    string
	CreatedTime int64
	Expression  Expression
}

// CalculationEntries list
type CalculationEntries []CalculationEntry

func (s CalculationEntries) Len() int {
	return len(s)
}

func (s CalculationEntries) Less(i, j int) bool {
	return s[i].CreatedTime < s[j].CreatedTime
}

func (s CalculationEntries) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
