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

func TestMileageShouldBeUpdated(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)
	mockRepo.EXPECT().Save(v).Return(nil)

	c := telemetry.NewUpdateMileage(mockRepo)
	schema := telemetry.MileageSchema{
		DeviceSerialNumber: test.ValidDeviceSerialNbr,
		Mileage:            15,
		Unit:               "km",
	}

	err := c.Handle(schema)
	if nil != err {
		t.Fatalf("Mileage should be updated")
	}
}

func TestNegativeValuesOfMileageAreNotAllowed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)

	c := telemetry.NewUpdateMileage(mockRepo)
	schema := telemetry.MileageSchema{
		DeviceSerialNumber: test.ValidDeviceSerialNbr,
		Mileage:            -15,
		Unit:               "km",
	}

	err := c.Handle(schema)
	if nil == err && reflect.TypeOf(err) != reflect.TypeOf(&domain.Error{}) {
		t.Fatalf("An error should be returned for negative mileage values")
	}
}

func TestErrorWhenVehicleIsNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v := new(domain.Vehicle)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	errF := domain.NewDomainError("not found")
	mockRepo.EXPECT().FindByDeviceSerialNumber("test").Return(v, errF).Times(1)

	c := telemetry.NewUpdateMileage(mockRepo)
	schema := telemetry.MileageSchema{
		DeviceSerialNumber: "test",
		Mileage:            10,
		Unit:               "km",
	}
	err := c.Handle(schema)
	if err != errF {
		t.Fatalf("Domain error should be thrown")
	}
}
