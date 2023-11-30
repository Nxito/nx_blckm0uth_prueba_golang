package main

import (

	"golang-gin-crud-api/router"  

)


func main() {
	r := router.SetupRouter()
	r.Run(":8081")
}

