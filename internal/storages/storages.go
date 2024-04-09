package storages

import (
	cStorage "goTest/internal/modules/car/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storages struct {
	cStorage.Carer
}

func NewStorages(pool *pgxpool.Pool) *Storages {
	return &Storages{
		Carer: cStorage.NewCarStorage(pool),
	}
}
