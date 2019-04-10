package repo

import (
	"errors"
	"time"

	"bitbucket.org/yesboss/sharingan/config"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"bitbucket.org/yesboss/sharingan/model"
)

type ExpenseRepo struct {
	Conn   *gorm.DB
	Config config.Config
}

func NewExpenseRepo(config config.Config, db *gorm.DB) (ExpenseRepo, error) {
	return ExpenseRepo{
		Config: config,
		Conn:   db,
	}, nil
}

func (r *ExpenseRepo) Create(ex model.Expense) model.Expense {
	r.Conn.Save(&ex)
	return ex
}

func (r *ExpenseRepo) GetByID(id string) (role model.Expense) {
	r.Conn.Where("id = ?", id).First(&role)
	return role
}

func (r *ExpenseRepo) TotalExpense(userID uuid.UUID, start time.Time, end time.Time) float64 {
	type Result struct {
		Total float64
	}

	var total Result

	r.Conn.
		Table("expenses").
		Select("sum(amount) total").
		Where("created_at >= ? AND created_at <= ? AND user_id = ?", start, end, userID).
		Scan(&total)

	return total.Total
}

func (r *ExpenseRepo) TotalExpenseDaily(userID uuid.UUID) float64 {
	type Result struct {
		Total float64
	}

	var total Result

	r.Conn.
		Table("expenses").
		Select("sum(amount) total").
		Where("user_id = ? AND date(created_at) = CURDATE()", userID).
		Scan(&total)

	return total.Total
}

func (r *ExpenseRepo) List(page int, limit int) (roles []model.Expense) {
	r.Conn.
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&roles)

	return roles
}

func (r *ExpenseRepo) GetMonthly(user model.User, start time.Time, end time.Time) model.MonthlyResponses {

	var m []model.MonthlyResponse
	var ms model.MonthlyResponses

	r.Conn.
		Table("expenses").
		Select("sum(amount) total, `type`").
		Where("created_at >= ? AND created_at <= ? AND user_id = ?", start, end, user.Model.ID).
		Group("type").
		Scan(&m)

	ms.Data = m
	ms.Total = r.TotalExpense(*user.Model.ID, start, end)

	batas := float64(user.Expense) / 100.0 * user.Salary
	ms.ExpensePercentage = ms.Total / batas * 100.0

	return ms
}

func (r *ExpenseRepo) GetDaily(user model.User) model.MonthlyResponses {

	var m []model.MonthlyResponse
	var ms model.MonthlyResponses

	r.Conn.
		Table("expenses").
		Select("sum(amount) total, type").
		Where("user_id = ? AND date(created_at) = CURDATE()", user.Model.ID).
		Group("type").
		Scan(&m)

	ms.Data = m
	ms.Total = r.TotalExpenseDaily(*user.Model.ID)

	batas := float64(user.Expense) / 100.0 * user.Salary
	ms.ExpensePercentage = ms.Total / batas * 100.0

	return ms
}

func (r *ExpenseRepo) ListBy(query model.Expense, page int, limit int) (users []model.Expense) {
	r.Conn.
		Offset((page - 1) * limit).
		Limit(limit).
		Where(query).
		Find(&users)

	return users
}

func (r *ExpenseRepo) Update(role *model.Expense) {
	r.Conn.Save(&role)
}

func (r *ExpenseRepo) Delete(role *model.Expense) (bool, error) {
	// WARNING When delete a record,
	// you need to ensure it's primary field has value, and GORM will use the primary key to delete the record,
	// if primary field's blank, GORM will delete all records for the model

	if role.ID == nil {
		return false, errors.New("ID cannot be empty")
	}

	r.Conn.Delete(&role)
	return true, nil
}
