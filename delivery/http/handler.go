package http

import (
	"net/http"
	"strconv"

	"bitbucket.org/yesboss/sharingan/controller"
	"bitbucket.org/yesboss/sharingan/model"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Handler struct {
	UserController    controller.UserController
	ExpenseController controller.ExpenseController
}

func New(u controller.UserController, ex controller.ExpenseController) *echo.Echo {
	e := echo.New()

	handler := Handler{
		UserController:    u,
		ExpenseController: ex,
	}

	e.GET("/ping", handler.ping)
	e.POST("/register", handler.register)
	e.POST("/budget_plan/:user_id", handler.budgetPlan)
	e.POST("/expense/:user_id", handler.addExpense)
	e.POST("/expense/check/:user_id", handler.checkExpense)
	e.POST("/monthly/:user_id", handler.monthly)
	e.POST("/daily/:user_id", handler.daily)

	return e
}

func (h *Handler) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Message string }{Message: "success"})
}

func (h *Handler) checkExpense(c echo.Context) error {
	userIDD := c.Param("user_id")
	userID := uuid.MustParse(userIDD)
	user := h.UserController.GetBy(model.User{Model: model.Model{ID: &userID}})
	amountS := c.FormValue("amount")
	amount, _ := strconv.ParseFloat(amountS, 64)
	over, amount := h.ExpenseController.CheckExpense(user, amount)

	res := model.ExpenseResponse{
		OverBudget:        over,
		TotalExpense:      amount,
		Salary:            user.Salary,
		ExpensePercentage: user.Expense,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) addExpense(c echo.Context) error {
	userIDD := c.Param("user_id")
	userID := uuid.MustParse(userIDD)
	user := h.UserController.GetBy(model.User{Model: model.Model{ID: &userID}})
	typed := c.FormValue("type")
	amountS := c.FormValue("amount")
	amount, _ := strconv.ParseFloat(amountS, 64)
	res := h.ExpenseController.AddExpense(user, amount, typed)
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) monthly(c echo.Context) error {
	userIDD := c.Param("user_id")
	userID := uuid.MustParse(userIDD)
	user := h.UserController.GetBy(model.User{Model: model.Model{ID: &userID}})
	response := h.ExpenseController.Monthly(user)
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) daily(c echo.Context) error {
	userIDD := c.Param("user_id")
	userID := uuid.MustParse(userIDD)
	user := h.UserController.GetBy(model.User{Model: model.Model{ID: &userID}})
	response := h.ExpenseController.Daily(user)
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) register(c echo.Context) error {
	username := c.FormValue("username")
	sessionID := c.FormValue("session_id")
	expenseS := c.FormValue("expense")
	salaryS := c.FormValue("salary")
	salaryDateS := c.FormValue("salary_date")

	expense, _ := strconv.Atoi(expenseS)
	salaryDate, _ := strconv.Atoi(salaryDateS)
	salary, _ := strconv.ParseFloat(salaryS, 64)

	user := model.User{
		Username:   username,
		SessionID:  sessionID,
		Expense:    expense, // percentage
		Salary:     salary,
		SalaryDate: salaryDate,
	}
	userCreated := h.UserController.Register(user)
	return c.JSON(http.StatusOK, userCreated)
}

func (h *Handler) budgetPlan(c echo.Context) error {
	expenseS := c.FormValue("expense")
	userIDD := c.Param("user_id")

	expense, _ := strconv.Atoi(expenseS)
	userID := uuid.MustParse(userIDD)

	userQuery := model.User{Model: model.Model{ID: &userID}}
	userData := model.User{Expense: expense}

	userUpdated := h.UserController.Update(userQuery, userData)
	return c.JSON(http.StatusOK, userUpdated)
}
