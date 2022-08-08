package name

import (
	"strconv"
)

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
