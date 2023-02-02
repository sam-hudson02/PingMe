package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	creds := getCreds()
	// Start Discord bot
	bot, err := NewBot(creds.token)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bot.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/send", func(c *gin.Context) {
		// get message from request body
		var req sendRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Printf("message is empty")
			c.JSON(400, gin.H{
				"message": "message is empty",
			})
			return
		}

		fmt.Printf("message: %s", req.Message)
		message := req.Message

		err := bot.Send(creds.user_id, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": "sent",
		})
	})

	r.POST("/build-notify", func(c *gin.Context) {
		// get build data from request body
		var req BuildRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Printf("build data invalid")
			c.JSON(400, gin.H{
				"message": "build data invalid",
			})
			return
		}

		fmt.Printf("build data: %s", req)
		err := bot.BuildPipelineNotify(creds.build_channel, req)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error sending build notification",
			})
			fmt.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": "sent",
		})
		return
	})

	fmt.Println("Starting server on port 5000")
	r.Run("0.0.0.0:5000") // listen and serve on
}

type sendRequest struct {
	Message string `json:"message" binding:"required"`
}

type BuildRequest struct {
	BuildID       string `json:"build_id" binding:"required"`
	BuildName     string `json:"build_name" binding:"required"`
	BuildURL      string `json:"build_url" binding:"required"`
	BuildDate     string `json:"build_date" binding:"required"`
	CommitAuthor  string `json:"commit_author" binding:"required"`
	CommitURL     string `json:"commit_url" binding:"required"`
	CommitTitle   string `json:"commit_title" binding:"required"`
	BuildResult   string `json:"build_result" binding:"required"`
	BuildDuration string `json:"build_duration" binding:"required"`
}

type creds struct {
	user_id       string
	token         string
	build_channel string
}

// get creds from .env file
func getCreds() creds {
	err := godotenv.Load("./secret/conf.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return creds{
		user_id:       os.Getenv("DISCORD_USER_ID"),
		token:         os.Getenv("DISCORD_TOKEN"),
		build_channel: os.Getenv("BUILD_CHANNEL_ID"),
	}
}
