package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

var db *gorm.DB

func initDB() {
	var err error
	// DSN (Data Source Name)
	dsn := "root:123456@tcp(db:3306)/meetyou?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移（创建表）
	db.AutoMigrate(&User{})

	// 插入一条数据（如果不存在）
	db.FirstOrCreate(&User{}, User{Name: "John Doe"})
}

func main() {
	initDB()
	router := gin.Default()
	router.GET("/", UserHandle)
	router.GET("/api/user", UserHandle)

	log.Println("Server starting on port 8080...")
	router.Run(":8080")
}

func UserHandle(c *gin.Context) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*") // 设置 CORS
	var user User
	result := db.First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
