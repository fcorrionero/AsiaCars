package telemetry

import (
	"fmt"
	"github.com/fcorrionero/europcar/domain"
)

type UpdateFuel struct {
	repo domain.VehicleRepository
}

type FuelSchema struct {
	DeviceSerialNumber string `json:"serial_number"`
	Fuel               int    `json:"fuel"`
	Type               string `json:"type"`
}

func NewUpdateFuel(vR domain.VehicleRepository) UpdateFuel {
	return UpdateFuel{repo: vR}
}

func (c UpdateFuel) Handle(data FuelSchema) error {
	v, err := c.repo.FindByDeviceSerialNumber(data.DeviceSerialNumber)
	if nil != err {
		return err
	}
	var level int
	switch data.Type {
	case "increment":
		level = v.GetFuelLevel() + data.Fuel
	case "decrement":
		level = v.GetFuelLevel() - data.Fuel
	default:
		return NewAppError(fmt.Sprintf("Operation %s not allowed", data.Type))
	}
	err = v.UpdateFuelLevel(level)
	if nil != err {
		return err
	}

	return c.repo.Save(v)
}
