package vehicle

import (
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
	str := r.PathValue("id")
	id, _ := strconv.ParseInt(str, 10, 64)
	bool, _ := d.store.Vehicle().Delete(r.Context(), id)
	print(bool)
	if bool {
		rw.WriteHeader(http.StatusNoContent)
	}
	if !bool {
		rw.WriteHeader(http.StatusNotFound)
	}
}
