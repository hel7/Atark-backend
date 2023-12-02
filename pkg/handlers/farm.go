package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) getUserFarms(c *gin.Context) {

}

func (h *Handlers) getFarmByID(c *gin.Context) {

}

func (h *Handlers) createFarm(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handlers) updateFarm(c *gin.Context) {

}

func (h *Handlers) deleteFarm(c *gin.Context) {

}

func (h *Handlers) getAnimalsOnFarm(c *gin.Context) {

}
func (h *Handlers) getAnimalByID(c *gin.Context) {

}
func (h *Handlers) addAnimalToFarm(c *gin.Context) {

}

func (h *Handlers) removeAnimalFromFarm(c *gin.Context) {

}

func (h *Handlers) getFeedingScheduleByFarm(c *gin.Context) {

}

func (h *Handlers) createFeedingSchedule(c *gin.Context) {

}

func (h *Handlers) updateFeedingSchedule(c *gin.Context) {

}

func (h *Handlers) deleteFeedingSchedule(c *gin.Context) {

}

func (h *Handlers) getFeedsOnFarm(c *gin.Context) {

}

func (h *Handlers) addFeedToFarm(c *gin.Context) {

}

func (h *Handlers) removeFeedFromFarm(c *gin.Context) {

}
func (h *Handlers) addFeedToAnimalSchedule(c *gin.Context) {

}
func (h *Handlers) removeFeedFromAnimalSchedule(c *gin.Context) {

}
func (h *Handlers) updateFeedInAnimalSchedule(c *gin.Context) {

}
func (h *Handlers) getAnimalFeedSchedule(c *gin.Context) {

}
func (h *Handlers) getAnalytics(c *gin.Context) {

}

func (h *Handlers) getAnalyticsByDate(c *gin.Context) {

}

func (h *Handlers) getAdminUsers(c *gin.Context) {

}

func (h *Handlers) createAdminUser(c *gin.Context) {

}

func (h *Handlers) getAdminUserByID(c *gin.Context) {

}

func (h *Handlers) updateAdminUser(c *gin.Context) {

}

func (h *Handlers) deleteAdminUser(c *gin.Context) {

}

func (h *Handlers) getUsers(c *gin.Context) {

}

func (h *Handlers) createUser(c *gin.Context) {

}

func (h *Handlers) getUserByID(c *gin.Context) {

}

func (h *Handlers) updateUser(c *gin.Context) {

}

func (h *Handlers) deleteUser(c *gin.Context) {

}

func (h *Handlers) backupData(c *gin.Context) {

}

func (h *Handlers) restoreData(c *gin.Context) {

}

func (h *Handlers) exportData(c *gin.Context) {

}

func (h *Handlers) importData(c *gin.Context) {

}
