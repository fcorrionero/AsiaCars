package memory

import (
	"github.com/fcorrionero/europcar/domain"
	"sync"
)

type memoryVehicle struct {
	lock    sync.Mutex
	Vehicle *domain.Vehicle
}

type VehicleRepository struct {
	Vehicles []memoryVehicle
}

func New() *VehicleRepository {
	return new(VehicleRepository)
}

func (vR *VehicleRepository) FindByChassisNumber(chNbr string) (*domain.Vehicle, error) {
	v := new(domain.Vehicle)
	for _, mV := range vR.Vehicles {
		if mV.Vehicle.GetChassisNumber() == chNbr {
			return mV.Vehicle, nil
		}
	}

	return v, domain.NewDomainError("vehicle not found")
}

func (vR *VehicleRepository) FindByDeviceSerialNumber(srlNbr string) (*domain.Vehicle, error) {
	v := new(domain.Vehicle)
	for _, mV := range vR.Vehicles {
		if mV.Vehicle.DeviceSerialNumber == srlNbr {
			return mV.Vehicle, nil
		}
	}

	return v, domain.NewDomainError("vehicle not found")
}

func (vR *VehicleRepository) Save(vehicle *domain.Vehicle) error {
	for i, mV := range vR.Vehicles {
		if (len(mV.Vehicle.DeviceSerialNumber) > 0 && mV.Vehicle.DeviceSerialNumber == vehicle.DeviceSerialNumber) || mV.Vehicle.GetChassisNumber() == vehicle.GetChassisNumber() {
			mV.lock.Lock()
			mV.Vehicle = vehicle
			mV.lock.Unlock()

			vR.Vehicles[i] = mV
			return nil
		}
	}
	mV := new(memoryVehicle)
	mV.Vehicle = vehicle
	vR.Vehicles = append(vR.Vehicles, *mV)
	return nil
}
