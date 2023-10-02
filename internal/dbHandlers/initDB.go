package dbHandlers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func OpenDBConnection() *pgx.Conn {
	connection_string := "postgres://postgres:password@localhost:5432/mathtestr"
	db, err := pgx.Connect(context.Background(), connection_string)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	return db
}
