package farmsage

type Animal struct {
	AnimalID    int    `json:"id" db:"AnimalID"`
	Name        string `json:"name" binding:"required" db:"Name"`
	Number      int    `json:"number" db:"Number"`
	DateOfBirth string `json:"date_of_birth" db:"DateOfBirth"`
	Sex         string `json:"sex" db:"Sex"`
	Age         int    `json:"age" db:"Age"`
	MedicalInfo string `json:"medical_info" db:"MedicalInfo"`
}

type Activity struct {
	ActivityID   int    `json:"id"`
	AnimalID     int    `json:"animalid"`
	ActivityType string `json:"activitytype"`
	StartTime    string `json:"starttime"`
	EndTime      string `json:"endtime"`
	Latitude     int    `json:"latitude"`
	Longitude    int    `json:"longitude"`
}

type FeedingSchedule struct {
	ScheduleID  int    `json:"id" db:"ScheduleID"`
	AnimalID    int    `json:"animalid" db:"AnimalID"`
	FeedID      int    `json:"feedid" db:"FeedID"`
	FeedingTime string `json:"feedingtime" db:"FeedingTime"`
}

type Feed struct {
	FeedID   int    `json:"id" db:"FeedID"`
	Name     string `json:"name" binding:"required" db:"Name"`
	Quantity int    `json:"quantity" db:"Quantity"`
}

type Biometrics struct {
	BiometricID   int     `json:"id"`
	AnimalID      int     `json:"animalid"`
	Pulse         int     `json:"pulse"`
	Temperature   float64 `json:"temperature"`
	Weight        float64 `json:"weight"`
	BreathingRate int     `json:"breathingrate"`
}

type Farm struct {
	FarmID   int    `json:"FarmID" db:"FarmID"`
	AnimalID int    `json:"AnimalID"`
	UserID   int    `json:"UserID"`
	Name     string `json:"Name" binding:"required" db:"Name"`
}
