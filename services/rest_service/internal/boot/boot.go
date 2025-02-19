package boot

import (
	"context"
	"os"
	"shared/cache/memcache"
	"shared/driver/mongodb"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitMemoryCache(ctx context.Context) *memcache.MemoryCache {
	return memcache.NewMemoryCache(ctx)
}

func InitViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func InitLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}

func GracefulShutdown(quit chan os.Signal, db *mongodb.MongoInstance, cancel context.CancelFunc) {
	<-quit
	logrus.Debug("Graceful shutdown...")
	if err := db.Client.Disconnect(context.Background()); err != nil {
		logrus.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
	logrus.Debug("Connection to MongoDB closed.")

	logrus.Debug("Canceling memory cache context...")
	cancel()

	os.Exit(0)
}
