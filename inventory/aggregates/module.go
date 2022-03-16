package aggregates

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetOne() string
	GetAll() []string
}

type Module struct {
	// my deps
	repository Repository

	// my props
	version string
}

func New(version string, repository Repository) *Module {
	return &Module{
		version:    version,
		repository: repository,
	}
}

func (m *Module) Boot(r *gin.Engine) error {

	// Create our route group
	grp := r.Group("/aggregate/:id")

	// Attach specific middlewares
	grp.Use(m.sanitize())

	// And then our own routes
	grp.GET("/getOne", m.getOne())
	grp.GET("/getAll", m.getAll())

	return nil
}

func (m *Module) getAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"all": m.repository.GetAll(),
		})
	}
}

func (m *Module) getOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"one": m.repository.GetOne(),
		})
	}
}

func (m *Module) sanitize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body == http.NoBody {
			return
		}

		if c.ContentType() == m.version {
			return
		}

		c.Error(fmt.Errorf(
			"Unsupported content type %q; expected client to send %q",
			c.ContentType(),
			m.version,
		))
	}
}
