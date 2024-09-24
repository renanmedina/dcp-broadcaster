package utils

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func MigrateDb(dir string) {
	db := GetDatabase()
	logger := GetApplicationLogger()

	logger.Info("Migrating database")
	err := db.Migrate(dir, GetConfigs().MIGRATIONS_PATH)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Migration success")
}

func (adapter *DatabaseAdapdater) GetMigrator(migrationsPath string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(adapter.db, &postgres.Config{})

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

	return migrator, nil
}

func (adapter *DatabaseAdapdater) Migrate(dir string, migrationsPath string) error {
	migrator, err := adapter.GetMigrator(migrationsPath)

	if err != nil {
		return err
	}

	if dir == "" || dir == "up" {
		err = migrator.Up()
	} else {
		err = migrator.Down()
	}

	if err != nil {
		return err
	}

	return nil
}
