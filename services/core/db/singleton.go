package db

import (
	"core/db/repository/user"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"pg"
)

var pool *pgxpool.Pool

var (
	userRepo *user.DBRepo
)

func makeRepo(p *pgxpool.Pool) {
	userRepo = user.NewDBRepo(p)
}

// Init initializes the database.
// Establishes a connection pool and migrates the data schema
func Init(cfg *pg.Config) {
	_pool, err := pg.MakePool(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to make pgx pool: %v", err))
	}

	pool = _pool

	makeRepo(pool)
}

func GetUserRepo() *user.DBRepo {
	return userRepo
}

// Close closes the connection pool
func Close() {
	if pool != nil {
		pool.Close()
	}
}
