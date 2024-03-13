package vehicle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/vgiudicelli1/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(rw, "Bip Boup ID Non Valide", http.StatusBadRequest)
	}
	isExist, _ := d.store.Vehicle().Delete(r.Context(), id)
	if isExist {
		rw.WriteHeader(http.StatusNoContent)
		fmt.Printf("Vehicle Deleted")
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Printf("Vehicle Not Found")
	}
}
