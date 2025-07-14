package articles

import (
	"context"
	"fmt"
	"heartbit/utils"
	"log"

	"github.com/jackc/pgx/v5"
)

var schema = `
	CREATE TABLE articles (
	  id   TEXT PRIMARY KEY,
	  --
	  title TEXT NOT NULL,
	  description  TEXT,
	  content TEXT NOT NULL,
	  --
	  author TEXT NOT NULL,
	  published BOOLEAN NOT NULL,
	  --
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

func InitializeDatabase() {
	ctx := context.Background()

	username := utils.GetEnvVariable("POSTGRES_USER")
	password := utils.GetEnvVariable("POSTGRES_PASSWORD")
	dbName := utils.GetEnvVariable("POSTGRES_DB")

	url := fmt.Sprintf("postgres://%s:%s@database:5432/%s", username, password, dbName)

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, schema)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Table initialized successfully")
}
