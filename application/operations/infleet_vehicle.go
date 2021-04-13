package operations

import "github.com/fcorrionero/europcar/domain"

type InFleetVehicle struct {
	repo domain.VehicleRepository
}

func NewInFleetVehicle(vR domain.VehicleRepository) InFleetVehicle {
	return InFleetVehicle{
		repo: vR,
	}
}

func (c InFleetVehicle) Handle(chNbr string, lcPlt string, cat string) error {
	v, err := domain.NewVehicle(chNbr, lcPlt, cat)
	if nil != err {
		return err
	}

	err = c.repo.Save(v)
	if nil != err {
		return err
	}

	return nil
}
