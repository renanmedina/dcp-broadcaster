package utils

import (
	"database/sql"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var dbMigrator *DatabaseMigrator
var dbConnection *gorm.DB

func init() {
	initDB()
}

func initDB() {
	configs := GetConfigs()
	openedDb, err := sql.Open("postgres", configs.DbConnectionInfo())

	openedDb.SetMaxOpenConns(20) // Sane default
	openedDb.SetMaxIdleConns(0)
	openedDb.SetConnMaxLifetime(time.Nanosecond)

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(gormPostgres.New(gormPostgres.Config{Conn: openedDb}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		panic(err)
	}

	if err = db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	dbConnection = db
	dbMigrator, err = NewDatabaseMigrator(openedDb, configs.MIGRATIONS_PATH)

	if err != nil {
		panic(err)
	}
}

func GetDatabaseConnection() *gorm.DB {
	return dbConnection
}

func GetDatabaseMigrator() *DatabaseMigrator {
	return dbMigrator
}
