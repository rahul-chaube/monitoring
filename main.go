package main

import "Monitoring/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
