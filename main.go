package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/crud_auth?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/signup", showSignupPage)
	r.POST("/signup", handleSignup)

	r.GET("/login", showLoginPage)
	r.POST("/login", handleLogin)

	Goprac()
	r.Run(":8083")

}

func showSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func handleSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := User{Username: username, Password: password}
	result := db.Create(&user)

	if result.Error != nil {
		c.String(http.StatusInternalServerError, "Signup failed")
		return
	}

	c.String(http.StatusOK, "Signup successful!")
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func handleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user User
	result := db.Where("username = ? AND password = ?", username, password).First(&user)

	if result.Error != nil {
		c.String(http.StatusUnauthorized, "Invalid login credentials")
		return
	}

	c.String(http.StatusOK, "Login successful!")
}
