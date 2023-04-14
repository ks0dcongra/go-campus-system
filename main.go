package main

import (
	"example1/app/http/middleware"
	database "example1/database"
	migration "example1/database/migrations"
	"example1/routes"
	"fmt"
	"log"

	// "log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	// 連接DB
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
		time.Sleep(3 * time.Second)
	}
	fmt.Println("Database connected!")

	// 連接伺服器
	mainServer := gin.New()

	// 定義router呼叫格式
	mainServer.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// 連接Router
	routes.ApiRoutes(mainServer)

	//Migration Init
	migration.Init()

	// 註冊Validator Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middleware.UserPasd)
	}

	// go func() {
	//     if err := mainServer.RunTLS(":443","./cert/server.pem", "./cert/server.key"); err != nil {
	//         log.Fatal("HTTPS service failed: ", err)
	//     }
	// }()
	if err := mainServer.Run(":8080"); err != nil {
		log.Fatal("HTTP service failed: ", err)
	}
}
