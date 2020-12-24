package usecase

// TestcaseStruct ..
type TestcaseStruct struct {
	Number     int    `json:"number"`
	Name       string `json:"name"`
	ExpectData Expect `json:"expect"`
}

// Expect ..
type Expect struct {
	ID int `json:"id"`
}

// Usecase ..
type Usecase interface {
	TestNaja(TestcaseStruct) Expect
}

// CaseUsecase ..
type CaseUsecase struct {
	usecase Usecase
}

// NewUsecase ..
func NewUsecase(usecase Usecase) *CaseUsecase {
	return &CaseUsecase{
		usecase: usecase,
	}
}

// TestNaja ..
func (uc *CaseUsecase) TestNaja(testcaseStruct TestcaseStruct) Expect {
	return uc.TestNaja(testcaseStruct)
}
