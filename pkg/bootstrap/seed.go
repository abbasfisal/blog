package bootstrap

import (
	"github.com/abbasfisal/blog/internal/database/seeder"
	"github.com/abbasfisal/blog/pkg/config"
	"github.com/abbasfisal/blog/pkg/database"
)

func Seed() {
	config.Set()
	database.Connect()

	seeder.Seed()
}
