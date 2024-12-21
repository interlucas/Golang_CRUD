package main

import (
	"github.com/interlucas/Golang_CRUD/src/controller"
	"github.com/interlucas/Golang_CRUD/src/model/repository"
	"github.com/interlucas/Golang_CRUD/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}