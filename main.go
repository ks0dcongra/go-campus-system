package main

import (
	database "example1/database"
	"example1/routes"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//DBConnect
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		database.UserName,
		database.Password,
		database.Addr,
		database.Port,
		database.Database,
	)
	var err error
	println(dsn)

	for {
		database.DB, err = database.DBInit(dsn)
		if err == nil {
			break
		}
		
		fmt.Println("Trying to connect database...")
		fmt.Println("DB Error===>", err)
		time.Sleep(3* time.Second)
	}
	fmt.Println("Database connected!")	

	mainServer := gin.New()
	mainServer.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	routes.ApiRoutes(mainServer)
	mainServer.Run(":9875")
}
