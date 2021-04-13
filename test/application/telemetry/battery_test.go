package telemetry

import (
	"github.com/fcorrionero/europcar/application/telemetry"
	"github.com/fcorrionero/europcar/domain"
	"github.com/fcorrionero/europcar/test"
	"github.com/fcorrionero/europcar/test/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestBatteryShouldBeUpdated(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)
	mockRepo.EXPECT().Save(v).Return(nil)

	c := telemetry.NewUpdateBattery(mockRepo)
	schema := telemetry.BatterySchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		BatteryLevel:       70,
	}

	err := c.Handle(schema)
	if nil != err {
		t.Fatalf("No errors should appear when updatting battery")
	}
}

func TestANegativeBatteryValueShouldReturnAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)
	c := telemetry.NewUpdateBattery(mockRepo)
	schema := telemetry.BatterySchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		BatteryLevel:       -10,
	}

	err := c.Handle(schema)
	if nil == err {
		t.Fatalf("An error should be thrown")
	}
}

func TestNotFoundVehicleShouldReturnAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v := new(domain.Vehicle)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, domain.NewDomainError("")).Times(1)
	c := telemetry.NewUpdateBattery(mockRepo)
	schema := telemetry.BatterySchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		BatteryLevel:       10,
	}

	err := c.Handle(schema)
	if nil == err {
		t.Fatalf("An error should be thrown")
	}
}
