package controller

import (
	"github.com/alrifqi/tokenomy-test/controller/name"
	usecase "github.com/alrifqi/tokenomy-test/usecase"
)

type Controller struct {
	NameController name.NameControllerIface
	uc             usecase.Usecase
}

func Init(
	uc usecase.Usecase,
) *Controller {
	nameController := name.InitNameController(uc.NameUsecase)
	return &Controller{
		NameController: nameController,
		uc:             uc,
	}
}
