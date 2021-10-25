package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"sync"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

)

type UserLogin struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}
type Token struct {
	token string
}
type Agent struct {
	ID         int
	PosX, PosY float32
}

var user []UserLogin


func main() {

	router := gin.Default()
	m := melody.New()

	router.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusAccepted, gin.H{"token": "tokenomg"})
	})
	router.GET("/ws/monitor", func(context *gin.Context){
		m.HandleRequest(context.Writer, context.Request)
	})
	router.GET("/ws/agent", func(context *gin.Context){
		fmt.Println("salut")
		m.HandleRequest(context.Writer, context.Request)

	})

	var monitor *melody.Session

	sessionLock := sync.Mutex{}
	sessionMap := map[*melody.Session]*Agent{}
	id := 0

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("connect", s.Request.URL.Path)
		sessionLock.Lock()

		if(s.Request.URL.Path == "/ws/monitor"){
			s.Set("isMonitor", true)
			monitor = s
		}else {
			newAgent := &Agent{
				ID: id,
				PosX: 0,
				PosY: 0,
			}
			sessionMap[s] = newAgent
			s.Set("id", newAgent.ID)
			fmt.Println("todo")
		}
		id += 1
		sessionLock.Unlock()
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		sessionLock.Lock()
		fmt.Println("Test HandleMessage")
		if (s != monitor){
			ag := sessionMap[s]
			var position map[string]float32
			if err := json.Unmarshal([]byte(msg), &position); err != nil {
				fmt.Println("err")
				return
			}
			ag.PosX = position["x"]
			ag.PosY = position["y"]
			fmt.Println(ag)
			agentUpdateBytes, err := json.Marshal(gin.H{"x" : ag.PosX, "y" : ag.PosY, "id" : ag.ID})
			if err != nil{
				fmt.Println("fail")
				return
			}
			monitor.Write(agentUpdateBytes)
		}
		sessionLock.Unlock()
	})
	router.Use(static.Serve("/", static.LocalFile("../client/public", false)))
	if err := router.Run(":8090"); err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}
}
