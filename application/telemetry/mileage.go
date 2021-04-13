package telemetry

import "github.com/fcorrionero/europcar/domain"

type UpdateMileage struct {
	repo domain.VehicleRepository
}

type MileageSchema struct {
	DeviceSerialNumber string `json:"serial_number"`
	Mileage            int    `json:"mileage"`
	Unit               string `json:"unit"`
}

func NewUpdateMileage(vR domain.VehicleRepository) UpdateMileage {
	return UpdateMileage{repo: vR}
}

func (c UpdateMileage) Handle(data MileageSchema) error {
	v, err := c.repo.FindByDeviceSerialNumber(data.DeviceSerialNumber)
	if nil != err {
		return err
	}

	err = v.UpdateCurrentMileage(data.Mileage)
	if nil != err {
		return err
	}

	return c.repo.Save(v)
}
