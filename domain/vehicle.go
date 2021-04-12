package domain

import (
	"errors"
	"regexp"
	"time"
)

type Vehicle struct {
	chassisNbr         string
	licensePlate       string
	brand              string
	category           string
	inFleetDate        time.Time
	deviceSerialNumber string
	batteryLevel       int
	fuelLevel          int
	currentMileage     int
}

func NewVehicle(chassisNbr string, licensePlate string, category string) (*Vehicle, error) {
	v := new(Vehicle)
	if !checkChassisNbr(chassisNbr) {
		return v, errors.New("incorrect chassis number")
	}

	if !checkLicensePlate(licensePlate) {
		return v, errors.New("incorrect license plate")
	}

	if !checkCategory(category) {
		return v, errors.New("incorrect category, does not match any valid acriss code")
	}

	v.chassisNbr = chassisNbr
	v.licensePlate = licensePlate
	v.category = category

	return v, nil
}

func checkLicensePlate(licensePlate string) bool {
	lPVal, _ := regexp.MatchString("^[A-Za-z0-9].*$", licensePlate)
	return lPVal
}

func checkChassisNbr(chassisNbr string) bool {
	if len(chassisNbr) > 17 {
		return false
	}
	cNVal, _ := regexp.MatchString("^[A-Za-z0-9]{17}$", chassisNbr)
	return cNVal
}

func checkCategory(category string) bool {
	if len(category) != 4 {
		return false
	}
	regex := "^[MNEHCDIJSRFGPULWOX][BCDWVLSTFJXPQZEMRHYNGK][MNCABD][RNDQHIECLSABMFVZUX]$"
	cVal, _ := regexp.MatchString(regex, category)

	return cVal
}
