package http

import (
	"encoding/json"
	"github.com/fcorrionero/europcar/application/telemetry"
	"io/ioutil"
	"log"
	"net/http"
)

type TelemetriesController struct {
	UpdateBattery  telemetry.UpdateBattery
	UpdateFuel     telemetry.UpdateFuel
	UpdateMileage  telemetry.UpdateMileage
	GetTelemetries telemetry.GetTelemetries
}

func NewTelemetriesController(
	uB telemetry.UpdateBattery,
	uF telemetry.UpdateFuel,
	uM telemetry.UpdateMileage,
	gT telemetry.GetTelemetries) TelemetriesController {
	return TelemetriesController{
		UpdateBattery:  uB,
		UpdateFuel:     uF,
		UpdateMileage:  uM,
		GetTelemetries: gT,
	}
}

func (c TelemetriesController) Battery(w http.ResponseWriter, r *http.Request) {
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

func (c TelemetriesController) Fuel(w http.ResponseWriter, r *http.Request) {
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

func (c TelemetriesController) Mileage(w http.ResponseWriter, r *http.Request) {
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

func (c TelemetriesController) Telemetries(w http.ResponseWriter, r *http.Request) {
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
