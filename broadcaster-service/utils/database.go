package utils

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

type DatabaseAdapdater struct {
	db *sql.DB
}

type DbRecordable interface {
	Persisted() bool
}

var dbAdapter *DatabaseAdapdater

func init() {
	initDB()
}

func initDB() {
	configs := GetConfigs()
	app := newNewRelicApp()
	app.WaitForConnection(5 * time.Second)
	openedDb, err := sql.Open("nrpostgres", configs.DbConnectionInfo())

	openedDb.SetMaxOpenConns(20) // Sane default
	openedDb.SetMaxIdleConns(0)
	openedDb.SetConnMaxLifetime(time.Nanosecond)

	if err != nil {
		panic(err)
	}

	dbAdapter = &DatabaseAdapdater{openedDb}
}

func GetDatabase() *DatabaseAdapdater {
	return dbAdapter
}

func (adapter *DatabaseAdapdater) GetConnection() *sql.DB {
	return adapter.db
}

func (adapter *DatabaseAdapdater) Insert(tableName string, fieldsAndValues map[string]interface{}) (bool, error) {
	columns := getKeys(fieldsAndValues)
	fieldValues := getValuesOrdered(columns, fieldsAndValues)

	_, errInsert := squirrel.
		Insert(tableName).
		Columns(columns...).
		Values(fieldValues...).
		RunWith(adapter.db).
		PlaceholderFormat(squirrel.Dollar).
		Exec()

	if errInsert != nil {
		return false, errInsert
	}

	return true, nil
}

func (adapter *DatabaseAdapdater) Select(fields string, tableName string, wheres interface{}) (*sql.Rows, error) {
	if fields == "" {
		fields = "*"
	}

	rows, errSelect := squirrel.Select(fields).
		From(tableName).
		Where(wheres).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(adapter.db).
		Query()

	if errSelect != nil {
		return nil, errSelect
	}

	return rows, nil
}

func (adapter *DatabaseAdapdater) SelectOne(fields string, tableName string, wheres interface{}) *squirrel.RowScanner {
	if fields == "" {
		fields = "*"
	}

	scanner := squirrel.Select(fields).
		From(tableName).
		Where(wheres).
		RunWith(adapter.db).
		PlaceholderFormat(squirrel.Dollar).
		QueryRow()

	return &scanner
}

func (adapter *DatabaseAdapdater) Update(tableName string, fieldsAndValues map[string]interface{}, wheres interface{}) (bool, error) {
	updateBuilder := squirrel.Update(tableName)

	for fieldName, fieldVal := range fieldsAndValues {
		updateBuilder = updateBuilder.Set(fieldName, fieldVal)
	}

	updateBuilder.Where(wheres)

	_, errUpdate := updateBuilder.RunWith(adapter.db).PlaceholderFormat(squirrel.Dollar).Exec()

	if errUpdate != nil {
		return false, errUpdate
	}

	return true, nil
}

func (adapter *DatabaseAdapdater) UpdateById(tableName string, id string, fieldsAndValues map[string]interface{}) (bool, error) {
	wheres := map[string]interface{}{
		"id": id,
	}

	return adapter.Update(tableName, fieldsAndValues, wheres)
}

func (adapter *DatabaseAdapdater) Delete(tableName string, wheres interface{}) (bool, error) {
	deleteBuilder := squirrel.Delete(tableName)
	deleteBuilder.Where(wheres)

	_, errDelete := deleteBuilder.RunWith(adapter.db).PlaceholderFormat(squirrel.Dollar).Exec()

	if errDelete != nil {
		return false, errDelete
	}

	return true, nil
}

func getKeys(mapVar map[string]interface{}) []string {
	keys := make([]string, len(mapVar))

	i := 0
	for key := range mapVar {
		keys[i] = key
		i++
	}

	return keys
}

func getValuesOrdered(columns []string, mapVar map[string]interface{}) []interface{} {
	vals := make([]interface{}, len(mapVar))

	i := 0
	for _, column := range columns {
		vals[i] = mapVar[column]
		i++
	}

	return vals
}
