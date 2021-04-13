package telemetry

import "github.com/fcorrionero/europcar/domain"

type UpdateBattery struct {
	repo domain.VehicleRepository
}

type BatterySchema struct {
	DeviceSerialNumber string `json:"serial_number"`
	BatteryLevel       int    `json:"battery"`
}

func NewUpdateBattery(vR domain.VehicleRepository) UpdateBattery {
	return UpdateBattery{
		repo: vR,
	}
}

func (c UpdateBattery) Handle(data BatterySchema) error {
	v, err := c.repo.FindByDeviceSerialNumber(data.DeviceSerialNumber)
	if nil != err {
		return err
	}
	err = v.UpdateBatteryLevel(data.BatteryLevel)
	if nil != err {
		return err
	}

	return c.repo.Save(v)
}
