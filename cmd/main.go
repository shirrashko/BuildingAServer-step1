package main

import (
	"buildingAServer/pkg/api"
	"fmt"
	"github.com/gin-gonic/gin" // need to run go get github.com/gin-gonic/gin in addition for it to be recognized
)

func main() {
	//create a new Gin server and start it\
	fmt.Println("start of the program")
	router := gin.Default()
	api.AnswerRequests(router)

	router.Run(":8080")
}

// A router is a traffic director, taking incoming requests and matching
// them to predefined routes. Each route is associated with a specific handler function that is
// executed when the corresponding URL and method match the request

//todo: find out how we decide on this number
//todo: need to check if error accord like the compiler suggests?
