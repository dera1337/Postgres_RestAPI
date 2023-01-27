package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Conn *pgxpool.Pool
)

func InitDB() error {
	bgCtx := context.Background()
	conn, err := pgxpool.New(bgCtx, "postgres://postgres:Paramore_12345@localhost:5432/postgres")
	if err != nil {
		return err
	}
	Conn = conn

	err = initTables(Conn)
	if err != nil {
		return err
	}

	return nil
}

func CloseConnection() {
	Conn.Close()
}

func initTables(conn *pgxpool.Pool) error {
	sqlFilePath := `C:\Users\yoris\Desktop\go-wokspace\src\exercisee\init.sql`

	c, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}
	sql := string(c)

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	return nil
}
