package usecase

type Usecase interface{}

type usecase struct{}

func InitUsecase() Usecase {
	return &usecase{}
}
