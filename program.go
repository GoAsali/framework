package main

import (
	"asalpolaki/infrastructure"
	"asalpolaki/routes"
	"fmt"
)

func main() {
	host := "0.0.0.0"
	port := 8000

	infrastructure.LoadEnv()
	if !infrastructure.NewDatabase() {
		fmt.Println("Database have problem")
		return
	}

	err := routes.NewRoutes().Run()

	if err != nil {
		fmt.Printf("Error in handling port - %s\n", err.Error())
	} else {
		fmt.Printf("Run in %s:%d", host, port)
	}
}
