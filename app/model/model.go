package model

/* Structs*/
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

 
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`

}

// Config representa la estructura de la configuración.
type Config struct {
	MaxQueue int         `json:"max_queue"` 
	Database DatabaseConfig `json:"database"`
}
