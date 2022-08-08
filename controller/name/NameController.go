package name

import (
	"net/http"
	"strings"

	nameUc "github.com/alrifqi/tokenomy-test/usecase/name"
	utils "github.com/alrifqi/tokenomy-test/utils"
)

type NameControllerIface interface {
	GetNumber(w http.ResponseWriter, r *http.Request)
}

type NameController struct {
	nameUsecase nameUc.NameUsecaseIface
	helper      utils.HelperIface
}

func (c *NameController) GetNumber(w http.ResponseWriter, r *http.Request) {
	var payload []string
	query := r.URL.Query()

	params := query["id"]
	if len(params) > 0 {
		payload = strings.Split(params[0], ",")
	}

	_, _ = c.validateGetName(payload)

	data, _ := c.nameUsecase.GetNumber(payload)
	c.helper.HttpResp(w, r, http.StatusOK, data)
}

func InitNameController(
	nameUsecase nameUc.NameUsecaseIface,
) NameControllerIface {
	return &NameController{
		nameUsecase: nameUsecase,
		helper:      utils.InitHelper(),
	}
}