package farmsage

import "errors"

type Animal struct {
	AnimalID    int    `json:"AnimalID" db:"AnimalID"`
	AnimalName  string `json:"AnimalName" binding:"required" db:"AnimalName"`
	Number      int    `json:"Number" db:"Number"`
	DateOfBirth string `json:"DateOfBirth" db:"DateOfBirth"`
	Sex         string `json:"Sex" db:"Sex"`
	Age         int    `json:"Age" db:"Age"`
	MedicalInfo string `json:"MedicalInfo" db:"MedicalInfo"`
	FarmName    string `json:"FarmName" db:"FarmName"`
}

type Activity struct {
	ActivityID   int    `json:"id"`
	AnimalID     int    `json:"animal_id"`
	ActivityType string `json:"activity_type"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Latitude     int    `json:"latitude"`
	Longitude    int    `json:"longitude"`
}

type FeedingSchedule struct {
	ScheduleID   int    `json:"ScheduleID" db:"ScheduleID"`
	AnimalID     int    `json:"AnimalID" db:"AnimalID"`
	FeedID       int    `json:"FeedID" db:"FeedID"`
	AnimalName   string `db:"AnimalName"`
	AnimalNumber int    `db:"Number"`
	FeedName     string `db:"FeedName"`
	FeedingTime  string `json:"FeedingTime" db:"FeedingTime"`
}

type Feed struct {
	FeedID   int    `json:"FeedID" db:"FeedID"`
	FeedName string `json:"FeedName" binding:"required" db:"FeedName"`
	Quantity int    `json:"Quantity" db:"Quantity"`
}

type Biometrics struct {
	BiometricID   int     `json:"id"`
	AnimalID      int     `json:"animal_id"`
	Pulse         int     `json:"pulse"`
	Temperature   float64 `json:"temperature"`
	Weight        float64 `json:"weight"`
	BreathingRate int     `json:"breathing_rate"`
}

type Farm struct {
	FarmID   int    `json:"FarmID" db:"FarmID"`
	AnimalID int    `json:"AnimalID"`
	UserID   int    `json:"UserID"`
	FarmName string `json:"FarmName" binding:"required" db:"FarmName"`
}
type FarmAnimal struct {
	FarmID   int `json:"FarmID" db:"FarmID"`
	AnimalID int `json:"AnimalID" db:"AnimalID"`
}
type UpdateFarmInput struct {
	FarmName *string `json:"FarmName"`
}
type UpdateFeedInput struct {
	FeedName *string `json:"FeedName"`
	Quantity *int    `json:"Quantity"`
}
type UpdateAnimalInput struct {
	AnimalName  *string `json:"AnimalName"`
	Number      *int    `json:"Number"`
	DateOfBirth *string `json:"DateOfBirth"`
	Sex         *string `json:"Sex"`
	Age         *int    `json:"Age"`
	MedicalInfo *string `json:"MedicalInfo"`
	FarmName    *string `json:"FarmName"`
}
type UpdateFeedingScheduleInput struct {
	ScheduleID   *int    `json:"ScheduleID"`
	AnimalID     *int    `json:"AnimalID"`
	FeedID       *int    `json:"FeedID"`
	AnimalName   *string `json:"AnimalName"`
	AnimalNumber *int    `json:"Number"`
	FeedName     *string `json:"FeedName"`
	FeedingTime  *string `json:"FeedingTime"`
}

func (i UpdateFarmInput) Validate() error {
	if i.FarmName == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}
func (i UpdateFeedInput) Validate() error {
	if i.FeedName == nil && i.Quantity == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}

func (i UpdateAnimalInput) Validate() error {
	if i.AnimalName == nil && i.Number == nil && i.DateOfBirth == nil && i.Sex == nil && i.Age == nil && i.MedicalInfo == nil && i.FarmName == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}

func (i UpdateFeedingScheduleInput) Validate() error {
	if i.ScheduleID == nil && i.AnimalID == nil && i.FeedID == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}
