package operations

import (
	"github.com/fcorrionero/europcar/domain"
)

type InstallVehicle struct {
	repo domain.VehicleRepository
}

type InstallSchema struct {
	DeviceSerialNumber string `json:"serial_number"`
	ChassisNumber      string `json:"chassis_number"`
}

func NewInstallVehicle(vR domain.VehicleRepository) InstallVehicle {
	return InstallVehicle{
		repo: vR,
	}
}

func (c InstallVehicle) Handle(data InstallSchema) error {
	v, err := c.repo.FindByChassisNumber(data.ChassisNumber)
	if nil != err {
		return err
	}
	v.DeviceSerialNumber = data.DeviceSerialNumber

	err = c.repo.Save(v)
	return err
}
