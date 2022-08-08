package name

import (
	"net/http"
	"strings"

	nameUc "github.com/alrifqi/tokenomy-test/usecase/name"
	utils "github.com/alrifqi/tokenomy-test/utils"
)

type NameControllerIface interface {
	GetName(w http.ResponseWriter, r *http.Request)
}

type NameController struct {
	nameUsecase nameUc.NameUsecaseIface
	helper      utils.HelperIface
}

func (c *NameController) GetName(w http.ResponseWriter, r *http.Request) {
	var payload []string
	query := r.URL.Query()

	params := query["id"]
	if len(params) > 0 {
		payload = strings.Split(params[0], ",")
	}

	formattedPayload, err := c.validateGetName(payload)
	if err != nil {
		c.helper.HttpRespError(w, r, &utils.CustomError{
			StatusCode: http.StatusBadRequest,
			Msg:        "invalid or empty ID: " + strings.Join(formattedPayload, ","),
		})
		return
	}

	data, err := c.nameUsecase.GetName(formattedPayload)
	if err != nil {
		c.helper.HttpRespError(w, r, err)
		return
	} else {
		c.helper.HttpResp(w, r, http.StatusOK, data)
		return
	}
}

func InitNameController(
	nameUsecase nameUc.NameUsecaseIface,
) NameControllerIface {
	return &NameController{
		nameUsecase: nameUsecase,
		helper:      utils.InitHelper(),
	}
}
