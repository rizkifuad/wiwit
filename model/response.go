package model

type ExpenseResponse struct {
	OverBudget        bool    `json:"over_budget"`
	Salary            float64 `json:"salary"`
	ExpensePercentage int     `json:"expense_percentage"`
	TotalExpense      float64 `json:"total_expense"`
}
