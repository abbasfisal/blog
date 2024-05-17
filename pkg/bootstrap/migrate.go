package bootstrap

import (
	"github.com/abbasfisal/blog/internal/database/migration"
	"github.com/abbasfisal/blog/pkg/config"
	"github.com/abbasfisal/blog/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()

	migration.Migrate()
}
