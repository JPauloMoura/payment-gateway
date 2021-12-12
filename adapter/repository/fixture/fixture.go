package fixture

import (
	"context"
	"database/sql"
	"io/fs"
	"log"

	"github.com/maragudk/migrate"
	_ "github.com/mattn/go-sqlite3"
)

// Up recebe o path do script sql para subir a tabela
func Up(migrationsDir fs.FS) *sql.DB {
	//abre a conexão com o banco de dados
	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}

	return db
}

// Down recebe o path do script sql para dropar a tabela
func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}

	db.Close()
}
