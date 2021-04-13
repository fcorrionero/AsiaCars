package http

import (
	"encoding/json"
	"fmt"
	"github.com/fcorrionero/europcar/application/operations"
	"io/ioutil"
	"log"
	"net/http"
)

type VehicleController struct {
	InFleetVehicle operations.InFleetVehicle
	InstallVehicle operations.InstallVehicle
}

func NewVehicleController(iFV operations.InFleetVehicle, iV operations.InstallVehicle) VehicleController {
	return VehicleController{
		InFleetVehicle: iFV,
		InstallVehicle: iV,
	}
}

func HandleRequests(apiPort string, controller VehicleController) error {
	http.HandleFunc("/hello", controller.Hello)
	http.HandleFunc("/infleet", controller.InFleet)
	http.HandleFunc("/install", controller.InstallDevice)
	err := http.ListenAndServe(":"+apiPort, nil)
	return err
}

func (c VehicleController) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func (c VehicleController) InFleet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	var schema operations.InFleetSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	err = c.InFleetVehicle.Handle(schema)
	if nil != err {
		log.Printf("Error registering vehicle: %v", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
}

func (c VehicleController) InstallDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	var schema operations.InstallSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}
	err = c.InstallVehicle.Handle(schema)
	if nil != err {
		log.Printf("Error installing vehicle: %v", err)
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
