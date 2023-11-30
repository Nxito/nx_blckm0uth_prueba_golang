package helpers

import (
	"golang-gin-crud-api/model" 
	
	"golang-gin-crud-api/conf"
	
	"database/sql"
	"fmt"

	

	_ "github.com/lib/pq"
)

/* Funciones  de Peticion a base de datos */
func GetSQLConfigString() (string){
	result, err := conf.GetSqlParams()
	if err != nil {
		// Manejar el error, por ejemplo, imprimirlo y salir
		fmt.Println("Error obteniendo parámetros SQL:", err)
	}
	return result
}
 
func GetDBPlayers() ([]model.Player, error) {
	
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "name", "id" FROM public.players`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Player
	for rows.Next() {
		var data model.Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func GetDBPlayerById(id string) ([]model.Player, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`SELECT "name", "id" FROM public.players where id = '%s'`, id)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Player
	for rows.Next() {
		var data model.Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func CreateDBPlayer(name string) ([]model.Player, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO public.players(name) VALUES ('%s') RETURNING id,name; `, name)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Player
	for rows.Next() {
		var data model.Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func UpdateDBPlayer(data model.Player) ([]model.Player, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE  public.players SET  name= '%s' WHERE id = '%s' RETURNING id,name; `, data.Name, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Player
	for rows.Next() {
		var data model.Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func DeleteDBPlayer(data model.Player) ([]model.Player, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM public.players WHERE id = '%s' RETURNING name; `, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Player
	for rows.Next() {
		var data model.Player
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func GetDBQueues() ([]model.Queue, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "name", "id", "max_players" FROM public.queues`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Queue
	for rows.Next() {
		var data model.Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func GetDBQueueById(id string) ([]model.Queue, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`SELECT "name", "id", "max_players" FROM public.queues where id = '%s'`, id)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Queue
	for rows.Next() {
		var data model.Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func CreateDBQueue(queue model.Queue) ([]model.Queue, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()
	
	result, err := conf.GetLimitQueue()
	if err != nil {
		fmt.Println("Error obteniendo parámetros SQL:", err)
		CheckError(err)
	}
 

	query := fmt.Sprintf(`
		WITH filecount AS (
			SELECT COUNT(*) AS fileamount FROM public.queues
		)
		INSERT INTO public.queues(name,max_players)
		SELECT '%s','%d'
		FROM filecount
		WHERE fileamount < %d
		
		RETURNING id, name, max_players;
	`, queue.Name, queue.MaxPlayers, result)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Queue
	for rows.Next() {
		var data model.Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func UpdateDBQueue(data model.Queue) ([]model.Queue, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE  public.queues SET  name= '%s' WHERE id = '%s' RETURNING id, name, max_players; `, data.Name, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Queue
	for rows.Next() {
		var data model.Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func DeleteDBQueue(data model.Queue) ([]model.Queue, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM public.queues WHERE id = '%s' RETURNING name;`, data.ID)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Queue
	for rows.Next() {
		var data model.Queue
		err := rows.Scan(&data.ID, &data.Name, &data.MaxPlayers)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func GetDBSessions() ([]model.Session, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "id_queue", "players", "status" FROM public.sessions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Session
	for rows.Next() {
		var data model.Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func GetDBSessionByStatus(status string) ([]model.Session, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
	CheckError(err)
	defer db.Close()

	query := fmt.Sprintf(`SELECT "id", "id_queue", "players", "status" FROM public.sessions where status = '%s'`, status)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Session
	for rows.Next() {
		var data model.Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func CreateDBSession(session model.Session) ([]model.Session, error) {
	db, err := sql.Open("postgres", GetSQLConfigString())
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

	var dataList []model.Session
	for rows.Next() {
		var data model.Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func UpdateDBSession(data model.Session) ([]model.Session, error) {

	db, err := sql.Open("postgres", GetSQLConfigString())
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

	var dataList []model.Session
	for rows.Next() {
		var data model.Session
		err := rows.Scan(&data.ID, &data.Queue, &data.Player, &data.Status)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

