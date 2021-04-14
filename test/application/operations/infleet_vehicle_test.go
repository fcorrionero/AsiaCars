package operations

import (
	"github.com/fcorrionero/europcar/application/operations"
	"github.com/fcorrionero/europcar/test"
	"github.com/fcorrionero/europcar/test/mocks"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestVehicleShouldBeAdded(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Now()
	v, _ := test.GetVehicleWithParams(test.ValidChassisNbr, test.ValidLicensePlate, test.ValidCategory, date)
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	mockRepo.EXPECT().Save(v).Return(nil).Times(1)

	schema := operations.InFleetSchema{
		ChassisNbr:   test.ValidChassisNbr,
		LicensePlate: test.ValidLicensePlate,
		Category:     test.ValidCategory,
		InFleetDate:  date,
	}
	c := operations.NewInFleetVehicle(mockRepo)
	err := c.Handle(schema)
	if nil != err {
		t.Errorf("Error adding vehicle %v", err)
	}
}

func TestVehicleShouldBeNotAdded(t *testing.T) {
	schema := operations.InFleetSchema{
		ChassisNbr:   "incorrect",
		LicensePlate: test.ValidLicensePlate,
		Category:     test.ValidCategory,
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockVehicleRepository(mockCtrl)
	c := operations.NewInFleetVehicle(mockRepo)
	err := c.Handle(schema)
	if nil == err {
		t.Errorf("An error should be returned")
	}
}
