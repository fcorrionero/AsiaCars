package telemetry

import "github.com/fcorrionero/europcar/domain"

type GetTelemetries struct {
	repo domain.VehicleRepository
}

type TelemetriesSchema struct {
	DeviceSerialNumber string
}

type TelResponseSchema struct {
	DeviceSerialNumber string `json:"serial_number"`
	BatteryLevel       int    `json:"battery"`
	FuelLevel          int    `json:"fuel"`
	Mileage            int    `json:"mileage"`
}

func NewGetTelemetries(vR domain.VehicleRepository) GetTelemetries {
	return GetTelemetries{repo: vR}
}

func (c GetTelemetries) Handle(data TelemetriesSchema) (TelResponseSchema, error) {
	res := TelResponseSchema{}
	v, err := c.repo.FindByDeviceSerialNumber(data.DeviceSerialNumber)
	if nil != err {
		return res, err
	}

	res.DeviceSerialNumber = v.DeviceSerialNumber
	res.FuelLevel = v.GetFuelLevel()
	res.BatteryLevel = v.GetBatteryLevel()
	res.Mileage = v.GetCurrentMilleage()
	return res, nil
}
