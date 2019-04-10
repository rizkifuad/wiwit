package model

type ExpenseResponse struct {
	OverBudget        bool    `json:"over_budget"`
	Salary            float64 `json:"salary"`
	ExpensePercentage int     `json:"expense_percentage"`
	TotalExpense      float64 `json:"total_expense"`
}

type MonthlyResponse struct {
	Total float64 `json:"total"`
	Type  string  `json:"type"`
}

type MonthlyResponses struct {
	Data              []MonthlyResponse `json:"data"`
	Total             float64           `json:"total"`
	ExpensePercentage float64           `json:"expense_percentage"`
}
