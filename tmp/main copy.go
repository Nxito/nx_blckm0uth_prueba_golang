package asd

import (
	"github.com/gin-gonic/gin"
)
func getPlayers(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		c.JSON(200, gin.H{
			"message": "Hello from Route 1 with parameter",
			"id":      id,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Hello from Route 1 with no parameter",
		})
	}
}

func createPlayers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Route 1!",
	})
}

func updatePlayers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Route 1!",
	})
}

func deletePlayers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Route 1!",
	})
}


func main() {
	router := gin.Default()
	router.GET("/api/players", getPlayers)
	router.GET("/api/players/:id", getPlayers)
	router.POST("/api/players", createPlayers)
	router.PUT("/api/players", updatePlayers)
	router.DELETE("/api/players", deletePlayers)


	// http.ListenAndServe(":8080", router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
