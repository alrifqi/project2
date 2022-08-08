package name

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/alrifqi/tokenomy-test/entity"
	nameUsecaseMock "github.com/alrifqi/tokenomy-test/usecase/mock/name"
	"github.com/alrifqi/tokenomy-test/utils"
)

var (
	mockNameUsecase nameUsecaseMock.NameUsecaseIfaceMock
	ctrl            *NameController
)

func provideTest(t *testing.T) func() {
	mockNameUsecase = nameUsecaseMock.InitNameUsecase()
	ctrl = &NameController{
		nameUsecase: mockNameUsecase,
		helper:      utils.InitHelper(),
	}
	return func() {}
}

func Test_GetName(t *testing.T) {
	tests := []struct {
		name                    string
		param                   http.Request
		urlParam                string
		nameUsecaseDataReturn   []entity.NameEntity
		nameUsecaseErrReturn    error
		expectingResult         string
		expectingRespStatusCode int
	}{
		{
			name: "Get all data with success",
			param: http.Request{
				URL: &url.URL{
					RawQuery: "",
				},
			},
			urlParam:                "/name",
			nameUsecaseDataReturn:   entity.NameDataDummy,
			nameUsecaseErrReturn:    nil,
			expectingResult:         "{\"code\":200,\"data\":[{\"id\":1,\"name\":\"A\"},{\"id\":2,\"name\":\"B\"},{\"id\":3,\"name\":\"C\"},{\"id\":4,\"name\":\"D\"}]}",
			expectingRespStatusCode: http.StatusOK,
		},
		{
			name: "Get data with 1 id",
			param: http.Request{
				URL: &url.URL{
					RawQuery: "",
				},
			},
			urlParam: "/name?id=1",
			nameUsecaseDataReturn: []entity.NameEntity{
				{
					Id:   1,
					Name: "A",
				},
			},
			nameUsecaseErrReturn:    nil,
			expectingResult:         "{\"code\":200,\"data\":[{\"id\":1,\"name\":\"A\"}]}",
			expectingRespStatusCode: http.StatusOK,
		},
		{
			name: "Get data error with not found id",
			param: http.Request{
				URL: &url.URL{
					RawQuery: "",
				},
			},
			urlParam:              "/name?id=1111",
			nameUsecaseDataReturn: []entity.NameEntity{},
			nameUsecaseErrReturn: &utils.CustomError{
				StatusCode: http.StatusNotFound,
				Msg:        "resource with ID not exist",
			},
			expectingResult:         "{\"code\":404,\"msg\":\"resource with ID not exist\"}",
			expectingRespStatusCode: http.StatusNotFound,
		},
		{
			name: "Get data error mixed found id and invalid id",
			param: http.Request{
				URL: &url.URL{
					RawQuery: "",
				},
			},
			urlParam: "/name?id=1,x",
			nameUsecaseDataReturn: []entity.NameEntity{
				{
					Id:   1,
					Name: "A",
				},
			},
			nameUsecaseErrReturn:    nil,
			expectingResult:         "{\"code\":400,\"msg\":\"invalid or empty ID: x\"}",
			expectingRespStatusCode: http.StatusBadRequest,
		},
		{
			name: "Error Internal Server Error because source data undefined",
			param: http.Request{
				URL: &url.URL{
					RawQuery: "",
				},
			},
			urlParam:              "/name",
			nameUsecaseDataReturn: []entity.NameEntity{},
			nameUsecaseErrReturn: &utils.CustomError{
				StatusCode: http.StatusInternalServerError,
				Msg:        "Internal Server Error",
			},
			expectingResult:         "{\"code\":500,\"msg\":\"Internal Server Error\"}",
			expectingRespStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		close := provideTest(t)
		defer close()

		t.Run(tc.name, func(t *testing.T) {
			nameUsecaseMock.GetNameResp = tc.nameUsecaseDataReturn
			nameUsecaseMock.GetNameRespErr = tc.nameUsecaseErrReturn
			req := httptest.NewRequest(http.MethodGet, tc.urlParam, nil)
			w := httptest.NewRecorder()
			ctrl.GetName(w, req)
			res := w.Result()
			defer res.Body.Close()
			data, _ := ioutil.ReadAll(res.Body)
			if !reflect.DeepEqual(tc.expectingRespStatusCode, res.StatusCode) {
				t.Errorf("expected (%b), got (%b)", res.StatusCode, tc.expectingRespStatusCode)
			}

			if !reflect.DeepEqual(tc.expectingResult, string(data)) {
				t.Errorf("expected (%s), got (%s)", tc.expectingResult, string(data))
			}
		})
	}
}
