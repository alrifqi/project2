package main

import (
	"fmt"
	"net/http"

	controller "github.com/alrifqi/tokenomy-test/controller"
	repository "github.com/alrifqi/tokenomy-test/repository"
	usecase "github.com/alrifqi/tokenomy-test/usecase"
)

func main() {
	repo := repository.Init()    // init repository
	uc := *usecase.Init(repo)    // init usecase
	ctrl := *controller.Init(uc) // init controller
	fmt.Println("Start Application")

	// name route
	http.HandleFunc("/name", ctrl.NameController.GetNumber)

	fmt.Println(http.ListenAndServe(":8080", nil))
}
