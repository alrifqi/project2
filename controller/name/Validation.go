package name

import (
	"strconv"
)

// Func for validate given query parameter
// will check if given parameter is invalid (criteria invalid: non-number parameter)
// it will return formatted/mapped parameter into string and error
func (v *NameController) validateGetName(payloads []string) ([]string, error) {
	var formattedPayload []string
	var invalidPayload []string
	var err error
	for _, p := range payloads {
		_, err = strconv.Atoi(p)
		if err != nil {
			invalidPayload = append(invalidPayload, p)
		}
		formattedPayload = append(formattedPayload, p)
	}

	if err != nil {
		return invalidPayload, err
	}
	return formattedPayload, nil
}
