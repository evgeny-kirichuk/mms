package main

import (
	"flag"
	"goapp/internal/log"
	"goapp/internal/scylla"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}


func main() {
	addr := flag.String("addr", ":3000", "http service address")
	flag.Parse()

	logger := log.CreateLogger("info")

	cluster := scylla.CreateCluster(gocql.Quorum, "catalog", "scylla-node1", "scylla-node2", "scylla-node3")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		logger.Fatal("unable to connect to scylla", zap.Error(err))
	}
	defer session.Close()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Hello World"})
	})

	app.Listen(*addr)

	// scylla.SelectQuery(session, logger)
	// insertQuery(session, logger)
	// scylla.SelectQuery(session, logger)
	// deleteQuery(session, logger)
	// scylla.SelectQuery(session, logger)
}

func insertQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Inserting Mike")
	if err := session.Query("INSERT INTO mutant_data (first_name,last_name,address,picture_location) VALUES ('Mike','Tyson','1515 Main St', 'http://www.facebook.com/mtyson')").Exec(); err != nil {
		logger.Error("insert catalog.mutant_data", zap.Error(err))
	}
}

func deleteQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Deleting Mike")
	if err := session.Query("DELETE FROM mutant_data WHERE first_name = 'Mike' and last_name = 'Tyson'").Exec(); err != nil {
		logger.Error("delete catalog.mutant_data", zap.Error(err))
	}
}
