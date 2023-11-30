package helpers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-crud-api/model" 

	"fmt"
)

/* Funciones  de respuesta al cliente */

// GetAllPlayers 	godoc
// @Summary			Get All Players
// @Description 	Returns a list of players
// @Tags			
// @Success			200 {object} response.Response{}
// @Router			/api/players [get]

func GetPlayers(c *gin.Context) {

	data, dberr := GetDBPlayers()
	CheckError(dberr)
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

func GetPlayerById(c *gin.Context) {
	id := c.Param("id")

	data, dberr := GetDBPlayerById(id)

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

	data, dberr := CreateDBPlayer(name)

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

	data, dberr := UpdateDBPlayer(player)

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

	data, dberr := DeleteDBPlayer(player)

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
func GetQueues(c *gin.Context) {

	data, dberr := GetDBQueues()
	CheckError(dberr)

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
func GetQueueById(c *gin.Context) {
	id := c.Param("id")

	data, dberr := GetDBQueueById(id)

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

	data, dberr := CreateDBQueue(queue)

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

	data, dberr := UpdateDBQueue(queue)

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

	data, dberr := DeleteDBQueue(queue)

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

func GetSessions(c *gin.Context) {

	data, dberr := GetDBSessions()
	CheckError(dberr)

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
func GetSessionByStatus(c *gin.Context) {
	status := c.Param("status")

	data, dberr := GetDBSessionByStatus(status)

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

	data, dberr := CreateDBSession(session)

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

	data, dberr := UpdateDBSession(session)

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

