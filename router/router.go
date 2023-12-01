package router
import (
	"github.com/gin-gonic/gin"
	h "golang-gin-crud-api/controller"
	docs "golang-gin-crud-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
/* Router */

// @title Golang Gin CRUD API
// @version	1.0
// @description A testing project fot BlackMouth
// @host 	localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/" 
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))


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

	// Routes for Sessions

	router.GET("/api/sessions", h.GetSessions)
	router.GET("/api/sessions/:status", h.GetSessionByStatus)
	router.POST("/api/sessions", h.CreateSession)
	router.PUT("/api/sessions", h.UpdateSession)
	return router
}
