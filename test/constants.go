package test

import "github.com/fcorrionero/europcar/domain"

const (
	ValidChassisNbr   = "A23DS6RW9WlK11D67"
	ValidLicensePlate = "4587JKM"
	ValidCategory     = "MBMR"
)

func GetVehicle() (*domain.Vehicle, error) {
	v, err := domain.NewVehicle(ValidChassisNbr, ValidLicensePlate, ValidCategory)

	return v, err
}

func GetVehicleWithParams(cN string, lP string, ct string) (*domain.Vehicle, error) {
	v, err := domain.NewVehicle(cN, lP, ct)

	return v, err
}
