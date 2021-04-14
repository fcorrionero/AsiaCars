package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fcorrionero/europcar/application/operations"
	"github.com/fcorrionero/europcar/application/telemetry"
	"github.com/fcorrionero/europcar/infrastructure/memory"
	"github.com/fcorrionero/europcar/infrastructure/ui/http"
	"github.com/fcorrionero/europcar/test"
	"github.com/joho/godotenv"
	"log"
	gHttp "net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	srv := startServer(httpServerExitDone)
	time.Sleep(500 * time.Millisecond)
	code := m.Run()
	err := srv.Shutdown(context.TODO())
	if nil != err {
		fmt.Println(err)
	}
	httpServerExitDone.Wait()
	os.Exit(code)
}

func startServer(wg *sync.WaitGroup) *gHttp.Server {
	srv := &gHttp.Server{Addr: ":8080"}
	err := godotenv.Load("../../../../.env")
	if nil != err {
		fmt.Println("Error loading env variables: " + err.Error())
		return srv
	}
	apiPort := os.Getenv("API_PORT")
	srv = &gHttp.Server{Addr: ":" + apiPort}
	vR := memory.New()
	inFleetVehicle := operations.NewInFleetVehicle(vR)
	installVehicle := operations.NewInstallVehicle(vR)
	updateBattery := telemetry.NewUpdateBattery(vR)
	updateFuel := telemetry.NewUpdateFuel(vR)
	updateMileage := telemetry.NewUpdateMileage(vR)
	getTelemetries := telemetry.NewGetTelemetries(vR)

	vC := http.NewVehicleController(inFleetVehicle, installVehicle, updateBattery, updateFuel, updateMileage, getTelemetries)
	gHttp.HandleFunc("/hello", vC.Hello)
	gHttp.HandleFunc("/infleet", vC.InFleet)
	gHttp.HandleFunc("/install", vC.InstallDevice)
	gHttp.HandleFunc("/battery", vC.Battery)
	gHttp.HandleFunc("/fuel", vC.Fuel)
	gHttp.HandleFunc("/mileage", vC.Mileage)
	gHttp.HandleFunc("/telemetries", vC.Telemetries)

	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != gHttp.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}

func TestHello(t *testing.T) {
	res, err := gHttp.Get("http://localhost:8888/hello")
	if nil != err {
		t.Fatalf(err.Error())
		return
	}
	if 200 != res.StatusCode {
		t.Fatalf("200 status code should be returned")
	}
}

func TestInFleet(t *testing.T) {
	values := map[string]string{
		"chassis_number": test.ValidChassisNbr,
		"license_plate":  test.ValidLicensePlate,
		"category":       test.ValidCategory,
	}
	jsonValue, _ := json.Marshal(values)
	res, err := gHttp.Post("http://localhost:8888/infleet", "application/json", bytes.NewBuffer(jsonValue))
	if nil != err {
		t.Fatalf(err.Error())
	}
	if 200 != res.StatusCode {
		t.Fatalf("200 status code should be returned")
	}
}

func TestBadMethodInFleet(t *testing.T) {
	res, err := gHttp.Get("http://localhost:8888/infleet")
	if nil != err {
		t.Fatalf(err.Error())
	}
	if 405 != res.StatusCode {
		t.Fatalf("405 status code should be returned")
	}
}
