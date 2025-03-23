package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/charlieroth/slot/app/domain/availabilityapp"
	"github.com/charlieroth/slot/business/domain/availabilitybus"
	"github.com/charlieroth/slot/business/domain/availabilitybus/stores/scheduledb"
	"github.com/charlieroth/slot/business/sdk/sqldb"
	"github.com/charlieroth/slot/foundation/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)

	ctx := context.Background()

	if err := run(ctx, &logger); err != nil {
		logger.Fatal().Err(err).Msg("startup failed")
	}
}

func run(ctx context.Context, logger *zerolog.Logger) error {
	logger.Info().Msgf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))

	// -----------------------------------------------------
	// Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// ------------------------------------------------------
	// App starting

	logger.Info().Msg("starting service")
	defer logger.Info().Msg("shutdown complete")

	// ------------------------------------------------------
	// Database support

	logger.Info().Msg("initializing database support")
	db, err := sqldb.Open(cfg.DB)
	if err != nil {
		return err
	}

	// ------------------------------------------------------
	// Create business packages
	logger.Info().Msg("initializing business packages")
	availabilityBus := availabilitybus.NewBusiness(scheduledb.NewStore(logger, db), logger)

	// ------------------------------------------------------
	// Start API service
	logger.Info().Msg("starting API service")
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	router := buildRoutes(availabilityBus)

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info().Msg("api router started")

		if err := router.Run(fmt.Sprintf(":%d", cfg.Web.Port)); err != nil {
			serverErrors <- err
		}
	}()

	// ------------------------------------------------------
	// Shutdown
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		logger.Info().Msgf("shutdown signal received: %s", sig)
		_, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		db.Close()
		return fmt.Errorf("shutdown")
	}
}

func buildRoutes(availabilityBus *availabilitybus.Business) *gin.Engine {
	router := gin.Default()

	router.GET("liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	availabilityapp.Routes(router, availabilityapp.Config{
		AvailabilityBus: availabilityBus,
	})
	return router
}
