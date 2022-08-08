package name

import (
	"net/http"

	"github.com/alrifqi/tokenomy-test/entity"
	"github.com/alrifqi/tokenomy-test/repository/name"
	"github.com/alrifqi/tokenomy-test/utils"
)

type NameUsecaseIface interface {
	GetNumber(param []string) ([]entity.NameEntity, error)
}
type NameUsecase struct {
	nameRepo name.NameRepositoryIface
}

func (u *NameUsecase) GetNumber(param []string) ([]entity.NameEntity, error) {
	formattedParam := map[string]string{}
	for _, p := range param {
		formattedParam[p] = p
	}
	data, err := u.nameRepo.GetName(formattedParam)
	if err != nil {
		return data, err
	}

	if len(data) < 1 {
		return data, &utils.DataNotFoundError{
			StatusCode: http.StatusNotFound,
			Msg:        "resource with ID not exist",
		}
	}
	return data, nil
}

func InitNameUsecase(
	nameRepo name.NameRepositoryIface,
) NameUsecaseIface {
	return &NameUsecase{
		nameRepo: nameRepo,
	}
}
