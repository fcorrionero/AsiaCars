package operations

import (
	"github.com/fcorrionero/europcar/domain"
)

type InFleetVehicle struct {
	repo domain.VehicleRepository
}

type InFleetSchema struct {
	ChassisNbr   string `json:"chassis_number"`
	LicensePlate string `json:"license_plate"`
	Category     string `json:"category"`
}

func NewInFleetVehicle(vR domain.VehicleRepository) InFleetVehicle {
	return InFleetVehicle{
		repo: vR,
	}
}

func (c InFleetVehicle) Handle(data InFleetSchema) error {
	v, err := domain.NewVehicle(data.ChassisNbr, data.LicensePlate, data.Category)
	if nil != err {
		return err
	}

	err = c.repo.Save(v)
	return err
}
