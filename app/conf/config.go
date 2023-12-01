package conf


import (
	"encoding/json"
	"fmt"
	"golang-gin-crud-api/model"
	"os"
)


func GetSqlParams() (string, error) {
    file, err := os.ReadFile("config.json")
    if err != nil {
        return "", err
    }

    var config model.Config
    err = json.Unmarshal(file, &config)
    if err != nil {
        return "", err
    }

    dbConfig := config.Database

    // Crear el string formateado
    formattedString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
 
    return formattedString, nil
}

func GetLimitQueue() (int, error) {
    file, err := os.ReadFile("config.json")
    if err != nil {
        return 0, err
    }

    var config model.Config
    err = json.Unmarshal(file, &config)
    if err != nil {
        return 0, err
    }

    return config.MaxQueue, nil
}
