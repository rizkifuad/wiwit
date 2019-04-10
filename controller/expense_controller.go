package controller

import (
	"time"

	"bitbucket.org/yesboss/sharingan/config"
	"bitbucket.org/yesboss/sharingan/model"
	"bitbucket.org/yesboss/sharingan/repo"
)

type ExpenseController struct {
	Config      config.Config
	ExpenseRepo repo.ExpenseRepo
}

func NewExpenseController(config config.Config, expenseRepo repo.ExpenseRepo) ExpenseController {
	return ExpenseController{Config: config, ExpenseRepo: expenseRepo}
}

func (c *ExpenseController) CheckExpense(user model.User, amount float64) (bool, float64) {
	todayDate := time.Now().Day()
	salaryDate := user.SalaryDate

	// month, _ = strconv.Atoi(time.Now().Month().String())

	expenseMonth := int(time.Now().Month())

	var exEnd time.Time

	if todayDate >= salaryDate {
		exEnd = time.Date(time.Now().Year(), time.Month(expenseMonth+1), salaryDate, 0, 0, 0, 0, time.UTC)
	} else {
		exEnd = time.Date(time.Now().Year(), time.Month(expenseMonth), salaryDate, 0, 0, 0, 0, time.UTC)
	}

	exStart := exEnd.AddDate(0, -1, 0)
	sumMonth := c.ExpenseRepo.TotalExpense(*user.Model.ID, exStart, exEnd) + amount

	batas := float64(user.Expense) / 100.0 * user.Salary

	if sumMonth > batas {
		return true, sumMonth - amount
	}

	return false, sumMonth - amount

}

func (c *ExpenseController) AddExpense(user model.User, amount float64, typed string) model.Expense {
	ex := model.Expense{
		UserID: *user.Model.ID,
		Amount: amount,
		Type:   typed,
	}

	return c.ExpenseRepo.Create(ex)
}

func (c *ExpenseController) Monthly(user model.User) model.MonthlyResponses {
	todayDate := time.Now().Day()
	salaryDate := user.SalaryDate

	// month, _ = strconv.Atoi(time.Now().Month().String())

	expenseMonth := int(time.Now().Month())

	var exEnd time.Time

	if todayDate >= salaryDate {
		exEnd = time.Date(time.Now().Year(), time.Month(expenseMonth+1), salaryDate, 0, 0, 0, 0, time.UTC)
	} else {
		exEnd = time.Date(time.Now().Year(), time.Month(expenseMonth), salaryDate, 0, 0, 0, 0, time.UTC)
	}

	exStart := exEnd.AddDate(0, -1, 0)

	return c.ExpenseRepo.GetMonthly(user, exStart, exEnd)
}

func (c *ExpenseController) Daily(user model.User) model.MonthlyResponses {
	return c.ExpenseRepo.GetDaily(user)
}

func (c *ExpenseController) List(page int, limit int) []model.Expense {
	return c.ExpenseRepo.List(page, limit)
}
