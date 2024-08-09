package main

import (
	"context"
	"github.com/activ-capital/partner-service/cmd/app/controllers/http"
	"github.com/activ-capital/partner-service/internal/configs"
	"github.com/activ-capital/partner-service/internal/database/clients"
	"github.com/activ-capital/partner-service/internal/entity"
	"github.com/activ-capital/partner-service/internal/usecase"
	pkghttp "github.com/activ-capital/partner-service/pkg/controller/http"
	pkgpostgres "github.com/activ-capital/partner-service/pkg/storage/postgres/pgx"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	cfg, err := configs.InitConfig()
	client, err := pkgpostgres.NewClient(ctx, cfg)
	if err != nil {
		panic(err)
	}

	database := clients.New(client)
	entities := entity.New(database)
	useCase := usecase.New(entities)
	srv := pkghttp.NewServer()
	controller := http.NewController(useCase, srv)

	wg.Add(1)
	go func() {
		defer wg.Done()

		quitCh := make(chan os.Signal, 1)
		signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		sig := <-quitCh
		log.Println(sig)
		cancelFunc()

		ctx, cancelFunc = context.WithTimeout(ctx, time.Second*10)
		defer cancelFunc()

		if err := controller.Shutdown(ctx); err != nil {
			log.Println()
		}

		log.Println("Server - < finished goroutines")
	}()
	if err = controller.Serve(ctx, &cfg); err != nil {
		panic(err)
	}
	wg.Wait()

	log.Println("Server - finished main goroutines")
}
