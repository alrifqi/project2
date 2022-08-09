package main

import (
	"fmt"
	"net/http"

	controller "github.com/alrifqi/tokenomy-test/controller"
	repository "github.com/alrifqi/tokenomy-test/repository"
	usecase "github.com/alrifqi/tokenomy-test/usecase"
)

func main() {
	repo := repository.Init()    // init repository module
	uc := *usecase.Init(repo)    // init usecase module and inject repo as dependencies of usecase module
	ctrl := *controller.Init(uc) // init controller module and inject usecase as dependencies of controller module
	fmt.Println("Start Application")

	// name route
	http.HandleFunc("/name", ctrl.NameController.GetName)

	http.ListenAndServe(":8080", nil)
}
