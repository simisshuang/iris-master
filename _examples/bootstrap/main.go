package main

import (
	"github.com/kataras/iris/v12/_examples/bootstrap/bootstrap"
	"github.com/kataras/iris/v12/_examples/bootstrap/middleware/identity"
	"github.com/kataras/iris/v12/_examples/bootstrap/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
