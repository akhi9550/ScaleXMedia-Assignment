package di

import (
	server "scalexmedia-assignment/pkg/api"
	"scalexmedia-assignment/pkg/api/handlers"
	"scalexmedia-assignment/pkg/config"
	"scalexmedia-assignment/pkg/db"
	"scalexmedia-assignment/pkg/helper"
	"scalexmedia-assignment/pkg/repository"
	"scalexmedia-assignment/pkg/usecase"
	"scalexmedia-assignment/pkg/utils/models"
)

var users = make(map[string]models.UserStruct)
var admin = make(map[string]models.AdminStruct)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	initializeData()
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	adminRepo := repository.NewAdminRepository(gormDB)
	helperAdmin := helper.NewAdminHelper(cfg)
	adminUseCase := usecase.NewAdminUseCase(adminRepo, helperAdmin)
	adminHandler := handlers.NewAdminHandler(adminUseCase, admin)

	helperUser := helper.NewUserHelper(cfg)
	userUsecase := usecase.NewUserUseCase(helperUser)
	userHandler := handlers.NewUserHandler(userUsecase, users)

	serverHTTP := server.NewServerHTTP(adminHandler, userHandler)

	return serverHTTP, nil
}

func initializeData() {
	users["akhil@gmail.com"] = models.UserStruct{Name: "Akhil", Password: "955055"}
	users["anjin@gmail.com"] = models.UserStruct{Name: "Ajin", Password: "955155"}
	users["amal@gmail.com"] = models.UserStruct{Name: "Amal", Password: "955255"}
	admin["libraryadmin@gmail.com"] = models.AdminStruct{Name: "Admin", Password: "999999"}
}
