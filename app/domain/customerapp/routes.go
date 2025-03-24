package customerapp

import (
	"github.com/charlieroth/slot/business/domain/userbus"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Config struct {
	Logger  *zerolog.Logger
	UserBus *userbus.Business
}

func Routes(r *gin.Engine, cfg Config) {
	api := newApp(cfg.UserBus)

	r.GET("/users", api.query())
	r.GET("/users/:id", api.queryByID())
	r.POST("/users", api.create())
	r.PUT("/users/:id", api.update())
	r.DELETE("/users/:id", api.delete())
}
