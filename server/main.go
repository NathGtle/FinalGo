package main

import (
	"log"

	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
)

type Agent struct {
	MemberID   string
	PosX, PosY uint
}

var agents []Agent

func AgentAddHandler(c *gin.Context) {
	agentID := c.Param("agentID")

	agents = append(agents, Agent{agentID, 100, 200})
}

func AgentReaderHandler(c *gin.Context) {
	c.JSON(200, agents)
}

func main() {
	router := gin.Default()
	router.GET("/join/:agentID", AgentAddHandler)

	router.GET("/agents", AgentReaderHandler)

	// if Allow DirectoryIndex
	// router.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	// router.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

	router.Use(static.Serve("/", static.LocalFile("/public/index.html", false)))
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	// Listen and Server in 0.0.0.0:8080
	if err := router.Run(":8090"); err != nil {
		log.Fatal("err")
	}
	// port := "8090"

	// log.Println("Serving on port:", port)
	// router.Run(":" + port)
}
