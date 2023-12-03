package handlers

import (
	"github.com/gin-gonic/gin"
	farmsage "github.com/hel7/Atark-backend"
	"net/http"
	"strconv"
)

type getAllFarms struct {
	Data []farmsage.Farm `json:"data"`
}

type FarmResponse struct {
	FarmID int    `json:"FarmID"`
	Name   string `json:"Name"`
}

func (h *Handlers) getUserFarms(c *gin.Context) {
	UserID, err := getUserID(c)
	if err != nil {
		return
	}
	farms, err := h.services.Farms.GetAll(UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var farmsResponse []FarmResponse
	for _, farm := range farms {
		farmResponse := FarmResponse{
			FarmID: farm.FarmID,
			Name:   farm.Name,
		}
		farmsResponse = append(farmsResponse, farmResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": farmsResponse})
}

func (h *Handlers) getFarmByID(c *gin.Context) {
	UserID, err := getUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("farmID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	farm, err := h.services.Farms.GetByID(UserID, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	farmResponse := FarmResponse{
		FarmID: farm.FarmID,
		Name:   farm.Name,
	}

	c.JSON(http.StatusOK, farmResponse)
}

func (h *Handlers) createFarm(c *gin.Context) {
	UserID, err := getUserID(c)
	if err != nil {
		return
	}

	var input farmsage.Farm
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.services.Farms.Create(UserID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
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
