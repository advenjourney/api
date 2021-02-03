package mysql

import (
	"database/sql"
	"log"

	"github.com/advenjourney/api/pkg/config"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"

	// initialize migrate file
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *sql.DB

func InitDB(config config.Config) {
	log.Println(config.Database.DSN)
	db, err := sql.Open("mysql", config.Database.DSN)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

func Migrate() {
	log.Println("debug1")
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("debug2")
	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/mysql/migrations",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
