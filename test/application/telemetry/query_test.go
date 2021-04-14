package telemetry

import (
	"github.com/fcorrionero/europcar/application/telemetry"
	"github.com/fcorrionero/europcar/domain"
	"github.com/fcorrionero/europcar/test"
	"github.com/fcorrionero/europcar/test/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestTelemetriesShouldBeObtained(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	v.DeviceSerialNumber = test.ValidDeviceSerialNbr
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByDeviceSerialNumber(v.DeviceSerialNumber).Return(v, nil).Times(1)

	q := telemetry.NewGetTelemetries(mockRepo)
	schema := telemetry.TelemetriesSchema{DeviceSerialNumber: test.ValidDeviceSerialNbr}

	d, err := q.Handle(schema)
	if nil != err {
		t.Fatalf("Telemetries should be obtained")
	}
	if d.DeviceSerialNumber != v.DeviceSerialNumber {
		t.Fatalf("Data not valid")
	}
}

func TestReturnErrorWhenThereIsNoVehicle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v := new(domain.Vehicle)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	errF := domain.NewDomainError("not found")
	mockRepo.EXPECT().FindByDeviceSerialNumber("test").Return(v, errF).Times(1)
	q := telemetry.NewGetTelemetries(mockRepo)
	schema := telemetry.TelemetriesSchema{DeviceSerialNumber: "test"}

	_, err := q.Handle(schema)
	if nil == err {
		t.Fatalf("Error should appear when there is no vehicle")
	}
}
