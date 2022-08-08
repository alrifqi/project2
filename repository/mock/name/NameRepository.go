package mock

import (
	"github.com/alrifqi/tokenomy-test/entity"
)

var (
	MockRepoReturn []entity.NameEntity
	MockRepoError  error
)

type NameRepositoryIfaceMock interface {
	GetName(param map[string]string) ([]entity.NameEntity, error)
}
type NameRepositoryMock struct{}

func (r *NameRepositoryMock) GetName(param map[string]string) ([]entity.NameEntity, error) {
	return MockRepoReturn, MockRepoError
}

func InitNameRepositoryMock() NameRepositoryIfaceMock {
	return &NameRepositoryMock{}
}
