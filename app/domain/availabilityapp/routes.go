package availabilityapp

import (
	"github.com/charlieroth/slot/business/domain/availabilitybus"
	"github.com/gin-gonic/gin"
)

type Config struct {
	AvailabilityBus *availabilitybus.Business
}

func Routes(router *gin.Engine, cfg Config) {
	app := newApp(cfg.AvailabilityBus)

	router.GET("/availability", app.query())
}
