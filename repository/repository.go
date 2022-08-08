package repository

import (
	"github.com/alrifqi/tokenomy-test/repository/name"
)

type Repository struct {
	NameRepo name.NameRepositoryIface
}

func Init() *Repository {
	nameRepo := name.InitNameRepository()
	return &Repository{
		NameRepo: nameRepo,
	}
}
