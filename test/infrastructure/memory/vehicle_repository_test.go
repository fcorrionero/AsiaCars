package memory

import (
	"github.com/fcorrionero/europcar/infrastructure/memory"
	"github.com/fcorrionero/europcar/test"
	"testing"
)

func TestAVehicleCanBeAdded(t *testing.T) {
	vR := memory.New()
	v, _ := test.GetVehicle()

	_ = vR.Save(v)

	if len(vR.Vehicles) != 1 {
		t.Errorf("A vehicle should be added")
	}
}

func TestNoDuplicatedVehiclesExists(t *testing.T) {
	vR := memory.New()
	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = "1"

	v2, _ := test.GetVehicle()
	v2.DeviceSerialNumber = "1"

	_ = vR.Save(v)
	_ = vR.Save(v2)

	if len(vR.Vehicles) != 1 {
		t.Errorf("Only one vehicle should exitst with the same serial number or chassis number")
	}
}

func TestVehicleShouldBeFoundByDevice(t *testing.T) {
	vR := memory.New()
	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = "1"
	_ = vR.Save(v)

	vF, err := vR.FindByDeviceSerialNumber("1")
	if err != nil || len(vF.DeviceSerialNumber) == 0 {
		t.Errorf("Vehicle should be found")
	}
}

func TestVehicleShouldNotBeFoundByDevice(t *testing.T) {
	vR := memory.New()
	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = "1"
	_ = vR.Save(v)

	_, err := vR.FindByDeviceSerialNumber("2")
	if err == nil {
		t.Errorf("Vehicle should not be found")
	}
}

func TestVehicleShouldBeFoundByChassisNbr(t *testing.T) {
	vR := memory.New()
	v, _ := test.GetVehicle()
	_ = vR.Save(v)

	_, err := vR.FindByChassisNumber(test.ValidChassisNbr)
	if err != nil {
		t.Errorf("Vehicle should be found")
	}
}

func TestVehicleShouldNotBeFoundByChassisNbr(t *testing.T) {
	vR := memory.New()
	v, _ := test.GetVehicle()
	_ = vR.Save(v)

	_, err := vR.FindByChassisNumber("NotValid")
	if err == nil {
		t.Errorf("Vehicle should not be found")
	}
}
