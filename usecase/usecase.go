package usecase

import (
	"context"

	"github.com/bayudava29/go-clean-arc/repository"
)

type Usecase interface {
	GetData(c context.Context)
}

type usecase struct {
	Repository repository.Repository
}

func InitUsecase(repository repository.Repository) Usecase {
	return &usecase{
		Repository: repository,
	}
}

func (usecase *usecase) GetData(c context.Context) {
	// Instruction for GetData Usecase
}
