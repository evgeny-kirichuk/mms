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

func SelectTables(session *gocql.Session, logger *zap.Logger) map[string]string {
	logger.Info("Displaying Results:")
	q := session.Query("SELECT table_name FROM system_schema.tables WHERE keyspace_name='catalog'")
	var tableName string
	it := q.Iter()
	res := make(map[string]string)
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
	}()
	for it.Scan(&tableName) {
		logger.Info("\t" + tableName)
		res[tableName] = tableName
	}

	return res
}