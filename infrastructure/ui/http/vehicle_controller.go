package http

import (
	"encoding/json"
	"fmt"
	"github.com/fcorrionero/europcar/application/operations"
	"github.com/fcorrionero/europcar/application/telemetry"
	"io/ioutil"
	"log"
	"net/http"
)

type VehicleController struct {
	InFleetVehicle operations.InFleetVehicle
	InstallVehicle operations.InstallVehicle
	UpdateBattery  telemetry.UpdateBattery
	UpdateFuel     telemetry.UpdateFuel
	UpdateMileage  telemetry.UpdateMileage
	GetTelemetries telemetry.GetTelemetries
}

func NewVehicleController(
	iFV operations.InFleetVehicle,
	iV operations.InstallVehicle,
	uB telemetry.UpdateBattery,
	uF telemetry.UpdateFuel,
	uM telemetry.UpdateMileage,
	gT telemetry.GetTelemetries) VehicleController {
	return VehicleController{
		InFleetVehicle: iFV,
		InstallVehicle: iV,
		UpdateBattery:  uB,
		UpdateFuel:     uF,
		UpdateMileage:  uM,
		GetTelemetries: gT,
	}
}

func HandleRequests(apiPort string, controller VehicleController) error {
	http.HandleFunc("/hello", controller.Hello)
	http.HandleFunc("/infleet", controller.InFleet)
	http.HandleFunc("/install", controller.InstallDevice)
	http.HandleFunc("/battery", controller.Battery)
	http.HandleFunc("/fuel", controller.Fuel)
	http.HandleFunc("/mileage", controller.Mileage)
	http.HandleFunc("/telemetries", controller.Telemetries)
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
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	var schema operations.InFleetSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}

	err = c.InFleetVehicle.Handle(schema)
	if nil != err {
		log.Printf("Error registering vehicle: %v", err)
		w.WriteHeader(400)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
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
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}

	var schema operations.InstallSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	err = c.InstallVehicle.Handle(schema)
	if nil != err {
		log.Printf("Error installing vehicle: %v", err)
		w.WriteHeader(404)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	w.WriteHeader(200)
}

func (c VehicleController) Battery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}

	var schema telemetry.BatterySchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	err = c.UpdateBattery.Handle(schema)
	if nil != err {
		log.Printf("Error updating battery value: %v", err)
		w.WriteHeader(400)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	w.WriteHeader(200)
}

func (c VehicleController) Fuel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}

	var schema telemetry.FuelSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	err = c.UpdateFuel.Handle(schema)
	if nil != err {
		log.Printf("Error updating fuel value: %v", err)
		w.WriteHeader(400)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	w.WriteHeader(200)
}

func (c VehicleController) Mileage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}

	var schema telemetry.MileageSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	err = c.UpdateMileage.Handle(schema)
	if nil != err {
		log.Printf("Error updating mileage: %v", err)
		w.WriteHeader(400)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	w.WriteHeader(200)
}

func (c VehicleController) Telemetries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	sN, ok := r.URL.Query()["serial_number"]
	if !ok || len(sN[0]) < 1 {
		log.Println("Url Param 'serial_number' is missing")
		w.WriteHeader(400)
		httpErr := NewHttpError("Url Param 'serial_number' is missing")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	schema := telemetry.TelemetriesSchema{DeviceSerialNumber: sN[0]}
	res, err := c.GetTelemetries.Handle(schema)
	if nil != err {
		w.WriteHeader(404)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	data, _ := json.Marshal(res)
	w.WriteHeader(200)
	_, err = w.Write(data)
	if nil != err {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	return
}
