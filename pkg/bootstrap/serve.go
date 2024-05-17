package bootstrap

import (
	"github.com/abbasfisal/blog/pkg/config"
	"github.com/abbasfisal/blog/pkg/database"
	"github.com/abbasfisal/blog/pkg/html"
	"github.com/abbasfisal/blog/pkg/routing"
	"github.com/abbasfisal/blog/pkg/sessions"
	"github.com/abbasfisal/blog/pkg/static"
)

func Serve() {
	config.Set()
	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())
	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()
	routing.Serve()

}
