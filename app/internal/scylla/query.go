package scylla

import (
	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func SelectQuery(session *gocql.Session, logger *zap.Logger) map[string]string {
	logger.Info("Displaying Results:")
	q := session.Query("SELECT first_name,last_name,address,picture_location FROM mutant_data")
	var firstName, lastName, address, pictureLocation string
	it := q.Iter()
	res := make(map[string]string)
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
	}()
	for it.Scan(&firstName, &lastName, &address, &pictureLocation) {
		logger.Info("\t" + firstName + " " + lastName + ", " + address + ", " + pictureLocation)
		res[firstName] = lastName
	}

	return res
}

func SelectTables(session *gocql.Session, logger *zap.Logger) struct {
	Tables    map[string]string `json:"tables"`
	Keyspaces map[string]string `json:"keyspaces"`
} {
	logger.Info("Displaying Results:")
	tablesQuery := session.Query("SELECT table_name FROM system_schema.tables")
	keyspacesQuery := session.Query("SELECT * FROM system_schema.keyspaces")
	var tableName string
	var keyspaceName string

	tables := tablesQuery.Iter()
	keyspaces := keyspacesQuery.Iter()
	res := struct {
		Tables    map[string]string `json:"tables"`
		Keyspaces map[string]string `json:"keyspaces"`
	}{}

	res.Tables = map[string]string{}
	res.Keyspaces = map[string]string{}

	defer func() {
		if err := tables.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
		if err := keyspaces.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
	}()

	for tables.Scan(&tableName) {
		logger.Info("\t" + "table: " + tableName)
		res.Tables[tableName] = tableName
	}
	for keyspaces.Scan(&keyspaceName, nil, nil) {
		logger.Info("\t" + "keyspace: " + keyspaceName)
		res.Keyspaces[keyspaceName] = keyspaceName
	}

	return res
}
