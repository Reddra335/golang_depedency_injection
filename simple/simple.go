package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {

	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(simpleRepo *SimpleRepository) (*SimpleService, error) {

	if simpleRepo.Error {
		return nil, errors.New("Failed Create Service")

	} else {
		return &SimpleService{
			SimpleRepository: simpleRepo,
		}, nil
	}

}
