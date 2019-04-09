package repo

import (
	"errors"

	"bitbucket.org/yesboss/sharingan/config"
	"github.com/jinzhu/gorm"

	"bitbucket.org/yesboss/sharingan/model"
)

type UserRepo struct {
	Conn   *gorm.DB
	Config config.Config
}

func NewUserRepo(config config.Config, db *gorm.DB) (UserRepo, error) {
	db.AutoMigrate(model.User{}, model.Expense{})
	return UserRepo{
		Config: config,
		Conn:   db,
	}, nil
}

func (r *UserRepo) Create(user model.User) model.User {
	r.Conn.Save(&user)
	return user
}

func (r *UserRepo) GetByID(id string) (role model.User) {
	r.Conn.Where("id = ?", id).First(&role)
	return role
}

func (r *UserRepo) GetBy(query model.User) model.User {
	var user model.User
	r.Conn.Where(query).First(&user)
	return user
}

func (r *UserRepo) List(page int, limit int) (roles []model.User) {
	r.Conn.
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&roles)

	return roles
}

func (r *UserRepo) ListBy(query model.User, page int, limit int) (users []model.User) {
	r.Conn.
		Offset((page - 1) * limit).
		Limit(limit).
		Where(query).
		Find(&users)

	return users
}

func (r *UserRepo) Update(query model.User, data model.User) model.User {
	r.Conn.
		Where(query).
		First(&query)

	r.Conn.Model(&query).Update(data)
	return query
}

func (r *UserRepo) Delete(role *model.User) (bool, error) {
	// WARNING When delete a record,
	// you need to ensure it's primary field has value, and GORM will use the primary key to delete the record,
	// if primary field's blank, GORM will delete all records for the model

	if role.ID == nil {
		return false, errors.New("ID cannot be empty")
	}

	r.Conn.Delete(&role)
	return true, nil
}
