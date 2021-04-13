package http

import (
	"fmt"
	"github.com/fcorrionero/europcar/application/operations"
	"net/http"
)

type VehicleController struct {
	inFleetVehicle operations.InFleetVehicle
	installVehicle operations.InstallVehicle
}

func NewVehicleController(iFV operations.InFleetVehicle, iV operations.InstallVehicle) VehicleController {
	return VehicleController{
		inFleetVehicle: iFV,
		installVehicle: iV,
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
}

func (c VehicleController) InstallDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
}
