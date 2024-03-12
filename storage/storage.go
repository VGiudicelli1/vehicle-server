package storage

import (
	"github.com/vgiudicelli1/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
