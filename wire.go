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

func InitializeVehicleController(repository domain.VehicleRepository) http.VehicleController {
	wire.Build(http.NewVehicleController, operations.NewInFleetVehicle, operations.NewInstallVehicle, telemetry.NewUpdateBattery)
	return http.VehicleController{}
}

func NewMemoryVehicleRepository() domain.VehicleRepository {
	mRepo := memory.New()
	return mRepo
}
