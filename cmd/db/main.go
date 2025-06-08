package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	m  *migrate.Migrate
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: db <command>")
		log.Println("Commands:")
		log.Println("  up       - Run migrations")
		log.Println("  down     - Rollback migrations")
		log.Println("  version  - Run migration to a specific version")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "up":
		runMigrations(m)
	case "down":
		rollbackMigrations(m)
	case "version":
		if len(os.Args) < 3 {
			log.Println("Usage: db version <version_number>")
			os.Exit(1)
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Printf("Invalid version number: %v\n", err)
			os.Exit(1)
		}
		runMigrationAtVersion(m, uint(version))
	default:
		log.Println("Invalid command")
		os.Exit(1)
	}
}

func init() {
	db = dataBaseConnection()
	defer db.Close()

	m = createMigrateInstance(db)
}

func dataBaseConnection() *sql.DB {
	log.Println("Connecting to database ...")
	connStr := "user=admin password=admin dbname=restaurant_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to postgres database")

	return db
}

func createMigrateInstance(db *sql.DB) *migrate.Migrate {
	log.Println("Creating migration instance ...")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Error creating database instance:", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}

	migrationsPath := filepath.Join(wd, "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver)
	if err != nil {
		log.Fatal("Error creating migration instance:", err)
	}

	log.Println("Successfully created migration instance")

	return m
}

func runMigrations(m *migrate.Migrate) {
	log.Println("Running migrations ...")
	if err := m.Run(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Error running migrations:", err)
	}

	log.Println("Migrations completed successfully")
}

func rollbackMigrations(m *migrate.Migrate) {
	log.Println("Rolling back migrations ...")
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Error rolling back migrations:", err)
	}

	log.Println("Rollback completed successfully")
}

func runMigrationAtVersion(m *migrate.Migrate, version uint) {
	log.Printf("Running migration to version %d ...\n", version)
	if err := m.Migrate(version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Error running migration:", err)
	}

	log.Printf("Migration to version %d completed successfully\n", version)
}
