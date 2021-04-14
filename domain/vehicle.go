package domain

import (
	"regexp"
	"time"
)

type Vehicle struct {
	chassisNbr         string
	licensePlate       string
	brand              string
	category           string
	inFleetDate        time.Time
	DeviceSerialNumber string
	batteryLevel       int
	fuelLevel          int
	currentMileage     int
}

func NewVehicle(chassisNbr string, licensePlate string, category string) (*Vehicle, error) {
	v := new(Vehicle)
	if !checkChassisNbr(chassisNbr) {
		return v, NewDomainError("incorrect chassis number")
	}

	if !checkLicensePlate(licensePlate) {
		return v, NewDomainError("incorrect license plate")
	}

	if !checkCategory(category) {
		return v, NewDomainError("incorrect category")
	}

	v.chassisNbr = chassisNbr
	v.licensePlate = licensePlate
	v.category = category
	//v.inFleetDate = time.Now()

	return v, nil
}

func checkLicensePlate(licensePlate string) bool {
	if len(licensePlate) == 0 {
		return false
	}
	lPVal, _ := regexp.MatchString("^[A-Za-z0-9]*$", licensePlate)
	return lPVal
}

func checkChassisNbr(chassisNbr string) bool {
	cNVal, _ := regexp.MatchString("^[A-Za-z0-9]{17}$", chassisNbr)
	return cNVal
}

func checkCategory(category string) bool {
	regex := "^[MNEHCDIJSRFGPULWOX][BCDWVLSTFJXPQZEMRHYNGK][MNCABD][RNDQHIECLSABMFVZUX]$"
	cVal, _ := regexp.MatchString(regex, category)

	return cVal
}

func (v *Vehicle) UpdateBatteryLevel(bL int) error {
	if bL < 0 {
		return NewDomainError("battery level can not be a negative number")
	}
	v.batteryLevel = bL
	return nil
}

func (v *Vehicle) UpdateFuelLevel(fL int) error {
	if fL < 0 {
		return NewDomainError("fuel level can not be a negative number")
	}
	v.fuelLevel = fL
	return nil
}

func (v *Vehicle) UpdateCurrentMileage(cM int) error {
	if cM < 0 {
		return NewDomainError("current mileage can not be a negative number")
	}
	v.currentMileage += cM
	return nil
}

func (v *Vehicle) GetChassisNumber() string {
	return v.chassisNbr
}

func (v *Vehicle) GetFuelLevel() int {
	return v.fuelLevel
}

func (v *Vehicle) GetBatteryLevel() int {
	return v.batteryLevel
}

func (v *Vehicle) GetCurrentMilleage() int {
	return v.currentMileage
}
