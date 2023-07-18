package repository

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func ConexaoDb() *bun.DB {
	dsn := "postgres://postgres:smarters123@localhost:5432/tesouro?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func CreateTables(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().IfNotExists().Model((*Users)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = db.NewCreateTable().IfNotExists().Model((*Caminho)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = db.NewCreateTable().IfNotExists().Model((*Pista)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
