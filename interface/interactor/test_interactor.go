package interactor

import (
	"fitcal-backend/domain"
	"fitcal-backend/interface/controllers/inputport"
	"log"
)

type TestInteractor struct {
	TestRepository inputport.TestRepository
}

func NewTestInteractor(r inputport.TestRepository) TestInteractor {
	return TestInteractor{r}
}

func (interactor *TestInteractor) TestInteractorTest() *domain.User {
	result, err := interactor.TestRepository.GetUserName()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
