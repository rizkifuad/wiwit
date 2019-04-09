//+build wireinject

package main

import (
	"bitbucket.org/yesboss/sharingan/common"
	"bitbucket.org/yesboss/sharingan/config"
	"bitbucket.org/yesboss/sharingan/controller"
	"bitbucket.org/yesboss/sharingan/delivery/grpc"
	"bitbucket.org/yesboss/sharingan/delivery/http"
	"bitbucket.org/yesboss/sharingan/repo"
	"github.com/google/wire"
)

func InitializeContainer() (App, error) {
	wire.Build(
		newApp,
		common.NewMysqlConnection,
		config.New,
		controller.NewUserController,
		controller.NewExpenseController,
		repo.NewUserRepo,
		repo.NewExpenseRepo,
		http.New,
		grpc.New)

	return App{}, nil
}
