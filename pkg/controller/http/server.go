package http

import (
	"context"
	"fmt"
	"github.com/activ-capital/partner-service/internal/models"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"time"
)

type server struct {
	srv *http.Server
}

type Server interface {
	Serve(context.Context, *models.Configuration, []Route) error
	Shutdown(ctx context.Context) error
}

func NewServer() Server {
	return &server{
		srv: &http.Server{
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 120 * time.Second,
		},
	}
}

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (s *server) Serve(ctx context.Context, cfg *models.Configuration, routes []Route) error {
	//s.srv.Addr = "0.0.0.0" + ":" + "8080"
	s.srv.Addr = cfg.Host + ":" + cfg.Port
	handler := httprouter.New()

	for _, route := range routes {
		handler.Handle(route.Method, route.Path, route.Handler)
	}
	//handler.ServeFiles("/photos/*filepath", http.Dir("./photos/"))
	fsHandler := http.StripPrefix("/photos/", http.FileServer(http.Dir("./photos/")))
	handler.Handler(http.MethodGet, "/photos/*filepath", fsHandler)

	s.srv.Handler = addCors(handler)
	fmt.Println("Serve Started on port:", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func addCors(router *httprouter.Router) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodDelete, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		MaxAge:           10,
		AllowCredentials: true,
	}).Handler(router)
}
