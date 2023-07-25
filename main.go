package main

import (
	"fmt"
	"testing-gin/config"
	"testing-gin/controller"
	"testing-gin/dao"
	"testing-gin/router"
	"testing-gin/service"
)

func main() {
	serveApplication()
}

func serveApplication() {
	// Database ORM
	dbOrm := config.DatabaseConnectionOrm()

	// Database Plain
	db := config.DatabaseConnection()

	// DAO
	testDao := dao.NewPersonDaoImpl(db, dbOrm)

	// Service
	testService := service.NewTestServiceImpl(testDao)

	// Controller
	testController := controller.NewTestController(testService)

	// Router
	routes := router.NewRouter(testController)

	routes.Run(":8082")

	fmt.Println("Server running on port 8082")
}
