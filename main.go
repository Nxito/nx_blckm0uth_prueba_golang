package main

import (

	"golang-gin-crud-api/router" 
	_ "golang-gin-crud-api/docs"  

)

// @title Survivor Battle Royale API 
// @version	1.0
// @description A testing project fot BlackMouth. Made with Golang & Gin 
// @host 	localhost:8080
// @BasePath /api
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}

