package usecase

import (
	"log"

	"github.com/raynaaliyya/enigma-laundry/entity"
	"github.com/raynaaliyya/enigma-laundry/repository"
	"github.com/raynaaliyya/enigma-laundry/validation"
)

type ServiceUsecase interface {
	Insert(service entity.Service)
	ViewAll() any
	ViewById(id string) any
	Update(service entity.Service)
	Delete(id string)
}

type serviceUsecase struct {
	serviceRepo repository.ServiceRepo
}

func (u *serviceUsecase) Insert(service entity.Service) {
	err := validation.ValidateService(service)
	if err != nil {
		log.Println(err)
	}

	u.serviceRepo.CreateService(service)
}

func (u *serviceUsecase) ViewAll() any {
	return u.serviceRepo.GetAllService()
}

func (u *serviceUsecase) ViewById(id string) any {
	return u.serviceRepo.GetServiceById(id)
}

func (u *serviceUsecase) Update(service entity.Service) {
	u.serviceRepo.UpdateService(service)
}

func (u *serviceUsecase) Delete(id string) {
	u.serviceRepo.DeleteService(id)
}

func NewServiceUsecase(serviceRepo repository.ServiceRepo) ServiceUsecase {
	return &serviceUsecase{
		serviceRepo: serviceRepo,
	}
}
