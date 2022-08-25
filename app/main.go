package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var pwd = os.Getenv("CHOPPA_PASSWORD")

func main() {
	db.connect()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", getChop)
	router.GET("/:path", getChop)
	router.POST("/", setChop)
	router.POST("/:path", setChop)

	router.Run("0.0.0.0:8080")
}

func getChop(c *gin.Context) {
	path := c.Param("path")

	userAgent := c.GetHeader("User-Agent")
	platform, _ := platformFromUserAgent(userAgent)
	url, err := db.getChop(path, platform)
	if err != nil {
		defaultUrl, sErr := db.getChop(path, 0)
		if sErr != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "path not found"})
			return
		}
		url = defaultUrl
	}
	c.Redirect(http.StatusFound, url)
}

func setChop(c *gin.Context) {
	if !authorize(c) {
		return
	}

	path := c.Param("path")
	url := c.PostForm("url")
	platformString := c.PostForm("platform")

	var platform int
	if platformString == "" {
		platform = 0
	} else {
		tempPlatform, pErr := platformFromString(platformString)
		if pErr != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "platform not found"})
			return
		}
		platform = tempPlatform
	}

	err := db.updateChop(path, platform, url)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

func authorize(c *gin.Context) bool {
	auth := c.PostForm("auth")
	if auth != pwd {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "authorization failed"})
		return false
	}
	return true
}
