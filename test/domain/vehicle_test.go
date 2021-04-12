package domain

import (
	"github.com/fcorrionero/europcar/domain"
	"reflect"
	"testing"
)

func TestVehicleShouldBeCreated(t *testing.T) {
	_, err := domain.NewVehicle("A23DS6RW9WlK11D67", "4587JKM", "MBMR")
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
		{"A23DS6RW9WlK11D67", nil},
		{"12345678901234567", nil},
		{"*****************", domain.NewDomainError("")},
		{"123*8$%&712123", domain.NewDomainError("")},
		{"", domain.NewDomainError("")},
	}

	for _, c := range chassisNbrs {
		_, err := domain.NewVehicle(c.value, "4587JKM", "MBMR")
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
		_, err := domain.NewVehicle("A23DS6RW9WlK11D67", l.value, "MBMR")
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
		{"MBMR", nil},
		{"3456", domain.NewDomainError("")},
		{"", domain.NewDomainError("")},
		{"NCNN", nil},
	}

	for _, c := range categories {
		_, err := domain.NewVehicle("A23DS6RW9WlK11D67", "4587JKM", c.value)
		if reflect.TypeOf(err) != reflect.TypeOf(c.expected) {
			t.Errorf("Expected (%s): %v, got: %v", c.value, reflect.TypeOf(c.expected), reflect.TypeOf(err))
		}
	}

}

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
			_, err := domain.NewVehicle("A23DS6RW9WlK11D67", "4587JKM", cmb2+string(c4))
			if nil != err {
				t.Errorf("error with valid acriss code")
			}
		}
	}
}
