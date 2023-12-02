package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hel7/Atark-backend/pkg/service"
)

type Handlers struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.registerUser)
		auth.POST("/login", h.loginUser)
	}

	api := router.Group("/api", h.userIndetity)
	{

		api.GET("/farms", h.getUserFarms)
		api.GET("/farms/:farmID", h.getFarmByID)
		api.POST("/farms", h.createFarm)
		api.PUT("/farms/:farmID", h.updateFarm)
		api.DELETE("/farms/:farmID", h.deleteFarm)

		farms := api.Group("/farms/:farmID")
		{
			farms.GET("/animals", h.getAnimalsOnFarm)
			farms.POST("/animals", h.addAnimalToFarm)
			farms.GET("/animals/:animalID", h.getAnimalByID)
			farms.DELETE("/animals/:animalID", h.removeAnimalFromFarm)

			farms.GET("/feeds", h.getFeedsOnFarm)
			farms.POST("/feeds", h.addFeedToFarm)
			farms.DELETE("/feeds/:feedID", h.removeFeedFromFarm)

			farms.POST("/animals/:animalID/feeds/:feedID/schedule", h.addFeedToAnimalSchedule)
			farms.DELETE("/animals/:animalID/feeds/:feedID/schedule", h.deleteFeedingSchedule)
			farms.PUT("/animals/:animalID/feeds/:feedID/schedule", h.updateFeedInAnimalSchedule)
			farms.GET("/animals/:animalID/feeds/:feedID/schedule", h.getAnimalFeedSchedule)
		}

		analytics := api.Group("/analytics")
		{
			analytics.GET("/", h.getAnalytics)
			analytics.GET("/:date", h.getAnalyticsByDate)
		}

		admin := api.Group("/admin")
		{
			users := api.Group("/users")
			{
				users.GET("/", h.getUsers)
				users.POST("/", h.createUser)
				users.GET("/:id", h.getUserByID)
				users.PUT("/:id", h.updateUser)
				users.DELETE("/:id", h.deleteUser)
			}

			data := admin.Group("/data")
			{
				data.POST("/backup", h.backupData)
				data.POST("/restore", h.restoreData)
				data.GET("/export", h.exportData)
				data.POST("/import", h.importData)
			}
		}
	}

	return router
}
