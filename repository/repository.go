package repository

type Repository interface{}

type repository struct{}

func InitRepository() Repository {
	return &repository{}
}
