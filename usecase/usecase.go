package usecase

import (
	"github.com/alrifqi/tokenomy-test/repository"
	nameUc "github.com/alrifqi/tokenomy-test/usecase/name"
)

type Usecase struct {
	NameUsecase nameUc.NameUsecaseIface
	repo        repository.Repository
}

func Init(
	repo *repository.Repository,
) *Usecase {
	nameUsecase := nameUc.InitNameUsecase(repo.NameRepo)
	return &Usecase{
		NameUsecase: nameUsecase,
		repo:        *repo,
	}
}
