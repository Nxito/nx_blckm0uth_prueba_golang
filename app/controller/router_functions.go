package controller

import (
	"github.com/gin-gonic/gin"
	"golang-gin-crud-api/model" 
	h "golang-gin-crud-api/helpers" 

	"fmt"
)

// GetAllPlayers 	godoc
// @Summary			Get All Players
// @Description 	Returns a list of players
// @Tags			player
// @Success			200 {object} model.Player
// @Success 		204
// @Router			/api/players [get]
func GetPlayers(c *gin.Context) {
	data, dberr := h.GetDBPlayers()
	h.CheckError(dberr)
	if len(data) <= 0 {
		fmt.Println("No data")
		c.JSON(204, gin.H{
			"message": "Id needed",
		})
	}
	c.JSON(200, gin.H{
		"message": "Success!",
		"data":    data,
	})
}

// GetPlayerById 	godoc
// @Summary			Get Player By Id
// @Description 	Returns a player requested by his id
// @Param			id path string true "Id needed to found the player"
// @Tags			player
// @Success			200 {object} model.Player
// @Success 		204
// @Router			/api/players/{:id} [get]
func GetPlayerById(c *gin.Context) {
	id := c.Param("id")

	data, dberr := h.GetDBPlayerById(id)

	if len(data) <= 0 {
		fmt.Println("No Data", dberr)
		c.JSON(204, gin.H{
			"message": "No se ha encontrado el Player",
			"id":      id,
		})
	}

	c.JSON(200, gin.H{
		"message": "Success!",
		"id":      id,
		"data":    data,
	})
}
// GetPlayer 		godoc
// @Summary			Create Player
// @Description 	Add a new player and returns it with his id
// @Param			name body string true "Name assigned to the new player"
// @Tags			player
// @Success			201 {object} model.Player
// @failure      	400 {string}  string "error"
// @Router			/api/players [post]
func CreatePlayer(c *gin.Context) {
	var player model.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	name := player.Name

	if name == "" {
		c.JSON(400, gin.H{"error": "name required"})
		return
	}

	data, dberr := h.CreateDBPlayer(name)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
			"name":    name,
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Success!",
			"name":    name,
			"data":    data,
		})
	}

}

// UpdatePlayer 	godoc
// @Summary			Update Player Name
// @Description 	Updates a  player's name by his id and returns it
// @Param			id body string true "id to find the player"
// @Param			name body string true "Name assigned to update the player"
// @Tags			player
// @Success			201 {object} model.Player
// @Success 		204
// @failure      	400 {string}  string "error"
// @Router			/api/players [put]
func UpdatePlayer(c *gin.Context) {
	var player model.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if player.ID == "" || player.Name == "" {
		c.JSON(400, gin.H{"error": "id and name required"})
		return
	}

	data, dberr := h.UpdateDBPlayer(player)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else if len(data) == 0 {

		c.JSON(204, gin.H{
			"message": "No data!",
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Updated!",
			"data":    data,
		})
	}

}

// DeletePlayer 	godoc
// @Summary			Delete Player
// @Description 	Delete a  player by his id
// @Param			id body string true "id to find and delete the Player"
// @Tags			player
// @Success			200 {object} model.Player
// @failure      	400 {string}  string "error"
// @Router			/api/players [delete]
func DeletePlayer(c *gin.Context) {
	var player model.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if player.ID == "" {
		c.JSON(400, gin.H{"error": "Id required"})
		return
	}

	data, dberr := h.DeleteDBPlayer(player)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else {

		c.JSON(200, gin.H{
			"message": "Deleted user with ID :  " + player.ID,
			"data":    data,
		})
	}

}


// GetAllQueues 	godoc
// @Summary			Get All Queues
// @Description 	Returns a list of queues
// @Tags			queue
// @Success			200 {object} model.Queue
// @Success 		204
// @Router			/api/queues [get]
func GetQueues(c *gin.Context) {

	data, dberr := h.GetDBQueues()
	h.CheckError(dberr)

	if len(data) <= 0 {
		fmt.Println("No Data")
		c.JSON(204, gin.H{
			"message": "Id needed",
		})
	}
	fmt.Println("Connected!")

	c.JSON(200, gin.H{
		"message": "Success!",
		"data":    data,
	})

}
// GetQueueById 	godoc
// @Summary			Get Queues By Id
// @Description 	Returns a queue requested by his id
// @Param			id path string true "Id needed to found the queue"
// @Tags			queue
// @Success			200 {object} model.Queue
// @Success 		204

// @Router			/api/queues/{:id} [get]
func GetQueueById(c *gin.Context) {
	id := c.Param("id")

	data, dberr := h.GetDBQueueById(id)

	if len(data) <= 0 {
		fmt.Println("No Data", dberr)
		c.JSON(204, gin.H{
			"message": "Queue not found",
			"id":      id,
		})
	}

	c.JSON(200, gin.H{
		"message": "Success!",
		"id":      id,
		"data":    data,
	})

}
// CreateQueue 	godoc
// @Summary			Create Queue
// @Description 	Add a new queue and returns it with his id. It cannot create more than the value specified on config.json: "max_queues"
// @Param			name body string true "Name assigned to the new player"
// @Param			max_players body int false "Value assigned to limit the payers in this queue. Predet: 2"
// @Tags			queue
// @Success			201 {object} model.Queue
// @failure      	400 {string}  string "error"
// @Router			/api/queues [post]
func CreateQueue(c *gin.Context) {
	var queue model.Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if queue.Name == "" {
		c.JSON(400, gin.H{"error": "name required"})
		return
	}
	if queue.MaxPlayers == 0 {
		c.JSON(400, gin.H{"error": "El parámetro 'maxPlayers' es requerido"})
		return
	}

	data, dberr := h.CreateDBQueue(queue)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message":     dberr,
			"name":        queue.Name,
			"max_players": queue.MaxPlayers,
		})
	} else if len(data) == 0 {

		c.JSON(400, gin.H{
			"message": "Limit queues exceeded",
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Success!",
			"data":    data,
		})
	}

}
// UpdateQueue 	godoc
// @Summary			Update Queue Name
// @Description 	Updates a  queue's name by his id and returns it
// @Param			id body string true "id to find the queue"
// @Param			name body string true "Name assigned to update the queue"
// @Tags			queue
// @Success			201 {object} model.Queue
// @Success 		204
// @failure      	400 {string}  string "error"
// @Router			/api/queues [put]
func UpdateQueue(c *gin.Context) {
	var queue model.Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if queue.ID == "" || queue.Name == "" {
		c.JSON(400, gin.H{"error": "id and name required"})
		return
	}

	data, dberr := h.UpdateDBQueue(queue)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else if len(data) == 0 {

		c.JSON(204, gin.H{
			"message": "No data!",
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Updated!",
			"data":    data,
		})
	}

}

// DeleteQueue 	godoc
// @Summary			Delete Queue
// @Description 	Delete a queue by his id
// @Param			id body string true "id to find and delete the queue"
// @Tags			queue
// @Success			201 {object} model.Queue
// @failure      	400 {string}  string "error"
// @Router			/api/queues [delete]
func DeleteQueue(c *gin.Context) {
	var queue model.Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if queue.ID == "" {
		c.JSON(400, gin.H{"error": "Id required"})
		return
	}

	data, dberr := h.DeleteDBQueue(queue)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else {

		c.JSON(200, gin.H{
			"message": "Deleted user with ID :  " + queue.ID,
			"data":    data,
		})
	}

}

// GetAllSessions 	godoc
// @Summary			Get All Sessions
// @Description 	Returns a list of sessions
// @Tags			session
// @Success			200 {object} model.Session
// @Success 		204 
// @Router			/api/sessions [get]
func GetSessions(c *gin.Context) {

	data, dberr := h.GetDBSessions()
	h.CheckError(dberr)

	if len(data) <= 0 {
		fmt.Println("No Data")
		c.JSON(204, gin.H{
			"message": "Id needed",
		})
	}
	fmt.Println("Connected!")

	c.JSON(200, gin.H{
		"message": "Success!",
		"data":    data,
	})

}
// GetSessionByStatus 	godoc
// @Summary			Get a list of sessions by status
// @Description 	Returns a list of sessions requested by his status: opened or closed
// @Param			status path string true "Id needed to found the sessions"
// @Tags			session
// @Success			200 {object} model.Session
// @Success 		204 
// @Router			/api/queues/{:status} [get]
func GetSessionByStatus(c *gin.Context) {
	status := c.Param("status")

	data, dberr := h.GetDBSessionByStatus(status)

	if len(data) <= 0 {
		fmt.Println("No Data", dberr)
		c.JSON(204, gin.H{
			"message": "Queue not found",
			"id":      status,
		})
	}

	c.JSON(200, gin.H{
		"message": "Success!",
		"id":      status,
		"data":    data,
	})

}
// CreateSession 	godoc
// @Summary			Create a Session
// @Description 	Creates a Session with one Player and an asociated queue. Returns the ID of the game session and the related data
// @Param			queue body string true "an Id to find the asociated queue"
// @Param			player body string true "an Id to add a player"
// @Tags			session
// @Success			201 {object} model.Session
// @failure      	400 {string}  string "error"
// @Router			/api/sessions [post]
func CreateSession(c *gin.Context) {
	var session model.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if session.Queue == "" {
		c.JSON(400, gin.H{"error": "Parameter 'queue' is required "})
		return
	}
	if session.Player == "" {
		c.JSON(400, gin.H{"error": "Parameter 'player'  is required "})
		return
	}
	session.Status = "opened"

	data, dberr := h.CreateDBSession(session)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
			"queue":   session.Queue,
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Success!",
			"data":    data,
		})
	}

}
// UpdateSession 	godoc
// @Summary			Update Session with a new Player
// @Description 	Updates the session adding a new Player. Returns the ID of the game session and the related data
// @Param			queue body string true "id to find the asociated queue"
// @Param			name body string true "Name assigned to update the queue"
// @Tags			session
// @Success			201 {object} model.Session
// @Success 		204
// @failure      	400 {string}  string "error"
// @Router			/api/sessions [put]
func UpdateSession(c *gin.Context) {
	var session model.Session

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if session.Queue == "" || session.Player == "" {
		c.JSON(400, gin.H{"error": " 'queue' and 'player' needed "})
		return
	}
	if session.Status != "closed" {
		session.Status = "opened"
	}

	data, dberr := h.UpdateDBSession(session)

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else if len(data) == 0 {

		c.JSON(204, gin.H{
			"message": "No data!",
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Updated! Remember: if closed update is gonna send latest data",
			"data":    data,
		})
	}

}

