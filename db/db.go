package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	Conn *pgx.Conn
)

func InitDB() error {
	bgCtx := context.Background()
	conn, err := pgx.Connect(bgCtx, "postgres://postgres:Paramore_12345@localhost:5432/postgres")
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

func CloseConnection() error {
	err := Conn.Close(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func initTables(conn *pgx.Conn) error {
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
