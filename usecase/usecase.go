package usecase

import (
	"github.com/alrifqi/tokenomy-test/repository"
	nameUc "github.com/alrifqi/tokenomy-test/usecase/name"
)

// Usecase type is representation of usecase attribute
type Usecase struct {
	NameUsecase nameUc.NameUsecaseIface
	repo        repository.Repository
}

// Func for initiate usecase module and to inject repository module into usecase
// Init function will register all usecase module
func Init(
	repo *repository.Repository,
) *Usecase {
	nameUsecase := nameUc.InitNameUsecase(repo.NameRepo)
	return &Usecase{
		NameUsecase: nameUsecase,
		repo:        *repo,
	}
}
