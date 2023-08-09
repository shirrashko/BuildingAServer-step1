package api

import "github.com/gin-gonic/gin"

func AnswerRequests(router *gin.Engine) {
	router.GET("/health", HealthCheck) // Associate the GET HTTP method and /health path with a handler function
}
