package router
import (
	"github.com/gin-gonic/gin"
	h "golang-gin-crud-api/helpers"
)
/* Router */
func SetupRouter() *gin.Engine {

	router := gin.Default()
	// Routes for Players

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.GET("/api/players", h.GetPlayers)
	router.GET("/api/players/:id", h.GetPlayerById)
	router.POST("/api/players", h.CreatePlayer)
	router.PUT("/api/players", h.UpdatePlayer)
	router.DELETE("/api/players", h.DeletePlayer)

	// Routes for Queues

	router.GET("/api/queues", h.GetQueues)
	router.GET("/api/queues/:id", h.GetQueueById)
	router.POST("/api/queues", h.CreateQueue)
	router.PUT("/api/queues", h.UpdateQueue)
	router.DELETE("/api/queues", h.DeleteQueue)

	// router.GET("/api/parameters/limitqueue", h.getLimitQueue)
	// router.PUT("/api/parameters/limitqueue", h.updateLimitQueue)

	// Routes for Sessions

	router.GET("/api/sessions", h.GetSessions)
	router.GET("/api/sessions/:status", h.GetSessionByStatus)
	router.POST("/api/sessions", h.CreateSession)
	router.PUT("/api/sessions", h.UpdateSession)
	return router
}
