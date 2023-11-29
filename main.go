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
    ID     string  `json:"id"`
    Name  string  `json:"name"`
	MaxPlayers  string  `json:"maxPlayers"`
}

type Session struct {
    ID     string  `json:"id"`
    Queue  string  `json:"queue"`
	Players  string  `json:"players"`
	Status  string  `json:"status"`
}

// var players = []Player{
//     {ID: "1",  Name: "John Smith"},
//     {ID: "2",  Name: "Gary Oak"},
//     {ID: "3",  Name: "Tomato49"},
// }

func main() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	router := gin.Default()
	router.GET("/api/players", func(c *gin.Context) {
		// open database
		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)
		data, dberr := getDBPlayers(db, "players")
		CheckError(dberr)
		// close database
		defer db.Close()

		if(len(data)<=0){
			fmt.Println("NO DATA")
			c.JSON(204, gin.H{
				"message": "Al parecer falta el id", 
			})
		}
		fmt.Println("Connected!")

		c.JSON(200, gin.H{
			"message": "Peticion realizada correctamente",
			"data": data ,
		})

	})
	router.GET("/api/players/:id", func(c *gin.Context) {
		id := c.Param("id")
		
			// open database
			db, err := sql.Open("postgres", psqlconn)
			CheckError(err)
			data, dberr := getDBPlayerById(db, "players",id)

			// close database
			defer db.Close()
			fmt.Println("Connected!")
			

			if(len(data)<=0){
				fmt.Println("No Data",dberr)
				c.JSON(204, gin.H{
					"message": "No se ha encontrado el Player",
					"id":      id,
				})
			}
			
			c.JSON(200, gin.H{
				"message": "Peticion realizada correctamente",
				"id":      id,
				"data": data ,
			})

	})
	router.POST("/api/players", func(c *gin.Context) { 
			var player Player
			if err := c.ShouldBindJSON(&player); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			name :=player.Name
			
			if name == "" {
				c.JSON(400 , gin.H{"error": "El parámetro 'name' es requerido"})
				return
			}
			// open database
			db, err := sql.Open("postgres", psqlconn)
			CheckError(err)
			data, dberr := createDBPlayer(db, "players", name)
			// close database
			
			fmt.Println("Connected!")
			

			defer db.Close()
			if(dberr != nil){
				c.JSON(400, gin.H{
					"message": dberr,
					"name":      name,
				})
			}else{
			
				c.JSON(201, gin.H{
					"message": "Peticion realizada correctamente",
					"name":      name,
					"data": data ,
				})
			} 


	})
	router.PUT("/api/players", func(c *gin.Context) { 
		var player Player
		if err := c.ShouldBindJSON(&player); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if player.ID == "" || player.Name =="" {
			c.JSON(400 , gin.H{"error": "Se requieren al menos 'id' y 'name'"})
			return
		}
		// open database
		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)
		data, dberr := updateDBPlayer(db, "players", player)
		// close database
		
		fmt.Println("Connected!")
		

		defer db.Close()
		if(dberr != nil){
			c.JSON(400, gin.H{
				"message": dberr,

			})
		}else{
		
			c.JSON(201, gin.H{
				"message": "Updated!",
				"data": data ,
			})
		} 


})
	router.DELETE("/api/players", func(c *gin.Context) { 
		var player Player
		if err := c.ShouldBindJSON(&player); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if player.ID == "" {
			c.JSON(400 , gin.H{"error": "Se requiere al menos un 'id'"})
			return
		}
		// open database
		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)
		data, dberr := deleteDBPlayer(db, "players", player)
		// close database
		
		fmt.Println("Connected!")
		

		defer db.Close()
		if(dberr != nil){
			c.JSON(400, gin.H{
				"message": dberr,

			})
		}else{
		
			c.JSON(200, gin.H{
				"message": "Se ha eliminado el usuario con id " + player.ID,
				"data": data ,
			})
		} 


})

	// http.ListenAndServe(":8080", router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "http://localhost:8080")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}


func getDBPlayers(db *sql.DB, tableName string) ([]Player, error) {
	rows, err := db.Query(fmt.Sprintf(`SELECT "name", "uuid" FROM "%s"`, tableName))
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

func getDBPlayerById(db *sql.DB, tableName string, id string) ([]Player, error) {
	query := fmt.Sprintf(`SELECT "name", "uuid" FROM "%s" where uuid = '%s'`, tableName,id)
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


func createDBPlayer(db *sql.DB, tableName string, name string) ([]Player, error) {
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


func updateDBPlayer(db *sql.DB, tableName string, data Player) ([]Player, error) {
	query := fmt.Sprintf(`UPDATE  public.players SET  name= '%s' WHERE uuid = '%s' RETURNING uuid,name;; `, data.Name,data.ID)
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

func deleteDBPlayer(db *sql.DB, tableName string, data Player) ([]Player, error) {
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
