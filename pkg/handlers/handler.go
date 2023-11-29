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
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	animals := router.Group("/animals")
	{
		animals.GET("/", h.GetAnimals)
		animals.POST("/", h.CreateAnimal)
		animals.GET("/:id", h.GetAnimalByID)
		animals.PUT("/:id", h.UpdateAnimal)
		animals.DELETE("/:id", h.DeleteAnimal)
	}

	farms := router.Group("/farms")
	{
		farms.GET("/", h.GetFarms)
		farms.POST("/", h.CreateFarm)
		farms.GET("/:id", h.GetFarmByID)
		farms.PUT("/:id", h.UpdateFarm)
		farms.DELETE("/:id", h.DeleteFarm)
	}

	feed := router.Group("/feed")
	{
		feed.GET("/", h.GetFeed)
		feed.POST("/", h.CreateFeed)
		feed.GET("/:id", h.GetFeedByID)
		feed.PUT("/:id", h.UpdateFeed)
		feed.DELETE("/:id", h.DeleteFeed)
	}

	feedingSchedule := router.Group("/feeding-schedule")
	{
		feedingSchedule.GET("/", h.GetFeedingSchedule)
		feedingSchedule.POST("/", h.CreateFeedingSchedule)
		feedingSchedule.GET("/:id", h.GetFeedingScheduleByID)
		feedingSchedule.PUT("/:id", h.UpdateFeedingSchedule)
		feedingSchedule.DELETE("/:id", h.DeleteFeedingSchedule)
	}

	analytics := router.Group("/analytics")
	{
		analytics.GET("/", h.GetAnalytics)
		analytics.GET("/:date", h.GetAnalyticsByDate)
	}

	admin := router.Group("/admin")
	{
		adminUsers := admin.Group("/users")
		{
			adminUsers.GET("/", h.GetUsers)
			adminUsers.POST("/", h.CreateUser)
			adminUsers.GET("/:id", h.GetUserByID)
			adminUsers.PUT("/:id", h.UpdateUser)
			adminUsers.DELETE("/:id", h.DeleteUser)
		}

		adminRoles := admin.Group("/roles")
		{
			adminRoles.GET("/", h.GetRoles)
			adminRoles.POST("/", h.CreateRole)
			adminRoles.GET("/:id", h.GetRoleByID)
			adminRoles.PUT("/:id", h.UpdateRole)
			adminRoles.DELETE("/:id", h.DeleteRole)
		}
		adminData := admin.Group("/data")
		{
			adminData.POST("/backup", h.CreateBackup)
			adminData.POST("/restore", h.RestoreBackup)
			adminData.GET("/export", h.ExportData)
			adminData.POST("/import", h.ImportData)
		}

	}
	return router
}
