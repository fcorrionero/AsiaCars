package main

import (
	"fmt"
	"github.com/fcorrionero/europcar/domain"
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
	err = HandleRequests(apiPort, vR)
	if nil != err {
		fmt.Println("Error handling requests: " + err.Error())
		return
	}
}

func HandleRequests(apiPort string, vR domain.VehicleRepository) error {
	oC := InitializeOperationsController(vR)
	tC := InitializeTelemetriesController(vR)
	gHttp.HandleFunc("/hello", oC.Hello)
	gHttp.HandleFunc("/infleet", oC.InFleet)
	gHttp.HandleFunc("/install", oC.InstallDevice)
	gHttp.HandleFunc("/battery", tC.Battery)
	gHttp.HandleFunc("/fuel", tC.Fuel)
	gHttp.HandleFunc("/mileage", tC.Mileage)
	gHttp.HandleFunc("/telemetries", tC.Telemetries)
	err := gHttp.ListenAndServe(":"+apiPort, nil)
	return err
}
