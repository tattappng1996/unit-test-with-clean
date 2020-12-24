package usecase

import "unit-test-with-clean/dto"

// Usecases ..
type Usecases interface {
	GetFirstCustomer(id int) (dto.Customer, error)
}

// CustomerRepository ..
type CustomerRepository interface {
	First(id int) (dto.Customer, error)
}

// Usecase ..
type Usecase struct {
	repo CustomerRepository
}

// NewUsecase ..
func NewUsecase(repo CustomerRepository) Usecases {
	return &Usecase{
		repo: repo,
	}
}

// GetFirstCustomer ..
func (uc *Usecase) GetFirstCustomer(id int) (dto.Customer, error) {
	customer, err := uc.repo.First(id)
	return customer, err
}
