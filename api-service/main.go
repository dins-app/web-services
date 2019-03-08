package main

import (
	"os"

	"github.com/labstack/echo"
	microclient "github.com/micro/go-micro/client"
	"github.com/team-morpheus/lasagna-msa/api-service/controllers"
	"github.com/team-morpheus/lasagna-msa/api-service/models"
	iPb "github.com/team-morpheus/lasagna-msa/identity-service/proto"
	irPb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

func main() {
	api := models.API{}

	// connect to services
	api.IrSvc = irPb.NewInternalRecipesService("lasagna.internal.recipes.service", microclient.DefaultClient)
	api.ISvc = iPb.NewIdentityService("lasagna.identity.service", microclient.DefaultClient)

	// create new instance of echo server
	api.Echo = echo.New()

	// register routes
	controllers.RegisterRecipesRoutes(&api)
	controllers.RegisterUsersRoutes(&api)

	// start http server wrapped with fatal helper func
	api.Echo.Logger.Fatal(api.Echo.Start(os.Getenv("PORT")))
}