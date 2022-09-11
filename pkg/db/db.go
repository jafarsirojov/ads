package db

import (
	"ads/pkg/config"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"runtime"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
}

type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
}

type dbConn struct {
	config *config.Config
	dbPool *pgxpool.Pool
	logger *zap.Logger
}

func New(params Params) Querier {
	db := &dbConn{
		config: params.Config,
		logger: params.Logger,
	}
	return db.getConnection()
}

func (db *dbConn) getConnection() Querier {

	db.logger.Info("DB: connection create here")

	var (
		err error
		ctx = context.Background()
	)

	connPoolConfig, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"user=%s password=%s host=%s port=5432 dbname=%s sslmode=allow pool_max_conns=%d",
			db.config.DBUser,
			db.config.DBPassword,
			db.config.DBHost,
			db.config.DBName,
			runtime.NumCPU()*2,
		),
	)

	if err != nil {
		log.Fatalf("could not parse connection pool config string: %s", err.Error())
		return nil
	}

	conn, err := pgxpool.ConnectConfig(ctx, connPoolConfig)
	if err != nil {
		log.Fatalf("could not provide database connection pool due to: %s", err.Error())
	}

	db.dbPool = conn

	return db
}

func (db *dbConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return db.getConn().Exec(ctx, sql, arguments...)
}

func (db *dbConn) Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error) {
	return db.getConn().Query(ctx, sql, optionsAndArgs...)
}

func (db *dbConn) QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row {
	return db.getConn().QueryRow(ctx, sql, optionsAndArgs...)
}

func (db *dbConn) Begin(ctx context.Context) (pgx.Tx, error) {
	return db.getConn().Begin(ctx)
}

func (db *dbConn) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return db.getConn().SendBatch(ctx, batch)
}

func (db *dbConn) getConn() *pgxpool.Pool {
	return db.dbPool
}
