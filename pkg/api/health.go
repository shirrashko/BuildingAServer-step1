package api

import (
	"buildingAServer/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	if pkg.HealthChecking() { // calling the bl
		c.JSON(http.StatusOK, gin.H{"status": "ok"}) // in JSON because requests and responses are being made using it
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"}) // status code 500 - internal server error
	}
}

// http.StatusOK is a constant from the net/http that stands for status code: 200, which means that
// the request was successful, but the meaning of success depends on the request
