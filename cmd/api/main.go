package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"

	"github.com/hndryno/go-lms-postgresql/internal/adapters/postgresql/db"
	"github.com/hndryno/go-lms-postgresql/internal/config"
	"github.com/hndryno/go-lms-postgresql/internal/customer"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	conn, err := pgx.Connect(
		ctx,
		cfg.DatabaseURL(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// Customer
	customerRepository := customer.NewRepository(queries)
	customerService := customer.NewService(customerRepository)
	customerHandler := customer.NewHandler(customerService)

	// Fiber
	app := fiber.New()

	api := app.Group("/api/v1")

	customer.RegisterRoutes(
		api,
		customerHandler,
	)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"code":    200,
			"success": true,
			"message": "OK",
		})
	})

	log.Printf(
		"server running on :%s",
		cfg.AppPort,
	)

	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}