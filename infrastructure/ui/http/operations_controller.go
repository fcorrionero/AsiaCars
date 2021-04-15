package http

import (
	"encoding/json"
	"fmt"
	"github.com/fcorrionero/europcar/application/operations"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type OperationsController struct {
	InFleetVehicle operations.InFleetVehicle
	InstallVehicle operations.InstallVehicle
}

func NewOperationsController(iFV operations.InFleetVehicle, iV operations.InstallVehicle) OperationsController {
	return OperationsController{
		InFleetVehicle: iFV,
		InstallVehicle: iV,
	}
}

func (c OperationsController) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to Asia Cars")
}

func (c OperationsController) InFleet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500) // Return 500 Internal Server Error.
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	var schema operations.InFleetSchema
	if err = json.Unmarshal(body, &schema); nil != err {
		w.WriteHeader(400) // Return 400 Bad Request.
		httpErr := NewHttpError("Request body can not be processed")
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}
	schema.InFleetDate = time.Now()

	err = c.InFleetVehicle.Handle(schema)
	if nil != err {
		w.WriteHeader(400)
		httpErr := NewHttpError(err.Error())
		errTxt, _ := json.Marshal(httpErr)
		w.Write(errTxt)
		return
	}

	w.WriteHeader(200)
}

func (c OperationsController) InstallDevice(w http.ResponseWriter, r *http.Request) {
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
