package operations

import (
	"github.com/fcorrionero/europcar/application/operations"
	"github.com/fcorrionero/europcar/domain"
	"github.com/fcorrionero/europcar/test"
	"github.com/fcorrionero/europcar/test/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestVehicleShouldBeInstalled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	v, _ := test.GetVehicle()
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByChassisNumber(test.ValidChassisNbr).Times(1).Return(v, nil)
	v.DeviceSerialNumber = "test-serial-number"
	mockRepo.EXPECT().Save(v).Times(1).Return(nil)

	c := operations.NewInstallVehicle(mockRepo)
	schema := operations.InstallSchema{
		DeviceSerialNumber: "test-serial-number",
		ChassisNumber:      test.ValidChassisNbr,
	}
	err := c.Handle(schema)
	if nil != err {
		t.Fatalf("Error installing vehicle, %v", err)
	}
}

func TestVehicleShouldNotBeInstalled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	v := new(domain.Vehicle)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().FindByChassisNumber(test.ValidChassisNbr).Times(1).Return(v, domain.NewDomainError(""))

	c := operations.NewInstallVehicle(mockRepo)
	schema := operations.InstallSchema{
		DeviceSerialNumber: "test-serial-number",
		ChassisNumber:      test.ValidChassisNbr,
	}

	err := c.Handle(schema)
	if nil == err {
		t.Fatalf("An error should be returned")
	}
}
