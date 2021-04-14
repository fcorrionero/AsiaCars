package telemetry

import (
	"github.com/fcorrionero/europcar/application/telemetry"
	"github.com/fcorrionero/europcar/domain"
	"github.com/fcorrionero/europcar/test"
	"github.com/fcorrionero/europcar/test/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestFuelShouldBeIncremented(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)
	mockRepo.EXPECT().Save(v).Return(nil)

	c := telemetry.NewUpdateFuel(mockRepo)
	schema := telemetry.FuelSchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		Fuel:               10,
		Type:               "increment",
	}

	err := c.Handle(schema)
	if nil != err {
		t.Fatalf("Fuel should be updated without errors")
	}

	if v.GetFuelLevel() != schema.Fuel {
		t.Fatalf("Erroneous fuel level")
	}
}

func TestFuelShouldBeDecremented(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	_ = v.UpdateFuelLevel(50)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)
	mockRepo.EXPECT().Save(v).Return(nil)

	c := telemetry.NewUpdateFuel(mockRepo)
	schema := telemetry.FuelSchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		Fuel:               10,
		Type:               "decrement",
	}

	err := c.Handle(schema)
	if nil != err {
		t.Fatalf("Fuel should be updated without errors")
	}

	if v.GetFuelLevel() != 40 {
		t.Fatalf("Erroneous fuel level")
	}
}

func TestErrorShouldBeThrownWithErroneousSchema(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)

	c := telemetry.NewUpdateFuel(mockRepo)
	schema := telemetry.FuelSchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		Fuel:               10,
		Type:               "error type",
	}

	err := c.Handle(schema)
	if reflect.TypeOf(err) != reflect.TypeOf(&telemetry.Error{}) {
		t.Fatalf("An application error should be thrown")
	}
}

func TestErrorShouldBeThrownWhenThereIsNoVehicle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v := new(domain.Vehicle)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	errF := domain.NewDomainError("not found")
	mockRepo.EXPECT().FindByDeviceSerialNumber("test").Return(v, errF).Times(1)

	c := telemetry.NewUpdateFuel(mockRepo)
	schema := telemetry.FuelSchema{
		DeviceSerialNumber: "test",
		Fuel:               10,
		Type:               "error type",
	}

	err := c.Handle(schema)
	if err != errF {
		t.Fatalf("Domain error should be thrown")
	}
}

func TestNegativeFuelShouldReturnAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	_ = v.UpdateFuelLevel(0)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)

	c := telemetry.NewUpdateFuel(mockRepo)
	schema := telemetry.FuelSchema{
		DeviceSerialNumber: v.DeviceSerialNumber,
		Fuel:               10,
		Type:               "decrement",
	}

	err := c.Handle(schema)
	if err == nil && reflect.TypeOf(err) != reflect.TypeOf(&domain.Error{}) {
		t.Fatalf("Negative fuel values are not allowed")
	}
}
