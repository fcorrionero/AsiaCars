package domain

type VehicleRepository interface {
	FindByChassisNumber(chNbr string) (*Vehicle, error)
	FindByDeviceSerialNumber(srlNbr string) (*Vehicle, error)
	Save(vehicle *Vehicle) error
}
