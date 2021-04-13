package domain

import (
	"github.com/fcorrionero/europcar/domain"
	"github.com/fcorrionero/europcar/test"
	"reflect"
	"testing"
)

func TestVehicleShouldBeCreated(t *testing.T) {
	_, err := domain.NewVehicle(test.ValidChassisNbr, test.ValidLicensePlate, test.ValidCategory)
	if nil != err {
		t.Fatalf(`Error creating vehicle: %v`, err)
	}
}

func TestOnlyValidChassisNbrAreAllowed(t *testing.T) {
	var chassisNbrs = []struct {
		value    string
		expected error
	}{
		{"12312123", domain.NewDomainError("")},
		{test.ValidChassisNbr, nil},
		{"12345678901234567", nil},
		{"*****************", domain.NewDomainError("")},
		{"123*8$%&712123", domain.NewDomainError("")},
		{"", domain.NewDomainError("")},
	}

	for _, c := range chassisNbrs {
		_, err := domain.NewVehicle(c.value, test.ValidLicensePlate, test.ValidCategory)
		if reflect.TypeOf(err) != reflect.TypeOf(c.expected) {
			t.Errorf("Expected: %v, got: %v", reflect.TypeOf(c.expected), reflect.TypeOf(err))
		}
	}
}

func TestOnlyValidLicensePlatesAreAllowed(t *testing.T) {
	var licensePlates = []struct {
		value    string
		expected error
	}{
		{"ZE45KI", nil},
		{"ZE-5KI", domain.NewDomainError("")},
		{"", domain.NewDomainError("")},
		{"333444555", nil},
	}

	for _, l := range licensePlates {
		_, err := domain.NewVehicle(test.ValidChassisNbr, l.value, test.ValidCategory)
		if reflect.TypeOf(err) != reflect.TypeOf(l.expected) {
			t.Errorf("Expected (%s): %v, got: %v", l.value, reflect.TypeOf(l.expected), reflect.TypeOf(err))
		}
	}
}

func TestOnlyValidCategoriesAreAllowed(t *testing.T) {
	var categories = []struct {
		value    string
		expected error
	}{
		{test.ValidCategory, nil},
		{"3456", domain.NewDomainError("")},
		{"", domain.NewDomainError("")},
		{"NCNN", nil},
	}

	for _, c := range categories {
		_, err := domain.NewVehicle(test.ValidChassisNbr, test.ValidLicensePlate, c.value)
		if reflect.TypeOf(err) != reflect.TypeOf(c.expected) {
			t.Errorf("Expected (%s): %v, got: %v", c.value, reflect.TypeOf(c.expected), reflect.TypeOf(err))
		}
	}

}

// Just for fun
func TestAllValidAcrissCodes(t *testing.T) {
	s1 := "MNEHCDIJSRFGPULWOX"
	s2 := "BCDWVLSTFJXPQZEMRHYNGK"
	s3 := "MNCABD"
	s4 := "RNDQHIECLSABMFVZUX"

	var comb1 []string
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			comb1 = append(comb1, string(c1)+string(c2))
		}
	}
	var comb2 []string
	for _, cmb1 := range comb1 {
		for _, c3 := range s3 {
			comb2 = append(comb2, cmb1+string(c3))
		}
	}

	for _, cmb2 := range comb2 {
		for _, c4 := range s4 {
			_, err := domain.NewVehicle(test.ValidChassisNbr, test.ValidLicensePlate, cmb2+string(c4))
			if nil != err {
				t.Errorf("error with valid acriss code")
			}
		}
	}
}

func TestBatteryLevelShouldBePositiveNumber(t *testing.T) {
	v, err := test.GetVehicle()
	if nil != err {
		t.Errorf("vehicle should be created")
	}

	err = v.UpdateBatteryLevel(10)
	if nil != err {
		t.Errorf("Battery level can be a positive number")
	}

	err = v.UpdateBatteryLevel(-1)
	if err == nil {
		t.Errorf("Battery level can't be negative")
	}
	if reflect.TypeOf(err) != reflect.TypeOf(&domain.Error{}) {
		t.Errorf("Domain error should be returned, Expected: %v; Got: %v", reflect.TypeOf(domain.Error{}), reflect.TypeOf(err))
	}
}

func TestFuelLevelShouldBePositiveNumber(t *testing.T) {
	v, err := test.GetVehicle()
	if nil != err {
		t.Errorf("vehicle should be created")
	}

	err = v.UpdateFuelLevel(10)
	if nil != err {
		t.Errorf("Fuel level can be a positive number")
	}

	err = v.UpdateFuelLevel(-1)
	if err == nil {
		t.Errorf("Fuel level can't be negative")
	}
	if reflect.TypeOf(err) != reflect.TypeOf(&domain.Error{}) {
		t.Errorf("Domain error should be returned, Expected: %v; Got: %v", reflect.TypeOf(domain.Error{}), reflect.TypeOf(err))
	}
}

func TestCurrentMileageShouldBePositiveNumber(t *testing.T) {
	v, err := test.GetVehicle()
	if nil != err {
		t.Errorf("vehicle should be created")
	}

	err = v.UpdateCurrentMileage(10)
	if nil != err {
		t.Errorf("Current mileage can be a positive number")
	}

	err = v.UpdateCurrentMileage(-1)
	if err == nil {
		t.Errorf("Current mileage can't be negative")
	}
	if reflect.TypeOf(err) != reflect.TypeOf(&domain.Error{}) {
		t.Errorf("Domain error should be returned, Expected: %v; Got: %v", reflect.TypeOf(domain.Error{}), reflect.TypeOf(err))
	}
}

func TestGetChassisNumber(t *testing.T) {
	v, _ := test.GetVehicle()

	if v.GetChassisNumber() != test.ValidChassisNbr {
		t.Errorf("Incorrect chassis number")
	}
}
