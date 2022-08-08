package name

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/alrifqi/tokenomy-test/entity"
	nameRepoMock "github.com/alrifqi/tokenomy-test/repository/mock/name"
	"github.com/alrifqi/tokenomy-test/utils"
)

var (
	mockNameRepository nameRepoMock.NameRepositoryIfaceMock
	mockNameUsecase    *NameUsecase
)

func provideTest(t *testing.T) func() {
	mockNameRepository = nameRepoMock.InitNameRepositoryMock()
	mockNameUsecase = &NameUsecase{
		nameRepo: mockNameRepository,
	}
	return func() {}
}

func Test_GetName(t *testing.T) {
	tests := []struct {
		name               string
		param              []string
		nameRepoDataReturn []entity.NameEntity
		nameRepoErrReturn  error
		expectingResult    []entity.NameEntity
		expectingErr       error
	}{
		{
			name:               "Success getting all name data",
			param:              []string{},
			nameRepoDataReturn: entity.NameDataDummy,
			nameRepoErrReturn:  nil,
			expectingErr:       nil,
			expectingResult:    entity.NameDataDummy,
		},
		{
			name:  "Success getting data with 1 filter",
			param: []string{"1"},
			nameRepoDataReturn: []entity.NameEntity{
				{
					Id:   1,
					Name: "A",
				},
			},
			nameRepoErrReturn: nil,
			expectingErr:      nil,
			expectingResult: []entity.NameEntity{
				{
					Id:   1,
					Name: "A",
				},
			},
		},
		{
			name:               "Success getting all name data",
			param:              []string{},
			nameRepoDataReturn: []entity.NameEntity{},
			nameRepoErrReturn:  nil,
			expectingErr: &utils.CustomError{
				StatusCode: http.StatusNotFound,
				Msg:        "resource with ID not exist",
			},
			expectingResult: []entity.NameEntity{},
		},
		{
			name:               "Error resource data from repository",
			param:              []string{},
			nameRepoDataReturn: []entity.NameEntity{},
			nameRepoErrReturn:  errors.New("Name source data is undefined"),
			expectingErr: &utils.CustomError{
				StatusCode: http.StatusInternalServerError,
				Msg:        "Internal Server Error",
			},
			expectingResult: []entity.NameEntity{},
		},
	}

	for _, tc := range tests {
		close := provideTest(t)
		defer close()

		t.Run(tc.name, func(t *testing.T) {
			nameRepoMock.MockRepoReturn = tc.nameRepoDataReturn
			nameRepoMock.MockRepoError = tc.nameRepoErrReturn
			data, err := mockNameUsecase.GetName(tc.param)
			if !reflect.DeepEqual(tc.expectingErr, err) {
				t.Errorf("expected (%b), got (%b)", tc.expectingErr, err)
			}
			if !reflect.DeepEqual(tc.expectingResult, data) {
				t.Errorf("expected (%b), got (%b)", tc.expectingErr, err)
			}
		})
	}
}
