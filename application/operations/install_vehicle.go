package operations

import "github.com/fcorrionero/europcar/domain"

type InstallVehicle struct {
	repo domain.VehicleRepository
}

func NewInstallVehicle(vR domain.VehicleRepository) InstallVehicle {
	return InstallVehicle{
		repo: vR,
	}
}
