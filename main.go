package main

import (
	"crypto/tls"
	"example1/app/http/middleware"
	database "example1/database"
	migration "example1/database/migrations"
	"example1/routes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
		time.Sleep(3 * time.Second)
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

	//Migration Init
	migration.Init()

	// 註冊Validator Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middleware.UserPasd)
	}

	// TLS連線
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	
	srv := &http.Server{
		Addr:      ":8080",
		TLSConfig: cfg,
		Handler:   mainServer,
	}

	// 啟動 TLS 服務器
	if err := srv.ListenAndServeTLS("cert/server.pem", "cert/server.key"); err != nil {
		panic(err)
	}

	// mainServer.Run(":8080")
}
