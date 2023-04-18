package main

import (
	"example1/app/http/middleware"
	database "example1/database"
	migration "example1/database/migrations"
	"example1/routes"
	"fmt"
	"log"
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

// Controller測試，借放。
// func Test_sample(t *testing.T) {
// 	// 首先一樣創建一個與main相同的router
// 	router := gin.New()
// 	router.GET("/hello", func(c *gin.Context) {
// 		fmt.Printf("c.Request.Header: %v\n", c.Request.Header)
// 	}, func(c *gin.Context) {
// 		c.JSON(200, gin.H{"hello": "world"})
// 	})

// 	// 建立測試用的 HTTP 回應紀錄器
// 	w := httptest.NewRecorder()
// 	// 建立測試用的 HTTP 請求
// 	req, _ := http.NewRequest("GET", "/hello", nil)
// 	req.Header.Set("Authorization", "Bearer testToken")

// 	// gin 有提供 ServeHTTP 的 function 來讓使用者模擬請求丟入的情況，將Request /hc的Response給conform到ResponseRecorder
// 	router.ServeHTTP(w, req)

// 	expectedStatus := http.StatusOK
// 	// 然後比較Response Code、Response Text等內容是否符合預期
// 	assert.Equal(t, expectedStatus, w.Code)

// 	// fmt.Println("Body: %+v", w.Body)
// }
