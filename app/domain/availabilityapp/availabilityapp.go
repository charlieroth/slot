package availabilityapp

import (
	"net/http"

	"github.com/charlieroth/slot/business/domain/availabilitybus"
	"github.com/gin-gonic/gin"
)

type app struct {
	availabilityBus *availabilitybus.Business
}

func newApp(availabilityBus *availabilitybus.Business) *app {
	return &app{
		availabilityBus: availabilityBus,
	}
}

func (a *app) query() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)
	}
}
