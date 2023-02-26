package repository

import (
	"context"
	"log"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository/sqlc"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	sqlc.Querier
	StartTX(ctx context.Context) (pgx.Tx, *sqlc.Queries, error)
}

func NewStore(ctx context.Context, psqlUri string) Store {
	dbConn, err := pgxpool.Connect(ctx, psqlUri)
	if err != nil {
		log.Fatal("failed to connecto to psql", err)
	}

	log.Panicln("successfully connected to postgres")
	return &postgresStore{
		Queries: sqlc.New(dbConn),
		DB:      dbConn,
	}
}

// postgresStore provides all functions to execute SQL queries and transactions
type postgresStore struct {
	*sqlc.Queries
	DB *pgxpool.Pool
}

// StartTX begins new transaction
func (store *postgresStore) StartTX(ctx context.Context) (pgx.Tx, *sqlc.Queries, error) {
	tx, err := store.DB.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.Serializable,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.NotDeferrable,
	})
	if err != nil {
		return nil, nil, err
	}

	q := sqlc.New(tx)
	return tx, q, nil
}
