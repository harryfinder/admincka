package http

import (
	"github.com/activ-capital/partner-service/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp models.Response
	)
	resp.Send(w, models.SUCCESS, "Pong")
	return
}
