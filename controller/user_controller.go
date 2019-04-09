package controller

import (
	"bitbucket.org/yesboss/sharingan/config"
	"bitbucket.org/yesboss/sharingan/model"
	"bitbucket.org/yesboss/sharingan/repo"
)

type UserController struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewUserController(config config.Config, userRepo repo.UserRepo) UserController {
	return UserController{Config: config, UserRepo: userRepo}
}

func (c *UserController) Register(user model.User) model.User {
	return c.UserRepo.Create(user)
}

func (c *UserController) GetBy(query model.User) model.User {
	return c.UserRepo.GetBy(query)
}

func (c *UserController) Update(query model.User, data model.User) model.User {
	return c.UserRepo.Update(query, data)
}
