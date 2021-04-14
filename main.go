package main

import (
	"fmt"
	"github.com/fcorrionero/europcar/infrastructure/ui/http"
	"github.com/joho/godotenv"
	gHttp "net/http"
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
	err = HandleRequests(apiPort, vC)
	if nil != err {
		fmt.Println("Error handling requests: " + err.Error())
		return
	}
}

func HandleRequests(apiPort string, controller http.VehicleController) error {
	gHttp.HandleFunc("/hello", controller.Hello)
	gHttp.HandleFunc("/infleet", controller.InFleet)
	gHttp.HandleFunc("/install", controller.InstallDevice)
	gHttp.HandleFunc("/battery", controller.Battery)
	gHttp.HandleFunc("/fuel", controller.Fuel)
	gHttp.HandleFunc("/mileage", controller.Mileage)
	gHttp.HandleFunc("/telemetries", controller.Telemetries)
	err := gHttp.ListenAndServe(":"+apiPort, nil)
	return err
}
