package controllers

import (
	"context"
	"github.com/activ-capital/partner-service/internal/models"
)

type Controller interface {
	Serve(context.Context, *models.Configuration) error
	Shutdown(ctx context.Context) error
}
