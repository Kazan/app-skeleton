package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kazan/app-skeleton/inventory/aggregates"
)

const V1 string = "application/vnd.lovelyplace.v1+json"

type Module interface {
	Boot(r *gin.Engine) error
}

func Load(r *gin.Engine) error {
	// Initialize dependencies
	repository := newRepository()

	// Configure the application modules to load
	mods := []Module{
		aggregates.New(V1, repository),
	}

	// Boot up all modules
	for _, mod := range mods {
		if err := mod.Boot(r); err != nil {
			panic(fmt.Sprintf("Failed to boot module: %s", err))
		}
	}

	return nil
}

// Left here becase... lazyness
func newRepository() *fakeRepository {
	return &fakeRepository{
		val:  "fixed",
		data: []string{"one", "two", "three"},
	}
}

type fakeRepository struct {
	data []string
	val  string
}

func (r *fakeRepository) GetOne() string {
	return r.val
}

func (r *fakeRepository) GetAll() []string {
	return r.data
}
