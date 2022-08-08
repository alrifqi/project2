package name

import (
	"strconv"
)

func (v *NameController) validateGetName(payloads []string) ([]int, error) {
	var formattedPayload []int
	for _, p := range payloads {
		converted, err := strconv.Atoi(p)
		if err != nil {
			return formattedPayload, err
		}
		formattedPayload = append(formattedPayload, converted)
	}
	return formattedPayload, nil
}
