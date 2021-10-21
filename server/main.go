package main

import (
	"log"

	"github.com/gin-contrib/static"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}
type Token struct {
	token string
}

var user []UserLogin

func UserReaderHandler(u *gin.Context) {
	fmt.Println("bonjour")
}

func main() {

	// router := gin.New()
	router := gin.Default()
	// m := melody.New()

	router.POST("/login", func(context *gin.Context) {
		user := UserLogin{}
		// using BindJson method to serialize body with struct
		if err := context.BindJSON(&user); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusAccepted, gin.H{"token": "tokenomg"})
		// context.JSON(http.StatusAccepted, &user)

	})

	router.Use(static.Serve("/", static.LocalFile("../client/public", false)))
	if err := router.Run(":8090"); err != nil {
		log.Fatal(err)
	}

}
