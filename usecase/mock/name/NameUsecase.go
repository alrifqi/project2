package mock

import (
	"github.com/alrifqi/tokenomy-test/entity"
	nameRepoMock "github.com/alrifqi/tokenomy-test/repository/mock/name"
)

var (
	GetNameResp    []entity.NameEntity
	GetNameRespErr error
)

type NameUsecaseIfaceMock interface {
	GetName(param []string) ([]entity.NameEntity, error)
}

type NameusecaseMock struct {
	nameRepo nameRepoMock.NameRepositoryMock
}

func (u *NameusecaseMock) GetName(param []string) ([]entity.NameEntity, error) {
	return GetNameResp, GetNameRespErr
}

func InitNameUsecase() NameUsecaseIfaceMock {
	return &NameusecaseMock{}
}
