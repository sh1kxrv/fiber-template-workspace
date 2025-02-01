package main

import (
	"common/cache/memcache"
	"common/driver/mongodb"
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"service/boot"
	"service/config"
	"syscall"

	// Swagger Docs
	_ "service/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

// @title <Service>
// @version 1.0
// @host 127.0.0.1:3001
// @securityDefinitions.apikey ApiKeyAuth
// @in Header
// @name Authorization
// @description Bearer Token Authortization
// @BasePath /
func main() {
	memoryCacheContext, cancel := context.WithCancel(context.Background())

	if err := boot.InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	boot.InitLogrus()

	cfg := config.InitConfig()

	logrus.Info("Database initialization")

	db, err := mongodb.NewMongoInstance(cfg.Database.ConnectionURL)
	if err != nil {
		logrus.Fatalf("Failed to initialize database: %s", err.Error())
	}

	memoryCache := boot.InitMemoryCache(memoryCacheContext)

	app := fiber.New(fiber.Config{
		ServerHeader: "<SHIKARU>",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(memcache.MemoryCacheMiddleware(memoryCache))
	app.Use(cors.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go boot.GracefulShutdown(quit, db, cancel)

	logrus.Fatal(app.Listen(":3000"))
}
