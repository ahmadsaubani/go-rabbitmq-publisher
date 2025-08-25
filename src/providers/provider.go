package providers

import (
	"publisher-topic/src/controllers"
	"publisher-topic/src/services"
)

type AppProvider struct {
	Services    services.ServiceProvider
	Controllers controllers.BaseController
}

func Register() AppProvider {
	allServices := services.InitServices()
	allControllers := controllers.InitControllers(allServices)

	return AppProvider{
		Services:    allServices,
		Controllers: allControllers,
	}
}
