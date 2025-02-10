// Package database contains the database functions for the service
package database

import (
	"context"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres, blank import needed
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JakeHumphries/investment-service/config"
	"github.com/JakeHumphries/investment-service/models"
)

// NewClient creates a new instance of the database client
func NewClient(ctx context.Context, config config.Config) (Client, error) {
	// Apply migrations
	m, err := migrate.New("file://"+config.MigrationsPath, config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	// Build connection pool
	pgxConf, err := pgxpool.ParseConfig(config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return nil, err
	}

	dbClient := NewPostgresClient(pool)

	return &dbClient, nil
}

// PostgresClient is the postgres implementation of the database client.
type PostgresClient struct {
	Database PgxInterface
}

// NewPostgresClient creates a new postgres client.
func NewPostgresClient(database PgxInterface) PostgresClient {
	return PostgresClient{
		Database: database,
	}
}

// Client is the interface for the database client, used mainly for mocking
//
//go:generate mockery --name Client
type Client interface {
	GetFunds(ctx context.Context) ([]models.Fund, error)
	GetFundByID(ctx context.Context, fundID string) (*models.Fund, error)

	CreateInvestment(ctx context.Context, investment *models.Investment) (*models.Investment, error)
	GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) ([]models.Investment, *string, error)
}

// QueryExecer is the interface for being able to query or exec against a SQL database
type QueryExecer interface {
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

// PgxInterface more or less matches the interface for Pgx pools
type PgxInterface interface {
	QueryExecer

	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
	Begin(context.Context) (pgx.Tx, error)
	Ping(context.Context) error
}
