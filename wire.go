//+build wireinject

package main

import (
	"github.com/fcorrionero/europcar/application/operations"
	"github.com/fcorrionero/europcar/application/telemetry"
	"github.com/fcorrionero/europcar/domain"
	"github.com/fcorrionero/europcar/infrastructure/memory"
	"github.com/fcorrionero/europcar/infrastructure/ui/http"
	"github.com/google/wire"
)

func InitializeVehicleRepository() domain.VehicleRepository {
	wire.Build(NewMemoryVehicleRepository)
	return &memory.VehicleRepository{}
}

func InitializeOperationsController(repository domain.VehicleRepository) http.OperationsController {
	wire.Build(
		http.NewOperationsController,
		operations.NewInFleetVehicle,
		operations.NewInstallVehicle,
	)
	return http.OperationsController{}
}

func InitializeTelemetriesController(repository domain.VehicleRepository) http.TelemetriesController {
	wire.Build(
		http.NewTelemetriesController,
		telemetry.NewUpdateBattery,
		telemetry.NewUpdateFuel,
		telemetry.NewUpdateMileage,
		telemetry.NewGetTelemetries,
	)
	return http.TelemetriesController{}
}

func NewMemoryVehicleRepository() domain.VehicleRepository {
	return memory.New()
}
