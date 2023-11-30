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
	MaxPlayers int `json:"maxPlayers"`
}

type Session struct {
	ID      string `json:"id"`
	Queue   string `json:"queue"`
	Players string `json:"players"`
	Status  string `json:"status"`
}
type Parameter struct {
	Name  string
	Value int
}
// var players = []Player{
//     {ID: "1",  Name: "John Smith"},
//     {ID: "2",  Name: "Gary Oak"},
//     {ID: "3",  Name: "Tomato49"},
// }

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)


func main() {

	// connection string

	router := gin.Default()
	// Routes for Players

	router.GET("/api/players",getPlayers);
	router.GET("/api/players/:id", getPlayerById )
	router.POST("/api/players",createPlayer)
	router.PUT("/api/players",updatePlayer)
	router.DELETE("/api/players",deletePlayer)

	// Routes for Queues
	
	router.GET("/api/queues", getQueues)
	router.GET("/api/queues/:id", getQueueById)
	router.POST("/api/queues",createQueue)
	router.PUT("/api/queues", updateQueue)
	router.DELETE("/api/queues",deleteQueue)

	// Routes for Sessions

	router.GET("/api/sessions", getSessions)
	router.GET("/api/sessions/:id", getSessionById)
	// router.POST("/api/sessions",createSession)
	// router.PUT("/api/sessions", updateSession)
	// router.DELETE("/api/sessions",deleteSession)



	// http.ListenAndServe(":8080", router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "http://localhost:8080")
}

func getPlayers(c *gin.Context ) {
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	fmt.Println("Connected!")
	data, dberr := getDBPlayers(db)
	CheckError(dberr)
	// close database
	defer db.Close()
	if len(data) <= 0 {
		fmt.Println("No data")
		c.JSON(204, gin.H{
			"message": "Al parecer falta el id",
		})
	}
	c.JSON(200, gin.H{
		"message": "Peticion realizada correctamente",
		"data":    data,
	})

}
func getPlayerById(c *gin.Context) {
	id := c.Param("id")

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := getDBPlayerById(db, id)

	// close database
	defer db.Close()
	fmt.Println("Connected!")

	if len(data) <= 0 {
		fmt.Println("No Data", dberr)
		c.JSON(204, gin.H{
			"message": "No se ha encontrado el Player",
			"id":      id,
		})
	}

	c.JSON(200, gin.H{
		"message": "Peticion realizada correctamente",
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
		c.JSON(400, gin.H{"error": "El parámetro 'name' es requerido"})
		return
	}
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := createDBPlayer(db,  name)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
			"name":    name,
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Peticion realizada correctamente",
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
		c.JSON(400, gin.H{"error": "Se requieren al menos 'id' y 'name'"})
		return
	}
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := updateDBPlayer(db, player)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else if  len(data) == 0 {

		c.JSON(204, gin.H{
			"message": "No data!",
		})
	}else {

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
		c.JSON(400, gin.H{"error": "Se requiere al menos un 'id'"})
		return
	}
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := deleteDBPlayer(db, player)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else {

		c.JSON(200, gin.H{
			"message": "Se ha eliminado el usuario con id " + player.ID,
			"data":    data,
		})
	}

}
func getQueues(c *gin.Context) {
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := getDBQueues(db)
	CheckError(dberr)
	// close database
	defer db.Close()

	if len(data) <= 0 {
		fmt.Println("No Data")
		c.JSON(204, gin.H{
			"message": "Al parecer falta el id",
		})
	}
	fmt.Println("Connected!")

	c.JSON(200, gin.H{
		"message": "Peticion realizada correctamente",
		"data":    data,
	})

}
func getQueueById(c *gin.Context) {
	id := c.Param("id")

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := getDBQueueById(db,  id)

	// close database
	defer db.Close()
	fmt.Println("Connected!")

	if len(data) <= 0 {
		fmt.Println("No Data", dberr)
		c.JSON(204, gin.H{
			"message": "No se ha encontrado la Cola",
			"id":      id,
		})
	}

	c.JSON(200, gin.H{
		"message": "Peticion realizada correctamente",
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
	name := queue.Name

	if name == "" {
		c.JSON(400, gin.H{"error": "El parámetro 'name' es requerido"})
		return
	}
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := createDBQueue(db,  name)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
			"name":    name,
		})
	} else if  len(data) == 0 {

		c.JSON(400, gin.H{
			"message": "Limit queues exceeded",
		})
	}else {

		c.JSON(201, gin.H{
			"message": "Peticion realizada correctamente",
			"name":    name,
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
		c.JSON(400, gin.H{"error": "Se requieren al menos 'id' y 'name'"})
		return
	}
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := updateDBQueue(db, queue)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else if  len(data) == 0 {

		c.JSON(204, gin.H{
			"message": "No data!",
		})
	}else{

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
		c.JSON(400, gin.H{"error": "Se requiere al menos un 'id'"})
		return
	}
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := deleteDBQueue(db,  queue)
	// close database

	fmt.Println("Connected!")

	defer db.Close()
	if dberr != nil {
		c.JSON(400, gin.H{
			"message": dberr,
		})
	} else {

		c.JSON(200, gin.H{
			"message": "Se ha eliminado el usuario con id " + queue.ID,
			"data":    data,
		})
	}

}

func getSessions(c *gin.Context) {
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := getDBSessions(db)
	CheckError(dberr)
	// close database
	defer db.Close()

	if len(data) <= 0 {
		fmt.Println("No Data")
		c.JSON(204, gin.H{
			"message": "Al parecer falta el id",
		})
	}
	fmt.Println("Connected!")

	c.JSON(200, gin.H{
		"message": "Peticion realizada correctamente",
		"data":    data,
	})

}
func getSessionById(c *gin.Context) {
	id := c.Param("id")

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	data, dberr := getDBSessionById(db,  id)

	// close database
	defer db.Close()
	fmt.Println("Connected!")

	if len(data) <= 0 {
		fmt.Println("No Data", dberr)
		c.JSON(204, gin.H{
			"message": "No se ha encontrado la Cola",
			"id":      id,
		})
	}

	c.JSON(200, gin.H{
		"message": "Peticion realizada correctamente",
		"id":      id,
		"data":    data,
	})

}
// func createSession(c *gin.Context) {
// 	var session Session
// 	if err := c.ShouldBindJSON(&session); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	name := session.Name

// 	if name == "" {
// 		c.JSON(400, gin.H{"error": "El parámetro 'name' es requerido"})
// 		return
// 	}
// 	// open database
// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)
// 	data, dberr := createDBSession(db,  name)
// 	// close database

// 	fmt.Println("Connected!")

// 	defer db.Close()
// 	if dberr != nil {
// 		c.JSON(400, gin.H{
// 			"message": dberr,
// 			"name":    name,
// 		})
// 	} else if  len(data) == 0 {

// 		c.JSON(400, gin.H{
// 			"message": "Limit queues exceeded",
// 		})
// 	}else {

// 		c.JSON(201, gin.H{
// 			"message": "Peticion realizada correctamente",
// 			"name":    name,
// 			"data":    data,
// 		})
// 	}

// }
// func updateSession(c *gin.Context) {
// 	var queue Session
// 	if err := c.ShouldBindJSON(&queue); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if queue.ID == "" || queue.Name == "" {
// 		c.JSON(400, gin.H{"error": "Se requieren al menos 'id' y 'name'"})
// 		return
// 	}
// 	// open database
// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)
// 	data, dberr := updateDBQueue(db, queue)
// 	// close database

// 	fmt.Println("Connected!")

// 	defer db.Close()
// 	if dberr != nil {
// 		c.JSON(400, gin.H{
// 			"message": dberr,
// 		})
// 	} else if  len(data) == 0 {

// 		c.JSON(204, gin.H{
// 			"message": "No data!",
// 		})
// 	}else{

// 		c.JSON(201, gin.H{
// 			"message": "Updated!",
// 			"data":    data,
// 		})
// 	}

// }
// func deleteQueue(c *gin.Context) {
// 	var queue Queue
// 	if err := c.ShouldBindJSON(&queue); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if queue.ID == "" {
// 		c.JSON(400, gin.H{"error": "Se requiere al menos un 'id'"})
// 		return
// 	}
// 	// open database
// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)
// 	data, dberr := deleteDBQueue(db,  queue)
// 	// close database

// 	fmt.Println("Connected!")

// 	defer db.Close()
// 	if dberr != nil {
// 		c.JSON(400, gin.H{
// 			"message": dberr,
// 		})
// 	} else {

// 		c.JSON(200, gin.H{
// 			"message": "Se ha eliminado el usuario con id " + queue.ID,
// 			"data":    data,
// 		})
// 	}

// }

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetLimitQueue(conn string) (Parameter, error) {
	db, err := sql.Open("postgres", conn)
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
func getDBPlayers(db *sql.DB) ([]Player, error) {
	rows, err := db.Query(`SELECT "name", "uuid" FROM public.players`)
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

func getDBPlayerById(db *sql.DB,  id string) ([]Player, error) {
	query := fmt.Sprintf(`SELECT "name", "uuid" FROM public.players where uuid = '%s'`, id)
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

func createDBPlayer(db *sql.DB, name string) ([]Player, error) {
	query := fmt.Sprintf(`INSERT INTO public.players(name) VALUES ('%s') RETURNING uuid,name; `, name)
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

func updateDBPlayer(db *sql.DB,  data Player) ([]Player, error) {
	query := fmt.Sprintf(`UPDATE  public.players SET  name= '%s' WHERE uuid = '%s' RETURNING uuid,name; `, data.Name, data.ID)
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

func deleteDBPlayer(db *sql.DB,  data Player) ([]Player, error) {
	query := fmt.Sprintf(`DELETE FROM public.players WHERE uuid = '%s'; `, data.ID)
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


func getDBQueues(db *sql.DB) ([]Queue, error) {
	rows, err := db.Query(`SELECT "name", "uuid", "max_players" FROM public.queues`)
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

func getDBQueueById(db *sql.DB, id string) ([]Queue, error) {
	query := fmt.Sprintf(`SELECT "name", "uuid", "max_players" FROM public.queues where uuid = '%s'`, id)
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

func createDBQueue(db *sql.DB,  name string) ([]Queue, error) {
	// limit := 10;
	
	query := fmt.Sprintf(`
		WITH filecount AS (
			SELECT COUNT(*) AS fileamount FROM public.queues
		)
		INSERT INTO public.queues(name)
		SELECT '%s'
		FROM filecount
		WHERE fileamount < (SELECT value::bigint FROM public.parameters WHERE name = 'max_queues')
		
		RETURNING uuid, name, max_players;
	`, name)
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

func updateDBQueue(db *sql.DB,data Queue) ([]Queue, error) {
	query := fmt.Sprintf(`UPDATE  public.queues SET  name= '%s' WHERE uuid = '%s' RETURNING uuid, name, max_players; `, data.Name, data.ID)
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

func deleteDBQueue(db *sql.DB,  data Queue) ([]Queue, error) {
	query := fmt.Sprintf(`DELETE FROM public.queues WHERE uuid = '%s'; `, data.ID)
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

func getDBSessions(db *sql.DB) ([]Queue, error) {
	rows, err := db.Query(`SELECT * FROM public.sessions`)
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

func getDBSessionById(db *sql.DB, id string) ([]Queue, error) {
	query := fmt.Sprintf(`SELECT * FROM public.sessions where uuid = '%s'`, id)
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

// func createDBQueue(db *sql.DB,  name string) ([]Queue, error) {
// 	// limit := 10;
	
// 	query := fmt.Sprintf(`
// 		WITH filecount AS (
// 			SELECT COUNT(*) AS fileamount FROM public.sessions
// 		)
// 		INSERT INTO public.queues(name)
// 		SELECT '%s'
// 		FROM filecount
// 		WHERE fileamount < (SELECT value::bigint FROM public.parameters WHERE name = 'max_queues')
		
// 		RETURNING uuid, name, max_players;
// 	`, name)
// 	fmt.Println(query)
 
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var dataList []Queue
// 	for rows.Next() {
// 		var data Queue
// 		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
// 		if err != nil {
// 			return nil, err
// 		}
// 		dataList = append(dataList, data)
// 	}

// 	return dataList, nil
// }

// func updateDBQueue(db *sql.DB,data Queue) ([]Queue, error) {
// 	query := fmt.Sprintf(`UPDATE  public.queues SET  name= '%s' WHERE uuid = '%s' RETURNING uuid, name, max_players; `, data.Name, data.ID)
// 	fmt.Println(query)

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var dataList []Queue
// 	for rows.Next() {
// 		var data Queue
// 		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
// 		if err != nil {
// 			return nil, err
// 		}
// 		dataList = append(dataList, data)
// 	}

// 	return dataList, nil
// }

// func deleteDBQueue(db *sql.DB,  data Queue) ([]Queue, error) {
// 	query := fmt.Sprintf(`DELETE FROM public.queues WHERE uuid = '%s'; `, data.ID)
// 	fmt.Println(query)

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var dataList []Queue
// 	for rows.Next() {
// 		var data Queue
// 		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
// 		if err != nil {
// 			return nil, err
// 		}
// 		dataList = append(dataList, data)
// 	}

// 	return dataList, nil
// }

