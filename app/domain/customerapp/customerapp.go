package customerapp

import (
	"net/http"

	"github.com/charlieroth/slot/business/domain/userbus"
	"github.com/gin-gonic/gin"
)

type app struct {
	userBus *userbus.Business
}

func newApp(userBus *userbus.Business) *app {
	return &app{
		userBus: userBus,
	}
}

func (a *app) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
	}
}

func (a *app) query() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
	}
}

func (a *app) queryByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
	}
}

func (a *app) update() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
	}
}

func (a *app) delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
	}
}
