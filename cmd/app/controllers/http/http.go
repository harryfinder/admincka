package http

import (
	"context"
	"github.com/activ-capital/partner-service/cmd/app/controllers"
	"github.com/activ-capital/partner-service/internal/models"
	"github.com/activ-capital/partner-service/internal/usecase"
	pkghttp "github.com/activ-capital/partner-service/pkg/controller/http"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type server struct {
	usecase usecase.Usecase
	srv     pkghttp.Server
}

func NewController(usecase usecase.Usecase, srv pkghttp.Server) controllers.Controller {
	return &server{usecase: usecase, srv: srv}
}

func adaptHandler(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}
func (s *server) Serve(ctx context.Context, config *models.Configuration) error {
	return s.srv.Serve(ctx, config, []pkghttp.Route{

		{Method: http.MethodGet, Path: "/swagger/*any", Handler: adaptHandler(httpSwagger.WrapHandler)},
		{Method: http.MethodGet, Path: "/ping", Handler: s.ping},
	})
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
