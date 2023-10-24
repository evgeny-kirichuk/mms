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
	addr := flag.String("addr", ":8000", "http service address")
	flag.Parse()

	logger := log.CreateLogger("info")

	cluster := scylla.CreateCluster(gocql.Quorum, "system", "scylla-node1", "scylla-node2", "scylla-node3")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		logger.Fatal("unable to connect to scylla", zap.Error(err))
	}
	defer session.Close()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Hello World!!!"})
	})

	apiv1.Get("/select", func(c *fiber.Ctx) error {
		res := scylla.SelectQuery(session, logger)
		return c.JSON(res)
	})

	apiv1.Get("/tabless", func(c *fiber.Ctx) error {
		res := scylla.SelectTables(session, logger)
		return c.JSON(res)
	})

	apiv1.Get("/keyspaces", func(c *fiber.Ctx) error {
		res := scylla.SelectKeyspaces(session, logger)
		return c.JSON(res)
	})

	app.Listen(*addr)
}
