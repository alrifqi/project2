package name

import (
	"errors"
	"strconv"

	"github.com/alrifqi/tokenomy-test/entity"
)

type NameRepository struct{}

type NameRepositoryIface interface {
	GetName(param map[string]string) ([]entity.NameEntity, error)
}

func (r *NameRepository) GetName(param map[string]string) ([]entity.NameEntity, error) {
	var data []entity.NameEntity
	data = entity.NameDataDummy

	if len(data) < 1 {
		return data, errors.New("Name source data is undefined")
	}

	if len(param) == 0 {
		return data, nil
	}

	data = r._contains(param, data)
	return data, nil
}

func (r *NameRepository) _contains(param map[string]string, nameDatas []entity.NameEntity) []entity.NameEntity {
	foundData := []entity.NameEntity{}
	for _, s := range nameDatas {
		idx := strconv.Itoa(s.Id)
		_, f := param[idx]
		if f {
			foundData = append(foundData, s)
		}
	}
	return foundData
}

func InitNameRepository() NameRepositoryIface {
	return &NameRepository{}
}
