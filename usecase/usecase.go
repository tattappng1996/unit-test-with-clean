package usecase

// Customer ..
type Customer struct {
	ID   int
	Name string
}

// Usecases ..
type Usecases interface {
	GetFirstCustomer(id int) (Customer, error)
}

// CustomerRepository ..
type CustomerRepository interface {
	First(id int) (Customer, error)
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
func (uc *Usecase) GetFirstCustomer(id int) (Customer, error) {
	customer, err := uc.repo.First(id)
	return customer, err
}
