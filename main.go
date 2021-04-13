package main

import (
	"fmt"
	"github.com/fcorrionero/europcar/infrastructure/ui/http"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if nil != err {
		fmt.Println("Error loading env variables: " + err.Error())
		return
	}
	apiPort := os.Getenv("API_PORT")
	vR := InitializeVehicleRepository()
	vC := InitializeVehicleController(vR)
	err = http.HandleRequests(apiPort, vC)
	if nil != err {
		fmt.Println("Error handling requests: " + err.Error())
		return
	}
}
