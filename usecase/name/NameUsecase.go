package name

import (
	"net/http"

	"github.com/alrifqi/tokenomy-test/entity"
	"github.com/alrifqi/tokenomy-test/repository/name"
	"github.com/alrifqi/tokenomy-test/utils"
)

// NameUsecaseIface for give required method in NameUsecase method
type NameUsecaseIface interface {
	GetName(param []string) ([]entity.NameEntity, error)
}

// NameUsecase is representation of attribute of NameUsecase
type NameUsecase struct {
	nameRepo name.NameRepositoryIface
}

// Func for processing name from source & give response to controller
func (u *NameUsecase) GetName(param []string) ([]entity.NameEntity, error) {
	formattedParam := map[string]string{}
	for _, p := range param {
		formattedParam[p] = p
	}
	data, err := u.nameRepo.GetName(formattedParam)
	if err != nil {
		return data, &utils.CustomError{
			StatusCode: http.StatusInternalServerError,
			Msg:        "Internal Server Error",
		}
	}

	if len(data) < 1 {
		return data, &utils.CustomError{
			StatusCode: http.StatusNotFound,
			Msg:        "resource with ID not exist",
		}
	}
	return data, nil
}

// Func for initiate NameUsecase Module
func InitNameUsecase(
	nameRepo name.NameRepositoryIface,
) NameUsecaseIface {
	return &NameUsecase{
		nameRepo: nameRepo,
	}
}
