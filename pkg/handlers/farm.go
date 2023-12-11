package handlers

import (
	"github.com/gin-gonic/gin"
	farmsage "github.com/hel7/Atark-backend"
	"github.com/sirupsen/logrus"

	"net/http"
	"strconv"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
type FarmResponse struct {
	FarmID int    `json:"farmID"`
	Name   string `json:"name"`
}
type AnimalResponse struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	DateOfBirth string `json:"dateOfBirth"`
	Sex         string `json:"sex"`
	Age         int    `json:"age"`
	MedicalInfo string `json:"medicalInfo"`
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
			Name:   farm.FarmName,
		}
		farmsResponse = append(farmsResponse, farmResponse)
	}

	c.JSON(http.StatusOK, gin.H{"farms": farmsResponse})
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
		Name:   farm.FarmName,
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
	createdFarmID, err := h.services.Farms.Create(UserID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"FarmID": createdFarmID,
	})
}

func (h *Handlers) updateFarm(c *gin.Context) {

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

	var animalsResponse []AnimalResponse
	for _, animal := range animals {
		animalResponse := AnimalResponse{
			Name:        animal.AnimalName,
			Number:      animal.Number,
			DateOfBirth: animal.DateOfBirth,
			Sex:         animal.Sex,
			Age:         animal.Age,
			MedicalInfo: animal.MedicalInfo,
		}
		animalsResponse = append(animalsResponse, animalResponse)
	}

	c.JSON(http.StatusOK, gin.H{"animals": animalsResponse})
}
func (h *Handlers) getAnimalByID(c *gin.Context) {
	UserID, err := getUserID(c)
	animalIDStr := c.Param("animalID")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	animal, err := h.services.Animals.GetByID(UserID, animalID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, animal)
}
func (h *Handlers) addAnimalToFarm(c *gin.Context) {
	UserID, err := getUserID(c)
	if err != nil {
		return
	}

	var input farmsage.Animal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	createdAnimalID, err := h.services.Animals.Create(UserID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"AnimalID": createdAnimalID,
	})
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

	feeds, err := h.services.Feed.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, feeds)
}

func (h *Handlers) addFeedToFarm(c *gin.Context) {
	var feed farmsage.Feed
	if err := c.ShouldBindJSON(&feed); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}

	createdFeedID, err := h.services.Feed.Create(feed)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"FeedID": createdFeedID,
	})
}

func (h *Handlers) removeFeedFromFarm(c *gin.Context) {

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
func (h *Handlers) removeAnimalFeedSchedule(c *gin.Context) {

}
func (h *Handlers) updateAnimalFeedSchedule(c *gin.Context) {

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

func (h *Handlers) getAnalytics(c *gin.Context) {

}

func (h *Handlers) getAnalyticsByDate(c *gin.Context) {

}

func (h *Handlers) updateAdminUser(c *gin.Context) {

}

func (h *Handlers) deleteAdminUser(c *gin.Context) {

}

func (h *Handlers) getUsers(c *gin.Context) {
	users, err := h.services.Admin.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var usersResponse []UserResponse
	for _, user := range users {
		userResponse := UserResponse{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{"users": usersResponse})
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
	userIDStr := c.Param("UserID")
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

	userResponse := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	c.JSON(http.StatusOK, userResponse)
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
