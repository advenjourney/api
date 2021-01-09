package mysql

import (
	"database/sql"
	"log"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"

	// initialize migrate file
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:dbpass@tcp(localhost)/advenjourney")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

func Migrate() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
