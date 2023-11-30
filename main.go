package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// conexion para la  DB
const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "password"
	dbname   = "mytestdb"
)

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Queue struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"maxPlayers"`
}

type Session struct {
	ID     string `json:"id"`
	Queue  string `json:"queue"`
	Player string `json:"player"`
	Status string `json:"status"`
}
type Parameter struct {
	Name  string
	Value int
}

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
func setupRouter() *gin.Engine {

	router := gin.Default()
	// Routes for Players

	router.GET("/api/players", getPlayers)
	router.GET("/api/players/:id", getPlayerById)
	router.POST("/api/players", createPlayer)
	router.PUT("/api/players", updatePlayer)
	router.DELETE("/api/players", deletePlayer)

	// Routes for Queues

	router.GET("/api/queues", getQueues)
	router.GET("/api/queues/:id", getQueueById)
	router.POST("/api/queues", createQueue)
	router.PUT("/api/queues", updateQueue)
	router.DELETE("/api/queues", deleteQueue)

	router.GET("/api/parameters/limitqueue", getLimitQueue)
	router.PUT("/api/parameters/limitqueue", updateLimitQueue)

	// Routes for Sessions

	router.GET("/api/sessions", getSessions)
	router.GET("/api/sessions/:status", getSessionByStatus)
	router.POST("/api/sessions", createSession)
	router.PUT("/api/sessions", updateSession)
	return router
}


func main() {
	r := setupRouter()
	r.Run(":8080")
}

func getPlayers(c *gin.Context) {

	data, dberr := getDBPlayers()
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
func getPlayerById(c *gin.Context) {
	id := c.Param("id")

	data, dberr := getDBPlayerById(id)

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

func createPlayer(c *gin.Context) {
	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	name := player.Name

	if name == "" {
		c.JSON(400, gin.H{"error": "name required"})
		return
	}

	data, dberr := createDBPlayer(name)

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
func updatePlayer(c *gin.Context) {
	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if player.ID == "" || player.Name == "" {
		c.JSON(400, gin.H{"error": "id and name required"})
		return
	}

	data, dberr := updateDBPlayer(player)

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
func deletePlayer(c *gin.Context) {
	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if player.ID == "" {
		c.JSON(400, gin.H{"error": "Id required"})
		return
	}

	data, dberr := deleteDBPlayer(player)

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
func getQueues(c *gin.Context) {

	data, dberr := getDBQueues()
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
func getQueueById(c *gin.Context) {
	id := c.Param("id")

	data, dberr := getDBQueueById(id)

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
func createQueue(c *gin.Context) {
	var queue Queue
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

	data, dberr := createDBQueue(queue)

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
func updateQueue(c *gin.Context) {
	var queue Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if queue.ID == "" || queue.Name == "" {
		c.JSON(400, gin.H{"error": "id and name required"})
		return
	}

	data, dberr := updateDBQueue(queue)

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
func deleteQueue(c *gin.Context) {
	var queue Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if queue.ID == "" {
		c.JSON(400, gin.H{"error": "Id required"})
		return
	}

	data, dberr := deleteDBQueue(queue)

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

func getSessions(c *gin.Context) {

	data, dberr := getDBSessions()
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
func getSessionByStatus(c *gin.Context) {
	status := c.Param("status")

	data, dberr := getDBSessionByStatus(status)

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
func createSession(c *gin.Context) {
	var session Session
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

	data, dberr := createDBSession(session)

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
func updateSession(c *gin.Context) {
	var session Session

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

	data, dberr := updateDBSession(session)

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

func getDBPlayers() ([]Player, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "name", "id" FROM public.players`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Player
	for rows.Next() {
		var data Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func getDBPlayerById(id string) ([]Player, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`SELECT "name", "id" FROM public.players where id = '%s'`, id)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Player
	for rows.Next() {
		var data Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func createDBPlayer(name string) ([]Player, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO public.players(name) VALUES ('%s') RETURNING id,name; `, name)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Player
	for rows.Next() {
		var data Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func updateDBPlayer(data Player) ([]Player, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE  public.players SET  name= '%s' WHERE id = '%s' RETURNING id,name; `, data.Name, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Player
	for rows.Next() {
		var data Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func deleteDBPlayer(data Player) ([]Player, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM public.players WHERE id = '%s' RETURNING name; `, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Player
	for rows.Next() {
		var data Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func getDBQueues() ([]Queue, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "name", "id", "max_players" FROM public.queues`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Queue
	for rows.Next() {
		var data Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func getDBQueueById(id string) ([]Queue, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`SELECT "name", "id", "max_players" FROM public.queues where id = '%s'`, id)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Queue
	for rows.Next() {
		var data Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func createDBQueue(queue Queue) ([]Queue, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`
		WITH filecount AS (
			SELECT COUNT(*) AS fileamount FROM public.queues
		)
		INSERT INTO public.queues(name,max_players)
		SELECT '%s','%d'
		FROM filecount
		WHERE fileamount < (SELECT value::bigint FROM public.parameters WHERE name = 'max_queues')
		
		RETURNING id, name, max_players;
	`, queue.Name, queue.MaxPlayers)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Queue
	for rows.Next() {
		var data Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func updateDBQueue(data Queue) ([]Queue, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE  public.queues SET  name= '%s' WHERE id = '%s' RETURNING id, name, max_players; `, data.Name, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Queue
	for rows.Next() {
		var data Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func deleteDBQueue(data Queue) ([]Queue, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM public.queues WHERE id = '%s' RETURNING name;`, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Queue
	for rows.Next() {
		var data Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func getDBSessions() ([]Session, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "id_queue", "players", "status" FROM public.sessions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Session
	for rows.Next() {
		var data Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func getDBSessionByStatus(status string) ([]Session, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`SELECT "id", "id_queue", "players", "status" FROM public.sessions where status = '%s'`, status)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Session
	for rows.Next() {
		var data Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func createDBSession(session Session) ([]Session, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`
	INSERT INTO public.sessions(
		id_queue, players, status)
	   VALUES (
		   '{%s}',
		   '{%s}',
		   '%s')
	RETURNING id, id_queue, players, status ;
	`, session.Queue, session.Player, session.Status)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Session
	for rows.Next() {
		var data Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func updateDBSession(data Session) ([]Session, error) {

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`
	UPDATE public.sessions
	SET players = CASE
					 WHEN status = 'opened' THEN players || '{%s}'
					 ELSE players
				  END,
		status = CASE
					WHEN status = 'opened' THEN '%s'
					ELSE status
				 END
	WHERE id_queue = '%s' 
	RETURNING id, id_queue, players, status;
	
	`, data.Player, data.Status, data.Queue)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []Session
	for rows.Next() {
		var data Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func getLimitQueue(c *gin.Context) {
	var parameter Parameter

	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, dberr := getDBLimitQueue()
	// close database

	fmt.Println("Connected!")

	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Updated!",
			"data":    data,
		})
	}

}
func getDBLimitQueue() (Parameter, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	var param Parameter
	row := db.QueryRow(`SELECT "name", "value" FROM public.parameters WHERE name = 'max_queues'`)

	err = row.Scan(&param.Name, &param.Value)
	if err != nil {
		return param, err
	}

	return param, nil
}

func updateLimitQueue(c *gin.Context) {
	var parameter Parameter

	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := updateDBLimitQueue(parameter)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Updated!",
			"data":    data,
		})
	}

}
func updateDBLimitQueue(data Parameter) (Parameter, error) {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	var param Parameter
	query := fmt.Sprintf(`UPDATE public.parameters SET  value = '%d' FROM  WHERE name = 'max_queues' RETURN name,value`, data.Value)

	row := db.QueryRow(query)

	err = row.Scan(&param.Name, &param.Value)
	if err != nil {
		return param, err
	}

	return param, nil
}
