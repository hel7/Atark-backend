package handlers

import (
	"github.com/gin-gonic/gin"
	farmsage "github.com/hel7/Atark-backend"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

	c.JSON(http.StatusOK, gin.H{"farms": farms})
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

	c.JSON(http.StatusOK, farm)
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

	createdFarmID, err := h.services.Farms.Create(UserID, input)
	if err != nil {
		if strings.Contains(err.Error(), "farm with this name already exists") {
			newErrorResponse(c, http.StatusBadRequest, "Farm with this name already exists")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"FarmID": createdFarmID,
	})
}

func (h *Handlers) updateFarm(c *gin.Context) {
	UserID, err := getUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("farmID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	var input farmsage.UpdateFarmInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Farms.Update(UserID, id, input)
	if err != nil {
		if strings.Contains(err.Error(), "farm with this name already exists") {
			newErrorResponse(c, http.StatusBadRequest, "Farm with this name already exists")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) deleteFarm(c *gin.Context) {
	UserID, err := getUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("farmID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.Farms.Delete(UserID, id)
	if err != nil {
		logrus.Error("Failed to delete farm:", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) addFeedToFarm(c *gin.Context) {
	var feed farmsage.Feed
	if err := c.ShouldBindJSON(&feed); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}

	createdFeedID, err := h.services.Feed.Create(feed)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			newErrorResponse(c, http.StatusConflict, "feed with this name already exists")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"FeedID": createdFeedID,
	})
}
func (h *Handlers) removeFeedFromFarm(c *gin.Context) {

	feedIDStr := c.Param("feedID")
	feedID, err := strconv.Atoi(feedIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid feed ID")
		return
	}

	err = h.services.Feed.Delete(feedID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
func (h *Handlers) getAnimalsOnFarm(c *gin.Context) {
	farmIDStr := c.Param("farmID")
	farmID, err := strconv.Atoi(farmIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	animals, err := h.services.Animals.GetAll(farmID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"animals": animals})
}

func (h *Handlers) getAnimalByID(c *gin.Context) {
	farmIDStr := c.Param("farmID")
	farmID, err := strconv.Atoi(farmIDStr)
	animalIDStr := c.Param("animalID")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	animal, err := h.services.Animals.GetByID(farmID, animalID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, animal)
}
func (h *Handlers) addAnimalToFarm(c *gin.Context) {
	farmIDStr := c.Param("farmID")
	farmID, err := strconv.Atoi(farmIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid farm ID")
		return
	}

	var input farmsage.Animal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	createdAnimalID, err := h.services.Animals.Create(farmID, input)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate entry for Animal.Number") {
			newErrorResponse(c, http.StatusBadRequest, "Duplicate entry for Animal.Number")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"AnimalID": createdAnimalID,
	})
}

func (h *Handlers) removeAnimalFromFarm(c *gin.Context) {
	farmIDStr := c.Param("farmID")
	farmID, err := strconv.Atoi(farmIDStr)
	animalIDStr := c.Param("animalID")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Animals.Delete(farmID, animalID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
func (h *Handlers) addAnimalFeedSchedule(c *gin.Context) {
	var feedingSchedule farmsage.FeedingSchedule
	if err := c.ShouldBindJSON(&feedingSchedule); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}

	createdFeedID, err := h.services.FeedingSchedule.Create(feedingSchedule)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"FeedingScheduleID": createdFeedID,
	})
}

func (h *Handlers) updateAnimalFeedSchedule(c *gin.Context) {
	scheduleIDStr := c.Param("scheduleID")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid scheduleID")
		return
	}

	var input farmsage.UpdateFeedingScheduleInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.FeedingSchedule.Update(scheduleID, input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
func (h *Handlers) getAnimalFeedSchedule(c *gin.Context) {
	animalIDStr := c.Param("animalID")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid animal ID"})
		return
	}

	feedingSchedule, err := h.services.FeedingSchedule.GetByID(animalID)
	if err != nil {
		logrus.Error("Failed to get feeding schedule: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feeding schedule"})
		return
	}

	c.JSON(http.StatusOK, feedingSchedule)

}
func (h *Handlers) deleteFeedingSchedule(c *gin.Context) {
	scheduleIDStr := c.Param("scheduleID")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid scheduleID")
		return
	}

	err = h.services.FeedingSchedule.Delete(scheduleID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) getFeedsOnFarm(c *gin.Context) {

	feeds, err := h.services.Feed.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, feeds)
}

func (h *Handlers) getUsers(c *gin.Context) {
	users, err := h.services.Admin.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handlers) createUser(c *gin.Context) {
	var user farmsage.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := h.services.Admin.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": userID})
}

func (h *Handlers) getUserByID(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid UserID")
		return
	}

	user, err := h.services.Admin.GetUserByID(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handlers) updateUser(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid UserID")
		return
	}

	var input farmsage.UpdateUserInput
	var userInput farmsage.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Admin.UpdateUser(userID, input, userInput)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) deleteUser(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid UserID")
		return
	}

	err = h.services.Admin.Delete(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) updateFeed(c *gin.Context) {
	FeedID := c.Param("feedID")
	id, err := strconv.Atoi(FeedID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input farmsage.UpdateFeedInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Feed.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			newErrorResponse(c, http.StatusConflict, "feed with this name already exists")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) updateAnimal(c *gin.Context) {
	animalID := c.Param("animalID")
	id, err := strconv.Atoi(animalID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input farmsage.UpdateAnimalInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Animals.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate entry for Animal.Number") {
			newErrorResponse(c, http.StatusBadRequest, "Duplicate entry for Animal.Number")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) backupData(c *gin.Context) {
	backupPath := "backup.sql"

	file, err := os.Create(backupPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create backup file",
		})
		return
	}
	defer file.Close()

	if err := h.services.Admin.BackupData(backupPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to backup data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data backup successful",
	})
}
func (h *Handlers) restoreData(c *gin.Context) {
	restorePath := "backup.sql"

	if err := h.services.Admin.RestoreData(restorePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to restore data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data restore successful",
	})
}

func (h *Handlers) exportData(c *gin.Context) {
	exportPath := "export.xlsx"

	if err := h.services.Admin.ExportData(exportPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data export successful"})
}
func (h *Handlers) importData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadPath := file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	if err := h.services.Admin.ImportData(uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
}
