package scylla

import (
	"fmt"

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
	Tables    map[string]map[string]interface{} `json:"tables"`
	Keyspaces map[string]map[string]interface{} `json:"keyspaces"`
} {
	logger.Info("Displaying Results:")
	tablesIt := session.Query("SELECT * FROM system_schema.tables").Iter()
	keyspacesIt := session.Query("SELECT * FROM system_schema.keyspaces").Iter()

	res := struct {
		Tables    map[string]map[string]interface{} `json:"tables"`
		Keyspaces map[string]map[string]interface{} `json:"keyspaces"`
	}{}

	defer func() {
		if err := tablesIt.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
		if err := keyspacesIt.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
	}()

	tablesValues := map[string]map[string]interface{}{}
	keyspacesValues := map[string]map[string]interface{}{}

	for {
		// New map each iteration
		row := make(map[string]interface{})
		if !tablesIt.MapScan(row) {
			break
		}
		// Do things with row
		if tableName, ok := row["table_name"]; ok {
			tablesValues[fmt.Sprintf("%v", tableName)] = row
		}
	}

	for {
		// New map each iteration
		row := make(map[string]interface{})
		if !keyspacesIt.MapScan(row) {
			break
		}
		// Do things with row
		if keyspaceName, ok := row["keyspace_name"]; ok {
			keyspacesValues[fmt.Sprintf("%v", keyspaceName)] = row
		}
	}

	res.Tables = tablesValues
	res.Keyspaces = keyspacesValues

	return res
}
