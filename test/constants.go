package test

import (
	"github.com/fcorrionero/europcar/domain"
	"time"
)

const (
	ValidChassisNbr      = "A23DS6RW9WlK11D67"
	ValidLicensePlate    = "4587JKM"
	ValidCategory        = "MBMR"
	ValidDeviceSerialNbr = "G-654789"
)

func GetVehicle() (*domain.Vehicle, error) {
	v, err := domain.NewVehicle(ValidChassisNbr, ValidLicensePlate, ValidCategory, time.Now())

	return v, err
}

func GetVehicleWithParams(cN string, lP string, ct string, t time.Time) (*domain.Vehicle, error) {
	v, err := domain.NewVehicle(cN, lP, ct, t)

	return v, err
}
