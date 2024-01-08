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
		auth.POST("/registerAdmin", h.registerAdmin)
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
			farms.GET("/feeds", h.getFeedsOnFarm)
			farms.POST("/feeds", h.addFeedToFarm)
			farms.PUT("/feeds/:feedID", h.updateFeed)
			farms.DELETE("/feeds/:feedID", h.removeFeedFromFarm)

			farms.GET("/animals", h.getAnimalsOnFarm)
			farms.POST("/animals", h.addAnimalToFarm)
			farms.GET("/animals/:animalID", h.getAnimalByID)
			farms.PUT("/animals/:animalID", h.updateAnimal)
			farms.DELETE("/animals/:animalID", h.removeAnimalFromFarm)

			farms.POST("/animals/:animalID/activity", h.addActivity)
			farms.GET("/animals/:animalID/activity", h.getActivityByAnimalID)
			farms.DELETE("/animals/:animalID/activity/:activityID", h.deleteActivityByAnimalID)

			farms.POST("/animals/:animalID/biometrics", h.addBiometrics)
			farms.GET("/animals/:animalID/biometrics", h.getBiometricsByAnimalID)
			farms.DELETE("/animals/:animalID/biometrics/:biometricID", h.deleteBiometricsByAnimalID)

			farms.POST("/animals/:animalID/feeds/:feedID/", h.addAnimalFeedSchedule)
			farms.DELETE("/animals/:animalID/schedule/:scheduleID", h.deleteFeedingSchedule)
			farms.PUT("/animals/:animalID/schedule/:scheduleID/", h.updateAnimalFeedSchedule)
			farms.GET("/animals/:animalID/schedule", h.getAnimalFeedSchedule)
		}

		admin := api.Group("/admin", h.adminRequired)
		{
			users := admin.Group("/users")
			{
				users.GET("/", h.getUsers)
				users.POST("/", h.createUser)
				users.GET("/:userID", h.getUserByID)
				users.PUT("/:userID", h.updateUser)
				users.DELETE("/:userID", h.deleteUser)
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
