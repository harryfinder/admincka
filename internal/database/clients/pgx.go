package clients

import (
	"github.com/activ-capital/partner-service/internal/database"
	pkgpostgres "github.com/activ-capital/partner-service/pkg/storage"
)

type db struct {
	postgres pkgpostgres.Database
}

func New(postgres pkgpostgres.Database) database.Database {
	return &db{postgres: postgres}
}
