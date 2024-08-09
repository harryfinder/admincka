package pgx

import (
	"context"
	"errors"
	"github.com/activ-capital/partner-service/internal/models"
	"github.com/activ-capital/partner-service/pkg/storage"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type postgresDatabase struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context, databaseDsn models.Configuration) (storage.Database, error) {
	poolConfig, err := pgxpool.ParseConfig(databaseDsn.PostgresDsn)
	if err != nil {
		log.Println(err)
		return nil, errors.New("pgxpool.ParseConfig ERROR: " + err.Error())
	}
	poolConfig.MaxConns = 10
	//poolConfig.AfterConnect = afterConnect

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Println(err)
		return nil, errors.New("pgxpool.ConnectConfig ERROR: " + err.Error())
	}
	return &postgresDatabase{pool: pool}, nil
}

func (c *postgresDatabase) Close() {
	c.pool.Close()
}

func (c *postgresDatabase) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...)
}

func (c *postgresDatabase) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return c.pool.QueryRow(ctx, sql, args...)
}

func (c *postgresDatabase) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.pool.Exec(ctx, sql, arguments...)
}
