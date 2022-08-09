package repository

import (
	"github.com/alrifqi/tokenomy-test/repository/name"
)

// Repository is representation of all child module in repository module
type Repository struct {
	NameRepo name.NameRepositoryIface
}

// Func to initiate Repository module
// this function will initiate all child repository module
func Init() *Repository {
	nameRepo := name.InitNameRepository()
	return &Repository{
		NameRepo: nameRepo,
	}
}
