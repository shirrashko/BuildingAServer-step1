package pkg

// should it gets c *gin.Context param?

func HealthChecking() bool {
	// In the real world the health-check function will also check connections to other resources that the server
	//depends on.
	return true
}
