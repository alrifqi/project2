package mock

import (
	"github.com/alrifqi/tokenomy-test/entity"
	nameRepoMock "github.com/alrifqi/tokenomy-test/repository/mock/name"
)

type NameUsecaseIfaceMock interface {
	GetName(param []string) ([]entity.NameEntity, error)
}

type NameusecaseMock struct {
	nameRepo nameRepoMock.NameRepositoryMock
}
