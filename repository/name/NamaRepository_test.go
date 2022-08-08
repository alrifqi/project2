package name

import (
	"reflect"
	"testing"

	"github.com/alrifqi/tokenomy-test/entity"
)

var (
	mockNameRepository *NameRepository
)

func provideTest(t *testing.T) func() {
	mockNameRepository = &NameRepository{}
	return func() {}
}

func Test_GetNumber(t *testing.T) {
	tests := []struct {
		name            string
		param           map[string]string
		expectingResult []entity.NameEntity
		expectingErr    error
	}{
		{
			name:            "Success getting all name data",
			param:           map[string]string{},
			expectingResult: entity.NameDataDummy,
			expectingErr:    nil,
		},
		{
			name: "Success getting data with filter id=1",
			param: map[string]string{
				"1": "1",
			},
			expectingResult: []entity.NameEntity{
				{
					Id:   1,
					Name: "A",
				},
			},
			expectingErr: nil,
		},
	}
	for _, tc := range tests {
		close := provideTest(t)
		defer close()

		t.Run(tc.name, func(t *testing.T) {
			data, err := mockNameRepository.GetName(tc.param)
			if !reflect.DeepEqual(tc.expectingErr, err) {
				t.Errorf("expected (%b), got (%b)", tc.expectingErr, err)
			}
			if !reflect.DeepEqual(tc.expectingResult, data) {
				t.Errorf("expected (%b), got (%b)", tc.expectingErr, err)
			}
		})
	}
}
