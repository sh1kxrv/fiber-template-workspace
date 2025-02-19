package main

import (
	"context"
	"os"
	"os/signal"
	"shared/cache/memcache"
	"shared/driver/mongodb"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	// docs
	"rest_service/internal/boot"
	"rest_service/internal/config"
	_ "rest_service/internal/docs"
	"rest_service/internal/module"

	"github.com/sirupsen/logrus"
)

// @title <SHIKARU TEMPLATE>
// @version 1.0
// @host 127.0.0.1:3000
// @securityDefinitions.apikey ApiKeyAuth
// @in Header
// @name Authorization
// @description Bearer Token authortization
// @BasePath /
func main() {
	memoryCacheContext, cancel := context.WithCancel(context.Background())

	if err := boot.InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	boot.InitLogrus()

	var cfg = config.InitConfig()

	logrus.Info("Database initialization")

	db, err := mongodb.NewMongoInstance(cfg.Database.ConnectionURL, cfg.Database.Name)
	if err != nil {
		logrus.Fatalf("Failed to connect to mongodb: %s", err.Error())
	}

	memoryCache := boot.InitMemoryCache(memoryCacheContext)

	var app = fiber.New(fiber.Config{
		ServerHeader: "<SHIKARU>",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(memcache.MemoryCacheMiddleware(memoryCache))
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	module.InitRouter(app, db)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go boot.GracefulShutdown(quit, db, cancel)

	logrus.Fatal(app.Listen(":3000"))
}
