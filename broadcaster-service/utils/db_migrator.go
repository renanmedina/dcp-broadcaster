package utils

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

type DatabaseMigrator struct {
	migratorConnection *migrate.Migrate
	migrationsPath     string
}

func MigrateDb(dir string) {
	migrator := GetDatabaseMigrator()
	logger := GetApplicationLogger()

	logger.Info("Migrating database")
	err := migrator.Migrate(dir)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Migration success")
}

func NewDatabaseMigrator(conn *sql.DB, migrationsPath string) (*DatabaseMigrator, error) {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})

	if err != nil {
		return nil, err
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)

	if err != nil {
		return nil, err
	}

	return &DatabaseMigrator{migrator, migrationsPath}, nil
}

func (migrator *DatabaseMigrator) Migrate(dir string) error {
	var err error

	if dir == "" || dir == "up" {
		err = migrator.migratorConnection.Up()
	} else {
		err = migrator.migratorConnection.Down()
	}

	if err != nil {
		return err
	}

	return nil
}
